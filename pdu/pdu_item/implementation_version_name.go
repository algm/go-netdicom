package pdu_item

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// PS3.7 Annex D.3.3.2.3
type ImplementationVersionNameSubItem subItemWithName

func decodeImplementationVersionNameSubItem(d *dicomio.Decoder, length uint16) *ImplementationVersionNameSubItem {
	return &ImplementationVersionNameSubItem{Name: DecodeSubItemWithName(d, length)}
}

func (v *ImplementationVersionNameSubItem) Write(e *dicomio.Encoder) {
	encodeSubItemWithName(e, ItemTypeImplementationVersionName, v.Name)
}

func (v *ImplementationVersionNameSubItem) String() string {
	return fmt.Sprintf("ImplementationVersionName{name: \"%s\"}", v.Name)
}
