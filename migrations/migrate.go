package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"

	"github.com/go-pg/pg/v10"
)

func Migrate(db *pg.DB, args ...string) error {
	var oldVersion, newVersion, err = migrations.Run(db, args...)
	if err != nil {
		fmt.Printf("migrations on db: %s\n", err)
		return err
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		if oldVersion == 0 {
			fmt.Println("No migrations have been applied")
		} else {
			fmt.Printf("%d migrations already applied\n", oldVersion)
		}
	}
	return nil
}
