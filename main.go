package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/Originate/exit"
)

type structGenData struct {
	Name   string
	Fields []fieldGenData
}

type fieldGenData struct {
	StructName            string
	FieldName             string
	FieldType             string
	FieldTypeDefaultValue string
}

type getterGenData struct {
	StructName            string
	FieldName             string
	FieldType             string
	FieldTypeDefaultValue string
}

var (
	fieldGetter = `
// Get{{ .FieldName }} do smth...
func (this *{{ .StructName }}) Get{{ .FieldName }}() {{ .FieldType }} {
	if this != nil {
		return this.{{ .FieldName }}
	}
	return {{ .FieldTypeDefaultValue }}
}
`

	fieldGetterTemplate = template.Must(template.New("getter").Parse(fieldGetter))
)

func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	exit.IfWrap(err, "parse file")

	out, err := os.Create(os.Args[2])
	exit.IfWrap(err, "create file")

	fmt.Fprintf(out, "package %s\n", node.Name.Name)

	allGettersGenData := []getterGenData{}

	singleGetterGenData := getterGenData{}
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			currStructTypeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			currStructType, ok := currStructTypeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			singleGetterGenData.StructName = currStructTypeSpec.Name.Name

			for _, field := range currStructType.Fields.List {
				singleGetterGenData.FieldName = field.Names[0].Name

				switch fieldType := field.Type.(type) {
				case *ast.StarExpr:
					singleGetterGenData.FieldType = "*" + fieldType.X.(*ast.Ident).Name
				case *ast.Ident:
					singleGetterGenData.FieldType = fieldType.Name
				case *ast.ArrayType:
					elt := fieldType.Elt
					switch v := elt.(type) {
					case *ast.Ident:
						singleGetterGenData.FieldType = "[]" + v.Name
					case *ast.StarExpr:
						singleGetterGenData.FieldType = "[]*" + v.X.(*ast.Ident).Name
					}
				case *ast.MapType:
					singleGetterGenData.FieldType = "map[" + fieldType.Key.(*ast.Ident).Name + "]"
					switch v := fieldType.Value.(type) {
					case *ast.StarExpr:
						singleGetterGenData.FieldType += "*" + v.X.(*ast.Ident).Name
					case *ast.Ident:
						singleGetterGenData.FieldType += v.Name
					}
				}

				currTypeDefaultValue := typeDefaultValue(singleGetterGenData.FieldType)
				singleGetterGenData.FieldTypeDefaultValue = currTypeDefaultValue

				allGettersGenData = append(allGettersGenData, singleGetterGenData)
			}
		}
	}

	for _, genData := range allGettersGenData {
		genGetter(out, &genData)
	}
}

func genGetter(w io.Writer, genData *getterGenData) {
	err := fieldGetterTemplate.Execute(w, genData)
	exit.IfWrapf(err, "gen getter for field %s", genData.FieldName)
}

func typeDefaultValue(typ string) (val string) {
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

	if strings.Contains(val, "[]") {
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
