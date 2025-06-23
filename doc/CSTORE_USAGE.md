# C-STORE Usage Documentation

## Table of Contents
1. [Overview](#overview)
2. [C-STORE Client (Service User)](#c-store-client-service-user)
3. [C-STORE Server (Service Provider)](#c-store-server-service-provider)
4. [Complete Examples](#complete-examples)
5. [Error Handling](#error-handling)
6. [Best Practices](#best-practices)
7. [Transfer Syntaxes](#transfer-syntaxes)
8. [Troubleshooting](#troubleshooting)

## Overview

C-STORE is a DICOM service used to store (send) DICOM objects from one Application Entity (AE) to another. This implementation supports both client-side (sending DICOM files) and server-side (receiving DICOM files) operations.

### Key Features
- Full DICOM C-STORE protocol implementation
- Support for all standard DICOM storage SOP classes
- Configurable transfer syntaxes
- Comprehensive error handling
- Thread-safe operations

## C-STORE Client (Service User)

The C-STORE client allows you to send DICOM files to a remote DICOM server.

### Basic Client Setup

```go
package main

import (
    "log"
    "github.com/grailbio/go-dicom"
    "github.com/mlibanori/go-netdicom"
    "github.com/mlibanori/go-netdicom/sopclass"
)

func main() {
    // Create ServiceUser with storage SOP classes
    su, err := netdicom.NewServiceUser(netdicom.ServiceUserParams{
        CalledAETitle:  "REMOTE_AE",     // Remote server's AE title
        CallingAETitle: "MY_CLIENT_AE",   // Your client's AE title
        SOPClasses:     sopclass.StorageClasses, // Support all storage classes
    })
    if err != nil {
        log.Fatal("Failed to create ServiceUser:", err)
    }
    defer su.Release() // Always release resources

    // Connect to DICOM server
    err = su.Connect("localhost:11112")
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }

    // Read DICOM file
    dataset, err := dicom.ReadDataSetFromFile("example.dcm", dicom.ReadOptions{})
    if err != nil {
        log.Fatal("Failed to read DICOM file:", err)
    }

    // Send file using C-STORE
    err = su.CStore(dataset)
    if err != nil {
        log.Fatal("C-STORE failed:", err)
    }

    log.Println("C-STORE completed successfully")
}
```

### ServiceUserParams Configuration

```go
params := netdicom.ServiceUserParams{
    CalledAETitle:  "DESTINATION_AE", // Target server's AE title
    CallingAETitle: "SOURCE_AE",      // Your AE title
    SOPClasses:     sopclass.StorageClasses, // SOP classes to support
    TransferSyntaxes: []string{       // Optional: specify transfer syntaxes
        "1.2.840.10008.1.2",     // Implicit VR Little Endian
        "1.2.840.10008.1.2.1",   // Explicit VR Little Endian  
        "1.2.840.10008.1.2.2",   // Explicit VR Big Endian
    },
}
```

### Connection Methods

```go
// Method 1: Connect to server by address
su.Connect("server.example.com:11112")

// Method 2: Use existing connection
conn, err := net.Dial("tcp", "server.example.com:11112")
if err == nil {
    su.SetConn(conn)
}
```

### Multiple File Transfer

```go
func storeMultipleFiles(su *netdicom.ServiceUser, filePaths []string) error {
    for _, filePath := range filePaths {
        dataset, err := dicom.ReadDataSetFromFile(filePath, dicom.ReadOptions{})
        if err != nil {
            return fmt.Errorf("failed to read %s: %v", filePath, err)
        }

        err = su.CStore(dataset)
        if err != nil {
            return fmt.Errorf("failed to store %s: %v", filePath, err)
        }
        
        log.Printf("Successfully stored: %s", filePath)
    }
    return nil
}
```

## C-STORE Server (Service Provider)

The C-STORE server receives DICOM files from remote clients.

### Basic Server Setup

```go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    
    "github.com/grailbio/go-dicom"
    "github.com/grailbio/go-dicom/dicomio"
    "github.com/grailbio/go-dicom/dicomtag"
    "github.com/mlibanori/go-netdicom"
    "github.com/mlibanori/go-netdicom/dimse"
)

func main() {
    // Define server parameters
    params := netdicom.ServiceProviderParams{
        AETitle: "MY_STORAGE_SCP",  // Server's AE title
        CStore:  onCStoreRequest,   // C-STORE handler function
    }

    // Create and start server
    provider, err := netdicom.NewServiceProvider(params, ":11112")
    if err != nil {
        log.Fatal("Failed to create provider:", err)
    }

    log.Println("DICOM C-STORE server started on port 11112")
    provider.Run() // This blocks forever
}

// C-STORE request handler
func onCStoreRequest(
    conn netdicom.ConnectionState,
    transferSyntaxUID string,
    sopClassUID string,
    sopInstanceUID string,
    data []byte) dimse.Status {
    
    log.Printf("Received C-STORE request:")
    log.Printf("  Transfer Syntax: %s", transferSyntaxUID)
    log.Printf("  SOP Class UID: %s", sopClassUID)
    log.Printf("  SOP Instance UID: %s", sopInstanceUID)
    log.Printf("  Data size: %d bytes", len(data))

    // Generate filename
    filename := fmt.Sprintf("received_%s.dcm", sopInstanceUID)
    filepath := filepath.Join("./incoming", filename)

    // Ensure directory exists
    err := os.MkdirAll("./incoming", 0755)
    if err != nil {
        log.Printf("Failed to create directory: %v", err)
        return dimse.Status{
            Status: dimse.StatusProcessingFailure,
            ErrorComment: "Failed to create storage directory",
        }
    }

    // Create file
    file, err := os.Create(filepath)
    if err != nil {
        log.Printf("Failed to create file: %v", err)
        return dimse.Status{
            Status: dimse.StatusProcessingFailure,
            ErrorComment: "Failed to create output file",
        }
    }
    defer file.Close()

    // Write DICOM file with proper header
    encoder := dicomio.NewEncoderWithTransferSyntax(file, transferSyntaxUID)
    
    // Write DICOM file header
    dicom.WriteFileHeader(encoder, []*dicom.Element{
        dicom.MustNewElement(dicomtag.TransferSyntaxUID, transferSyntaxUID),
        dicom.MustNewElement(dicomtag.MediaStorageSOPClassUID, sopClassUID),
        dicom.MustNewElement(dicomtag.MediaStorageSOPInstanceUID, sopInstanceUID),
    })
    
    // Write the actual DICOM data
    encoder.WriteBytes(data)
    
    if err := encoder.Error(); err != nil {
        log.Printf("Failed to write DICOM file: %v", err)
        return dimse.Status{
            Status: dimse.StatusProcessingFailure,
            ErrorComment: "Failed to encode DICOM data",
        }
    }

    log.Printf("Successfully stored file: %s", filepath)
    return dimse.Success
}
```

### Advanced Server Configuration

```go
params := netdicom.ServiceProviderParams{
    AETitle:   "ADVANCED_SCP",
    CStore:    onCStoreRequest,
    CEcho:     onCEchoRequest,     // Optional: Handle C-ECHO
    TLSConfig: &tls.Config{...},   // Optional: Enable TLS
}
```

### Database Integration Example

```go
func onCStoreRequestWithDB(
    conn netdicom.ConnectionState,
    transferSyntaxUID string,
    sopClassUID string,
    sopInstanceUID string,
    data []byte) dimse.Status {

    // Parse DICOM data to extract metadata
    reader := bytes.NewReader(data)
    decoder := dicomio.NewDecoderWithTransferSyntax(reader, transferSyntaxUID)
    
    var patientID, studyUID, seriesUID string
    
    // Read DICOM elements to extract metadata
    for !decoder.IsEOF() {
        element, err := dicom.ReadElement(decoder, &dicom.ReadOptions{})
        if err != nil {
            break
        }
        
        switch element.Tag {
        case dicomtag.PatientID:
            patientID, _ = element.GetString()
        case dicomtag.StudyInstanceUID:
            studyUID, _ = element.GetString()
        case dicomtag.SeriesInstanceUID:
            seriesUID, _ = element.GetString()
        }
    }

    // Store file with organized directory structure
    dir := filepath.Join("./dicom_storage", patientID, studyUID, seriesUID)
    err := os.MkdirAll(dir, 0755)
    if err != nil {
        return dimse.Status{Status: dimse.StatusProcessingFailure}
    }

    filename := fmt.Sprintf("%s.dcm", sopInstanceUID)
    fullPath := filepath.Join(dir, filename)

    // Save file (implementation similar to basic example)
    // ... file saving code ...

    // Insert metadata into database
    err = insertIntoDatabase(patientID, studyUID, seriesUID, sopInstanceUID, fullPath)
    if err != nil {
        log.Printf("Database insertion failed: %v", err)
        // File saved but DB failed - you might want to handle this differently
    }

    return dimse.Success
}
```

## Complete Examples

### Complete Client Example

```go
package main

import (
    "flag"
    "log"
    "path/filepath"
    "strings"

    "github.com/grailbio/go-dicom"
    "github.com/mlibanori/go-netdicom"
    "github.com/mlibanori/go-netdicom/sopclass"
)

var (
    serverAddr = flag.String("server", "localhost:11112", "DICOM server address")
    aeTitle    = flag.String("ae", "CLIENT_AE", "Client AE title")
    remoteAE   = flag.String("remote-ae", "SERVER_AE", "Server AE title")
    inputFile  = flag.String("file", "", "DICOM file to send")
)

func main() {
    flag.Parse()
    
    if *inputFile == "" {
        log.Fatal("Please specify a DICOM file with -file")
    }

    // Create service user
    su, err := netdicom.NewServiceUser(netdicom.ServiceUserParams{
        CalledAETitle:  *remoteAE,
        CallingAETitle: *aeTitle,
        SOPClasses:     sopclass.StorageClasses,
    })
    if err != nil {
        log.Fatal("Failed to create ServiceUser:", err)
    }
    defer su.Release()

    // Connect to server  
    log.Printf("Connecting to %s", *serverAddr)
    su.Connect(*serverAddr)

    // Test connection with C-ECHO
    err = su.CEcho()
    if err != nil {
        log.Fatal("C-ECHO failed:", err)
    }
    log.Println("C-ECHO successful")

    // Read and send DICOM file
    log.Printf("Reading DICOM file: %s", *inputFile)
    dataset, err := dicom.ReadDataSetFromFile(*inputFile, dicom.ReadOptions{})
    if err != nil {
        log.Fatal("Failed to read DICOM file:", err)
    }

    log.Println("Sending C-STORE request...")
    err = su.CStore(dataset)
    if err != nil {
        log.Fatal("C-STORE failed:", err)
    }

    log.Printf("Successfully sent %s", filepath.Base(*inputFile))
}
```

### Complete Server Example with Multiple Services

```go
package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"
    "sync"

    "github.com/mlibanori/go-netdicom"
    "github.com/mlibanori/go-netdicom/dimse"
)

var (
    port       = flag.String("port", "11112", "Port to listen on")
    aeTitle    = flag.String("ae", "STORAGE_SCP", "Server AE title")
    storageDir = flag.String("dir", "./storage", "Directory to store DICOM files")
)

type DicomServer struct {
    storageDir string
    mu         sync.Mutex
    fileCount  int
}

func (ds *DicomServer) onCEcho(conn netdicom.ConnectionState) dimse.Status {
    log.Println("Received C-ECHO request")
    return dimse.Success
}

func (ds *DicomServer) onCStore(
    conn netdicom.ConnectionState,
    transferSyntaxUID string,
    sopClassUID string,
    sopInstanceUID string,
    data []byte) dimse.Status {

    ds.mu.Lock()
    ds.fileCount++
    fileNum := ds.fileCount
    ds.mu.Unlock()

    log.Printf("[%d] C-STORE request received", fileNum)
    log.Printf("[%d] SOP Instance UID: %s", fileNum, sopInstanceUID)
    log.Printf("[%d] Data size: %d bytes", fileNum, len(data))

    // Ensure storage directory exists
    err := os.MkdirAll(ds.storageDir, 0755)
    if err != nil {
        log.Printf("[%d] Failed to create storage directory: %v", fileNum, err)
        return dimse.Status{Status: dimse.StatusProcessingFailure}
    }

    // Create unique filename
    filename := fmt.Sprintf("image_%06d_%s.dcm", fileNum, sopInstanceUID)
    fullPath := filepath.Join(ds.storageDir, filename)

    // Save DICOM file
    err = ds.saveDicomFile(fullPath, transferSyntaxUID, sopClassUID, sopInstanceUID, data)
    if err != nil {
        log.Printf("[%d] Failed to save file: %v", fileNum, err)
        return dimse.Status{Status: dimse.StatusProcessingFailure}
    }

    log.Printf("[%d] Successfully saved: %s", fileNum, filename)
    return dimse.Success
}

func (ds *DicomServer) saveDicomFile(filename, transferSyntaxUID, sopClassUID, sopInstanceUID string, data []byte) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := dicomio.NewEncoderWithTransferSyntax(file, transferSyntaxUID)
    
    // Write DICOM header
    dicom.WriteFileHeader(encoder, []*dicom.Element{
        dicom.MustNewElement(dicomtag.TransferSyntaxUID, transferSyntaxUID),
        dicom.MustNewElement(dicomtag.MediaStorageSOPClassUID, sopClassUID),
        dicom.MustNewElement(dicomtag.MediaStorageSOPInstanceUID, sopInstanceUID),
    })
    
    encoder.WriteBytes(data)
    return encoder.Error()
}

func main() {
    flag.Parse()

    dicomServer := &DicomServer{
        storageDir: *storageDir,
    }

    params := netdicom.ServiceProviderParams{
        AETitle: *aeTitle,
        CEcho:   dicomServer.onCEcho,
        CStore:  dicomServer.onCStore,
    }

    provider, err := netdicom.NewServiceProvider(params, ":"+*port)
    if err != nil {
        log.Fatal("Failed to create provider:", err)
    }

    log.Printf("DICOM server started")
    log.Printf("AE Title: %s", *aeTitle)
    log.Printf("Port: %s", *port)
    log.Printf("Storage Directory: %s", *storageDir)
    log.Println("Waiting for connections...")

    provider.Run()
}
```

## Error Handling

### Common Error Scenarios

```go
func robustCStore(su *netdicom.ServiceUser, dataset *dicom.DataSet) error {
    // Verify required DICOM elements exist
    sopClassElem, err := dataset.FindElementByTag(dicomtag.MediaStorageSOPClassUID)
    if err != nil {
        return fmt.Errorf("missing SOP Class UID: %v", err)
    }
    
    sopInstanceElem, err := dataset.FindElementByTag(dicomtag.MediaStorageSOPInstanceUID)
    if err != nil {
        return fmt.Errorf("missing SOP Instance UID: %v", err)
    }

    // Attempt C-STORE with retries
    maxRetries := 3
    for attempt := 1; attempt <= maxRetries; attempt++ {
        err = su.CStore(dataset)
        if err == nil {
            return nil // Success
        }
        
        log.Printf("C-STORE attempt %d failed: %v", attempt, err)
        
        if attempt < maxRetries {
            // Wait before retry
            time.Sleep(time.Duration(attempt) * time.Second)
        }
    }
    
    return fmt.Errorf("C-STORE failed after %d attempts: %v", maxRetries, err)
}
```

### Server Error Responses

```go
func onCStoreWithValidation(
    conn netdicom.ConnectionState,
    transferSyntaxUID string,
    sopClassUID string,
    sopInstanceUID string,
    data []byte) dimse.Status {

    // Validate SOP Class UID
    if !isValidSOPClass(sopClassUID) {
        return dimse.Status{
            Status: dimse.StatusSOPClassNotSupported,
            ErrorComment: "Unsupported SOP Class",
        }
    }

    // Validate data size
    if len(data) == 0 {
        return dimse.Status{
            Status: dimse.StatusInvalidArgumentValue,
            ErrorComment: "Empty DICOM data",
        }
    }

    // Check disk space
    if !hasSufficientDiskSpace(len(data)) {
        return dimse.Status{
            Status: dimse.StatusStorageOutOfResources,
            ErrorComment: "Insufficient disk space",
        }
    }

    // Proceed with storage...
    return dimse.Success
}
```

## Best Practices

### Client Best Practices

1. **Always Release Resources**
   ```go
   su, err := netdicom.NewServiceUser(params)
   if err != nil {
       return err  
   }
   defer su.Release() // Always call Release
   ```

2. **Test Connection First**
   ```go
   err = su.CEcho()
   if err != nil {
       return fmt.Errorf("server not responding: %v", err)
   }
   ```

3. **Validate DICOM Files**
   ```go
   dataset, err := dicom.ReadDataSetFromFile(filepath, dicom.ReadOptions{})
   if err != nil {
       return fmt.Errorf("invalid DICOM file: %v", err)
   }
   
   // Verify required elements exist
   if _, err := dataset.FindElementByTag(dicomtag.MediaStorageSOPClassUID); err != nil {
       return fmt.Errorf("missing SOP Class UID: %v", err)
   }
   ```

4. **Handle Large Files Efficiently**
   ```go
   // For large files, consider using streaming
   options := dicom.ReadOptions{
       DropPixelData: false, // Include pixel data
       ReturnTags:    nil,   // Read all tags
   }
   dataset, err := dicom.ReadDataSetFromFile(filepath, options)
   ```

### Server Best Practices

1. **Validate Incoming Data**
   ```go
   func onCStore(...) dimse.Status {
       // Validate SOP Class
       if !isValidSOPClass(sopClassUID) {
           return dimse.Status{Status: dimse.StatusSOPClassNotSupported}
       }
       
       // Validate Instance UID format
       if !isValidUID(sopInstanceUID) {
           return dimse.Status{Status: dimse.StatusInvalidArgumentValue}
       }
       
       // Continue processing...
   }
   ```

2. **Organize File Storage**
   ```go
   // Create hierarchical directory structure
   // /storage/PatientID/StudyUID/SeriesUID/InstanceUID.dcm
   dir := filepath.Join(baseDir, patientID, studyUID, seriesUID)
   os.MkdirAll(dir, 0755)
   ```

3. **Implement Proper Logging**
   ```go
   func onCStore(...) dimse.Status {
       start := time.Now()
       defer func() {
           duration := time.Since(start)
           log.Printf("C-STORE completed in %v", duration)
       }()
       
       log.Printf("C-STORE started: SOP Instance %s", sopInstanceUID)
       // Process...
   }
   ```

4. **Handle Concurrent Requests**
   ```go
   type Server struct {
       mu sync.RWMutex
       activeConnections map[net.Conn]bool
   }
   
   func (s *Server) onCStore(...) dimse.Status {
       s.mu.Lock()
       defer s.mu.Unlock()
       
       // Thread-safe operations...
   }
   ```

## Transfer Syntaxes

### Supported Transfer Syntaxes

The library supports standard DICOM transfer syntaxes:

```go
// Common transfer syntaxes
var supportedTransferSyntaxes = []string{
    "1.2.840.10008.1.2",     // Implicit VR Little Endian
    "1.2.840.10008.1.2.1",   // Explicit VR Little Endian
    "1.2.840.10008.1.2.2",   // Explicit VR Big Endian
    "1.2.840.10008.1.2.4.50", // JPEG Baseline
    "1.2.840.10008.1.2.4.90", // JPEG 2000
    "1.2.840.10008.1.2.5",   // RLE Lossless
}
```

### Configuring Transfer Syntaxes

```go
params := netdicom.ServiceUserParams{
    CalledAETitle:  "SERVER_AE",
    CallingAETitle: "CLIENT_AE", 
    SOPClasses:     sopclass.StorageClasses,
    TransferSyntaxes: []string{
        "1.2.840.10008.1.2.1", // Explicit VR Little Endian (preferred)
        "1.2.840.10008.1.2",   // Implicit VR Little Endian (fallback)
    },
}
```

## Troubleshooting

### Common Issues

1. **Connection Refused**
   ```
   Error: dial tcp: connection refused
   Solution: Verify server is running and port is correct
   ```

2. **Association Rejected**
   ```
   Error: Association rejected
   Solution: Check AE titles and SOP classes match server configuration
   ```

3. **Transfer Syntax Not Supported**
   ```
   Error: Transfer syntax not supported
   Solution: Add required transfer syntax to ServiceUserParams
   ```

4. **SOP Class Not Supported**
   ```
   Error: SOP class not found in context
   Solution: Include appropriate storage SOP classes in client configuration
   ```

### Debug Logging

Enable debug logging for troubleshooting:

```go
import "github.com/grailbio/go-dicom/dicomlog"

func init() {
    dicomlog.SetLevel(2) // 0=errors, 1=info, 2=debug
}
```

### Network Diagnostics

```go
// Test basic connectivity
conn, err := net.DialTimeout("tcp", "server:11112", 10*time.Second)
if err != nil {
    log.Fatal("Cannot connect to server:", err)
}
conn.Close()

// Test DICOM association
su, _ := netdicom.NewServiceUser(params)
defer su.Release()
su.Connect("server:11112")

err = su.CEcho()
if err != nil {
    log.Fatal("DICOM association failed:", err)
}
```

### File Validation

```go
func validateDicomFile(filepath string) error {
    dataset, err := dicom.ReadDataSetFromFile(filepath, dicom.ReadOptions{})
    if err != nil {
        return fmt.Errorf("not a valid DICOM file: %v", err)
    }

    // Check required elements
    requiredTags := []dicomtag.Tag{
        dicomtag.MediaStorageSOPClassUID,
        dicomtag.MediaStorageSOPInstanceUID,
        dicomtag.TransferSyntaxUID,
    }

    for _, tag := range requiredTags {
        if _, err := dataset.FindElementByTag(tag); err != nil {
            return fmt.Errorf("missing required tag %s: %v", tag, err)
        }
    }

    return nil
}
```

---

This documentation provides comprehensive coverage of C-STORE functionality for both client and server implementations. For additional examples, see the `sampleclient` and `sampleserver` directories in the source code.