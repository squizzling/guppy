package ftypes

import (
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

// Helper types to reduce boilerplate

type ThingOrNone[T itypes.Object] struct {
	None  *primitive.ObjectNone
	Thing T
}

func NewThingOrNoneNone[T itypes.Object]() ThingOrNone[T] {
	return ThingOrNone[T]{
		None: primitive.NewObjectNone(),
	}
}

func NewThingOrNoneThing[T itypes.Object](thing T) ThingOrNone[T] {
	return ThingOrNone[T]{
		Thing: thing,
	}
}
