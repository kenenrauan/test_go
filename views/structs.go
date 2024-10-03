package views

type PostRequest struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}
