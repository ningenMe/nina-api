package domainmodel

import "time"

type Contribution struct {
	ContributedAt time.Time `db:"contributed_at"`
	Organization  string    `db:"organization"`
	Repository    string    `db:"repository"`
	User          string    `db:"user"`
	Status        string    `db:"status"`
}

//TODO 消す
type ContributionSummaryKey struct {
	Date          string
	User          string
	Status        string
}
type ContributionSummary struct {
	Count         int
	ContributionSummaryKey
}
//TODO 消す

type ContributionSumKey struct {
	Date          string
	User          string
	Status        string
}
type ContributionSum struct {
	Count         int
	ContributionSumKey
}

