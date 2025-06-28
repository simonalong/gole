package test

import (
	"context"
	"database/sql"
	"fmt"
	driverMysql "github.com/go-sql-driver/mysql"
	"github.com/qustavo/sqlhooks/v2"
	"testing"
	"time"
)

func TestGormHook1(t *testing.T) {
	sql.Register("mysqlHook", sqlhooks.Wrap(&driverMysql.MySQLDriver{}, &Hooks{}))

	// Connect to the registered wrapped driver
	db, _ := sql.Open("mysqlHook", "test")

	// Do you're stuff
	db.Exec("CREATE TABLE t (id INTEGER, text VARCHAR(16))")
	db.Exec("INSERT into t (text) VALUES(?), (?)", "foo", "bar")
	db.Query("SELECT id, text FROM t")
}

// Hooks satisfies the sqlhook.Hooks interface
type Hooks struct{}

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	fmt.Printf("> %s %q", query, args)
	return context.WithValue(ctx, "begin", time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin := ctx.Value("begin").(time.Time)
	fmt.Printf(". took: %s\n", time.Since(begin))
	return ctx, nil
}
