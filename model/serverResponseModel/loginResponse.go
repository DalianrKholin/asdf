package serverResponseModel

type Token struct {
	Token      string `json:"Auth Token"`
	AccessType string `json:"User Type"`
}
