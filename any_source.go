package main

import "fmt"

func main() {
    var nameOfStore = "Nike store"
    fmt.Println("Welcome to our", nameOfStore, "!")
    fmt.Println("Register to continue: ")

    var userName string
    var userSurname string
    var userEmail string
    var userPassword string

    fmt.Println("Enter your name: ")
    fmt.Scan(&userName)
    fmt.Println("Enter your surname: ")
    fmt.Scan(&userSurname)
    fmt.Println("Enter your email: ")
    fmt.Scan(&userEmail)
    fmt.Println("Create a password: ")
    fmt.Scan(&userPassword)

    fmt.Printf("Hello, %v %v! Your account is created.\n", userName, userSurname)
}