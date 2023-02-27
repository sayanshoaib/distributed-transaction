package apimodels

type ReqProductCreate struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"price"`
}
