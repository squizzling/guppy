package ftypes

import (
	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

// Helper types to reduce boilerplate

type ThingOrNone[T itypes.Object] struct {
	None  *interpreter.ObjectNone
	Thing T
}

func NewThingOrNoneNone[T itypes.Object]() ThingOrNone[T] {
	return ThingOrNone[T]{
		None: interpreter.NewObjectNone(),
	}
}

func NewThingOrNoneThing[T itypes.Object](thing T) ThingOrNone[T] {
	return ThingOrNone[T]{
		Thing: thing,
	}
}
