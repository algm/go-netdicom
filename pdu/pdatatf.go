package pdu

import (
	"bytes"
	"encoding/binary"

	"github.com/grailbio/go-dicom/dicomio"
)

type PDataTf struct {
	Items []PresentationDataValueItem
}

func (PDataTf) Read(d *dicomio.Decoder) PDU {
	pdu := &PDataTf{}
	for !d.EOF() {
		item := ReadPresentationDataValueItem(d)
		if d.Error() != nil {
			break
		}
		pdu.Items = append(pdu.Items, item)
	}
	return pdu
}

func (pdu *PDataTf) Write() ([]byte, error) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	for _, item := range pdu.Items {
		item.Write(e)
	}
	return e.Bytes(), e.Error()
}

func (pdu *PDataTf) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("P_DATA_TF{items: [")
	for i, item := range pdu.Items {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(item.String())
	}
	buf.WriteString("]}")
	return buf.String()
}
