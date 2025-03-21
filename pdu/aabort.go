package pdu

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

type AbortReasonType byte

const (
	AbortReasonNotSpecified             AbortReasonType = 0
	AbortReasonUnexpectedPDU            AbortReasonType = 2
	AbortReasonUnrecognizedPDUParameter AbortReasonType = 3
	AbortReasonUnexpectedPDUParameter   AbortReasonType = 4
	AbortReasonInvalidPDUParameterValue AbortReasonType = 5
)

type AAbort struct {
	Source SourceType
	Reason AbortReasonType
}

func (AAbort) Read(d *dicomio.Decoder) PDU {
	pdu := &AAbort{}
	d.Skip(2)
	pdu.Source = SourceType(d.ReadByte())
	pdu.Reason = AbortReasonType(d.ReadByte())
	return pdu
}

func (pdu *AAbort) Write() ([]byte, error) {
	return []byte{
		0,
		0,
		byte(pdu.Source),
		byte(pdu.Reason),
	}, nil
}

func (pdu *AAbort) String() string {
	return fmt.Sprintf("A_ABORT{source:%v reason:%v}", pdu.Source, pdu.Reason)
}
