package domainmodel

type Blog struct {
	Url   string `db:"url"`
	Date  string `db:"date"`
	Type  string `db:"type"`
	Title string `db:"title"`
}
