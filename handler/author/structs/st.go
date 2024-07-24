package structs

import (
	"time"
)

type Author struct {
	Author_id  int       `json:"authorid"`
	Name       string    `json:"name"`
	Birth_date string    `json:"birthdate"`
	Biography  string    `json:"biography"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
}
