package main

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "strconv"
)
var store = sessions.NewCookieStore([]byte("super-secret"))
var db *sql.DB

type Products struct {
    Product_id int64
    Category_id int
    Product_name string
    Price int
    Size string
    Color string
    Rating int
    Description string
    Photo string
}
type Customer struct {
    Customer_id int64
    Name string
    Surname string
    Email_address string
    Phone_number string
    Password string
}

func login(w http.ResponseWriter, r *http.Request){
     t, err := template.ParseFiles("templates/login.html")
      if err != nil {
           fmt.Fprintf(w, err.Error())
      }
     t.ExecuteTemplate(w, "login", nil)
}
func logout(w http.ResponseWriter, r *http.Request){
     t, err := template.ParseFiles("templates/index.html")
      if err != nil {
           fmt.Fprintf(w, err.Error())
      }
     t.ExecuteTemplate(w, "index", nil)
}
func register(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("templates/register.html")
       if err != nil {
            fmt.Fprintf(w, err.Error())
       }
      t.ExecuteTemplate(w, "register", nil)
 }
func desc(w http.ResponseWriter, r *http.Request){
       vars := mux.Vars(r)
       y, e := strconv.Atoi(vars["Product_id"])
       if e != nil {
            fmt.Fprintf(w, "HHKJ")
       }

       t, err := template.ParseFiles("templates/desc.html", "templates/header.html")
               if err != nil {
                    fmt.Fprintf(w, err.Error())
               }
               db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
                      if err != nil {
                         panic(err)
                      }
                   defer db.Close()

               sel, err := db.Query(fmt.Sprintf("SELECT * FROM `product` WHERE `product_id` = %d", y))
               if err != nil {
                    fmt.Fprintf(w, err.Error())
                    return
               }
               defer sel.Close()

               var prods []Products
               for sel.Next() {
                  var P Products
                  err = sel.Scan(&P.Product_id, &P.Category_id, &P.Product_name, &P.Price, &P.Size, &P.Color, &P.Rating, &P.Description, &P.Photo)
                  if err != nil {
                       panic(err)
                  }
                  prods = append(prods, P)
               }

               t.ExecuteTemplate(w, "desc", prods)
  }


func index(w http.ResponseWriter, r *http.Request){
       t, err := template.ParseFiles("templates/index.html", "templates/header.html")
       if err != nil {
            fmt.Fprintf(w, err.Error())
       }
      t.ExecuteTemplate(w, "index", nil)
 }
func filter(w http.ResponseWriter, r *http.Request){
       t, err := template.ParseFiles("templates/filter.html",  "templates/header.html")
               if err != nil {
                    fmt.Fprintf(w, err.Error())
               }
               db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
                      if err != nil {
                         panic(err)
                      }
                   defer db.Close()

               sel, err := db.Query(fmt.Sprintf("SELECT * FROM `product`"))
               if err != nil {
                    fmt.Println("Ayaaau")
                    fmt.Fprintf(w, err.Error())
                    return
               }
               defer sel.Close()

               var prods []Products
               for sel.Next() {
                  var P Products
                  err = sel.Scan(&P.Product_id, &P.Category_id, &P.Product_name, &P.Price, &P.Size, &P.Color, &P.Rating, &P.Description, &P.Photo)
                  if err != nil {
                       panic(err)
                  }
                  prods = append(prods, P)
               }
               t.ExecuteTemplate(w, "filter", prods)
  }
 func products(w http.ResponseWriter, r *http.Request){
       t, err := template.ParseFiles("templates/products.html", "templates/header.html")
        if err != nil {
             fmt.Fprintf(w, err.Error())
        }
        db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
               if err != nil {
                  panic(err)
               }
            defer db.Close()

        sel, err := db.Query(fmt.Sprintf("SELECT * FROM `product`"))
        if err != nil {
             fmt.Println("Ayaaau")
             fmt.Fprintf(w, err.Error())
             return
        }
        defer sel.Close()

        var prods []Products
        for sel.Next() {
           var P Products
           err = sel.Scan(&P.Product_id, &P.Category_id, &P.Product_name, &P.Price, &P.Size, &P.Color, &P.Rating, &P.Description, &P.Photo)
           if err != nil {
                panic(err)
           }
           prods = append(prods, P)
        }

        t.ExecuteTemplate(w, "products", prods)
  }

func save_reg(w http.ResponseWriter, r *http.Request){
    t, err2 := template.ParseFiles("templates/register.html",  "templates/header.html")
       if err2 != nil {
            fmt.Fprintf(w, err2.Error())
       }
    name := r.FormValue("name")
    surname := r.FormValue("surname")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    password := r.FormValue("password")

    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
           if err != nil {
              panic(err)
           }
        defer db.Close()

    insert, err := db.Query(fmt.Sprintf("Insert into `customer` (`name`, `surname`,`email_address`, `phone_number`, `password`) Values ('%s', '%s', '%s', '%s', '%s')", name, surname, email, phone, password))
    if err != nil {
         fmt.Println("Ayaaau")
         t.ExecuteTemplate(w, "register", "something is not right")
         return
    }
    defer insert.Close()

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
func save_log(w http.ResponseWriter, r *http.Request){
     t, err2 := template.ParseFiles("templates/login.html", "templates/index.html", "templates/header.html")
       if err2 != nil {
            fmt.Fprintf(w, err2.Error())
       }
    Email := r.FormValue("email")
    Password := r.FormValue("password")
    fmt.Println("email:", Email, "password:", Password)

    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
           if err != nil {
              panic(err)
           }
        defer db.Close()

    var hash string
    stmt := "SELECT password FROM customer WHERE email_address = ?"
    row := db.QueryRow(stmt, Email)
    erro := row.Scan(&hash)
    fmt.Println("hash:", hash)

    if erro != nil {
     fmt.Println("Ayaaau")
     t.ExecuteTemplate(w, "login", "check email and password")
     return
    }

    if hash == Password {

     session, _ := store.Get(r, "session")
     session.Values["customer_id"] = true
     session.Save(r, w)
     t.Lookup("header")
     t.ExecuteTemplate(w, "index", Email)
     return
    }

    fmt.Println("incorrect password")
    t.ExecuteTemplate(w, "login", "check email and password")
}

func search(w http.ResponseWriter, r *http.Request){
        t, err := template.ParseFiles("templates/products.html",  "templates/header.html")
        if err != nil {
             fmt.Fprintf(w, err.Error())
        }
        name := r.FormValue("search")

        db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
               if err != nil {
                  panic(err)
               }
            defer db.Close()

        rows, err := db.Query("SELECT * FROM product WHERE product_name LIKE ?", "%" + name + "%")
        if err != nil {
             fmt.Println("Ayaaau")
             fmt.Fprintf(w, err.Error())
             return
        }
        defer rows.Close()

        var prod []Products
        for rows.Next() {
           var P Products
           err = rows.Scan(&P.Product_id, &P.Category_id, &P.Product_name, &P.Price, &P.Size, &P.Color, &P.Rating, &P.Description, &P.Photo)
           if err != nil {
                panic(err)
           }
           prod = append(prod, P)
        }

        t.ExecuteTemplate(w, "products", prod)
  }

 func filter_filter(w http.ResponseWriter, r *http.Request){
         t, err := template.ParseFiles("templates/filter.html",  "templates/header.html")
         if err != nil {
              fmt.Fprintf(w, err.Error())
         }
         minval := r.FormValue("minval")
         maxval := r.FormValue("maxval")
         rating := r.FormValue("rating")

         db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/golang")
                if err != nil {
                   panic(err)
                }
             defer db.Close()


         rows, err := db.Query("SELECT * FROM product WHERE price >= ? && price <= ? && rating >= ?;", minval, maxval, rating)

         if err != nil {

              fmt.Fprintf(w, err.Error())
              return
         }
         defer rows.Close()

         var prod []Products
         for rows.Next() {
            var P Products
            err = rows.Scan(&P.Product_id, &P.Category_id, &P.Product_name, &P.Price, &P.Size, &P.Color, &P.Rating, &P.Description, &P.Photo)
            if err != nil {
                 panic(err)
            }
            prod = append(prod, P)
         }
         t.ExecuteTemplate(w, "filter", prod)
   }



func handleFunc (){
    r := mux.NewRouter()
    r.HandleFunc("/", index)
    r.HandleFunc("/login", login)
    r.HandleFunc("/register", register)
    r.HandleFunc("/products", products)
    r.HandleFunc("/save_reg", save_reg)
    r.HandleFunc("/logout", logout)
    r.HandleFunc("/filter", filter)
    r.HandleFunc("/desc/{Product_id}", desc)
    r.HandleFunc("/filter_filter", filter_filter)
    r.HandleFunc("/save_log", save_log)
    r.HandleFunc("/search", search)

    fileServer := http.FileServer(http.Dir("./static/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
    http.ListenAndServe(":8080", r)
}
func main() {
    handleFunc()
    fmt.Println("dvfjdf")
}
