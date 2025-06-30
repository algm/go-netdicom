package dimse

import (
	"fmt"
	"io"

	"github.com/algm/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CFindRsp struct {
	AffectedSOPClassUID       string
	MessageIDBeingRespondedTo MessageID
	CommandDataSetType        CommandDataSetType
	Status                    Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CFindRsp) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID)
	if err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to create AffectedSOPClassUID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo)
	if err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to create MessageIDBeingRespondedTo element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, uint16(v.CommandDataSetType))
	if err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	statusElems, err := v.Status.ToElements()
	if err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to create status elements: %w", err)
	}
	elems = append(elems, statusElems...)

	elems = append(elems, v.Extra...)

	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CFindRsp.Encode: failed to encode elements: %w", err)
	}
	return nil
}

func (v *CFindRsp) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CFindRsp) CommandField() uint16 {
	return CommandFieldCFindRsp
}

func (v *CFindRsp) GetMessageID() MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CFindRsp) GetStatus() *Status {
	return &v.Status
}

func (v *CFindRsp) String() string {
	return fmt.Sprintf("CFindRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.Status)
}

func (CFindRsp) decode(d *MessageDecoder) (*CFindRsp, error) {
	v := &CFindRsp{}
	var err error

	v.AffectedSOPClassUID, err = d.GetString(commandset.AffectedSOPClassUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cFindRsp.decode: failed to decode AffectedSOPClassUID: %w", err)
	}

	v.MessageIDBeingRespondedTo, err = d.GetUInt16(commandset.MessageIDBeingRespondedTo, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cFindRsp.decode: failed to decode MessageIDBeingRespondedTo: %w", err)
	}

	v.CommandDataSetType, err = d.GetCommandDataSetType()
	if err != nil {
		return nil, fmt.Errorf("cFindRsp.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.Status, err = d.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("cFindRsp.decode: failed to decode Status: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
