# Create main.go
cat <<EOF > main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Connect to MySQL
    dsn := "myuser:mypassword@tcp(mysql.default.svc.cluster.local:3306)/mydatabase"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query the database
    var result string
    err = db.QueryRow("SELECT 'Hello from MySQL!'").Scan(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Fprintf(w, result)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
EOF

# Build and test the app locally
go run main.go
