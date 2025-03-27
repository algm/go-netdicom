package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CMoveRsp struct {
	AffectedSOPClassUID            string
	MessageIDBeingRespondedTo      dimse.MessageID
	CommandDataSetType             uint16
	NumberOfRemainingSuboperations uint16
	NumberOfCompletedSuboperations uint16
	NumberOfFailedSuboperations    uint16
	NumberOfWarningSuboperations   uint16
	Status                         dimse.Status
	Extra                          []*dicom.Element // Unparsed elements
}

func (v *CMoveRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCMoveRsp)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	if v.NumberOfRemainingSuboperations != 0 {
		elems = append(elems, dimse.NewElement(commandset.NumberOfRemainingSuboperations, v.NumberOfRemainingSuboperations))
	}
	if v.NumberOfCompletedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(commandset.NumberOfCompletedSuboperations, v.NumberOfCompletedSuboperations))
	}
	if v.NumberOfFailedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(commandset.NumberOfFailedSuboperations, v.NumberOfFailedSuboperations))
	}
	if v.NumberOfWarningSuboperations != 0 {
		elems = append(elems, dimse.NewElement(commandset.NumberOfWarningSuboperations, v.NumberOfWarningSuboperations))
	}
	elems = append(elems, dimse.NewStatusElements(v.Status)...)
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CMoveRsp) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CMoveRsp) CommandField() int {
	return 32801
}

func (v *CMoveRsp) GetMessageID() dimse.MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CMoveRsp) GetStatus() *dimse.Status {
	return &v.Status
}

func (v *CMoveRsp) String() string {
	return fmt.Sprintf("CMoveRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v NumberOfRemainingSuboperations:%v NumberOfCompletedSuboperations:%v NumberOfFailedSuboperations:%v NumberOfWarningSuboperations:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.NumberOfRemainingSuboperations, v.NumberOfCompletedSuboperations, v.NumberOfFailedSuboperations, v.NumberOfWarningSuboperations, v.Status)
}

func (CMoveRsp) decode(d *dimse.MessageDecoder) *CMoveRsp {
	v := &CMoveRsp{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(commandset.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.NumberOfRemainingSuboperations = d.GetUInt16(commandset.NumberOfRemainingSuboperations, dimse.OptionalElement)
	v.NumberOfCompletedSuboperations = d.GetUInt16(commandset.NumberOfCompletedSuboperations, dimse.OptionalElement)
	v.NumberOfFailedSuboperations = d.GetUInt16(commandset.NumberOfFailedSuboperations, dimse.OptionalElement)
	v.NumberOfWarningSuboperations = d.GetUInt16(commandset.NumberOfWarningSuboperations, dimse.OptionalElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}
