package interpreter

// Helper types to reduce boilerplate

type ThingOrNone[T Object] struct {
	None  *ObjectNone
	Thing T
}
