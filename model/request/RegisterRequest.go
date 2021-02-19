package request

//RegisterRequest ...
type RegisterRequest struct {
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	BirthPlace string `json:"birth_place"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
