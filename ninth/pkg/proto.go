package pkg

import (
	"encoding/binary"
	"errors"
)

/*
	Package Length	4 bytes
	Header Length	2 bytes
	Protocol Version	2 bytes
	Operation	4 bytes
	SequenceId	4 bytes
	Body
*/

const (
	HeadSize               = 16
	packLengthFieldSize    = 4
	headerLengthFieldSize  = 2
	protocVersionFieldSize = 2
	operationCodeFieldSize = 4
	sequenceIdFieldSize    = 4
)

var (
	ErrPackageLength = errors.New("error package")
)

type Pack struct {
	PackageLength   uint32
	HeadLength      uint16
	ProtocolVersion uint16
	Operation       uint32
	SequenceId      uint32
	Body            []byte
}

func NewPack(version uint16, code, seq uint32, content []byte) Pack {
	return Pack{
		PackageLength:   uint32(len(content) + HeadSize),
		HeadLength:      HeadSize,
		ProtocolVersion: version,
		Operation:       code,
		SequenceId:      seq,
		Body:            content,
	}
}

// Encode 将消息编码
func Encode(pack Pack) []byte {
	res := make([]byte, pack.PackageLength)

	// 写入消息头
	binary.BigEndian.PutUint32(
		res[:headerLengthStart()],
		pack.PackageLength,
	)
	// set header length
	binary.BigEndian.PutUint16(
		res[headerLengthStart():protocolVersionStart()],
		pack.HeadLength,
	)
	// set protocol version
	binary.BigEndian.PutUint16(
		res[protocolVersionStart():operationCodeStart()],
		pack.ProtocolVersion,
	)
	// set operation code
	binary.BigEndian.PutUint32(
		res[operationCodeStart():sequenceIdStart()],
		pack.Operation,
	)
	// set sequence id
	binary.BigEndian.PutUint32(
		res[sequenceIdStart():sequenceIdStart()+sequenceIdFieldSize],
		pack.SequenceId,
	)
	// set body
	copy(res[HeadSize:], pack.Body)

	return res
}

// Decode 将消息解码
func Decode(msg []byte) (Pack, error) {

	if len(msg) < HeadSize+1 {
		return Pack{}, ErrPackageLength
	}
	// get package length
	packageLength := binary.BigEndian.Uint32(msg[:headerLengthStart()])
	// get header length
	headerLength := binary.BigEndian.Uint16(msg[headerLengthStart():protocolVersionStart()])
	// get protocol version
	protocolVersion := binary.BigEndian.Uint16(msg[protocolVersionStart():operationCodeStart()])
	// get operation code
	operationCode := binary.BigEndian.Uint32(msg[operationCodeStart():sequenceIdStart()])
	// get sequence id
	sequenceId := binary.BigEndian.Uint32(msg[sequenceIdStart() : sequenceIdStart()+sequenceIdFieldSize])
	// get data
	body := msg[HeadSize:]
	return Pack{
		PackageLength:   packageLength,
		HeadLength:      headerLength,
		ProtocolVersion: protocolVersion,
		Operation:       operationCode,
		SequenceId:      sequenceId,
		Body:            body,
	}, nil
}

// position of header length field start
func headerLengthStart() int {
	return packLengthFieldSize
}

// position of protocol version field start
func protocolVersionStart() int {
	return headerLengthStart() + headerLengthFieldSize
}

// postion of operation code field start
func operationCodeStart() int {
	return protocolVersionStart() + protocVersionFieldSize
}

// position of sequence id field start
func sequenceIdStart() int {
	return operationCodeStart() + operationCodeFieldSize
}

func PackageLengthSize() int {
	return packLengthFieldSize
}
