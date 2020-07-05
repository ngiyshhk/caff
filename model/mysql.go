package model

import "fmt"

type Mysql struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (m *Mysql) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.User, m.Pass, m.Host, m.Port, m.DBName)
}
