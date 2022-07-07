package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	once            sync.Once
	postgresConnStr string = os.Getenv("DATABASE_URL")
)

// postgres GORM postgresql driver implementation
type postgres struct {
	conn *gorm.DB
}

// newPostgresDB stablishes a new postgres connection
func (p *postgres) newPostgresDB() {
	once.Do(func() {
		var err error
		p.conn, err = gorm.Open("postgres", postgresConnStr)
		if err != nil {
			log.Fatalf("ERROR: can't open db: %v", err)
		}

		fmt.Println("Connected to PostgreSQL")
	})
}

// Pool get the connection to the postgres db
func (p *postgres) Pool() *gorm.DB {
	p.newPostgresDB()
	return p.conn
}

// closeDBConnection closes the conection to postgresql db
func (p *postgres) closeDBConnection() {
	p.conn.Close()
	log.Println("Connection to PostgreSQL closed")
}
