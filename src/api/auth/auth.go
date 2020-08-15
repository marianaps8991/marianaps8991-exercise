package auth

import (
	"api/database"
	"api/models"
	"api/security"
	"api/utils/channels"
)

func SignIn(username, password string) (string, error) {
	user := models.User{}
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db, err := database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(user.UserId, user.Username, user.Role)
	}
	return "", err
}
