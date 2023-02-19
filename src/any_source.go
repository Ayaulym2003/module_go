package main

import (
  "fmt"
  "net/http"
)

type Registration struct {
    name string
    surname string
    email string
    password string
    phoneNumber int
}

type Authorization struct {
    email string
    password string
}
func index_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hii everyone! We are Ayaulym, Zhibek, Diana and Zhanna")
}
func main() {
    http.HandleFunc("/", index_handler)
    http.ListenAndServe(":8000", nil)

    var nameOfStore = "Nike store"
    fmt.Println("Welcome to our", nameOfStore, "!")
    fmt.Println("Register to continue: ")

    var userName string
    var userSurname string
    var userEmail string
    var userPassword string
    var phoneNumber int

    fmt.Println("Enter your name: ")
    fmt.Scan(&userName)
    fmt.Println("Enter your surname: ")
    fmt.Scan(&userSurname)
    fmt.Println("Enter your email: ")
    fmt.Scan(&userEmail)
    fmt.Println("Create a password: ")
    fmt.Scan(&userPassword)
    fmt.Println("Enter your phone number:")
    fmt.Scan(&phoneNumber)

    user := Registration {userName, userSurname, userEmail, userPassword, phoneNumber}

    fmt.Printf("Hello, %v %v! Your account is created.\n", user.name, user.surname)
    fmt.Println("Now log in, please!")

    fmt.Println(user)


}