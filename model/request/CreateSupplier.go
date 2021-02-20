package request

//CreateSupplierRequest ...
type CreateSupplierRequest struct {
	SupplierName string `json:"supplier_name"`
	Address      string `json:"address"`
	Telepon      string `json:"telepon"`
}
