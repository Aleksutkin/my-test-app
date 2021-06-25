package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`
			CREATE TABLE users (
			  id	SERIAL	PRIMARY KEY,
			  name  VARCHAR(60) NOT NULL
			);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec("DROP TABLE users")
		return err
	})
}
