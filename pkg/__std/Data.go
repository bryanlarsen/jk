// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package __std

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Data is data. 
type Data struct {
	_tab flatbuffers.Table
}

func GetRootAsData(buf []byte, offset flatbuffers.UOffsetT) *Data {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Data{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Data) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Data) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Data) Bytes(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Data) BytesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DataStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DataAddBytes(builder *flatbuffers.Builder, bytes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(bytes), 0)
}
func DataStartBytesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
