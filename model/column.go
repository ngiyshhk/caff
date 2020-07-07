package model

import (
	"fmt"

	"github.com/ngiyshhk/caff/utils"
)

// Column .
type Column struct {
	TableName              string `gorm:"column:TABLE_NAME"`
	ColumnName             string `gorm:"column:COLUMN_NAME"`
	IsNullable             string `gorm:"column:IS_NULLABLE"`
	DataType               string `gorm:"column:DATA_TYPE"`
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
}

func (c *Column) pointerPrefix() string {
	if c.IsNullable == "YES" {
		return "*"
	}
	return ""
}

func (c *Column) golangDataType() string {
	switch c.DataType {
	case "char", "varchar", "text", "enum", "json":
		return "string"
	case "int", "bigint", "smallint", "tinyint":
		return "int"
	case "decimal":
		return "float64"
	case "datetime", "date":
		return "time.Time"
	default:
		return "string"
	}
}

// String .
func (c *Column) String() string {
	return fmt.Sprintf(
		"%s %s `gorm:\"column:%s\"`",
		utils.ToCamelCase(c.ColumnName),
		c.pointerPrefix()+c.golangDataType(),
		c.ColumnName,
	)
}
