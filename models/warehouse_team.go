package models

type WarehouseTeam struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
}

type EmployeeReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateEmployeeRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginEmployeeRes struct {
	AccessToken string
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}
