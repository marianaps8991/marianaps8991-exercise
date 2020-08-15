package auto

import (
	"log"

	"api/database"
	"api/models"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}, &models.Family{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.Family{}, &models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.User{}).AddForeignKey("user_id", "families(family_id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {

		err = db.Debug().Model(&models.Family{}).Create(&familys[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		users[i].FamilyID = familys[i].FamilyId

		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}

	}
}
