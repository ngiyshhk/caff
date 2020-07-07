package gateway

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/ngiyshhk/caff/consts"
	"github.com/ngiyshhk/caff/model"
)

// Src .
type Src struct {
	schema *model.Schema
}

// NewSrc .
func NewSrc(schema *model.Schema) *Src {
	return &Src{schema: schema}
}

// WriteConstsContextKey .
func (s *Src) WriteConstsContextKey() error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/consts", 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/consts/context_key.go")
	if err != nil {
		return err
	}

	tpl, err := template.New("consts_context_key").Parse(consts.ConstsContextKeyTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteModelMysql .
func (s *Src) WriteModelMysql() error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/model", 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/model/mysql.go")
	if err != nil {
		return err
	}

	tpl, err := template.New("model_mysql").Parse(consts.ModelMysqlTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteInfraMysql .
func (s *Src) WriteInfraMysql() error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/infrastructure", 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/infrastructure/mysql.go")
	if err != nil {
		return err
	}

	tpl, err := template.New("infra_mysql").Parse(consts.InfraMysqlTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteUtilsContext .
func (s *Src) WriteUtilsContext() error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/utils", 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/utils/context.go")
	if err != nil {
		return err
	}

	tpl, err := template.New("utils_context").Parse(consts.UtilContextTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteGoMod .
func (s *Src) WriteGoMod() error {
	if err := os.MkdirAll(s.schema.ParentDir(), 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/go.mod")
	if err != nil {
		return err
	}

	tpl, err := template.New("gomod").Parse(consts.GoModTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteGoSum .
func (s *Src) WriteGoSum() error {
	if err := os.MkdirAll(s.schema.ParentDir(), 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/go.sum")
	if err != nil {
		return err
	}

	tpl, err := template.New("gomod").Parse(consts.GoSumTemplate)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
}

// WriteModel .
func (s *Src) WriteModel(packageName string, columns []model.Column) error {
	var res []string
	res = append(res, fmt.Sprintf("package %s", packageName))
	res = append(res, "")
	if useTime(columns) {
		res = append(res, "import \"time\"")
		res = append(res, "")
	}
	res = append(res, fmt.Sprintf("// %s .", s.schema.SchemaName()))
	res = append(res, fmt.Sprintf("type %s struct {", s.schema.SchemaName()))
	for _, column := range columns {
		res = append(res, fmt.Sprintf("	%s", column.String()))
	}
	res = append(res, "}")
	res = append(res, "")
	res = append(res, fmt.Sprintf("func (%s) TableName() string {", s.schema.SchemaName()))
	res = append(res, fmt.Sprintf("	return \"%s\"", s.schema.TableName))
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

// Write .
func (s *Src) Write(layer *model.Template) error {
	if err := os.MkdirAll(s.schema.ParentDir()+"/"+layer.PackageNameFunc(s.schema), 0777); err != nil {
		return err
	}

	file, err := os.Create(s.schema.ParentDir() + "/" + layer.PackageNameFunc(s.schema) + "/" + s.schema.TableName + ".go")
	if err != nil {
		return err
	}

	tpl, err := template.New(s.schema.TableName).Parse(layer.Value)
	if err != nil {
		return err
	}

	return tpl.Execute(file, s.schema)
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
