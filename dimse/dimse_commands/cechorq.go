package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CEchoRq struct {
	MessageID          dimse.MessageID
	CommandDataSetType uint16
	Extra              []*dicom.Element // Unparsed elements
}

func (v *CEchoRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCEchoRq)))
	elems = append(elems, dimse.NewElement(commandset.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CEchoRq) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CEchoRq) CommandField() int {
	return 48
}

func (v *CEchoRq) GetMessageID() dimse.MessageID {
	return v.MessageID
}

func (v *CEchoRq) GetStatus() *dimse.Status {
	return nil
}

func (v *CEchoRq) String() string {
	return fmt.Sprintf("CEchoRq{MessageID:%v CommandDataSetType:%v}}", v.MessageID, v.CommandDataSetType)
}

func (CEchoRq) decode(d *dimse.MessageDecoder) *CEchoRq {
	v := &CEchoRq{}
	v.MessageID = d.GetUInt16(commandset.MessageID, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}
