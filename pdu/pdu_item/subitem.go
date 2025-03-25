package pdu_item

import (
	"bytes"
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// SubItem is the interface for DUL items, such as ApplicationContextItem and
// TransferSyntaxSubItem.
type SubItem interface {
	fmt.Stringer

	// Write serializes the item.
	Write(*dicomio.Encoder)
}

// Possible Type field values for SubItem.
const (
	ItemTypeApplicationContext           = 0x10
	ItemTypePresentationContextRequest   = 0x20
	ItemTypePresentationContextResponse  = 0x21
	ItemTypeAbstractSyntax               = 0x30
	ItemTypeTransferSyntax               = 0x40
	ItemTypeUserInformation              = 0x50
	ItemTypeUserInformationMaximumLength = 0x51
	ItemTypeImplementationClassUID       = 0x52
	ItemTypeAsynchronousOperationsWindow = 0x53
	ItemTypeRoleSelection                = 0x54
	ItemTypeImplementationVersionName    = 0x55
)

func DecodeSubItem(d *dicomio.Decoder) SubItem {
	itemType := d.ReadByte()
	d.Skip(1)
	length := d.ReadUInt16()
	switch itemType {
	case ItemTypeApplicationContext:
		return decodeApplicationContextItem(d, length)
	case ItemTypeAbstractSyntax:
		return decodeAbstractSyntaxSubItem(d, length)
	case ItemTypeTransferSyntax:
		return decodeTransferSyntaxSubItem(d, length)
	case ItemTypePresentationContextRequest:
		return decodePresentationContextItem(d, itemType, length)
	case ItemTypePresentationContextResponse:
		return decodePresentationContextItem(d, itemType, length)
	case ItemTypeUserInformation:
		return decodeUserInformationItem(d, length)
	case ItemTypeUserInformationMaximumLength:
		return decodeUserInformationMaximumLengthItem(d, length)
	case ItemTypeImplementationClassUID:
		return decodeImplementationClassUIDSubItem(d, length)
	case ItemTypeAsynchronousOperationsWindow:
		return decodeAsynchronousOperationsWindowSubItem(d, length)
	case ItemTypeRoleSelection:
		return decodeRoleSelectionSubItem(d, length)
	case ItemTypeImplementationVersionName:
		return decodeImplementationVersionNameSubItem(d, length)
	default:
		d.SetError(fmt.Errorf("Unknown item type: 0x%x", itemType))
		return nil
	}
}

func encodeSubItemHeader(e *dicomio.Encoder, itemType byte, length uint16) {
	e.WriteByte(itemType)
	e.WriteZeros(1)
	e.WriteUInt16(length)
}

func SubItemListString(items []SubItem) string {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i, subitem := range items {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(subitem.String())
	}
	buf.WriteString("]")
	return buf.String()
}

type subItemWithName struct {
	// Type byte
	Name string
}

func encodeSubItemWithName(e *dicomio.Encoder, itemType byte, name string) {
	encodeSubItemHeader(e, itemType, uint16(len(name)))
	// TODO: handle unicode properly
	e.WriteBytes([]byte(name))
}

func DecodeSubItemWithName(d *dicomio.Decoder, length uint16) string {
	return d.ReadString(int(length))
}

type SubItemUnsupported struct {
	Type byte
	Data []byte
}

func (item *SubItemUnsupported) Write(e *dicomio.Encoder) {
	encodeSubItemHeader(e, item.Type, uint16(len(item.Data)))
	// TODO: handle unicode properly
	e.WriteBytes(item.Data)
}

func (item *SubItemUnsupported) String() string {
	return fmt.Sprintf("SubitemUnsupported{type: 0x%0x data: %dbytes}",
		item.Type, len(item.Data))
}
