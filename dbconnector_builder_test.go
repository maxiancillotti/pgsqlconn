package pgsqlconn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
var (
	testdbcBuilder = &dbConnBuilder{
		hostname: "hostname",
		port:     1234,
		user:     "user",
		password: "password",
		dbname:   "dbname",
		sslmode:  "SSLMODE",
		debug:    true,
	}
)
*/
func TestNewBuilder(t *testing.T) {

	Builder := NewBuilder()
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	// Testing Defaults
	assert.Equal(t, 5432, builder.port)
	assert.Equal(t, "disable", builder.sslmode)
	assert.Equal(t, false, builder.debug)
}

func TestSetHostname(t *testing.T) {

	expectedHostname := "hostname"

	Builder := NewBuilder().SetHostname(expectedHostname)
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	assert.Equal(t, expectedHostname, builder.hostname)
}

func TestSetPort(t *testing.T) {

	expectedPort := 1234

	Builder := NewBuilder().SetPort(expectedPort)
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	assert.Equal(t, expectedPort, builder.port)
}

func TestSetDatabaseName(t *testing.T) {

	expectedDBName := "dbName"

	Builder := NewBuilder().SetDatabaseName(expectedDBName)
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	assert.Equal(t, expectedDBName, builder.dbname)
}

func TestSetCredentials(t *testing.T) {

	expectedUser := "user"
	expectedPW := "pw"

	Builder := NewBuilder().SetCredentials(expectedUser, expectedPW)
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	assert.Equal(t, expectedUser, builder.user)
	assert.Equal(t, expectedPW, builder.password)
}

func TestSetSSLMode(t *testing.T) {

	expectedSSLModes := []string{"require", "verify-ca", "verify-full", "disable", "default"}

	for _, sslMode := range expectedSSLModes {

		t.Run(sslMode, func(t *testing.T) {
			Builder := NewBuilder().SetSSLMode(sslMode)
			builder, ok := Builder.(*dbConnBuilder)
			assert.True(t, ok)

			if sslMode != "default" {
				assert.Equal(t, sslMode, builder.sslmode)
			} else {
				assert.Equal(t, "disable", builder.sslmode)
			}
		})
	}
}

func TestEnableDebug(t *testing.T) {
	expectedDebugMode := true

	Builder := NewBuilder().EnableDebug()
	builder, ok := Builder.(*dbConnBuilder)
	assert.True(t, ok)

	assert.Equal(t, expectedDebugMode, builder.debug)
}
