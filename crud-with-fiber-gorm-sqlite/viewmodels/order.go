package viewmodels

import "time"

type Order struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Product   Product   `json:"product"`
	User      User      `json:"user"`
}
