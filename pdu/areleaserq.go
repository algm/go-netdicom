package pdu

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

type AReleaseRq struct {
}

func (AReleaseRq) Read(d *dicomio.Decoder) PDU {
	pdu := &AReleaseRq{}
	d.Skip(4)
	return pdu
}

func (pdu *AReleaseRq) Write() ([]byte, error) {
	return []byte{0, 0, 0, 0}, nil
}

func (pdu *AReleaseRq) String() string {
	return fmt.Sprintf("A_RELEASE_RQ(%v)", *pdu)
}
