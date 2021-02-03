package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kelvin-mai/go-anon-board/models"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					&models.Thread{},
					&models.Reply{},
				); err != nil {
					return err
				}
				if err := tx.Exec("ALTER TABLE threads ADD CONSTRAINT fk_ FOREIGN KEY (thread_id) REFERENCES threads (id)").Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					&models.Thread{},
					&models.Reply{},
				)
			},
		},
	})

	return m.Migrate()
}
