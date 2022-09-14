package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"log"
	"time"
)

type ContributionRepository struct{}

func (ContributionRepository) GetList() []*domainmodel.Contribution {
	rows, err := NingenmeMysql.Queryx(`SELECT contributed_at, organization, repository, user, status FROM github_contribution`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var list []*domainmodel.Contribution
	for rows.Next() {
		c := &domainmodel.Contribution{}
		if err = rows.StructScan(c); err != nil {
			log.Fatalln(err)
		}
		list = append(list, c)
	}

	return list
}


func (ContributionRepository) InsertList(contributionList []*domainmodel.Contribution) {
	if len(contributionList) == 0 {
		return
	}
	_, err := NingenmeMysql.NamedExec(`INSERT INTO github_contribution (contributed_at, organization, repository, user, status) 
                                 VALUES (:contributed_at, :organization, :repository, :user, :status)`, contributionList)
	if err != nil {
		log.Fatalln(err)
	}
}

func (ContributionRepository) Delete(startAt time.Time, endAt time.Time) {
	_, err := NingenmeMysql.NamedExec(`DELETE FROM github_contribution WHERE contributed_at BETWEEN :start_at AND :end_at)`,
		map[string]interface{}{
		    "start_at": startAt.Format(time.RFC3339),
			"end_at": endAt.Format(time.RFC3339),
		})
	if err != nil {
		log.Fatalln(err)
	}
}
