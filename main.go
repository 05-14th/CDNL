package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "time"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "CodeDemonz24"
    password = "legendaryCodeGodz"
    dbname   = "cdnldb"
)

var db *sql.DB

func initDB() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    
    var err error
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Unable to ping database: %v", err)
    }

    log.Println("Successfully connected to database")
}

func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        email := r.FormValue("email")
        
        if email == "" {
            http.Error(w, "Email is required", http.StatusBadRequest)
            return
        }

        // Load the Philippine timezone
        loc, err := time.LoadLocation("Asia/Manila")
        if err != nil {
            http.Error(w, "Unable to load timezone", http.StatusInternalServerError)
            return
        }

        // Create the timestamp in Philippine time
        timestamp := time.Now().In(loc).Format("060102150405") // yymmddhhMMss format

        _, err = db.Exec("INSERT INTO emails (email, timestamp) VALUES ($1, $2)", email, timestamp)
        if err != nil {
            http.Error(w, "Unable to save email to database", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/news", http.StatusSeeOther)
    } else {
        tmpl := template.Must(template.ParseFiles("index.html"))
        tmpl.Execute(w, nil)
    }
}

func successHandler(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("news.html"))
    tmpl.Execute(w, nil)
}

func main() {
    initDB()
    defer db.Close()

    fs := http.FileServer(http.Dir("./static"))
	  http.Handle("/static/", http.StripPrefix("/static/", fs))
	  
	  stylefs := http.FileServer(http.Dir("./style"))
	  http.Handle("/style/", http.StripPrefix("/style/", stylefs))

    http.HandleFunc("/", submitEmailHandler)
    http.HandleFunc("/news", successHandler)

    fmt.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
