package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../cookie"
)

type User struct {
	UserId   int
	UserName string
	Mail     string
	Age      int
	Image    string
	Address  string
	Animal   string
	Word     string
}

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, userName, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Println("CheckSessionID", err)
			}
			mail, age, err := GetUserInfos(uid)
			if err != nil {
				fmt.Println("GetUserInfos", err)
			}

			userImage, userAddress, userAnimal, userWord, err := GetOptUserDetails(uid)
			if err != nil {
				fmt.Println("GetOptUserDetails", err)
			}

			u := User{UserName: userName, Mail: mail, Age: age, Image: userImage, Address: userAddress, Animal: userAnimal, Word: userWord}
			t, _ := template.ParseFiles("./views/users.gtpl")
			t.Execute(w, u)
		} else {
			http.NotFound(w, nil)
		}

	} else {
		t, _ := template.ParseFiles("./views/error.gtpl")
		t.Execute(w, nil)
	}
}

func GetOptUserDetails(uid int) (usersimage string, usersaddress string, usersanimal string, usersword string, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var userImage, address, animal, word string

	res, err := db.Query("select userImage,address,animal,word from vulnapp.userdetails where uid=?", uid)
	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		if err := res.Scan(&userImage, &address, &animal, &word); err != nil {
			fmt.Println(err)
		}
	}

	return userImage, address, animal, word, nil

}

func GetUserInfos(uid int) (userMail string, userAge int, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	defer db.Close()

	var mail string
	var age int

	res, err := db.Query("select mail,age from vulnapp.user where id=?", uid)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}

	for res.Next() {
		if err := res.Scan(&mail, &age); err != nil {
			fmt.Println(err)
			return "", 0, err
		}
	}

	return mail, age, nil
}

func GetUserMoreDetails(uid int) (userImage string, Address string, Animal string, Word string, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", err
	}
	defer db.Close()

	var getUserImage, getUserAddress, getUserAnimal, getUserWord string

	res, err := db.Query("select userimage,address,animal,word from vulnapp.userdetails where uid=?", uid)
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", err
	}

	for res.Next() {
		if err := res.Scan(&getUserImage, &getUserAddress, &getUserAnimal, &getUserWord); err != nil {
			fmt.Println(err)
			return "", "", "", "", err
		}
	}
	return getUserImage, getUserAddress, getUserAnimal, getUserWord, nil
}

func ShowUserModifyPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, userName, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Println("CheckSessionID", err)
			}
			mail, age, err := GetUserInfos(uid)
			if err != nil {
				fmt.Println("GetUserInfos", err)
			}
			userImage, userAddress, userAnimal, userWord, err := GetUserMoreDetails(uid)
			if err != nil {
				fmt.Println("GetUserMoreDetails", err)
			}

			u := User{UserName: userName, Mail: mail, Age: age, Image: userImage, Address: userAddress, Animal: userAnimal, Word: userWord}

			fmt.Println(u)
			t, _ := template.ParseFiles("./views/users_modify.gtpl")
			t.Execute(w, u)
		}

	} else {
		http.NotFound(w, nil)
	}
}

func ShowEditConfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			userName := r.FormValue("username")
			age := r.FormValue("age")
			mail := r.FormValue("mail")
			fmt.Println("mail : ", mail)
			address := r.FormValue("address")
			animal := r.FormValue("animal")
			word := r.FormValue("word")

			userAge, err := strconv.Atoi(age)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			u := User{UserName: userName, Mail: mail, Age: userAge, Address: address, Animal: animal, Word: word}

			t, _ := template.ParseFiles("./views/users_confirm.gtpl")
			t.Execute(w, u)
		}
	} else {
		http.NotFound(w, nil)
	}
}