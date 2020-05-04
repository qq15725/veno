package dsn

import "fmt"

type Mysql struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Charset  string
}

func (mysql *Mysql) Driver() string {
	return "mysql"
}

func (mysql *Mysql) DataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.Database,
		mysql.Charset,
	)
}
