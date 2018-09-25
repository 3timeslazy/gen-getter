package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

type getterGenData struct {
	StructName       string
	Name             string
	Type             string
	TypeDefaultValue string
}

var (
	importText = "\n" +
		"import (" +
		"\t{{ range . }}\n" +
		"\t{{ . }}" +
		"{{- end }}\n" +
		")\n"

	getterText = "\n" +
		"{{ range . }}" +
		"\n// Get{{ .Name }} returns  {{ .Name }} of the struct\n" +
		"func (this{{ .StructName }} *{{ .StructName }}) Get{{ .Name }}() {{ .Type }} {\n" +
		"\tif this{{ .StructName }} != nil {\n" +
		"\t\treturn this{{ .StructName }}.{{ .Name }}\n" +
		"\t}\n" +
		"\treturn {{ .TypeDefaultValue }}\n" +
		"}\n" +
		"{{ end }}\n"

	importTemplate = template.Must(template.New("import").Parse(importText))
	getterTemplate = template.Must(template.New("getter").Parse(getterText))
)

func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	fatal(err, "parse file")

	out, err := os.Create(os.Args[2])
	fatal(err, "create file")

	fmt.Fprintf(out, "package %s\n", node.Name.Name)

	allGettersGenData := []getterGenData{}
	imports := []string{}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			if impt, ok := spec.(*ast.ImportSpec); ok {
				imports = append(imports, impt.Path.Value)
				continue
			}

			currStructTypeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			currStructType, ok := currStructTypeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			currStructName := currStructTypeSpec.Name.Name

			for _, field := range currStructType.Fields.List {
				Name := field.Names[0].Name
				Type := getType(field.Type, "").(string)
				TypeDefaultValue := getTypeDefaultValue(Type)

				allGettersGenData = append(allGettersGenData, getterGenData{
					StructName:       currStructName,
					Name:             Name,
					Type:             Type,
					TypeDefaultValue: TypeDefaultValue,
				})
			}
		}
	}

	err = importTemplate.Execute(out, imports)
	fatal(err, "failed to make imports")

	err = getterTemplate.Execute(out, allGettersGenData)
	fatal(err, "failed to make getters")
}

func getType(dst interface{}, result string) interface{} {
	switch Type := dst.(type) {
	case string:
		return result + Type
	case *ast.Ident:
		return getType(Type.Name, result)
	case *ast.StarExpr:
		typ := getType(Type.X, "").(string)
		return getType("*"+typ, result)
	case *ast.ArrayType:
		length := ""
		if Type.Len != nil {
			length = getType(Type.Len, "").(string)
		}
		result += "[" + length + "]"
		return getType(Type.Elt, result)
	case *ast.MapType:
		key := getType(Type.Key, "").(string)
		value := getType(Type.Value, "").(string)
		return getType("map["+key+"]"+value, result)
	case *ast.SelectorExpr:
		typ := Type.Sel.Name
		pckg := getType(Type.X, "").(string)
		return pckg + "." + typ
	case *ast.BasicLit:
		return getType(Type.Value, result)
	default:
		panic(fmt.Sprintf("unsupported type %T"))
	}
}

func getTypeDefaultValue(typ string) (val string) {
	basicTypes := map[string]string{
		"int":        "0",
		"int8":       "0",
		"int16":      "0",
		"int32":      "0",
		"int64":      "0",
		"uint":       "0",
		"uint8":      "0",
		"uint16":     "0",
		"uint32":     "0",
		"uint64":     "0",
		"uintptr":    "0",
		"byte":       "0",
		"rune":       "0",
		"float32":    "0.0",
		"float64":    "0.0",
		"complex64":  "(0 + 0i)",
		"complex128": "(0 + 0i)",
		"string":     `""`,
		"bool":       "false",
	}

	if v, ok := basicTypes[typ]; ok {
		return v
	}

	val = typ + "{}"

	if strings.HasPrefix(val, "[") {
		return val
	}

	if strings.Contains(val, "map") {
		return val
	}

	if strings.Contains(val, "*") {
		val = strings.Replace(val, "*", "&", -1)
	}

	return val
}

func fatal(err error, msg string, args ...interface{}) {
	if err != nil {
		log.Fatalf(fmt.Sprintf(msg, args...)+": %s", err.Error())
	}
}
