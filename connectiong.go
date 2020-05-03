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

type defaultDialer struct {
	d net.Dialer
}

func (d defaultDialer) Dial(network, address string) (net.Conn, error) {
	return d.d.Dial(network, address)
}

func (d defaultDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}
func (d defaultDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return d.d.DialContext(ctx, network, address)
}

//Open - implementation of the Open method to implement the ability to work with database / sql
func (d *Driver) Open(name string) (driver.Conn, error) {
	return DialConnection(defaultDialer{}, name)
}

//DialConnection - opens a new connection to the database using a dialer
func DialConnection(d Dialer, dsn string) (driver.Conn, error) {
	return nil, nil
}

//register driver for database/sql
func init() {
	sql.Register("firebird", &Driver{})
}
