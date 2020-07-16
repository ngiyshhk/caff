package gateway

import (
	"fmt"
	"os"
	"strings"

	"github.com/ngiyshhk/caff/model"
)

// MainSrc .
type MainSrc struct {
	schemas []*model.Schema
}

// NewMainSrc .
func NewMainSrc(schemas []*model.Schema) *MainSrc {
	return &MainSrc{schemas: schemas}
}

// Write .
func (m *MainSrc) Write() error {
	s := m.schemas[0]
	var res []string
	res = append(res, "package main")
	res = append(res, "")
	res = append(res, "import (")
	res = append(res, `	"context"`)
	res = append(res, `	"fmt"`)
	res = append(res, `	"github.com/gin-gonic/gin"`)
	res = append(res, `	_ "github.com/go-sql-driver/mysql"`)
	res = append(res, fmt.Sprintf(`	"%s/%s"`, s.RepositoryName, s.ControllerPackageName))
	res = append(res, fmt.Sprintf(`	"%s/%s"`, s.RepositoryName, s.GatewayPackageName))
	res = append(res, fmt.Sprintf(`	"%s/infrastructure"`, s.RepositoryName))
	res = append(res, fmt.Sprintf(`	"%s/model"`, s.RepositoryName))
	res = append(res, fmt.Sprintf(`	"%s/%s"`, s.RepositoryName, s.UsecasePackageName))
	res = append(res, `	"os"`)
	res = append(res, `)`)
	res = append(res, ``)
	res = append(res, `func main() {`)
	res = append(res, `	mysqlConfig := model.NewMysql("root", "", "localhost", 3306, "hoge")`)
	res = append(res, ``)
	res = append(res, `	if err := infrastructure.InitDB(mysqlConfig.String()); err != nil {`)
	res = append(res, `		fmt.Println(err)`)
	res = append(res, `		os.Exit(1)`)
	res = append(res, `	}`)
	res = append(res, ``)
	res = append(res, `	r := gin.Default()`)
	for _, schema := range m.schemas {
		res = append(res, ``)
		res = append(res, fmt.Sprintf(`	%sGateway := %s.New%s()`, schema.LowerSchemaName(), schema.GatewayPackageName, schema.SchemaName()))
		res = append(res, fmt.Sprintf(`	%sUsecase := %s.New%s(%sGateway)`, schema.LowerSchemaName(), schema.UsecasePackageName, schema.SchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sController := %s.New%s(%sUsecase)`, schema.LowerSchemaName(), schema.ControllerPackageName, schema.SchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter := r.Group("%ss")`, schema.LowerSchemaName(), schema.LowerSnakeSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter.GET("", toGin(%sController.List))`, schema.LowerSchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter.GET(":id", toGin(%sController.Get))`, schema.LowerSchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter.POST("", toGin(%sController.Post))`, schema.LowerSchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter.PUT(":id", toGin(%sController.Put))`, schema.LowerSchemaName(), schema.LowerSchemaName()))
		res = append(res, fmt.Sprintf(`	%sRooter.DELETE(":id", toGin(%sController.Delete))`, schema.LowerSchemaName(), schema.LowerSchemaName()))
	}
	res = append(res, ``)
	res = append(res, `	r.Run()`)
	res = append(res, "}")
	res = append(res, ``)
	res = append(res, `func toGin(f func(context.Context)) func(*gin.Context) {`)
	res = append(res, `	return func(gc *gin.Context) {`)
	res = append(res, `		f(gc)`)
	res = append(res, `	}`)
	res = append(res, `}`)

	if err := os.MkdirAll(s.ParentDir, 0777); err != nil {
		return err
	}

	file, err := os.Create(s.ParentDir + "/main.go")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(strings.Join(res, "\n")))
	return err
}
