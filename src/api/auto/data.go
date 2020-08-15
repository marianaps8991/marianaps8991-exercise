package auto

import "api/models"

var users = []models.User{
	models.User{Name: "Mariana", Age: 22, Username: "shiny", Password: "shiny", Family: familys[0], Role: "admin"},
}

var familys = []models.Family{
	models.Family{Name: "Santos", MaxPerson: 2, CurrentSize: 1},
}
