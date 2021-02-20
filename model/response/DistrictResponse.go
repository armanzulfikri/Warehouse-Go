package response

import "time"

//DistrictResponse ...
type DistrictResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"district_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
