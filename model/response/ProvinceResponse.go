package response

import "time"

//ProvinceResponse ...
type ProvinceResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"province_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
