package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ngiyshhk/caff/consts"
	"github.com/ngiyshhk/caff/gateway"
	"github.com/ngiyshhk/caff/infrastructure"
	"github.com/ngiyshhk/caff/model"
	"github.com/ngiyshhk/caff/utils"
	"os"
)

func main() {
	mysqlConfig := model.NewMysql()

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
	for _, table := range tables {
		args := &model.Schema{
			TableName:              table,
			SchemaName:             utils.ToCamelCase(table),
			RepositoryName:         "github.com/ngiyshhk/caffed",
			ModelPackageName:       "model",
			GatewayPackageName:     "gateway",
			IGatewayPackageName:    "igateway",
			IControllerPackageName: "icontroller",
			ControllerPackageName:  "controller",
			IUsecasePackageName:    "iusecase",
			UsecasePackageName:     "usecase",
		}

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
	}
}
