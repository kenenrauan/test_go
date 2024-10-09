package views

type PostRequest struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}
