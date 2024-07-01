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
    tableName = "emails"
)

func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
    var db *sql.DB
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

func countRows(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(DISTINCT email) FROM " + tableName).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close() // Close the connection after use

	rowCount, err := countRows(db)
	if err != nil {
		log.Println("Error counting rows:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		RowCount int
	}{
		RowCount: rowCount,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
    fs := http.FileServer(http.Dir("./static"))
	  http.Handle("/static/", http.StripPrefix("/static/", fs))
	  
	  stylefs := http.FileServer(http.Dir("./style"))
	  http.Handle("/style/", http.StripPrefix("/style/", stylefs))

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/submit", submitEmailHandler)
    http.HandleFunc("/news", successHandler)

    fmt.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
