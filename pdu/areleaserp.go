package pdu

import (
	"fmt"

	"github.com/grailbio/go-dicom/dicomio"
)

type AReleaseRp struct {
}

func (AReleaseRp) Read(d *dicomio.Decoder) PDU {
	pdu := &AReleaseRp{}
	d.Skip(4)
	return pdu
}

func (pdu *AReleaseRp) Write() ([]byte, error) {
	return []byte{0, 0, 0, 0}, nil
}

func (pdu *AReleaseRp) String() string {
	return fmt.Sprintf("A_RELEASE_RP(%v)", *pdu)
}
