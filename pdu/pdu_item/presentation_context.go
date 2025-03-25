package pdu_item

//go:generate stringer -type PresentationContextResult

import (
	"encoding/binary"
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// Result of abstractsyntax/transfersyntax handshake during A-ACCEPT.  P3.8,
// 90.3.3.2, table 9-18.
type PresentationContextResult byte

const (
	PresentationContextAccepted                                    PresentationContextResult = 0
	PresentationContextUserRejection                               PresentationContextResult = 1
	PresentationContextProviderRejectionNoReason                   PresentationContextResult = 2
	PresentationContextProviderRejectionAbstractSyntaxNotSupported PresentationContextResult = 3
	PresentationContextProviderRejectionTransferSyntaxNotSupported PresentationContextResult = 4
)

// P3.8 9.3.2.2, 9.3.3.2
type PresentationContextItem struct {
	Type      byte // ItemTypePresentationContext*
	ContextID byte
	// 1 byte reserved

	// Result is meaningful iff Type=0x21, zero else.
	Result PresentationContextResult

	// 1 byte reserved
	Items []SubItem // List of {Abstract,Transfer}SyntaxSubItem
}

func decodePresentationContextItem(d *dicomio.Decoder, itemType byte, length uint16) *PresentationContextItem {
	v := &PresentationContextItem{Type: itemType}
	d.PushLimit(int64(length))
	defer d.PopLimit()
	v.ContextID = d.ReadByte()
	d.Skip(1)
	v.Result = PresentationContextResult(d.ReadByte())
	d.Skip(1)
	for !d.EOF() {
		item := DecodeSubItem(d)
		if d.Error() != nil {
			break
		}
		v.Items = append(v.Items, item)
	}
	if v.ContextID%2 != 1 {
		d.SetError(fmt.Errorf("PresentationContextItem ID must be odd, but found %x", v.ContextID))
	}
	return v
}

func (v *PresentationContextItem) Write(e *dicomio.Encoder) {
	if v.Type != ItemTypePresentationContextRequest &&
		v.Type != ItemTypePresentationContextResponse {
		panic(*v)
	}

	itemEncoder := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	for _, s := range v.Items {
		s.Write(itemEncoder)
	}
	if err := itemEncoder.Error(); err != nil {
		e.SetError(err)
		return
	}
	itemBytes := itemEncoder.Bytes()
	encodeSubItemHeader(e, v.Type, uint16(4+len(itemBytes)))
	e.WriteByte(v.ContextID)
	e.WriteZeros(3)
	e.WriteBytes(itemBytes)
}

func (v *PresentationContextItem) String() string {
	itemType := "rq"
	if v.Type == ItemTypePresentationContextResponse {
		itemType = "ac"
	}
	return fmt.Sprintf("PresentationContext%s{id: %d result: %d, items:%s}",
		itemType, v.ContextID, v.Result, SubItemListString(v.Items))
}
