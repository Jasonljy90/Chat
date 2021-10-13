package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Language  string
}

// Insert new user account information into database
func insertRecord(userName, passWord, firstName, lastName, language string) {
	_, err := db.Exec("INSERT INTO MYSTOREDBFOODPANDA.Users VALUES (?,?,?,?,?)", userName, passWord, firstName, lastName, language)
	if err != nil {
		Error.Println("Error inserting record into database")
		fmt.Println(err)
	} else {
		fmt.Println("New user account information added to database successfully")
	}
}

// Delete user account information from database
func deleteRecord(userName string) int {
	results, err := db.Exec("DELETE FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		Error.Println("Error deleting record from database")
		fmt.Println(err)
		return 0
	} else {
		fmt.Println("User account information deleted from database successfully")
		rows, _ := results.RowsAffected()
		return int(rows)
	}
}

// Update user account password in database
func changePasswordRecord(userName string, password []byte) {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
	} else {
		for results.Next() {
			var person User
			err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User new password updated in database successfully")
				_, err := db.Exec("UPDATE MYSTOREDBFOODPANDA.Users set password = ?, firstname = ?, lastname = ?, language = ? where Username=?", password, &person.FirstName, &person.LastName, &person.Language, &person.UserName)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// Update user account password in database
func changeLanguageRecord(userName string, language string) {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
	} else {
		for results.Next() {
			var person User
			err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User new language preference updated in database successfully")
				_, err := db.Exec("UPDATE MYSTOREDBFOODPANDA.Users set language = ?, firstname = ?, lastname = ? where Username=?", language, &person.FirstName, &person.LastName, &person.UserName)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// Check if user exists in database
func checkUserExists(userName string) bool {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for results.Next() {
		var person User
		err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
		if err != nil {
			fmt.Println(err)
		} else {
			return true
		}
	}
	return false
}

// get the hashed password of the user in string type
func getPasswordOfUser(userName string) string {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for results.Next() {
		var person User
		err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
		if err != nil {
			fmt.Println(err)
			return ""
		} else {
			return person.Password
		}
	}
	return ""
}

// get the first name of user in string type
func getFirstNameOfUser(userName string) string {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for results.Next() {
		var person User
		err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
		if err != nil {
			fmt.Println(err)
			return ""
		} else {
			return person.FirstName
		}
	}
	return ""
}

// hash the given password using bcrypt()
func hashPassword(password string) []byte {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost); err != nil {
		fmt.Println(err)
		return nil
	} else {
		return hash
	}
}

//                   saved in the db        user supplied
func verifyPassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

// Open data base to insert new user account information
func userSignupDataBase(un, pw, fn, ln, language string) {
	// Get environment variables
	// getEnvVars()
	// mysqlAccount := os.Getenv("MySql_Account")

	// db, err := sql.Open("mysql", mysqlAccount)
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println("Database opened to add new user account information.")
	// }
	// defer db.Close()

	//----adding new user---
	userName := un
	password := pw
	firstName := fn
	lastName := ln
	insertRecord(userName, string(hashPassword(password)), firstName, lastName, language)
}

// Open data base to remove user account information
func userDeleteDataBase(userName string) int {
	// Get environment variables
	// getEnvVars()
	// mysqlAccount := os.Getenv("MySql_Account")

	// db, err := sql.Open("mysql", mysqlAccount)
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println("Database opened to delete user account.")
	// }
	// defer db.Close()

	//----deleting user---
	return deleteRecord(userName)
}

// Open data base to update user account password
func userChangePasswordDataBase(userName string, password string) {
	hashedPassword := hashPassword(password)
	// Get environment variables
	// getEnvVars()
	// mysqlAccount := os.Getenv("MySql_Account")

	// db, err := sql.Open("mysql", mysqlAccount)
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println("Database opened to change user password.")
	// }
	// defer db.Close()

	//----change password of user---
	changePasswordRecord(userName, hashedPassword)
}

// Checks if the entered username and password match an account in database. If yes login successful, else unsuccessful
func authenticatingUserFromDataBase(un string, pw string) bool {
	// Get environment variables
	// getEnvVars()
	// mysqlAccount := os.Getenv("MySql_Account")

	// db, err := sql.Open("mysql", mysqlAccount)
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println("Database opened to authenticate user from database.")
	// }
	// defer db.Close()

	//---authenticating user---
	userName := un
	password := pw

	// retrieve the user's saved password (in string); hashed
	userSavedPassword := getPasswordOfUser(userName)

	// check if the password saved in the db match the user's supplied password
	return verifyPassword([]byte(userSavedPassword), password)
}
