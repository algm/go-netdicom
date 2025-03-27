package dimse_commands

import (
	"fmt"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/dimse"
)

type CStoreRsp struct {
	AffectedSOPClassUID       string
	MessageIDBeingRespondedTo dimse.MessageID
	CommandDataSetType        uint16
	AffectedSOPInstanceUID    string
	Status                    dimse.Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CStoreRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(commandset.CommandField, uint16(CommandFieldCStoreRsp)))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(commandset.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewElement(commandset.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID))
	elems = append(elems, dimse.NewStatusElements(v.Status)...)
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CStoreRsp) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CStoreRsp) CommandField() int {
	return 32769
}

func (v *CStoreRsp) GetMessageID() dimse.MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CStoreRsp) GetStatus() *dimse.Status {
	return &v.Status
}

func (v *CStoreRsp) String() string {
	return fmt.Sprintf("CStoreRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v AffectedSOPInstanceUID:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.AffectedSOPInstanceUID, v.Status)
}

func (CStoreRsp) decode(d *dimse.MessageDecoder) *CStoreRsp {
	v := &CStoreRsp{}
	v.AffectedSOPClassUID = d.GetString(commandset.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(commandset.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(commandset.CommandDataSetType, dimse.RequiredElement)
	v.AffectedSOPInstanceUID = d.GetString(commandset.AffectedSOPInstanceUID, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}
