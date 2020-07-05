package model

import (
	"github.com/ngiyshhk/caff/utils"
	"strings"
)

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
	SchemaName             string
}

func (s Schema) LowerInitialSchemaName() string {
	return strings.ToLower(string(s.SchemaName[0]))
}

func (s Schema) LowerSchemaName() string {
	return s.LowerInitialSchemaName() + s.SchemaName[1:]
}

func (s Schema) LowerSnakeSchemaName() string {
	return utils.ToSnakeCase(s.SchemaName)
}

func (s Schema) LastModelPackageName() string {
	res := strings.Split(s.ModelPackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastIGatewayPackageName() string {
	res := strings.Split(s.IGatewayPackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastGatewayPackageName() string {
	res := strings.Split(s.GatewayPackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastIControllerPackageName() string {
	res := strings.Split(s.IControllerPackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastControllerPackageName() string {
	res := strings.Split(s.ControllerPackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastIUsecasePackageName() string {
	res := strings.Split(s.IUsecasePackageName, "/")
	return res[len(res)-1]
}

func (s Schema) LastUsecasePackageName() string {
	res := strings.Split(s.UsecasePackageName, "/")
	return res[len(res)-1]
}

func (s Schema) ParentDir() string {
	return strings.Join([]string{".", s.RepositoryName}, "/")
}
