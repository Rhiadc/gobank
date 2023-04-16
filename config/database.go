package config

import "fmt"

const (
	defaultDBHost     = "localhost"
	defaultDBUser     = "root"
	defaultDBPassword = "secret"
	defaultDBName     = "gobank"
	defaultDBPort     = "5432"
	defaultDBSSLMode  = "disable"
	defaultDBDriver   = "postgres"
)

type Database struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	Driver   string
}

func (db Database) GetURIConnection() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", db.Driver, db.User, db.Password, db.Host, db.Port, db.DBName, db.SSLMode)
}
