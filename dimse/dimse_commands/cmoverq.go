package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CMoveRq struct {
	AffectedSOPClassUID string
	MessageID           dimse.MessageID
	Priority            uint16
	MoveDestination     string
	CommandDataSetType  uint16
	Extra               []*dicom.Element // Unparsed elements
}

func (v *CMoveRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCMoveRq)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(commandset.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(commandset.MoveDestination, v.MoveDestination))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CMoveRq) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CMoveRq) CommandField() int {
	return 33
}

func (v *CMoveRq) GetMessageID() dimse.MessageID {
	return v.MessageID
}

func (v *CMoveRq) GetStatus() *dimse.Status {
	return nil
}

func (v *CMoveRq) String() string {
	return fmt.Sprintf("CMoveRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v MoveDestination:%v CommandDataSetType:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.MoveDestination, v.CommandDataSetType)
}

func (CMoveRq) decode(d *dimse.MessageDecoder) *CMoveRq {
	v := &CMoveRq{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(commandset.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(commandset.Priority, dimse.RequiredElement)
	v.MoveDestination = d.GetString(commandset.MoveDestination, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}
