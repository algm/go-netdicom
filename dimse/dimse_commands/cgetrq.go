package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CGetRq struct {
	AffectedSOPClassUID string
	MessageID           dimse.MessageID
	Priority            uint16
	CommandDataSetType  uint16
	Extra               []*dicom.Element // Unparsed elements
}

func (v *CGetRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCGetRq)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(commandset.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CGetRq) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CGetRq) CommandField() int {
	return 16
}

func (v *CGetRq) GetMessageID() dimse.MessageID {
	return v.MessageID
}

func (v *CGetRq) GetStatus() *dimse.Status {
	return nil
}

func (v *CGetRq) String() string {
	return fmt.Sprintf("CGetRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v CommandDataSetType:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.CommandDataSetType)
}

func (CGetRq) decode(d *dimse.MessageDecoder) *CGetRq {
	v := &CGetRq{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(commandset.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(commandset.Priority, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}
