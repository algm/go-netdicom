package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CFindRsp struct {
	AffectedSOPClassUID       string
	MessageIDBeingRespondedTo dimse.MessageID
	CommandDataSetType        uint16
	Status                    dimse.Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CFindRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCFindRsp)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewStatusElements(v.Status)...)
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CFindRsp) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CFindRsp) CommandField() int {
	return 32800
}

func (v *CFindRsp) GetMessageID() dimse.MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CFindRsp) GetStatus() *dimse.Status {
	return &v.Status
}

func (v *CFindRsp) String() string {
	return fmt.Sprintf("CFindRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.Status)
}

func (CFindRsp) decode(d *dimse.MessageDecoder) *CFindRsp {
	v := &CFindRsp{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(commandset.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}
