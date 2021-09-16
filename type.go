package spdg

import "context"

/*
Type is an interface that can be given to other objects you are requesting data
from. It allows the other object to either provide the data requested, or some exception
when it decides not to. The requested object will not have to be bothered after that and
handling exceptions can be done from the Type the calling object owns.
*/
type Type interface {
	State() (Status, Reason)
	Peek(Value)
	Poke(Value)
}

type ProtoType struct {
	ctx   context.Context
	value Value
}

func NewProtoType(ctx context.Context) Type {
	return ProtoType{
		ctx:   ctx,
		value: NewProtoValue(ctx),
	}
}

func (t ProtoType) State() (Status, Reason) {
	return OK, BUSY
}

func (t ProtoType) Peek(value Value) {
	value = t.value
}

func (t ProtoType) Poke(value Value) {
	t.value = value
}
