package pdu

import (
	"encoding/binary"
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// Defines A_ASSOCIATE_AC. P3.8 9.3.2 and 9.3.3
type AAssociateAC struct {
	ProtocolVersion uint16
	// Reserved uint16
	CalledAETitle  string // For .._AC, the value is copied from A_ASSOCIATE_RQ
	CallingAETitle string // For .._AC, the value is copied from A_ASSOCIATE_RQ
	Items          []SubItem
}

func (AAssociateAC) Read(d *dicomio.Decoder) PDU {
	pdu := &AAssociateAC{}
	pdu.ProtocolVersion = d.ReadUInt16()
	d.Skip(2) // Reserved
	pdu.CalledAETitle = d.ReadString(16)
	pdu.CallingAETitle = d.ReadString(16)
	d.Skip(8 * 4)
	for !d.EOF() {
		item := decodeSubItem(d)
		if d.Error() != nil {
			break
		}
		pdu.Items = append(pdu.Items, item)
	}
	if pdu.CalledAETitle == "" || pdu.CallingAETitle == "" {
		d.SetError(fmt.Errorf("A_ASSOCIATE.{Called,Calling}AETitle must not be empty, in %v", pdu.String()))
	}
	return pdu
}

func (pdu *AAssociateAC) Write() ([]byte, error) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	if pdu.CalledAETitle == "" || pdu.CallingAETitle == "" {
		panic(*pdu)
	}
	e.WriteUInt16(pdu.ProtocolVersion)
	e.WriteZeros(2) // Reserved
	e.WriteString(fillString(pdu.CalledAETitle, 16))
	e.WriteString(fillString(pdu.CallingAETitle, 16))
	e.WriteZeros(8 * 4)
	for _, item := range pdu.Items {
		item.Write(e)
	}
	return e.Bytes(), e.Error()
}

func (pdu *AAssociateAC) String() string {
	return fmt.Sprintf("A_ASSOCIATE_AC{version:%v called:'%v' calling:'%v' items:%s}",
		pdu.ProtocolVersion,
		pdu.CalledAETitle, pdu.CallingAETitle, subItemListString(pdu.Items))
}
