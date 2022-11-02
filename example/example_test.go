// This is only a handy package to test an actual connection to a database.
// See the comments and change the data below.
package example

import (
	"testing"

	"github.com/maxiancillotti/pgsqlconn"
)

var conn = pgsqlconn.NewBuilder().
	// COMPLETE WITH VALID INPUTS
	SetHostname("localhost").
	SetDatabaseName("postgres").
	SetCredentials("postgres", "PASSWORD").
	EnableDebug().
	Build().
	OpenConn()

func TestActualConnection(t *testing.T) {
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM test_table") // COMPLETE WITH A VALID QUERY

	if err != nil {
		t.Fatal("table SELECT failed:", err.Error())
	}

	defer rows.Close()
	for rows.Next() {

		// CHANGE THE VARIABLES TO SCAN ACTUAL DATABASE OUTPUT
		var id int
		var description string

		err = rows.Scan(&id, &description) // ALSO HERE
		if err != nil {
			t.Fatal("SELECT result rows scan failed:", err.Error())
		}

		t.Log(id, description) // AND HERE
	}
	err = rows.Err()
	if err != nil {
		t.Fatal("Rows scan returned an error:", err.Error())
	}
}
