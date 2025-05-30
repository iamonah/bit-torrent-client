package message

import (
	"encoding/binary"
	"io"
)

type messageID uint8

const (
	MsgChoke         messageID = 0
	MsgUnchoke       messageID = 1
	MsgInterested    messageID = 2
	MsgNotInterested messageID = 3
	MsgHave          messageID = 4 //telling other connected peers that i now have a piece
	MsgBitfield      messageID = 5
	MsgRequest       messageID = 6
	MsgPiece         messageID = 7
	MsgCancel        messageID = 8
)

type Message struct {
	ID      messageID
	Payload []byte
}

//<length prefix><message ID = 5><bitfield>
//interprets 'nil' as a keep-alive message

func (m *Message) Serialize() []byte {
	if m == nil {
		return make([]byte, 4)
	}
	//length of payload and message id in bytes
	length := uint32(len(m.Payload) + 1)
	buf := make([]byte, 4+length)
	binary.BigEndian.PutUint32(buf[0:4], length)
	buf[4] = byte(m.ID)
	copy(buf[5:], m.Payload)
	return buf
}

func Read(r io.Reader) (*Message, error) {
	bufLength := make([]byte, 4)
	_, err := io.ReadFull(r, bufLength)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(bufLength)
	//keep-alive message
	if length == 0 {
		return nil, nil
	}

	msgBuff := make([]byte, length+1)
	_, err = io.ReadFull(r, msgBuff)
	if err != nil {
		return nil, err
	}
	
	msg := Message{
		ID:      messageID(msgBuff[0]),
		Payload: msgBuff[1:],
	}
	return &msg, nil
}
