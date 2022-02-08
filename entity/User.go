package entity

type User struct {
	Name        string `json:"name"`
	Age         string `json:"age"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"-"`
}
