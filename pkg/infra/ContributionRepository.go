package infra

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ContributionRepository struct{}

type Contribution struct {
	Time   time.Time `db:"contributed_at"`
	Org    string    `db:"organization"`
	Repo   string    `db:"repository"`
	User   string    `db:"user"`
	Status string    `db:"status"`
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

func (ningenmeMysql NingenmeMysql) getConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/ningenme", ningenmeMysql.User,ningenmeMysql.Password,ningenmeMysql.Host,ningenmeMysql.Port)
}

func (ContributionRepository) GetList() []*Contribution {
	fmt.Println(ningenmeMysql.getConfig())

	db, err := sqlx.Open("mysql",  ningenmeMysql.getConfig() + "?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Queryx(`SELECT contributed_at, organization, repository, user, status FROM github_contribution`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var list []*Contribution
	for rows.Next() {
		c := &Contribution{}
		if err = rows.StructScan(c); err != nil {
			log.Fatalln(err)
		}
		list = append(list, c)
	}

	return list
}


func (ContributionRepository) InsertList(contributionList []Contribution) {

	db, err := sqlx.Open("mysql",  ningenmeMysql.getConfig() + "?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.NamedExec(`INSERT INTO github_contribution (contributed_at, organization, repository, user, status) 
                                 VALUES (:contributed_at, :organization, :repository, :user, :status)`, contributionList)
	if err != nil {
		log.Fatalln(err)
	}
}
