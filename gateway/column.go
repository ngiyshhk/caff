package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/ngiyshhk/caff/model"
)

type Column struct {
	db *gorm.DB
}

func NewColumn(db *gorm.DB) *Column {
	return &Column{db: db}
}

func (c *Column) ListTables(schemaName string) ([]string, error) {
	type res struct {
		TableName string
	}
	var tmpResult []res
	c.db.Table("information_schema.COLUMNS").
		Select("DISTINCT(TABLE_NAME) AS table_name").
		Where("TABLE_SCHEMA = ?", schemaName).
		Find(&tmpResult)
	if c.db.Error != nil {
		return nil, c.db.Error
	}

	var result []string
	for _, r := range tmpResult {
		result = append(result, r.TableName)
	}
	return result, nil
}

func (c *Column) ListColumns(schemaName, tableName string) ([]model.Column, error) {
	var result []model.Column
	c.db.Table("information_schema.COLUMNS").
		Where("TABLE_SCHEMA = ? AND TABLE_NAME = ?", schemaName, tableName).
		Find(&result)
	if c.db.Error != nil {
		return nil, c.db.Error
	}
	return result, nil
}
