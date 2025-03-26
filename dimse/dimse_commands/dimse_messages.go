package dimse_commands

// Code generated from generate_dimse_messages.py. DO NOT EDIT.

import (
	"fmt"
	"encoding/binary"
	"github.com/mlibanori/go-netdicom/dimse"
	"github.com/mlibanori/go-netdicom/pdu"
	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/grailbio/go-dicom/dicomtag"
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
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(1)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(dicomtag.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID))
	if v.MoveOriginatorApplicationEntityTitle != "" {
		elems = append(elems, dimse.NewElement(dicomtag.MoveOriginatorApplicationEntityTitle, v.MoveOriginatorApplicationEntityTitle))
	}
	if v.MoveOriginatorMessageID != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.MoveOriginatorMessageID, v.MoveOriginatorMessageID))
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

func decodeCStoreRq(d *dimse.MessageDecoder) *CStoreRq {
	v := &CStoreRq{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(dicomtag.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(dicomtag.Priority, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.AffectedSOPInstanceUID = d.GetString(dicomtag.AffectedSOPInstanceUID, dimse.RequiredElement)
	v.MoveOriginatorApplicationEntityTitle = d.GetString(dicomtag.MoveOriginatorApplicationEntityTitle, dimse.OptionalElement)
	v.MoveOriginatorMessageID = d.GetUInt16(dicomtag.MoveOriginatorMessageID, dimse.OptionalElement)
	v.Extra = d.UnparsedElements()
	return v
}

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
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32769)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID))
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

func decodeCStoreRsp(d *dimse.MessageDecoder) *CStoreRsp {
	v := &CStoreRsp{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(dicomtag.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.AffectedSOPInstanceUID = d.GetString(dicomtag.AffectedSOPInstanceUID, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}

type CFindRq struct {
	AffectedSOPClassUID string
	MessageID           dimse.MessageID
	Priority            uint16
	CommandDataSetType  uint16
	Extra               []*dicom.Element // Unparsed elements
}

func (v *CFindRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(dicomtag.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CFindRq) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CFindRq) CommandField() int {
	return 32
}

func (v *CFindRq) GetMessageID() dimse.MessageID {
	return v.MessageID
}

func (v *CFindRq) GetStatus() *dimse.Status {
	return nil
}

func (v *CFindRq) String() string {
	return fmt.Sprintf("CFindRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v CommandDataSetType:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.CommandDataSetType)
}

func decodeCFindRq(d *dimse.MessageDecoder) *CFindRq {
	v := &CFindRq{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(dicomtag.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(dicomtag.Priority, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}

type CFindRsp struct {
	AffectedSOPClassUID       string
	MessageIDBeingRespondedTo dimse.MessageID
	CommandDataSetType        uint16
	Status                    dimse.Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CFindRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32800)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
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

func decodeCFindRsp(d *dimse.MessageDecoder) *CFindRsp {
	v := &CFindRsp{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(dicomtag.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}

type CGetRq struct {
	AffectedSOPClassUID string
	MessageID           dimse.MessageID
	Priority            uint16
	CommandDataSetType  uint16
	Extra               []*dicom.Element // Unparsed elements
}

func (v *CGetRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(16)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(dicomtag.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
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

func decodeCGetRq(d *dimse.MessageDecoder) *CGetRq {
	v := &CGetRq{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(dicomtag.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(dicomtag.Priority, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}

type CGetRsp struct {
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

func (v *CGetRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32784)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
	if v.NumberOfRemainingSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfRemainingSuboperations, v.NumberOfRemainingSuboperations))
	}
	if v.NumberOfCompletedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfCompletedSuboperations, v.NumberOfCompletedSuboperations))
	}
	if v.NumberOfFailedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfFailedSuboperations, v.NumberOfFailedSuboperations))
	}
	if v.NumberOfWarningSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfWarningSuboperations, v.NumberOfWarningSuboperations))
	}
	elems = append(elems, dimse.NewStatusElements(v.Status)...)
	elems = append(elems, v.Extra...)
	dimse.EncodeElements(e, elems)
}

func (v *CGetRsp) HasData() bool {
	return v.CommandDataSetType != dimse.CommandDataSetTypeNull
}

func (v *CGetRsp) CommandField() int {
	return 32784
}

func (v *CGetRsp) GetMessageID() dimse.MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CGetRsp) GetStatus() *dimse.Status {
	return &v.Status
}

func (v *CGetRsp) String() string {
	return fmt.Sprintf("CGetRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v NumberOfRemainingSuboperations:%v NumberOfCompletedSuboperations:%v NumberOfFailedSuboperations:%v NumberOfWarningSuboperations:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.NumberOfRemainingSuboperations, v.NumberOfCompletedSuboperations, v.NumberOfFailedSuboperations, v.NumberOfWarningSuboperations, v.Status)
}

func decodeCGetRsp(d *dimse.MessageDecoder) *CGetRsp {
	v := &CGetRsp{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(dicomtag.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.NumberOfRemainingSuboperations = d.GetUInt16(dicomtag.NumberOfRemainingSuboperations, dimse.OptionalElement)
	v.NumberOfCompletedSuboperations = d.GetUInt16(dicomtag.NumberOfCompletedSuboperations, dimse.OptionalElement)
	v.NumberOfFailedSuboperations = d.GetUInt16(dicomtag.NumberOfFailedSuboperations, dimse.OptionalElement)
	v.NumberOfWarningSuboperations = d.GetUInt16(dicomtag.NumberOfWarningSuboperations, dimse.OptionalElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}

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
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(33)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(dicomtag.Priority, v.Priority))
	elems = append(elems, dimse.NewElement(dicomtag.MoveDestination, v.MoveDestination))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
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

func decodeCMoveRq(d *dimse.MessageDecoder) *CMoveRq {
	v := &CMoveRq{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageID = d.GetUInt16(dicomtag.MessageID, dimse.RequiredElement)
	v.Priority = d.GetUInt16(dicomtag.Priority, dimse.RequiredElement)
	v.MoveDestination = d.GetString(dicomtag.MoveDestination, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}

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
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32801)))
	elems = append(elems, dimse.NewElement(dicomtag.AffectedSOPClassUID, v.AffectedSOPClassUID))
	elems = append(elems, dimse.NewElement(dicomtag.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
	if v.NumberOfRemainingSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfRemainingSuboperations, v.NumberOfRemainingSuboperations))
	}
	if v.NumberOfCompletedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfCompletedSuboperations, v.NumberOfCompletedSuboperations))
	}
	if v.NumberOfFailedSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfFailedSuboperations, v.NumberOfFailedSuboperations))
	}
	if v.NumberOfWarningSuboperations != 0 {
		elems = append(elems, dimse.NewElement(dicomtag.NumberOfWarningSuboperations, v.NumberOfWarningSuboperations))
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

func decodeCMoveRsp(d *dimse.MessageDecoder) *CMoveRsp {
	v := &CMoveRsp{}
	v.AffectedSOPClassUID = d.GetString(dicomtag.AffectedSOPClassUID, dimse.RequiredElement)
	v.MessageIDBeingRespondedTo = d.GetUInt16(dicomtag.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.NumberOfRemainingSuboperations = d.GetUInt16(dicomtag.NumberOfRemainingSuboperations, dimse.OptionalElement)
	v.NumberOfCompletedSuboperations = d.GetUInt16(dicomtag.NumberOfCompletedSuboperations, dimse.OptionalElement)
	v.NumberOfFailedSuboperations = d.GetUInt16(dicomtag.NumberOfFailedSuboperations, dimse.OptionalElement)
	v.NumberOfWarningSuboperations = d.GetUInt16(dicomtag.NumberOfWarningSuboperations, dimse.OptionalElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}

type CEchoRq struct {
	MessageID          dimse.MessageID
	CommandDataSetType uint16
	Extra              []*dicom.Element // Unparsed elements
}

func (v *CEchoRq) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(48)))
	elems = append(elems, dimse.NewElement(dicomtag.MessageID, v.MessageID))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
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

func decodeCEchoRq(d *dimse.MessageDecoder) *CEchoRq {
	v := &CEchoRq{}
	v.MessageID = d.GetUInt16(dicomtag.MessageID, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Extra = d.UnparsedElements()
	return v
}

type CEchoRsp struct {
	MessageIDBeingRespondedTo dimse.MessageID
	CommandDataSetType        uint16
	Status                    dimse.Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CEchoRsp) Encode(e *dicomio.Encoder) {
	elems := []*dicom.Element{}
	elems = append(elems, dimse.NewElement(dicomtag.CommandField, uint16(32816)))
	elems = append(elems, dimse.NewElement(dicomtag.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo))
	elems = append(elems, dimse.NewElement(dicomtag.CommandDataSetType, v.CommandDataSetType))
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

func decodeCEchoRsp(d *dimse.MessageDecoder) *CEchoRsp {
	v := &CEchoRsp{}
	v.MessageIDBeingRespondedTo = d.GetUInt16(dicomtag.MessageIDBeingRespondedTo, dimse.RequiredElement)
	v.CommandDataSetType = d.GetUInt16(dicomtag.CommandDataSetType, dimse.RequiredElement)
	v.Status = d.GetStatus()
	v.Extra = d.UnparsedElements()
	return v
}

const CommandFieldCStoreRq = 1
const CommandFieldCStoreRsp = 32769
const CommandFieldCFindRq = 32
const CommandFieldCFindRsp = 32800
const CommandFieldCGetRq = 16
const CommandFieldCGetRsp = 32784
const CommandFieldCMoveRq = 33
const CommandFieldCMoveRsp = 32801
const CommandFieldCEchoRq = 48
const CommandFieldCEchoRsp = 32816

func DecodeMessageForType(d *dimse.MessageDecoder, commandField uint16) dimse.Message {
	switch commandField {
	case 0x1:
		return decodeCStoreRq(d)
	case 0x8001:
		return decodeCStoreRsp(d)
	case 0x20:
		return decodeCFindRq(d)
	case 0x8020:
		return decodeCFindRsp(d)
	case 0x10:
		return decodeCGetRq(d)
	case 0x8010:
		return decodeCGetRsp(d)
	case 0x21:
		return decodeCMoveRq(d)
	case 0x8021:
		return decodeCMoveRsp(d)
	case 0x30:
		return decodeCEchoRq(d)
	case 0x8030:
		return decodeCEchoRsp(d)
	default:
		d.SetError(fmt.Errorf("Unknown DIMSE command 0x%x", commandField))
		return nil
	}
}

// CommandAssembler is a helper that assembles a DIMSE command message and data
// payload from a sequence of P_DATA_TF PDUs.
type CommandAssembler struct {
	contextID      byte
	commandBytes   []byte
	command        dimse.Message
	dataBytes      []byte
	readAllCommand bool

	readAllData bool
}

func (a *CommandAssembler) AddDataPDU(pdu *pdu.PDataTf) (byte, dimse.Message, []byte, error) {
	for _, item := range pdu.Items {
		if a.contextID == 0 {
			a.contextID = item.ContextID
		} else if a.contextID != item.ContextID {
			return 0, nil, nil, fmt.Errorf("Mixed context: %d %d", a.contextID, item.ContextID)
		}
		if item.Command {
			a.commandBytes = append(a.commandBytes, item.Value...)
			if item.Last {
				if a.readAllCommand {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 command chunks with the Last bit set")
				}
				a.readAllCommand = true
			}
		} else {
			a.dataBytes = append(a.dataBytes, item.Value...)
			if item.Last {
				if a.readAllData {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 data chunks with the Last bit set")
				}
				a.readAllData = true
			}
		}
	}
	if !a.readAllCommand {
		return 0, nil, nil, nil
	}
	if a.command == nil {
		d := dicomio.NewBytesDecoder(a.commandBytes, nil, dicomio.UnknownVR)
		a.command = ReadMessage(d)
		if err := d.Finish(); err != nil {
			return 0, nil, nil, err
		}
	}
	if a.command.HasData() && !a.readAllData {
		return 0, nil, nil, nil
	}
	contextID := a.contextID
	command := a.command
	dataBytes := a.dataBytes
	*a = CommandAssembler{}
	return contextID, command, dataBytes, nil
	// TODO(saito) Verify that there's no unread items after the last command&data.
}

// ReadMessage constructs a typed dimse.Message object, given a set of
// dicom.Elements,
func ReadMessage(d *dicomio.Decoder) dimse.Message {
	// A DIMSE message is a sequence of Elements, encoded in implicit
	// LE.
	//
	// TODO(saito) make sure that's the case. Where the ref?
	var elems []*dicom.Element
	d.PushTransferSyntax(binary.LittleEndian, dicomio.ImplicitVR)
	defer d.PopTransferSyntax()
	for !d.EOF() {
		elem := dicom.ReadElement(d, dicom.ReadOptions{})
		if d.Error() != nil {
			break
		}
		elems = append(elems, elem)
	}

	// Convert elems[] into a golang struct.
	dd := dimse.NewMessageDecoder(elems)
	commandField := dd.GetUInt16(dicomtag.CommandField, dimse.RequiredElement)
	if dd.Error() != nil {
		d.SetError(dd.Error())
		return nil
	}
	v := DecodeMessageForType(dd, commandField)
	if dd.Error() != nil {
		d.SetError(dd.Error())
		return nil
	}
	return v
}