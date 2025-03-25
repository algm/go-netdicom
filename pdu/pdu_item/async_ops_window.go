package pdu_item

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// PS3.7 Annex D.3.3.3.1
type AsynchronousOperationsWindowSubItem struct {
	MaxOpsInvoked   uint16
	MaxOpsPerformed uint16
}

func decodeAsynchronousOperationsWindowSubItem(d *dicomio.Decoder, length uint16) *AsynchronousOperationsWindowSubItem {
	return &AsynchronousOperationsWindowSubItem{
		MaxOpsInvoked:   d.ReadUInt16(),
		MaxOpsPerformed: d.ReadUInt16(),
	}
}

func (v *AsynchronousOperationsWindowSubItem) Write(e *dicomio.Encoder) {
	encodeSubItemHeader(e, ItemTypeAsynchronousOperationsWindow, 2*2)
	e.WriteUInt16(v.MaxOpsInvoked)
	e.WriteUInt16(v.MaxOpsPerformed)
}

func (v *AsynchronousOperationsWindowSubItem) String() string {
	return fmt.Sprintf("AsynchronousOpsWindow{invoked: %d performed: %d}",
		v.MaxOpsInvoked, v.MaxOpsPerformed)
}
