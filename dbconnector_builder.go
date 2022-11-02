package pgsqlconn

type DBConnectorBuilder interface {
	SetHostname(hostname string) DBConnectorBuilder

	SetPort(port int) DBConnectorBuilder

	//SetInstance(instance string) DBConnectorBuilder

	SetDatabaseName(dbname string) DBConnectorBuilder

	SetCredentials(user string, password string) DBConnectorBuilder

	/*
		Valid values for sslmode are:

		disable - No SSL (DEFAULT when no value is set or when an invalid value is passed)
		require - Always SSL (skip verification)
		verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
		verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
	*/
	SetSSLMode(mode string) DBConnectorBuilder

	EnableDebug() DBConnectorBuilder

	Build() DBConnector
}

type dbConnBuilder struct {
	hostname string
	port     int
	//instance string
	user     string
	password string
	dbname   string
	sslmode  string
	debug    bool
}

// NewBuiler returns a DBConnecterBuilder that you can configure to
// build a database connector that can open a database connection.
func NewBuilder() DBConnectorBuilder {
	return &dbConnBuilder{
		port:    5432,
		sslmode: "disable",
		debug:   false,
	}
}

func (b *dbConnBuilder) Build() DBConnector {
	return &dbConn{
		dbconfig: b,
	}
}

func (b *dbConnBuilder) SetHostname(hostname string) DBConnectorBuilder {
	b.hostname = hostname
	return b
}

func (b *dbConnBuilder) SetPort(port int) DBConnectorBuilder {
	b.port = port
	return b
}

/*
func (b *dbConnBuilder) SetInstance(instance string) DBConnectorBuilder {
	b.instance = instance
	return b
}
*/
func (b *dbConnBuilder) SetDatabaseName(dbname string) DBConnectorBuilder {
	b.dbname = dbname
	return b
}

func (b *dbConnBuilder) SetCredentials(user, password string) DBConnectorBuilder {
	b.user = user
	b.password = password
	return b
}

func (b *dbConnBuilder) SetSSLMode(mode string) DBConnectorBuilder {

	switch mode {
	case "require", "verify-ca", "verify-full":
		b.sslmode = mode
	default:
		b.sslmode = "disable"
	}
	return b
}

func (b *dbConnBuilder) EnableDebug() DBConnectorBuilder {
	b.debug = true
	return b
}
