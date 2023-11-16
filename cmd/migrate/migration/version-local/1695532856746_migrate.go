package version_local

import (
	"goods/cmd/migrate/migration"
	"goods/cmd/migrate/migration/models"
	"goods/common/config"
	common "goods/models"
	"gorm.io/gorm"
	"runtime"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1695532856746Test)
}

func _1695532856746Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		// TODO: 这里开始写入要变更的内容

		// TODO: 例如 修改表字段 使用过程中请删除此段代码
		//err := tx.Migrator().RenameColumn(&models.SysConfig{}, "config_id", "id")
		//if err != nil {
		// 	return err
		//}

		// TODO: 例如 新增表结构 使用过程中请删除此段代码
		err := tx.Debug().Migrator().AutoMigrate(
			new(models.SysApi),
			new(models.SysDept),
			new(models.SysGoods),
			new(models.SysJob),
			new(models.SysMenu),
			new(models.SysRole),
			new(models.SysUser),
		)
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
