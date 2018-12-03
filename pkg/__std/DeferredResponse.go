// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package __std

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Response encodes the _synchronous_ response to a message expecting a deferred value or values.
type DeferredResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsDeferredResponse(buf []byte, offset flatbuffers.UOffsetT) *DeferredResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DeferredResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DeferredResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DeferredResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DeferredResponse) RetvalType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DeferredResponse) MutateRetvalType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *DeferredResponse) Retval(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func DeferredResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DeferredResponseAddRetvalType(builder *flatbuffers.Builder, retvalType byte) {
	builder.PrependByteSlot(0, retvalType, 0)
}
func DeferredResponseAddRetval(builder *flatbuffers.Builder, retval flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(retval), 0)
}
func DeferredResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}