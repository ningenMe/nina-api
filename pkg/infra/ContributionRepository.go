package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ContributionRepository struct{}

type Contribution struct {
	Time   time.Time
	Org    string
	Repo   string
	User   string
	Status string
}

type NingenmeMysql struct {
	User string
	Password string
	Host string
	Port string
}

var ningenmeMysql = NingenmeMysql{
	User: os.Getenv("NINGENME_MYSQL_MASTER_USER_USERNAME"),
	Password: os.Getenv("NINGENME_MYSQL_MASTER_USER_PASSWORD"),
	Host: os.Getenv("NINGENME_MYSQL_HOST"),
	Port: os.Getenv("NINGENME_MYSQL_PORT"),
}

func (ContributionRepository) GetList() []*Contribution {
	conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/ningenme", ningenmeMysql.User,ningenmeMysql.Password,ningenmeMysql.Host,ningenmeMysql.Port)

	db, err := sql.Open("mysql", conf + "?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT contributed_at, organization, repository, user, status FROM github_contribution")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var list []*Contribution
	for rows.Next() {
		c := &Contribution{}
		if err = rows.Scan(&c.Time, &c.Org, &c.Repo, &c.User, &c.Status); err != nil {
			log.Fatalln(err)
		}

		fmt.Println(c)
		list = append(list, c)
	}

	return list
}

