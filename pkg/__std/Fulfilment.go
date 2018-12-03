// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package __std

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Resolution represents the _asynchronous_ fulfilment of a deferred value (possibly one of many values).
type Fulfilment struct {
	_tab flatbuffers.Table
}

func GetRootAsFulfilment(buf []byte, offset flatbuffers.UOffsetT) *Fulfilment {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Fulfilment{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Fulfilment) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Fulfilment) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Fulfilment) Serial() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Fulfilment) MutateSerial(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Fulfilment) ValueType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Fulfilment) MutateValueType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Fulfilment) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func FulfilmentStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FulfilmentAddSerial(builder *flatbuffers.Builder, serial uint64) {
	builder.PrependUint64Slot(0, serial, 0)
}
func FulfilmentAddValueType(builder *flatbuffers.Builder, valueType byte) {
	builder.PrependByteSlot(1, valueType, 0)
}
func FulfilmentAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(value), 0)
}
func FulfilmentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
