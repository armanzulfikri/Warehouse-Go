package response

// StatisticResponse model
type StatisticResponse struct {
	Warehouse   int `json:"warehouse"`
	Product     int `json:"product"`
	Supplier    int `json:"supplier"`
	Transaction int `json:"transaction"`
}
