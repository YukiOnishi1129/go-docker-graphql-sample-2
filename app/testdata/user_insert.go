package testdata

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type createUserData struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func CreateUserData(con *sql.DB) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	insertDataList := [...]*createUserData{
		{
			ID:       1,
			Name:     "太郎",
			Email:    "taro@gmail.com",
			Password: string(hashPassword),
		},
		{
			ID:       2,
			Name:     "次郎",
			Email:    "jiro@gmail.com",
			Password: string(hashPassword),
		},
		{
			ID:       3,
			Name:     "花子",
			Email:    "hanako@gmail.com",
			Password: string(hashPassword),
		},
	}

	for _, insertData := range insertDataList {
		ins, err = con.Prepare("INSERT INTO users (id, name, email, password) VALUES (?,?,?,?)")
		if err != nil {
			return err
		}
		_, err = ins.Exec(insertData.ID, insertData.Name, insertData.Email, insertData.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
