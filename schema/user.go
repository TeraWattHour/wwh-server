package schema

import "time"

type User struct {
	ID            string
	Name          string
	Discriminator uint
	Email         string
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
