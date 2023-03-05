package main

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func login(w http.ResponseWriter, r *http.Request){
     t, err := template.ParseFiles("templates/login.html")
      if err != nil {
           fmt.Fprintf(w, err.Error())
      }
     t.ExecuteTemplate(w, "login", nil)
}
func register(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("templates/register.html")
       if err != nil {
            fmt.Fprintf(w, err.Error())
       }
      t.ExecuteTemplate(w, "register", nil)
 }
func index(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("templates/index.html")
       if err != nil {
            fmt.Fprintf(w, err.Error())
       }
      t.ExecuteTemplate(w, "index", nil)
 }
func save_reg(w http.ResponseWriter, r *http.Request){
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    password := r.FormValue("password")

    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")

    if err != nil {
        panic(err)
    }
    defer db.Close()

    insert, err := db.Query(fmt.Sprintf("Insert into `register` (`email`, `phone`, `password`) Values ('%s', '%s', '%s')", email, phone, password))
    if err != nil {
         panic(err)
    }
    defer insert.Close()

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
func save_log(w http.ResponseWriter, r *http.Request){
     t, err2 := template.ParseFiles("templates/login.html", "templates/index.html")
       if err2 != nil {
            fmt.Fprintf(w, err2.Error())
       }
    email := r.FormValue("email")
    password := r.FormValue("password")
    fmt.Println("email:", email, "password:", password)

    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    var hash string
    stmt := "SELECT password FROM register WHERE email = ?"
    row := db.QueryRow(stmt, email)
    erro := row.Scan(&hash)
    fmt.Println("hash:", hash)

    if erro != nil {
     fmt.Println("Ayaaau")
     t.ExecuteTemplate(w, "login", "check email and password")
     return
    }
    if hash == password {
     t.ExecuteTemplate(w, "index", "You have successfully logged in!")
     return
    }
    fmt.Println("incorrect password")
    t.ExecuteTemplate(w, "login", "check email and password")

}
func handleFunc (){
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
    http.HandleFunc("/save_reg", save_reg)
    http.HandleFunc("/save_log", save_log)
    http.ListenAndServe(":8080", nil)
}
func main() {
    handleFunc()
    fmt.Println("dvfjdf")
}