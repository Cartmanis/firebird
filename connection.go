package firebird

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net"
	"time"
)

//Driver is the Firebird database driver.
type Driver struct{}

// Dialer is the dialer interface. It can be used to obtain more control over
// how pq creates network connections.
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

//DialerContext dial network connection from use context
type DialerContext interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

type defaultDialer struct {
	d net.Dialer
}

//Dial implementation for interface Dialer
func (d defaultDialer) Dial(network, address string) (net.Conn, error) {
	return d.d.Dial(network, address)
}

//DialTimeout implementation for interface Dialer
func (d defaultDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}

//DialContext implementation for interface DialerContext
func (d defaultDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return d.d.DialContext(ctx, network, address)
}

//Open - implementation of the Open method to implement the ability to work with database / sql
func (d *Driver) Open(name string) (driver.Conn, error) {
	return DialConnection(defaultDialer{}, name)
}

//DialConnection - opens a new connection to the database using a dialer
func DialConnection(d Dialer, dsn string) (driver.Conn, error) {
	conn, err := NewConnector(dsn)
	if err != nil {
		return nil, err
	}
	conn.dialer = d
	return conn.open(context.Background())
}

//conn - object connection from Db. Need implementation Prepare, Begin, Close
type conn struct {
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, nil
}

func (c *Connector) open(ctx context.Context) (*conn, error) {
	return nil, nil
}

//register driver for database/sql
func init() {
	sql.Register("firebird", &Driver{})
}
