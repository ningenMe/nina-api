package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ningenme/nina-api/pkg/domainmodel"
)

type BlogRepository struct{}

func (BlogRepository) GetList() []*domainmodel.Blog {
	rows, err := NingenmeMysql.Queryx(`SELECT url, date, type, title FROM blog WHERE type != 'DIARY' ORDER BY date DESC`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var list []*domainmodel.Blog
	for rows.Next() {
		c := &domainmodel.Blog{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	return list
}
