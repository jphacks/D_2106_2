package domain

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Introduction    string `json:"introduction"`
}
