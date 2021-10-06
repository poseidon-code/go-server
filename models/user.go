package models

type User struct {
    ID          int             `json:"id"`
    Name        UserName        `json:"name"`
    Age         int             `json:"age"`
}

type UserName struct {
    FirstName   string          `json:"firstname"`
    LastName    string          `json:"lastname"`
}