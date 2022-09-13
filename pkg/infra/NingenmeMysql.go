package infra

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type MysqlConfig struct {
	User string
	Password string
	Host string
	Port string
}

func (c MysqlConfig) GetConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/ningenme", c.User,c.Password,c.Host,c.Port) + "?parseTime=true&loc=Asia%2FTokyo"
}

var NingenmeMysqlConfig = MysqlConfig{
	User: os.Getenv("NINGENME_MYSQL_MASTER_USER_USERNAME"),
	Password: os.Getenv("NINGENME_MYSQL_MASTER_USER_PASSWORD"),
	Host: os.Getenv("NINGENME_MYSQL_HOST"),
	Port: os.Getenv("NINGENME_MYSQL_PORT"),
}


var NingenmeMysql *sqlx.DB