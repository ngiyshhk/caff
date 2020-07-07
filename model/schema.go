package model

import (
	"strings"

	"github.com/ngiyshhk/caff/utils"
)

// Schema .
type Schema struct {
	RepositoryName         string
	ModelPackageName       string
	IGatewayPackageName    string
	GatewayPackageName     string
	IControllerPackageName string
	ControllerPackageName  string
	IUsecasePackageName    string
	UsecasePackageName     string
	TableName              string
}

// SchemaWithTableName .
func SchemaWithTableName(schema *Schema, tableName string) *Schema {
	return &Schema{
		RepositoryName:         schema.RepositoryName,
		ModelPackageName:       schema.ModelPackageName,
		IGatewayPackageName:    schema.IGatewayPackageName,
		GatewayPackageName:     schema.GatewayPackageName,
		IControllerPackageName: schema.IControllerPackageName,
		ControllerPackageName:  schema.ControllerPackageName,
		IUsecasePackageName:    schema.IUsecasePackageName,
		UsecasePackageName:     schema.UsecasePackageName,
		TableName:              tableName,
	}
}

// SchemaName .
func (s Schema) SchemaName() string {
	return utils.ToCamelCase(s.TableName)
}

// LowerInitialSchemaName .
func (s Schema) LowerInitialSchemaName() string {
	return strings.ToLower(string(s.SchemaName()[0]))
}

// LowerSchemaName .
func (s Schema) LowerSchemaName() string {
	return s.LowerInitialSchemaName() + s.SchemaName()[1:]
}

// LowerSnakeSchemaName .
func (s Schema) LowerSnakeSchemaName() string {
	return utils.ToSnakeCase(s.SchemaName())
}

// LastModelPackageName .
func (s Schema) LastModelPackageName() string {
	res := strings.Split(s.ModelPackageName, "/")
	return res[len(res)-1]
}

// LastIGatewayPackageName .
func (s Schema) LastIGatewayPackageName() string {
	res := strings.Split(s.IGatewayPackageName, "/")
	return res[len(res)-1]
}

// LastGatewayPackageName .
func (s Schema) LastGatewayPackageName() string {
	res := strings.Split(s.GatewayPackageName, "/")
	return res[len(res)-1]
}

// LastIControllerPackageName .
func (s Schema) LastIControllerPackageName() string {
	res := strings.Split(s.IControllerPackageName, "/")
	return res[len(res)-1]
}

// LastControllerPackageName .
func (s Schema) LastControllerPackageName() string {
	res := strings.Split(s.ControllerPackageName, "/")
	return res[len(res)-1]
}

// LastIUsecasePackageName .
func (s Schema) LastIUsecasePackageName() string {
	res := strings.Split(s.IUsecasePackageName, "/")
	return res[len(res)-1]
}

// LastUsecasePackageName .
func (s Schema) LastUsecasePackageName() string {
	res := strings.Split(s.UsecasePackageName, "/")
	return res[len(res)-1]
}

// ParentDir .
func (s Schema) ParentDir() string {
	return strings.Join([]string{".", s.RepositoryName}, "/")
}
