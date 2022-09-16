package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"time"
)

type ContributionRepository struct{}

func (ContributionRepository) GetList() []*domainmodel.Contribution {
	rows, err := NingenmeMysql.Queryx(`SELECT contributed_at, organization, repository, user, status FROM github_contribution`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var list []*domainmodel.Contribution
	for rows.Next() {
		c := &domainmodel.Contribution{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	return list
}


func (ContributionRepository) InsertList(contributionList []*domainmodel.Contribution) {

	chunkSize := 1000

	for _, partitionedList := range domainmodel.PartitionedList[domainmodel.Contribution](contributionList, chunkSize) {
		_, err := NingenmeMysql.NamedExec(`INSERT INTO github_contribution (contributed_at, organization, repository, user, status) 
                                 VALUES (:contributed_at, :organization, :repository, :user, :status)`, partitionedList)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 2)
	}
}

func (ContributionRepository) Delete(startAt time.Time, endAt time.Time) {
	_, err := NingenmeMysql.NamedExec(`DELETE FROM github_contribution WHERE contributed_at BETWEEN :start_at AND :end_at`,
		map[string]interface{}{
		    "start_at": startAt.Format(time.RFC3339),
			"end_at": endAt.Format(time.RFC3339),
		})
	if err != nil {
		fmt.Println(err)
	}
}
