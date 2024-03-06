package repositories

import (
	"fmt"
	"time"

	storage "github.com/kakuzops/echo-crud/cmd/config"
	"github.com/kakuzops/echo-crud/cmd/models"
)

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")
	db := storage.GetDB()
	sqlStatement := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = id
	return user, nil
}
