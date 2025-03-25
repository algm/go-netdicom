package pdu

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// P3.8 D.1
type UserInformationMaximumLengthItem struct {
	MaximumLengthReceived uint32
}

func (v *UserInformationMaximumLengthItem) Write(e *dicomio.Encoder) {
	encodeSubItemHeader(e, ItemTypeUserInformationMaximumLength, 4)
	e.WriteUInt32(v.MaximumLengthReceived)
}

func decodeUserInformationMaximumLengthItem(d *dicomio.Decoder, length uint16) *UserInformationMaximumLengthItem {
	if length != 4 {
		d.SetError(fmt.Errorf("UserInformationMaximumLengthItem must be 4 bytes, but found %dB", length))
	}
	return &UserInformationMaximumLengthItem{MaximumLengthReceived: d.ReadUInt32()}
}

func (v *UserInformationMaximumLengthItem) String() string {
	return fmt.Sprintf("UserInformationMaximumlengthItem{%d}",
		v.MaximumLengthReceived)
}
