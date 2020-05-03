package firebird

import (
	"database/sql"
	"database/sql/driver"
)

//Driver is the Firebird database driver.
type Driver struct{}

//register driver for database/sql
func init() {
	sql.Register("firebird", &Driver{})
}

//Open - implementation of the Open method to implement the ability to work with database / sql
func (d *Driver) Open(name string) (driver.Conn, error) {
	return nil, nil
}
