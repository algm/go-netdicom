package pdu

import (
	"encoding/binary"
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// Defines A_ASSOCIATE_RQ. P3.8 9.3.2 and 9.3.3
type AAssociateRQ struct {
	ProtocolVersion uint16
	// Reserved uint16
	CalledAETitle  string
	CallingAETitle string
	Items          []SubItem
}

func (AAssociateRQ) Read(d *dicomio.Decoder) PDU {
	pdu := &AAssociateRQ{}
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

func (pdu *AAssociateRQ) Write() ([]byte, error) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	if pdu.CalledAETitle == "" || pdu.CallingAETitle == "" {
		panic(*pdu)
	}
	e.WriteUInt16(pdu.ProtocolVersion)
	e.WriteZeros(2) // Reserved
	e.WriteString(fillString(pdu.CalledAETitle))
	e.WriteString(fillString(pdu.CallingAETitle))
	e.WriteZeros(8 * 4)
	for _, item := range pdu.Items {
		item.Write(e)
	}
	return e.Bytes(), e.Error()
}

func (pdu *AAssociateRQ) String() string {
	return fmt.Sprintf("A_ASSOCIATE_RQ{version:%v called:'%v' calling:'%v' items:%s}",
		pdu.ProtocolVersion,
		pdu.CalledAETitle, pdu.CallingAETitle, subItemListString(pdu.Items))
}
