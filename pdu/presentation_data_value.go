package pdu

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// P3.8 9.3.2.2.1 & 9.3.2.2.2
type PresentationDataValueItem struct {
	// Length: 2 + len(Value)
	ContextID byte

	// P3.8, E.2: the following two fields encode a single byte.
	Command bool // Bit 7 (LSB): 1 means command 0 means data
	Last    bool // Bit 6: 1 means last fragment. 0 means not last fragment.

	// Payload, either command or data
	Value []byte
}

func ReadPresentationDataValueItem(d *dicomio.Decoder) PresentationDataValueItem {
	item := PresentationDataValueItem{}
	length := d.ReadUInt32()
	item.ContextID = d.ReadByte()
	header := d.ReadByte()
	item.Command = (header&1 != 0)
	item.Last = (header&2 != 0)
	item.Value = d.ReadBytes(int(length - 2)) // remove contextID and header
	return item
}

func (v *PresentationDataValueItem) Write(e *dicomio.Encoder) {
	var header byte
	if v.Command {
		header |= 1
	}
	if v.Last {
		header |= 2
	}
	e.WriteUInt32(uint32(2 + len(v.Value)))
	e.WriteByte(v.ContextID)
	e.WriteByte(header)
	e.WriteBytes(v.Value)
}

func (v *PresentationDataValueItem) String() string {
	return fmt.Sprintf("PresentationDataValue{context: %d, cmd:%v last:%v value: %d bytes}", v.ContextID, v.Command, v.Last, len(v.Value))
}
