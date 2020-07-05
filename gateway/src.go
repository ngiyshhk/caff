package gateway

import (
	"fmt"
	"github.com/ngiyshhk/caff/model"
	"os"
	"strings"
	"text/template"
)

type Src struct {
	schema *model.Schema
}

func NewSrc(schema *model.Schema) *Src {
	return &Src{schema: schema}
}

func (s *Src) WriteModel(packageName string, columns []model.Column) error {
	var res []string
	res = append(res, fmt.Sprintf("package %s", packageName))
	res = append(res, "")
	if useTime(columns) {
		res = append(res, "import \"time\"")
		res = append(res, "")
	}
	res = append(res, fmt.Sprintf("// %s .", s.schema.SchemaName))
	res = append(res, fmt.Sprintf("type %s struct {", s.schema.SchemaName))
	for _, column := range columns {
		res = append(res, fmt.Sprintf("	%s", column.String()))
	}
	res = append(res, "}")

	if err := os.MkdirAll(s.schema.ParentDir()+"/"+s.schema.LastModelPackageName(), 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/" + s.schema.LastModelPackageName() + "/" + s.schema.TableName + ".go")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(strings.Join(res, "\n")))
	return err
}

func (s *Src) Write(layer *model.Template) error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/"+layer.PackageNameFunc(s.schema), 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/" + layer.PackageNameFunc(s.schema) + "/" + s.schema.TableName + ".go")
	if err != nil {
		return err
	}

	tpl, err := template.New(s.schema.TableName).Parse(layer.Value)
	err = tpl.Execute(file, s.schema)
	return err
}

func useTime(columns []model.Column) bool {
	for _, column := range columns {
		switch column.DataType {
		case "datetime", "date":
			return true
		}
	}
	return false
}
