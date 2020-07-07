package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ngiyshhk/caff/consts"
	"github.com/ngiyshhk/caff/gateway"
	"github.com/ngiyshhk/caff/infrastructure"
	"github.com/ngiyshhk/caff/model"
	"os"
)

func main() {
	mysqlConfig := model.NewMysql("root", "", "localhost", 3306, "a")

	if err := infrastructure.InitDB(mysqlConfig.String()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	columnGateway := gateway.NewColumn(infrastructure.GetDB())
	tables, err := columnGateway.ListTables(mysqlConfig.DBName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	schema := &model.Schema{
		RepositoryName:         "github.com/ngiyshhk/caffed",
		ModelPackageName:       "model",
		GatewayPackageName:     "gateway",
		IGatewayPackageName:    "igateway",
		IControllerPackageName: "icontroller",
		ControllerPackageName:  "controller",
		IUsecasePackageName:    "iusecase",
		UsecasePackageName:     "usecase",
	}

	{
		srcGateway := gateway.NewSrc(schema)
		if err := srcGateway.WriteConstsContextKey(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := srcGateway.WriteGoMod(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := srcGateway.WriteGoSum(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := srcGateway.WriteUtilsContext(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := srcGateway.WriteModelMysql(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := srcGateway.WriteInfraMysql(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var schemas []*model.Schema
	for _, table := range tables {
		args := model.SchemaWithTableName(schema, table)

		columns, err := columnGateway.ListColumns(mysqlConfig.DBName, table)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		srcGateway := gateway.NewSrc(args)
		if err := srcGateway.WriteModel(args.LastModelPackageName(), columns); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, layer := range consts.Templates {
			err = srcGateway.Write(layer)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		schemas = append(schemas, args)
	}

	mainSrcGateway := gateway.NewMainSrc(schemas)
	if err := mainSrcGateway.Write(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
