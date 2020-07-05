package consts

import "github.com/ngiyshhk/caff/model"

const IGatewayTemplate = `
package {{.LastIGatewayPackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} interface {
	Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error)
	List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error)
	Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Delete(ctx context.Context, id string) error
}
`

const GatewayTemplate = `
package {{.LastGatewayPackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.IGatewayPackageName}}"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
	"{{.RepositoryName}}/utils"
)

// {{.SchemaName}} .
type {{.SchemaName}} struct {
}

// New{{.SchemaName}} .
func New{{.SchemaName}}() {{.IGatewayPackageName}}.{{.SchemaName}} {
	return &{{.SchemaName}}{}
}

// Get .
func ({{.LowerInitialSchemaName}} *{{.LowerSchemaName}}) Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result := &{{.ModelPackageName}}.{{.SchemaName}}{}
	err = db.Where("id = ?", id).First(result).Error
	return result, err
}

// List .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var result []*{{.ModelPackageName}}.{{.SchemaName}}
	err = db.Find(&result).Error
	return result, err
}

// Create .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Create({{.LowerSchemaName}}).Error
}

// Update .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Update({{.LowerSchemaName}}).Error
}

// Delete .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Delete(ctx context.Context, id string) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(&{{.ModelPackageName}}.{{.SchemaName}}{}).Error
}
`

const IControllerTemplate = `
package {{.LastIControllerPackageName}}

import "github.com/gin-gonic/gin"

type {{.SchemaName}} interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
`

const IUsecaseTemplate = `
package {{.LastIUsecasePackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} interface {
	Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error)
	List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error)
	Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Delete(ctx context.Context, id string) error
}
`

const UsecaseTemplate = `
package {{.LastUsecasePackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.IGatewayPackageName}}"
	"{{.RepositoryName}}/{{.IUsecasePackageName}}"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} struct {
	{{.LowerSchemaName}}Gateway {{.IGatewayPackageName}}.{{.SchemaName}}
}

func New{{.SchemaName}}({{.LowerSchemaName}}Gateway {{.IGatewayPackageName}}.{{.SchemaName}}) {{.IUsecasePackageName}}.{{.SchemaName}}) {
	return &{{.SchemaName}}{ {{.LowerSchemaName}}Gateway: {{.LowerSchemaName}}Gateway }
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}), error) {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Get(ctx, id)
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.List(ctx)
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Create(ctx, {{.LowerSchemaName}})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Update(ctx, {{.SchemaName}})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Delete(ctx context.Context, id string) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Delete(ctx, id)
}
`

var Templates = []*model.Template{
	{
		Type:  ControllerTemplateType,
		Value: "ControllerTemplateType",
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.ControllerPackageName
		},
	},
	{
		Type:  IControllerTemplateType,
		Value: IControllerTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IControllerPackageName
		},
	},
	{
		Type:  GatewayTemplateType,
		Value: GatewayTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.GatewayPackageName
		},
	},
	{
		Type:  IGatewayTemplateType,
		Value: IGatewayTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IGatewayPackageName
		},
	},
	{
		Type:  UsecaseTemplateType,
		Value: UsecaseTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.UsecasePackageName
		},
	},
	{
		Type:  IUsecaseTemplateType,
		Value: IUsecaseTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IUsecasePackageName
		},
	},
}
