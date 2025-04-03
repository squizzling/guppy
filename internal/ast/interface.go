package ast

import (
	"strings"
)

type Interfaces []Interface

func (is Interfaces) IsConcrete(name string) bool {
	name = strings.TrimLeft(name, "*")
	for _, i := range is {
		for _, t := range i.Nodes {
			if i.Name+t.Name == name {
				return true
			}
		}
	}
	return false
}

func (is Interfaces) IsConcreteArray(name string) bool {
	if strings.HasPrefix(name, "[]") {
		return is.IsConcrete(name[2:])
	}
	return false
}

func (is Interfaces) IsInterfaceArray(name string) bool {
	if strings.HasPrefix(name, "[]") {
		return is.IsInterface(name[2:])
	}
	return false
}

func (is Interfaces) IsInterface(name string) bool {
	for _, i := range is {
		if i.Name == name {
			return true
		}
	}
	return false
}

type Interface struct {
	Name  string
	Nodes []Node
}

type Node struct {
	Name        string
	NewConcrete bool
	Fields      []Field
}

type Field struct {
	Name string
	Type string
}
