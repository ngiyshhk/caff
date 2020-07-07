package model

import "fmt"

// Mysql .
type Mysql struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

// NewMysql .
func NewMysql(user, pass, host string, port int, dbName string) *Mysql {
	return &Mysql{
		User:   user,
		Pass:   pass,
		Host:   host,
		Port:   port,
		DBName: dbName,
	}
}

// String .
func (m *Mysql) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.User, m.Pass, m.Host, m.Port, m.DBName)
}
