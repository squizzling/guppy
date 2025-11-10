package interpreter

import (
	"guppy/pkg/interpreter/itypes"
)

// Helper types to reduce boilerplate

type ThingOrNone[T itypes.Object] struct {
	None  *ObjectNone
	Thing T
}

func NewThingOrNoneNone[T itypes.Object]() ThingOrNone[T] {
	return ThingOrNone[T]{
		None: NewObjectNone(),
	}
}

func NewThingOrNoneThing[T itypes.Object](thing T) ThingOrNone[T] {
	return ThingOrNone[T]{
		Thing: thing,
	}
}
