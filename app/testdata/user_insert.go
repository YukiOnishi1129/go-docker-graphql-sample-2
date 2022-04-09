package testdata

import (
	"database/sql"
)

type createUserData struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func CreateUserData(con *sql.DB) error {
	insertDataList := [...]*createUserData{
		{
			ID:       1,
			Name:     "太郎",
			Email:    "taro@gmail.com",
			Password: "$2a$10$45R3cE3b/S2q7HMG/2gau.8L2y3cCiQ0lp48.YdeCXOXuxqxMWdLS",
		},
		{
			ID:       2,
			Name:     "次郎",
			Email:    "jiro@gmail.com",
			Password: "$2a$10$45R3cE3b/S2q7HMG/2gau.8L2y3cCiQ0lp48.YdeCXOXuxqxMWdLS",
		},
		{
			ID:       3,
			Name:     "花子",
			Email:    "hanako@gmail.com",
			Password: "$2a$10$45R3cE3b/S2q7HMG/2gau.8L2y3cCiQ0lp48.YdeCXOXuxqxMWdLS",
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
