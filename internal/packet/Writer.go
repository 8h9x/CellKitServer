package packet

import (
	"encoding/binary"
	"image/color"
	"math"
)

type Writer struct {
	buf    []byte
	offset int
}

func NewWriter(buf []byte) *Writer {
	return &Writer{
		buf:    buf,
		offset: 0,
	}
}

func (w *Writer) WriteUint8(n uint8) {
	w.buf[w.offset] = n
	w.offset++
}

func (w *Writer) WriteInt8(n int8) {
	w.WriteUint8(uint8(n))
}

func (w *Writer) WriteUint16(n uint16) {
	binary.LittleEndian.PutUint16(w.buf[w.offset:], n)
	w.offset += 2
}

func (w *Writer) WriteInt16(n int16) {
	w.WriteUint16(uint16(n))
}

func (w *Writer) WriteUint32(n uint32) {
	binary.LittleEndian.PutUint32(w.buf[w.offset:], n)
	w.offset += 4
}

func (w *Writer) WriteInt32(n int32) {
	w.WriteUint32(uint32(n))
}

func (w *Writer) WriteUint64(n uint64) {
	binary.LittleEndian.PutUint64(w.buf[w.offset:], n)
	w.offset += 8
}

func (w *Writer) WriteInt64(n int64) {
	w.WriteUint64(uint64(n))
}

func (w *Writer) WriteFloat32(n float32) {
	binary.LittleEndian.PutUint32(w.buf[w.offset:], math.Float32bits(n))
	w.offset += 4
}

func (w *Writer) WriteFloat64(n float64) {
	binary.LittleEndian.PutUint64(w.buf[w.offset:], math.Float64bits(n))
	w.offset += 8
}

func (w *Writer) WriteStringUTF8(s string) {
	bytes := []byte(s)
	copy(w.buf[w.offset:], bytes)
	w.offset += len(bytes)
	w.buf[w.offset] = 0
	w.offset++
}

func (w *Writer) WriteColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	w.WriteUint8(uint8(r))
	w.WriteUint8(uint8(g))
	w.WriteUint8(uint8(b))
}

func (w *Writer) Buffer() []byte {
	result := make([]byte, w.offset)
	copy(result, w.buf[:w.offset])
	return result
}

func (w *Writer) Bytes() []byte {
	return w.buf[:w.offset]
}
