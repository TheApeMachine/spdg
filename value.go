package spdg

import "context"

/*
Value is a proxy object from Layer to Type that acts as a accessible for the composite that makes up Type.
There is probably no need to ever use this object stnad-alone or directly and mostly called while compiling
a type to layers.
*/
type Value interface {
	State() (Status, Reason)
	Peek(Layer)
	Poke(Layer)
}

type ProtoValue struct {
	layers []Layer
}

func NewProtoValue(ctx context.Context) Value {
	return ProtoValue{
		layers: make([]Layer, 0),
	}
}

func (value ProtoValue) State() (Status, Reason) {
	return OK, BUSY
}

func (value ProtoValue) Peek(layer Layer) {
}

func (value ProtoValue) Poke(layer Layer) {
}
