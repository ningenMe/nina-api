package domainmodel

import "time"

type Contribution struct {
	ContributedAt time.Time `db:"contributed_at"`
	Organization  string    `db:"organization"`
	Repository    string    `db:"repository"`
	User          string    `db:"user"`
	Status        string    `db:"status"`
}

