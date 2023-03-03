package models

import (
	"github.com/EleisonC/vending-machine/config"
	// "database/sql"
	_ "github.com/mattn/go-sqlite3"
	// "log"
	"golang.org/x/crypto/bcrypt"
)



func CreateNewUser(newUser *UserModel) error {
	config.ConnectDB()
	stmt, err := config.DB.Prepare("INSERT INTO usertable (username, password, deposit, role) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	passByte := []byte(newUser.Password)

	hashPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser.Password = string(hashPassword)

	_, err = stmt.Exec(newUser.Username, newUser.Password, newUser.Deposit, newUser.Role)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(userUpdate *UserModel, userId string) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE users SET username=?, role=? WHERE id=?")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(userUpdate.Username, userUpdate.Role, userId)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

	return count, nil
}

func FindUserByUsername(username string, user *UserModeldb) (error) {
	config.ConnectDB()
	err := config.DB.QueryRow("SELECT * FROM usertable WHERE username=?", username).Scan(&user.Id, &user.Username, &user.Password, &user.Deposit, &user.Role)
	if err != nil {
        return err
    }
	return nil
}

func FindUserById(userId string) error {
	var user UserModeldb
	err := config.DB.QueryRow("SELECT * FROM users WHERE id=?", userId).Scan(&user)
	if err != nil {
        return err
    }
	return nil
}


func DeleteUser(userId string) (int64, error) {
	queryStmt := `
		DELETE FROM users WHERE id=?
	`

	// if err := FindUserByUsername(userId); err != nil {
	// 	return 0, err
	// }

	stmt, err := config.DB.Prepare(queryStmt)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(userId)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}