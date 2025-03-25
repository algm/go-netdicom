package pdu_item

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

type TransferSyntaxSubItem subItemWithName

func decodeTransferSyntaxSubItem(d *dicomio.Decoder, length uint16) *TransferSyntaxSubItem {
	return &TransferSyntaxSubItem{Name: DecodeSubItemWithName(d, length)}
}

func (v *TransferSyntaxSubItem) Write(e *dicomio.Encoder) {
	encodeSubItemWithName(e, ItemTypeTransferSyntax, v.Name)
}

func (v *TransferSyntaxSubItem) String() string {
	return fmt.Sprintf("TransferSyntax{name: \"%s\"}", v.Name)
}
