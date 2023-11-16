package version

import (
	"goods/cmd/migrate/migration"
	"goods/common/config"
	common "goods/models"
	"runtime"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		err := tx.Migrator().AutoMigrate(
			new(common.SysDept),
		)
		if err != nil {
			return err
		}
		//if err := models.InitDb(tx); err != nil {
		//  return err
		//}
		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
