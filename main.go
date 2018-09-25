package main

import (
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
		")"

	getterText = "\n" +
		"{{ range . }}" +
		`{{ if ne .Type "" }}` +
		"\n// Get{{ .Name }} returns field {{ .Name }} of the struct\n" +
		"func (this{{ .StructName }} *{{ .StructName }}) Get{{ .Name }}() {{ .Type }} {\n" +
		"\tif this{{ .StructName }} != nil {\n" +
		"\t\treturn this{{ .StructName }}.{{ .Name }}\n" +
		"\t}\n" +
		"\treturn {{ .TypeDefaultValue }}\n" +
		"}\n" +
		"{{ end }}" +
		"{{ end }}"

	importTemplate = template.Must(template.New("import").Parse(importText))
	getterTemplate = template.Must(template.New("getter").Parse(getterText))
)

func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	fatal(err, "parse file")

	out, err := os.Create(os.Args[2])
	fatal(err, "create file")

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
				fname := field.Names[0].Name
				ftype := getType(field.Type, "").(string)
				ftypeDefaultValue := getTypeDefaultValue(ftype)

				allGettersGenData = append(allGettersGenData, getterGenData{
					StructName:       currStructName,
					Name:             fname,
					Type:             ftype,
					TypeDefaultValue: ftypeDefaultValue,
				})
			}
		}
	}

	_, err = out.WriteString("package " + node.Name.Name)
	fatal(err, "failed to write package")

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
		return getType("", "")
	}
}

func getTypeDefaultValue(typ string) string {
	numTypes := []string{"int", "int8", "int16", "int32",
		"int64", "uint", "uint8", "uint16", "uint32",
		"uint64", "uintptr", "byte", "rune", "float32",
		"float32", "float64", "complex64", "complex128",
	}

	for _, numType := range numTypes {
		if typ == numType {
			return "0"
		}
	}
	if typ == "string" {
		return `""`
	}
	if typ == "bool" {
		return "false"
	}

	val := typ + "{}"

	if strings.HasPrefix(val, "[") {
		return val
	}
	if strings.Contains(val, "map") {
		return val
	}

	if strings.Contains(val, "*") {
		return strings.Replace(val, "*", "&", -1)
	}

	return val
}

func fatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
