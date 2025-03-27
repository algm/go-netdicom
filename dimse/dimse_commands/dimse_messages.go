package dimse_commands

// Code generated from generate_dimse_messages.py. DO NOT EDIT.

import (
	"fmt"
	"encoding/binary"
	"github.com/mlibanori/go-netdicom/dimse"
	"github.com/mlibanori/go-netdicom/pdu"
	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomio"
	"github.com/mlibanori/go-netdicom/commandset"
)

const (
CommandFieldCStoreRq	int	= 0x0001
CommandFieldCStoreRsp int	= 0x8001
CommandFieldCFindRq		int	= 0x0020
CommandFieldCFindRsp	int	= 0x8020
CommandFieldCGetRq		int	= 0x0010
CommandFieldCGetRsp		int	= 0x8010
CommandFieldCMoveRq		int	= 0x0021
CommandFieldCMoveRsp 	int	= 0x8021
CommandFieldCEchoRq 	int	= 0x0030
CommandFieldCEchoRsp 	int	= 0x8030
)

func DecodeMessageForType(d *dimse.MessageDecoder, commandField uint16) dimse.Message {
	switch commandField {
	case 0x0001:
		return CStoreRq{}.decode(d)
	case 0x8001:
		return CStoreRsp{}.decode(d)
	case 0x0020:
		return CFindRq{}.decode(d)
	case 0x8020:
		return CFindRsp{}.decode(d)
	case 0x0010:
		return CGetRq{}.decode(d)
	case 0x8010:
		return CGetRsp{}.decode(d)
	case 0x0021:
		return CMoveRq{}.decode(d)
	case 0x8021:
		return CMoveRsp{}.decode(d)
	case 0x0030:
		return CEchoRq{}.decode(d)
	case 0x8030:
		return CEchoRsp{}.decode(d)
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
	commandField := dd.GetUInt16(commandset.MessageID, dimse.RequiredElement)
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