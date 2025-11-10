package interpreter

// Helper types to reduce boilerplate

type ThingOrNone[T Object] struct {
	None  *ObjectNone
	Thing T
}

func NewThingOrNoneNone[T Object]() ThingOrNone[T] {
	return ThingOrNone[T]{
		None: NewObjectNone(),
	}
}

func NewThingOrNoneThing[T Object](thing T) ThingOrNone[T] {
	return ThingOrNone[T]{
		Thing: thing,
	}
}
