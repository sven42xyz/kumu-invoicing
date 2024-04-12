package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLDB represents the MySQL database connection
type MySQLDB struct {
    DB *sql.DB
}

// NewMySQLDB creates a new MySQL database connection
func NewMySQLDB(dataSourceName string) (*MySQLDB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }
    return &MySQLDB{DB: db}, nil
}

// Close closes the database connection
func (mdb *MySQLDB) Close() error {
    return mdb.DB.Close()
}

// User represents the user model
type User struct {
    ID       int
    Username string
    Password []byte
    Secret   string
}

// GetUserByUsername retrieves a user from the database by username
func (mdb *MySQLDB) GetUserByUsername(username string) (*User, error) {
    row := mdb.DB.QueryRow("SELECT user_id, username, password FROM kumu_auth.users WHERE username = ?", username)
    user := &User{}
    err := row.Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// CreateUser inserts a new user into the database
func (mdb *MySQLDB) CreateUser(username, password string) error {
    _, err := mdb.DB.Exec("INSERT INTO kumu_auth.users (username, password) VALUES (?, ?)", username, password)

    log.Println(err)

    return err
}
