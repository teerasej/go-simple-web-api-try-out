package main

type UserModel struct {
	UserId   int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Cart     []Product `json:"cart"`
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
