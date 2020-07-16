package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ngiyshhk/caff/consts"
	"github.com/ngiyshhk/caff/gateway"
	"github.com/ngiyshhk/caff/infrastructure"
	"github.com/ngiyshhk/caff/model"
	"github.com/urfave/cli"
	"os"
)

func run(args []string) error {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "parentDir",
			Usage: "output parent directory",
			Value: ".",
		},
		cli.StringFlag{
			Name:     "repo",
			Usage:    "output repository name",
			Required: true,
		},
		cli.StringFlag{
			Name:  "table",
			Usage: "output target table",
		},
		cli.StringFlag{
			Name:     "mysql_user",
			Required: true,
		},
		cli.StringFlag{
			Name:  "mysql_pass",
			Value: "",
		},
		cli.StringFlag{
			Name:  "mysql_host",
			Value: "localhost",
		},
		cli.IntFlag{
			Name:  "mysql_port",
			Value: 3306,
		},
		cli.StringFlag{
			Name:     "mysql_dbname",
			Required: true,
		},
	}
	app.Action = exec
	return app.Run(args)
}

func exec(c *cli.Context) error {
	mysqlConfig := model.NewMysql(
		c.String("mysql_user"),
		c.String("mysql_pass"),
		c.String("mysql_host"),
		c.Int("mysql_port"),
		c.String("mysql_dbname"),
	)

	if err := infrastructure.InitDB(mysqlConfig.String()); err != nil {
		return err
	}

	columnGateway := gateway.NewColumn(infrastructure.GetDB())
	tables, err := columnGateway.ListTables(mysqlConfig.DBName)
	if err != nil {
		return err
	}

	schema := &model.Schema{
		ParentDir:              c.String("parentDir"),
		RepositoryName:         c.String("repo"),
		ModelPackageName:       "entity",
		GatewayPackageName:     "gateway",
		IGatewayPackageName:    "igateway",
		IControllerPackageName: "icontroller",
		ControllerPackageName:  "controller",
		IUsecasePackageName:    "iusecase",
		UsecasePackageName:     "usecase",
	}

	{
		srcGateway := gateway.NewSrc(schema)
		for _, f := range []func() error{
			srcGateway.WriteConstsContextKey,
			srcGateway.WriteGoMod,
			srcGateway.WriteGoSum,
			srcGateway.WriteUtilsContext,
			srcGateway.WriteModelMysql,
			srcGateway.WriteInfraMysql,
		} {
			if err := f(); err != nil {
				return err
			}
		}
	}

	var schemas []*model.Schema
	for _, table := range tables {
		args := model.SchemaWithTableName(schema, table)

		columns, err := columnGateway.ListColumns(mysqlConfig.DBName, table)
		if err != nil {
			return err
		}

		srcGateway := gateway.NewSrc(args)
		if err := srcGateway.WriteModel(args.LastModelPackageName(), columns); err != nil {
			return err
		}

		for _, layer := range consts.Templates {
			err = srcGateway.Write(layer)
			if err != nil {
				return err
			}
		}

		schemas = append(schemas, args)
	}

	mainSrcGateway := gateway.NewMainSrc(schemas)
	return mainSrcGateway.Write()
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
