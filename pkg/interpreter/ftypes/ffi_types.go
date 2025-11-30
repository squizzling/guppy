package ftypes

import (
	"github.com/squizzling/guppy/pkg/interpreter"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
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

type ThingOrMissing[T itypes.Object] struct {
	Missing *interpreter.ObjectMissing
	Thing   T
}

func NewThingOrMissingNone[T itypes.Object]() ThingOrMissing[T] {
	return ThingOrMissing[T]{
		Missing: interpreter.NewObjectMissing(),
	}
}

func NewThingOrMissingThing[T itypes.Object](thing T) ThingOrMissing[T] {
	return ThingOrMissing[T]{
		Thing: thing,
	}
}
