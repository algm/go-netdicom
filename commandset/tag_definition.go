package commandset

import "github.com/grailbio/go-dicom/dicomtag"

var (
	CommandGroupLength                   = dicomtag.Tag{Group: 0x0000, Element: 0x0000}
	AffectedSOPClassUID                  = dicomtag.Tag{Group: 0x0000, Element: 0x0002}
	RequestedSOPClassUID                 = dicomtag.Tag{Group: 0x0000, Element: 0x0003}
	CommandField                         = dicomtag.Tag{Group: 0x0000, Element: 0x0100}
	MessageID                            = dicomtag.Tag{Group: 0x0000, Element: 0x0110}
	MessageIDBeingRespondedTo            = dicomtag.Tag{Group: 0x0000, Element: 0x0120}
	MoveDestination                      = dicomtag.Tag{Group: 0x0000, Element: 0x0600}
	Priority                             = dicomtag.Tag{Group: 0x0000, Element: 0x0700}
	CommandDataSetType                   = dicomtag.Tag{Group: 0x0000, Element: 0x0800}
	Status                               = dicomtag.Tag{Group: 0x0000, Element: 0x0900}
	OffendingElement                     = dicomtag.Tag{Group: 0x0000, Element: 0x0901}
	ErrorComment                         = dicomtag.Tag{Group: 0x0000, Element: 0x0902}
	ErrorID                              = dicomtag.Tag{Group: 0x0000, Element: 0x0903}
	AffectedSOPInstanceUID               = dicomtag.Tag{Group: 0x0000, Element: 0x1000}
	RequestedSOPInstanceUID              = dicomtag.Tag{Group: 0x0000, Element: 0x1001}
	EventTypeID                          = dicomtag.Tag{Group: 0x0000, Element: 0x1002}
	AttributeIdentifierList              = dicomtag.Tag{Group: 0x0000, Element: 0x1005}
	ActionTypeID                         = dicomtag.Tag{Group: 0x0000, Element: 0x1008}
	NumberOfRemainingSuboperations       = dicomtag.Tag{Group: 0x0000, Element: 0x1020}
	NumberOfCompletedSuboperations       = dicomtag.Tag{Group: 0x0000, Element: 0x1021}
	NumberOfFailedSuboperations          = dicomtag.Tag{Group: 0x0000, Element: 0x1022}
	NumberOfWarningSuboperations         = dicomtag.Tag{Group: 0x0000, Element: 0x1023}
	MoveOriginatorApplicationEntityTitle = dicomtag.Tag{Group: 0x0000, Element: 0x1030}
	MoveOriginatorMessageID              = dicomtag.Tag{Group: 0x0000, Element: 0x1031}
)
