package dsn

type Connector interface {
	Driver() string
	DataSourceName() string
}
