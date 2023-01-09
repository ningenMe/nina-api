package infra

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func GetMysqlConfig(dbName string) *mysql.Config {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	c := mysql.NewConfig()
	return &mysql.Config{
		DBName:    dbName,
		User:      os.Getenv("NINGENME_MYSQL_MASTER_USER_USERNAME"),
		Passwd:    os.Getenv("NINGENME_MYSQL_MASTER_USER_PASSWORD"),
		Addr:      os.Getenv("NINGENME_MYSQL_HOST") + ":" + os.Getenv("NINGENME_MYSQL_PORT"),
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
		MaxAllowedPacket: c.MaxAllowedPacket,
		AllowNativePasswords: c.AllowNativePasswords,
		CheckConnLiveness: c.CheckConnLiveness,
	}
}

var NingenmeMysql *sqlx.DB
var ComproMysql *sqlx.DB