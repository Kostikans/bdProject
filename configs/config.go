package configs

import "os"

const Port = ":5000"

var BdConfig postgresConfig
var PrefixPath string

type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

func Init() {

	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBNameBDProject"),
	}
}
