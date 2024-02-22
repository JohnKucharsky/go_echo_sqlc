package serializer

import (
	"github.com/JohnKucharsky/go_echo_sqlc/internal/database"
	"github.com/JohnKucharsky/go_echo_sqlc/utils"
	"time"
)

type UserBody struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name"`
}

type User struct {
	ID        int32     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserBody
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		UpdatedAt: dbUser.UpdatedAt,
		UserBody: UserBody{
			FirstName: dbUser.FirstName,
			LastName:  utils.SqlNullStringToString(dbUser.LastName),
		},
	}
}

func UserBodyToUserCreate(body UserBody) database.CreateUserParams {
	return database.CreateUserParams{
		FirstName: body.FirstName,
		LastName:  utils.StringToSqlNullString(body.LastName),
		UpdatedAt: time.Now().Local(),
	}
}

func UserBodyToUserUpdate(body UserBody, id int32) database.UpdateUserParams {
	return database.UpdateUserParams{
		FirstName: body.FirstName,
		LastName:  utils.StringToSqlNullString(body.LastName),
		UpdatedAt: time.Now().Local(),
		ID:        id,
	}
}

func DatabaseUsersToUsers(dbUsers []database.User) []User {
	var users []User

	for _, user := range dbUsers {
		users = append(users, DatabaseUserToUser(user))
	}

	return users
}
