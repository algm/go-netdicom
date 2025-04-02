package dimse

import (
	"bytes"
	"fmt"

	"github.com/mlibanori/go-netdicom/pdu"
	"github.com/suyashkumar/dicom"
)

// CommandAssembler is a helper that assembles a DIMSE command message and data
// payload from a sequence of P_DATA_TF PDUs.
type CommandAssembler struct {
	contextID      byte
	commandBytes   []byte
	command        Message
	dataBytes      []byte
	readAllCommand bool

	readAllData bool
}

// AddDataPDU is to be called for each P_DATA_TF PDU received from the
// network. If the fragment is marked as the last one, AddDataPDU returns
// <SOPUID, TransferSyntaxUID, payload, nil>.  If it needs more fragments, it
// returns <"", "", nil, nil>.  On error, it returns a non-nil error.
func (commandAssembler *CommandAssembler) AddDataPDU(pdu *pdu.PDataTf) (byte, Message, []byte, error) {
	for _, item := range pdu.Items {
		if commandAssembler.contextID == 0 {
			commandAssembler.contextID = item.ContextID
		} else if commandAssembler.contextID != item.ContextID {
			return 0, nil, nil, fmt.Errorf("mixed context: %d %d", commandAssembler.contextID, item.ContextID)
		}
		if item.Command {
			commandAssembler.commandBytes = append(commandAssembler.commandBytes, item.Value...)
			if item.Last {
				if commandAssembler.readAllCommand {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 command chunks with the Last bit set")
				}
				commandAssembler.readAllCommand = true
			}
		} else {
			commandAssembler.dataBytes = append(commandAssembler.dataBytes, item.Value...)
			if item.Last {
				if commandAssembler.readAllData {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 data chunks with the Last bit set")
				}
				commandAssembler.readAllData = true
			}
		}
	}
	if !commandAssembler.readAllCommand {
		return 0, nil, nil, nil
	}
	if commandAssembler.command == nil {
		ioReader := bytes.NewReader(commandAssembler.commandBytes)
		parser, err := dicom.Parse(ioReader, int64(ioReader.Len()), nil, dicom.SkipPixelData(), dicom.SkipMetadataReadOnNewParserInit())
		if err != nil {
			return 0, nil, nil, fmt.Errorf("P_DATA_TF: failed to parse command bytes: %w", err)
		}
		commandAssembler.command, err = ReadMessage(&parser)
		if err != nil {
			return 0, nil, nil, err
		}
	}
	if commandAssembler.command.HasData() && !commandAssembler.readAllData {
		return 0, nil, nil, nil
	}
	contextID := commandAssembler.contextID
	command := commandAssembler.command
	dataBytes := commandAssembler.dataBytes
	*commandAssembler = CommandAssembler{}
	return contextID, command, dataBytes, nil
	// TODO(saito) Verify that there's no unread items after the last command&data.
}
