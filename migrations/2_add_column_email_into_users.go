package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`
			ALTER TABLE users ADD COLUMN email VARCHAR(100) DEFAULT NULL;
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`
			ALTER TABLE users DROP COLUMN email;
		`)
		return err
	})
}
