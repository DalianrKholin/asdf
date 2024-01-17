package serverResponseModel

type AddItemResponse struct {
	Name        string `json:"Item name"`
	Description string `json:"Item Description"`
	Amount      int    `json:"Amount of Item"`
}
