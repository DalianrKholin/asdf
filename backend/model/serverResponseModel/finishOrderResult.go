package serverResponseModel

type FinishOrderServerResult struct {
	UserUpdated string `json:"User Updated"`
	OrderStatus bool   `json:"finished"`
	OrderId     string `json:"Order Id"`
}
