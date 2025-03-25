package pdu_item

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

// PS3.7 Annex D.3.3.4
type RoleSelectionSubItem struct {
	SOPClassUID string
	SCURole     byte
	SCPRole     byte
}

func decodeRoleSelectionSubItem(d *dicomio.Decoder, length uint16) *RoleSelectionSubItem {
	uidLen := d.ReadUInt16()
	return &RoleSelectionSubItem{
		SOPClassUID: d.ReadString(int(uidLen)),
		SCURole:     d.ReadByte(),
		SCPRole:     d.ReadByte(),
	}
}

func (v *RoleSelectionSubItem) Write(e *dicomio.Encoder) {
	encodeSubItemHeader(e, ItemTypeRoleSelection, uint16(2+len(v.SOPClassUID)+1*2))
	e.WriteUInt16(uint16(len(v.SOPClassUID)))
	e.WriteString(v.SOPClassUID)
	e.WriteByte(v.SCURole)
	e.WriteByte(v.SCPRole)
}

func (v *RoleSelectionSubItem) String() string {
	return fmt.Sprintf("RoleSelection{sopclassuid: %v, scu: %v, scp: %v}", v.SOPClassUID, v.SCURole, v.SCPRole)
}
