package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	"guppy/internal/ast"
	"guppy/internal/ast/ast-flow"
)

func main() {
	defineAst(astflow.Package, astflow.Interfaces)
}

func defineAst(packageName string, interfaces ast.Interfaces) {
	fmt.Printf("package %s\n", packageName)
	fmt.Printf("\n")
	fmt.Printf("import (\n")
	fmt.Printf("\t\"guppy/internal/parser/tokenizer\"\n")
	fmt.Printf(")\n")

	for _, iface := range interfaces {
		fmt.Printf("\n")
		fmt.Printf("type Visitor%s interface {\n", iface.Name)
		for _, t := range iface.Nodes {
			typeName := iface.Name + t.Name
			fmt.Printf("\tVisit%s(%s %s) (any, error)\n", typeName, genReceiverName(typeName), typeName)
		}
		fmt.Printf("}\n")
		fmt.Printf("\n")
		fmt.Printf("type %s interface {\n", iface.Name)
		fmt.Printf("\tAccept(v%s Visitor%s) (any, error)\n", genReceiverName(iface.Name), iface.Name)
		fmt.Printf("}\n")
		for _, t := range iface.Nodes {
			defineType(iface.Name, t)
		}
	}
}

func defineType(interfaceName string, t ast.Node) {
	fmt.Printf("\n")
	structName := interfaceName + t.Name
	fmt.Printf("type %s struct {\n", structName)
	maxNameLen := 0
	for _, field := range t.Fields {
		maxNameLen = max(maxNameLen, len(field.Name))
	}
	for _, field := range t.Fields {
		fmt.Printf("\t%-*s %s\n", maxNameLen, field.Name, field.Type)
	}
	fmt.Printf("}\n")
	fmt.Printf("\n")
	receiverName := genReceiverName(structName)
	fmt.Printf("func New%s(", structName)
	for idx, field := range t.Fields {
		if idx == 0 { // makes gofmt happy
			fmt.Printf("\n")
		}
		fmt.Printf("\t%s %s,\n", field.Name, field.Type)
	}

	typeName := structName
	if t.NewConcrete {
		typeName = "&" + structName
		fmt.Printf(") *%s {\n", structName)
	} else {
		fmt.Printf(") %s {\n", interfaceName)
	}

	if len(t.Fields) == 0 {
		fmt.Printf("\treturn %s{}\n", typeName)
	} else {
		fmt.Printf("\treturn %s{\n", typeName)
		for _, field := range t.Fields {
			fmt.Printf("\t\t%-*s %s,\n", maxNameLen+1, field.Name+":", field.Name)
		}
		fmt.Printf("\t}\n")
	}

	fmt.Printf("}\n")
	fmt.Printf("\n")

	parameterType := "Visitor" + interfaceName
	parameterName := genReceiverName(parameterType)
	if parameterName == receiverName {
		parameterName = parameterName[:1] + parameterName
	}
	fmt.Printf("func (%s %s) Accept(%s %s) (any, error) {\n", receiverName, structName, parameterName, parameterType)
	fmt.Printf("\treturn %s.Visit%s(%s)\n", parameterName, structName, receiverName)
	fmt.Printf("}\n")
}

func genReceiverName(n string) string {
	return strings.ToLower(n[0:1] + string(slices.DeleteFunc([]rune(n[1:]), unicode.IsLower)))
}
