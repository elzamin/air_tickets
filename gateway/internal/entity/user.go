package entity

type User struct {
	Id      string	`json:"id"`
	Name    string	`json:"name"`
	Age     int		`json:"age"`
	Address string	`json:"address"`
	Work    string	`json:"work"`
}
