package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ningenme/nina-api/pkg/domainmodel"
)

type CategoryRepository struct{}

func (CategoryRepository) GetList() []*domainmodel.Category {
	rows, err := ComproMysql.Queryx(`SELECT category_id, category_display_name, category_system_name, category_order FROM category ORDER BY category_order ASC`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var list []*domainmodel.Category
	for rows.Next() {
		c := &domainmodel.Category{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	return list
}


func (CategoryRepository) Upsert(category *domainmodel.Category) {

	_, err := ComproMysql.NamedExec(`REPLACE INTO category (category_id, category_display_name, category_system_name, category_order) 
                                 VALUES (:category_id, :category_display_name, :category_system_name, :category_order)`, category)
	if err != nil {
		fmt.Println(err)
	}
}

func (CategoryRepository) Delete(categoryId string) {
	_, err := ComproMysql.NamedExec(`DELETE FROM category WHERE category_id = :categoryId`,
		map[string]interface{}{
		    "categoryId": categoryId,
		})
	if err != nil {
		fmt.Println(err)
	}
}
