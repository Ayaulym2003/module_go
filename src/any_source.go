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

func handleFunc (){
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
    http.HandleFunc("/save_reg", save_reg)
    http.ListenAndServe(":8080", nil)
}
func main() {
    handleFunc()
    fmt.Println("dvfjdf")
}