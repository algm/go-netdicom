package fuzzpdu

import (
	"bytes"
	"encoding/binary"
	"flag"

	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/dimse/dimse_commands"
	"github.com/mlibanori/go-netdicom/pdu"
)

func init() {
	flag.Parse()
}

func Fuzz(data []byte) int {
	in := bytes.NewBuffer(data)
	if len(data) == 0 || data[0] <= 0xc0 {
		pdu.ReadPDU(in, 4<<20) // nolint: errcheck
	} else {
		d := dicomio.NewDecoder(in, binary.LittleEndian, dicomio.ExplicitVR)
		dimse_commands.ReadMessage(d)
	}
	return 0
}
