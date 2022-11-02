package pgsqlconn

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"sync"

	//Database PosgreSQL Driver Initialization
	_ "github.com/lib/pq"
	//https://pkg.go.dev/github.com/lib/pq
)

const (
	driverName = "postgres"
)

type DBConnector interface {
	OpenConn() *sql.DB
}

type dbConn struct {
	dbconfig *dbConnBuilder
	connOnce sync.Once
	conn     *sql.DB
}

// OpenConn opens and returns a singleton connection to the database
// previously indicated to builder.
func (db *dbConn) OpenConn() *sql.DB {

	db.connOnce.Do(func() {
		connString := db.getConnString()
		conn, err := sql.Open(driverName, connString)
		if err != nil {
			log.Panicln("Open connection to database failed: ", err.Error())
		}
		err = conn.Ping()
		if err != nil {
			log.Panicln("Connection to database failed: ", err.Error())
		}
		db.conn = conn
	})
	return db.conn
}

func (db *dbConn) getConnString() string {

	query := url.Values{}
	query.Add("sslmode", db.dbconfig.sslmode)

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(db.dbconfig.user, db.dbconfig.password),
		Host:   fmt.Sprintf("%s:%d", db.dbconfig.hostname, db.dbconfig.port),
		//Path:     db.dbconfig.instance,
		Path:     db.dbconfig.dbname,
		RawQuery: query.Encode(),
	}

	connString := u.String()

	if db.dbconfig.debug {
		fmt.Printf(" hostname:%s\n", db.dbconfig.hostname)
		fmt.Printf(" port:%d\n", db.dbconfig.port)
		//fmt.Printf(" instance:%s\n", db.dbconfig.instance)
		fmt.Printf(" user:%s\n", db.dbconfig.user)
		fmt.Printf(" password:%s\n", db.dbconfig.password)
		fmt.Printf(" dbname:%s\n", db.dbconfig.dbname)
		fmt.Printf(" connString:%s\n", connString)
	}
	return connString
}
