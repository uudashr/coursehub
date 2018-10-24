// +build integration

package mysql_test

import (
	"database/sql"
	"flag"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (
	scripts    = flag.String("scripts", "file://migrations", "The location of migration scripts.")
	dbUser     = flag.String("db-user", "coursehub", "Database username")
	dbPassword = flag.String("db-password", "coursehubsecret", "Database password")
	dbAddress  = flag.String("db-address", "localhost:3306", "Database address")
	dbName     = flag.String("db-name", "coursehub_test", "Database name")
)

const driverName = "mysql"

type dbFixture struct {
	t  *testing.T
	db *sql.DB
}

func (s *dbFixture) tearDown() {
	if err := s.db.Close(); err != nil {
		s.t.Error("failed closing db:", err)
	}
}

func setupDBFixture(t *testing.T) *dbFixture {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&clientFoundRows=true&parseTime=true&loc=Local", *dbUser, *dbPassword, *dbAddress, *dbName)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		t.Fatal("err:", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatal("err:", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		t.Fatal("err:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(*scripts, driverName, driver)
	if err != nil {
		t.Fatal("err:", err)
	}

	if err := m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			t.Error("failed execute migration down scripts", err)
		}
	}

	if err := m.Drop(); err != nil {
		t.Error("failed execute migration pre-drop:", err)
	}

	if err := m.Up(); err != nil {
		t.Error("failed execute migration up scripts:", err)
	}

	return &dbFixture{
		t:  t,
		db: db,
	}
}
