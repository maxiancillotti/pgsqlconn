package pgsqlconn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testdbc = &dbConn{
		dbconfig: &dbConnBuilder{
			hostname: "hostname",
			port:     1234,
			user:     "user",
			password: "password",
			dbname:   "dbname",
			sslmode:  "SSLMODE",
			debug:    true,
		},
	}
)

func TestGetConnString(t *testing.T) {

	expectedConnString := "postgres://user:password@hostname:1234/dbname?sslmode=SSLMODE"
	connString := testdbc.getConnString()

	assert.Equal(t, expectedConnString, connString)
}
