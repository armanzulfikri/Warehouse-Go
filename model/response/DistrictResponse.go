package response

import "time"

//DistrictResponse ...
type DistrictResponse struct {
	ID        uint
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
