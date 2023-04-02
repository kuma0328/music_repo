package database

import "database/sql"

type Conn struct {
	DB *sql.DB
}