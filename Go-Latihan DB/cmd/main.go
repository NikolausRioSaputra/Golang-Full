package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type dbConfig struct {
	Database string
}

type User struct {
	Id    int    `field:"id"`
	Name  string `field:"name"`
	Email string `field:"email"`
	Age   int    `field:age`
}

func NewConnection(database string) dbConfig {
	return dbConfig{
		Database: database,
	}
}

func GetDb() *sql.DB {
	driver := "pgx"
	connString := fmt.Sprintf("postgres://postgres:root@localhost:5432/GoLatihan_DB")

	DB, err := sql.Open(driver, connString)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully connected to the database")
	return DB
}

func InsertUser(db *sql.DB, user User) (isSuccess bool, err error) {
	_, err = db.Exec(`INSERT INTO GoLatihan_DB(name,age) values ($1,$2)`, user.Name, user.Age)

	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteUser(db *sql.DB, id int) error {
	res, err := db.Exec(`DELETE FROM users WHERE id= $1`, id)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("delete success", aff)
	return nil
}

func GetAll(db *sql.DB) (users []User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser() {
	
}

func main() {
	db := GetDb()
	defer db.Close()

	// Get all user
	users, err := GetAll(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, user := range users {
		fmt.Println(user)
	}

	// Insert user untuk masukan data
	var users1 User
	users1.Name = "Niko"
	users1.Email = "nikoblue@gmail.com"
	users1.Age = 20

	isSuccess, err := InsertUser(db, users1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if isSuccess {
			fmt.Println("Insert success")
		}
	}

	// Delete user
	err = DeleteUser(db, users1.Id)
	if err != nil {
		fmt.Println(err.Error())
	}

}
