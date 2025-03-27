package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CEchoRsp struct {
	MessageIDBeingRespondedTo dimse.MessageID
	CommandDataSetType        uint16
	Status                    dimse.Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CEchoRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCEchoRsp)))
	elems = append(elems, dimse.NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewStatusElements(v.Status)...)
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CEchoRsp) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CEchoRsp) CommandField() int {
	return 32816
}

func (v *CEchoRsp) GetMessageID() dimse.MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CEchoRsp) GetStatus() *dimse.Status {
	return &v.Status
}

func (v *CEchoRsp) String() string {
	return fmt.Sprintf("CEchoRsp{MessageIDBeingRespondedTo:%v CommandDataSetType:%v Status:%v}}", v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.Status)
}

func (CEchoRsp) decode(d *dimse.MessageDecoder) *CEchoRsp {
	v := &CEchoRsp{}
	v.MessageIDBeingRespondedTo = d.GetUInt16(commandset.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}
