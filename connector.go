package firebird

//Connector satisfies the database/sql/driver
type Connector struct {
	opts   *options
	dialer Dialer
}

type options struct {
	Host             string
	Port             string
	ExtraFloatDigits string
}

// NewConnector returns a connector for the firebird driver
func NewConnector(dsn string) (*Connector, error) {
	return nil, nil
}
