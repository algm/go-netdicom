package dimse

//go:generate ./generate_dimse_messages.py
//go:generate stringer -type StatusCode

// Implements message types defined in P3.7.
//
// http://dicom.nema.org/medical/dicom/current/output/pdf/part07.pdf

import (
	"encoding/binary"
	"fmt"
	"sort"

	dicom "github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/grailbio/go-dicom/dicomlog"
	"github.com/grailbio/go-dicom/dicomtag"
	"github.com/mlibanori/go-netdicom/commandset"
)

// Message defines the common interface for all DIMSE message types.
type Message interface {
	fmt.Stringer // Print human-readable description for debugging.
	Encode(*dicomio.Encoder)
	// GetMessageID extracts the message ID field.
	GetMessageID() MessageID
	// CommandField returns the command field value of this message.
	CommandField() int
	// GetStatus returns the the response status value. It is nil for request message
	// types, and non-nil for response message types.
	GetStatus() *Status
	// HasData is true if we expect P_DATA_TF packets after the command packets.
	HasData() bool
}

// Status represents a result of a DIMSE call.  P3.7 C defines list of status
// codes and error payloads.
type Status struct {
	// Status==StatusSuccess on success. A non-zero value on error.
	Status StatusCode

	// Optional error payloads.
	ErrorComment string // Encoded as (0000,0902)
}

// Helper class for extracting values from a list of DicomElement.
type MessageDecoder struct {
	elems  []*dicom.Element
	parsed []bool // true if this element was parsed into a message field.
	err    error
}

// NewMessageDecoder creates a new MessageDecoder instance.
func NewMessageDecoder(elems []*dicom.Element) *MessageDecoder {
	return &MessageDecoder{
		elems:  elems,
		parsed: make([]bool, len(elems)),
		err:    nil,
	}
}

type isOptionalElement int

const (
	RequiredElement isOptionalElement = iota
	OptionalElement
)

func (d *MessageDecoder) SetError(err error) {
	if d.err == nil {
		d.err = err
	}
}

func (d *MessageDecoder) Error() error {
	return d.err
}

// Find an element with the given tag. If optional==OptionalElement, returns nil
// if not found.  If optional==RequiredElement, sets d.err and return nil if not found.
func (d *MessageDecoder) findElement(tag dicomtag.Tag, optional isOptionalElement) *dicom.Element {
	for i, elem := range d.elems {
		if elem.Tag == tag {
			dicomlog.Vprintf(3, "dimse.findElement: Return %v for %s", elem, tag.String())
			d.parsed[i] = true
			return elem
		}
	}
	if optional == RequiredElement {
		d.SetError(fmt.Errorf("dimse.findElement: Element %s not found during DIMSE decoding", dicomtag.DebugString(tag)))
	}
	return nil
}

// Return the list of elements that did not match any of the prior getXXX calls.
func (d *MessageDecoder) UnparsedElements() (unparsed []*dicom.Element) {
	for i, parsed := range d.parsed {
		if !parsed {
			unparsed = append(unparsed, d.elems[i])
		}
	}
	return unparsed
}

func (d *MessageDecoder) GetStatus() (s Status) {
	s.Status = StatusCode(d.GetUInt16(commandset.Status, RequiredElement))
	s.ErrorComment = d.GetString(commandset.ErrorComment, OptionalElement)
	return s
}

// Find an element with "tag", and extract a string value from it. Errors are reported in d.err.
func (d *MessageDecoder) GetString(tag dicomtag.Tag, optional isOptionalElement) string {
	e := d.findElement(tag, optional)
	if e == nil {
		return ""
	}
	v, err := e.GetString()
	if err != nil {
		d.SetError(err)
	}
	return v
}

// Find an element with "tag", and extract a uint16 from it. Errors are reported in d.err.
func (d *MessageDecoder) GetUInt16(tag dicomtag.Tag, optional isOptionalElement) uint16 {
	e := d.findElement(tag, optional)
	if e == nil {
		return 0
	}
	v, err := e.GetUInt16()
	if err != nil {
		d.SetError(err)
	}
	return v
}

// Encode the given elements. The elements are sorted in ascending tag order.
func EncodeElements(e *dicomio.Encoder, elems []*dicom.Element) {
	sort.Slice(elems, func(i, j int) bool {
		return elems[i].Tag.Compare(elems[j].Tag) < 0
	})
	for _, elem := range elems {
		dicom.WriteElement(e, elem)
	}
}

// Create a list of elements that represent the dimse status. The list contains
// multiple elements for non-ok status.
func NewStatusElements(s Status) []*dicom.Element {
	elems := []*dicom.Element{NewElement(commandset.Status, uint16(s.Status))}
	if s.ErrorComment != "" {
		elems = append(elems, NewElement(commandset.ErrorComment, s.ErrorComment))
	}
	return elems
}

// Create a new element. The value type must match the tag's.
func NewElement(tag dicomtag.Tag, v interface{}) *dicom.Element {
	return &dicom.Element{
		Tag:             tag,
		VR:              "", // autodetect
		UndefinedLength: false,
		Value:           []interface{}{v},
	}
}

// CommandDataSetTypeNull indicates that the DIMSE message has no data payload,
// when set in dicom.TagCommandDataSetType. Any other value indicates the
// existence of a payload.
const CommandDataSetTypeNull uint16 = 0x101

// CommandDataSetTypeNonNull indicates that the DIMSE message has a data
// payload, when set in dicom.TagCommandDataSetType.
const CommandDataSetTypeNonNull uint16 = 1

// Success is an OK status for a call.
var Success = Status{Status: StatusSuccess}

// StatusCode represents a DIMSE service response code, as defined in P3.7
type StatusCode uint16

const (
	StatusSuccess               StatusCode = 0
	StatusCancel                StatusCode = 0xFE00
	StatusSOPClassNotSupported  StatusCode = 0x0112
	StatusInvalidArgumentValue  StatusCode = 0x0115
	StatusInvalidAttributeValue StatusCode = 0x0106
	StatusInvalidObjectInstance StatusCode = 0x0117
	StatusUnrecognizedOperation StatusCode = 0x0211
	StatusNotAuthorized         StatusCode = 0x0124
	StatusPending               StatusCode = 0xff00

	// C-STORE-specific status codes. P3.4 GG4-1
	CStoreOutOfResources              StatusCode = 0xa700
	CStoreCannotUnderstand            StatusCode = 0xc000
	CStoreDataSetDoesNotMatchSOPClass StatusCode = 0xa900

	// C-FIND-specific status codes.
	CFindUnableToProcess StatusCode = 0xc000

	// C-MOVE/C-GET-specific status codes.
	CMoveOutOfResourcesUnableToCalculateNumberOfMatches StatusCode = 0xa701
	CMoveOutOfResourcesUnableToPerformSubOperations     StatusCode = 0xa702
	CMoveMoveDestinationUnknown                         StatusCode = 0xa801
	CMoveDataSetDoesNotMatchSOPClass                    StatusCode = 0xa900

	// Warning codes.
	StatusAttributeValueOutOfRange StatusCode = 0x0116
	StatusAttributeListError       StatusCode = 0x0107
)

// EncodeMessage serializes the given message. Errors are reported through e.Error()
func EncodeMessage(e *dicomio.Encoder, v Message) {
	// DIMSE messages are always encoded Implicit+LE. See P3.7 6.3.1.
	subEncoder := dicomio.NewBytesEncoder(binary.LittleEndian, dicomio.ImplicitVR)
	v.Encode(subEncoder)
	if err := subEncoder.Error(); err != nil {
		e.SetError(err)
		return
	}
	bytes := subEncoder.Bytes()
	e.PushTransferSyntax(binary.LittleEndian, dicomio.ImplicitVR)
	defer e.PopTransferSyntax()
	dicom.WriteElement(e, NewElement(commandset.CommandGroupLength, uint32(len(bytes))))
	e.WriteBytes(bytes)
}

// AddDataPDU is to be called for each P_DATA_TF PDU received from the
// network. If the fragment is marked as the last one, AddDataPDU returns
// <SOPUID, TransferSyntaxUID, payload, nil>.  If it needs more fragments, it
// returns <"", "", nil, nil>.  On error, it returns a non-nil error.

type MessageID = uint16
