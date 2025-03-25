package pdu_item

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

type AbstractSyntaxSubItem subItemWithName

func decodeAbstractSyntaxSubItem(d *dicomio.Decoder, length uint16) *AbstractSyntaxSubItem {
	return &AbstractSyntaxSubItem{Name: DecodeSubItemWithName(d, length)}
}

func (v *AbstractSyntaxSubItem) Write(e *dicomio.Encoder) {
	encodeSubItemWithName(e, ItemTypeAbstractSyntax, v.Name)
}

func (v *AbstractSyntaxSubItem) String() string {
	return fmt.Sprintf("AbstractSyntax{name: \"%s\"}", v.Name)
}
