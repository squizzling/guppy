package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	"guppy/internal/gen"
)

func main() {
	defineAst(gen.Package, gen.Interfaces)
}

func defineAst(packageName string, interfaces []gen.Interface) {
	fmt.Printf("package %s\n", packageName)
	fmt.Printf("\n")
	fmt.Printf("import (\n")
	fmt.Printf("\t\"fmt\"\n")
	fmt.Printf("\t\"strings\"\n")
	fmt.Printf(")\n")
	fmt.Printf("\n")
	fmt.Printf("type DebugWriter struct {\n")
	fmt.Printf("\tdepth int\n")
	fmt.Printf("}\n")
	fmt.Printf("\n")
	fmt.Printf("func (dw DebugWriter) p() string {\n")
	fmt.Printf("\treturn strings.Repeat(\" \", 2*dw.depth)\n")
	fmt.Printf("}\n")
	fmt.Printf("\n")
	fmt.Printf("func (dw *DebugWriter) i() {\n")
	fmt.Printf("\tdw.depth++\n")
	fmt.Printf("}\n")
	fmt.Printf("\n")
	fmt.Printf("func (dw *DebugWriter) o() {\n")
	fmt.Printf("\tdw.depth--\n")
	fmt.Printf("}\n")
	fmt.Printf("\n")
	fmt.Printf("func s(a any, err error) string {\n")
	fmt.Printf("\treturn a.(string)\n")
	fmt.Printf("}\n")

	for _, iface := range interfaces {
		for _, t := range iface.Nodes {
			fmt.Printf("\n")
			typeName := iface.Name + t.Name
			if typeName == "dw" {
				typeName = typeName + "2"
			}
			receiverName := genReceiverName(typeName)
			fmt.Printf("func (dw DebugWriter) Visit%s(%s %s) (any, error) {\n", typeName, receiverName, typeName)
			//fmt.Printf("\tfmt.Printf(\"Entering %s\\n\")\n", typeName)
			fmt.Printf("\t_s := \"%s(\\n\"\n", typeName)
			fmt.Printf("\tdw.i()\n")
			for idx, field := range t.Fields {
				if gen.IsInterface(field.Type) {
					fmt.Printf("\tif %s.%s != nil {\n", receiverName, field.Name)
					fmt.Printf("\t\t_s += dw.p() + \"%s: \" + s(%s.%s.Accept(dw)) // IsInterface\n", field.Name, receiverName, field.Name)
					fmt.Printf("\t} else {\n")
					fmt.Printf("\t\t_s += dw.p() + \"%s: nil\\n\"\n", field.Name)
					fmt.Printf("\t}\n")
				} else if gen.IsConcrete(field.Type) {
					fmt.Printf("\t_s += dw.p() + \"%s: \" + s(%s.%s.Accept(dw)) // IsConcrete\n", field.Name, receiverName, field.Name)
				} else if gen.IsInterfaceArray(field.Type) || gen.IsConcreteArray(field.Type) { // Done
					fmt.Printf("\tif %s.%s == nil {\n", receiverName, field.Name)
					fmt.Printf("\t\t_s += dw.p() + \"%s: nil\\n\"\n", field.Name)
					fmt.Printf("\t} else if len(%s.%s) == 0 {\n", receiverName, field.Name)
					fmt.Printf("\t\t_s += dw.p() + \"%s: []\\n\"\n", field.Name)
					fmt.Printf("\t} else {\n")
					fmt.Printf("\t\t_s += dw.p() + \"%s: [\\n\"\n", field.Name)
					fmt.Printf("\t\tdw.i()\n")
					fmt.Printf("\t\tfor _, _r := range %s.%s {\n", receiverName, field.Name)
					fmt.Printf("\t\t\t_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray\n")
					fmt.Printf("\t\t}\n")
					fmt.Printf("\t\tdw.o()\n")
					fmt.Printf("\t\t_s += dw.p() + \"]\\n\"\n")
					fmt.Printf("\t}\n")
				} else if field.Type == "string" { // Done
					fmt.Printf("\t_s += dw.p() + \"%s: string(\" + %s.%s + \")\\n\"\n", field.Name, receiverName, field.Name)
				} else if field.Type == "bool" { // Done
					fmt.Printf("\t_s += dw.p() + \"%s: bool(\" + fmt.Sprintf(\"%%t\", %s.%s) + \")\\n\"\n", field.Name, receiverName, field.Name)
				} else if field.Type == "[]string" { // Done
					fmt.Printf("\tif %s.%s == nil {\n", receiverName, field.Name)
					fmt.Printf("\t\t_s += dw.p() + \"%s: nil\\n\"\n", field.Name)
					fmt.Printf("\t} else if len(%s.%s) == 0 {\n", receiverName, field.Name)
					fmt.Printf("\t\t_s += dw.p() + \"%s: []\\n\"\n", field.Name)
					fmt.Printf("\t} else {\n")
					fmt.Printf("\t\t_s += dw.p() + \"%s: [\\n\"\n", field.Name)
					fmt.Printf("\t\tdw.i()\n")
					fmt.Printf("\t\tfor _, _r := range %s.%s {\n", receiverName, field.Name)
					fmt.Printf("\t\t\t_s += dw.p() + _r + \"\\n\" // []string\n")
					fmt.Printf("\t\t}\n")
					fmt.Printf("\t\tdw.o()\n")
					fmt.Printf("\t\t_s += dw.p() + \"]\\n\"\n")
					fmt.Printf("\t}\n")
				} else {
					fmt.Printf("\t// TODO: %d %s %s\n", idx, field.Name, field.Type)
					fmt.Printf("\t_s += dw.p() + fmt.Sprintf(\"%s: %%T(%%v)\\n\", %s.%s, %s.%s)\n", field.Name, receiverName, field.Name, receiverName, field.Name)
				}
			}
			fmt.Printf("\tdw.o()\n")
			fmt.Printf("\t_s += dw.p() + \")\\n\"\n")
			//fmt.Printf("\tfmt.Printf(\"Leaving %s: %%s\\n\", _s)\n", typeName)
			fmt.Printf("\treturn _s, nil\n")
			fmt.Printf("}\n")
		}
	}
}

func defineType(interfaceName string, nodeType string, fields []gen.Field) {
	fmt.Printf("\n")
	structName := interfaceName + nodeType
	fmt.Printf("type %s struct {\n", structName)
	maxNameLen := 0
	for _, field := range fields {
		maxNameLen = max(maxNameLen, len(field.Name))
	}
	for _, field := range fields {
		fmt.Printf("\t%-*s %s\n", maxNameLen, field.Name, field.Type)
	}
	fmt.Printf("}\n")
	fmt.Printf("\n")
	receiverName := genReceiverName(structName)
	fmt.Printf("func New%s(", structName)
	for idx, field := range fields {
		if idx == 0 { // makes gofmt happy
			fmt.Printf("\n")
		}
		fmt.Printf("\t%s %s,\n", field.Name, field.Type)
	}
	fmt.Printf(") %s {\n", interfaceName)

	if len(fields) == 0 {
		fmt.Printf("\treturn &%s{}\n", structName)
	} else {
		fmt.Printf("\treturn &%s{\n", structName)
		for _, field := range fields {
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
	fmt.Printf("func (%s %s) Accept(%s %s) any {\n", receiverName, structName, parameterName, parameterType)
	fmt.Printf("\treturn %s.Visit%s(%s)\n", parameterName, structName, receiverName)
	fmt.Printf("}\n")
}

func genReceiverName(n string) string {
	return strings.ToLower(n[0:1] + string(slices.DeleteFunc([]rune(n[1:]), unicode.IsLower)))
}
