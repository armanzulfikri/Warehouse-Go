package response

import "time"

//UserResponse ...
type UserResponse struct {
	ID         uint
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	BirthPlace string `json:"birth_place"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
