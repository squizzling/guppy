package interpreter

type ObjectMissing struct {
	Object
}

func NewObjectMissing() Object {
	return &ObjectMissing{
		NewObject(map[string]Object{}),
	}
}

func (om *ObjectMissing) Repr() string {
	return "Missing"
}
