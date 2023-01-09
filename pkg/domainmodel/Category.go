package domainmodel

import "crypto/rand"

const (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	chars_size = len(chars)
)


type Category struct {
	CategoryId string `db:"category_id"`
	CategoryDisplayName string `db:"category_display_name"`
	CategorySystemName string `db:"category_system_name"`
	CategoryOrder int32 `db:"category_order"`
}

func GetNewCategoryId() string {
	return "category_" + getRandomString(6);
}

//https://qiita.com/nakaryooo/items/7d269525a288c4b3ecda
func getRandomString(digit uint32) string {

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	var result string
	for _, v := range b {
		result += string(chars[int(v)%chars_size])
	}
	return result
}