package packet

import (
	"bytes"
	"encoding/binary"
)

type Reader struct {
	data   []byte
	offset int
}

func NewReader(data []byte) *Reader {
	return &Reader{
		data:   data,
		offset: 0,
	}
}

func (r *Reader) ReadUint8() uint8 {
	if r.offset >= len(r.data) {
		return 0
	}
	val := r.data[r.offset]
	r.offset++
	return val
}

func (r *Reader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

func (r *Reader) ReadUint16() uint16 {
	if r.offset+2 > len(r.data) {
		r.offset = len(r.data)
		return 0
	}
	val := binary.LittleEndian.Uint16(r.data[r.offset : r.offset+2])
	r.offset += 2
	return val
}

func (r *Reader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

func (r *Reader) ReadUint32() uint32 {
	if r.offset+4 > len(r.data) {
		r.offset = len(r.data)
		return 0
	}
	val := binary.LittleEndian.Uint32(r.data[r.offset : r.offset+4])
	r.offset += 4
	return val
}

func (r *Reader) ReadInt32() int32 {
	return int32(r.ReadUint32())
}

func (r *Reader) ReadUint64() uint64 {
	if r.offset+8 > len(r.data) {
		r.offset = len(r.data)
		return 0
	}
	val := binary.LittleEndian.Uint64(r.data[r.offset : r.offset+8])
	r.offset += 8
	return val
}

func (r *Reader) ReadInt64() int64 {
	return int64(r.ReadUint64())
}

func (r *Reader) ReadString() string {
	end := bytes.IndexByte(r.data[r.offset:], 0)

	if end == -1 {
		r.offset += len(r.data)
		return string(r.data[r.offset:])
	}
	str := string(r.data[r.offset : r.offset+end])

	r.offset += end + 1
	return str
}
