package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CStoreRq struct {
	AffectedSOPClassUID                  string
	MessageID                            dimse.MessageID
	Priority                             uint16
	CommandDataSetType                   uint16
	AffectedSOPInstanceUID               string
	MoveOriginatorApplicationEntityTitle string
	MoveOriginatorMessageID              dimse.MessageID
	Extra                                []*dicom.Element // Unparsed elements
}

func (v *CStoreRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCStoreRq)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(commandset.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID))
	if v.MoveOriginatorApplicationEntityTitle != "" {
		elems = append(elems, dimse.NewElement(commandset.MoveOriginatorApplicationEntityTitle, v.MoveOriginatorApplicationEntityTitle))
	}
	if v.MoveOriginatorMessageID != 0 {
		elems = append(elems, dimse.NewElement(commandset.MoveOriginatorMessageID, v.MoveOriginatorMessageID))
	}
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CStoreRq) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CStoreRq) CommandField() int {
	return 1
}

func (v *CStoreRq) GetMessageID() dimse.MessageID {
	return v.MessageID
}

func (v *CStoreRq) GetStatus() *dimse.Status {
	return nil
}

func (v *CStoreRq) String() string {
	return fmt.Sprintf("CStoreRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v CommandDataSetType:%v AffectedSOPInstanceUID:%v MoveOriginatorApplicationEntityTitle:%v MoveOriginatorMessageID:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.CommandDataSetType, v.AffectedSOPInstanceUID, v.MoveOriginatorApplicationEntityTitle, v.MoveOriginatorMessageID)
}

func (CStoreRq) decode(d *dimse.MessageDecoder) *CStoreRq {
	v := &CStoreRq{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(commandset.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(commandset.Priority, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.AffectedSOPInstanceUID = d.GetString(commandset.AffectedSOPInstanceUID, dimse.RequiredElement)
	v.MoveOriginatorApplicationEntityTitle = d.GetString(commandset.MoveOriginatorApplicationEntityTitle, dimse.OptionalElement)
	v.MoveOriginatorMessageID = d.GetUInt16(commandset.MoveOriginatorMessageID, dimse.OptionalElement)
	v.Extra = d.UnparsedElements()
	return v
}
