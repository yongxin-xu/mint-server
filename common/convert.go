package mintcommon

import (
	"bytes"
	"encoding/binary"
)

func BytesToUint16(b []byte) uint16 {
	buf := bytes.NewBuffer(b)

	var x uint16
	if err := binary.Read(buf, binary.BigEndian, &x); err != nil {
		panic(err)
	}
	return x
}

func Uint16ToBytes(n uint16) []byte {
	x := uint16(n)

	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, x); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func BytesConcatenate(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}