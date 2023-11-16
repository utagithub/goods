package migrate

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"goods/cmd/migrate/migration"
	_ "goods/cmd/migrate/migration/version-local"
	"goods/common/config"
	"goods/common/config/source/file"
	"goods/common/database"
	"goods/common/runtime"
	"goods/common/tools"
	"goods/models"
	"strconv"
	"text/template"
	"time"
)

var (
	configYml string
	generate  bool
	goods     bool
	host      string
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "goods migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

// fixme 在您看不见代码的时候运行迁移，我觉得是不安全的，所以编译后最好不要去执行迁移
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
	StartCmd.PersistentFlags().BoolVarP(&goods, "goods", "a", false, "generate goods migration file")
	StartCmd.PersistentFlags().StringVarP(&host, "domain", "d", "*", "select tenant host")
}

func run() {

	if !generate {
		fmt.Println(`start init`)
		//1. 读取配置
		config.Setup(
			file.NewSource(file.WithPath(configYml)),
			initDB,
		)
	} else {
		fmt.Println(`generate migration file`)
		_ = genFile()
	}
}

func initDB() {
	//3. 初始化数据库链接
	database.Setup()
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func migrateModel() error {
	if host == "" {
		host = "*"
	}
	db := runtime.App.GetDbByKey(host)
	if config.DatabasesConfig[host].Driver == "mysql" {
		//初始化数据库时候用
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	err := db.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"
	if goods {
		m["Package"] = "version"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if goods {
		tools.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
	} else {
		tools.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")
	}
	return nil
}
