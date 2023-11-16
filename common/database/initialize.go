package database

import (
  "goods/common/casbin"
  "goods/common/config"
  "goods/common/global"
  gormLogger "goods/common/gorm/logger"
  log "goods/common/logger"
  "goods/common/runtime"
  "goods/common/tools"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/gorm/schema"
  "time"
)

// Setup 配置数据库
func Setup() {
  for k := range config.DatabasesConfig {
    setupSimpleDatabase(k, config.DatabasesConfig[k])
  }
}

func setupSimpleDatabase(host string, c *config.Database) {
  if global.Driver == "" {
    global.Driver = c.Driver
  }
  log.Infof("%s => %s", host, tools.Green(c.Source))
  registers := make([]ResolverConfigure, len(c.Registers))
  for i := range c.Registers {
    registers[i] = NewResolverConfigure(
      c.Registers[i].Sources,
      c.Registers[i].Replicas,
      c.Registers[i].Policy,
      c.Registers[i].Tables)
  }
  resolverConfig := NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
  db, err := resolverConfig.Init(&gorm.Config{
    NamingStrategy: schema.NamingStrategy{
      SingularTable: true,
    },
    Logger: gormLogger.New(
      logger.Config{
        SlowThreshold: time.Second,
        Colorful:      true,
        LogLevel:      logger.LogLevel(log.DefaultLogger.Options().Level.LevelForGorm()),
      },
    ),
  }, opens[c.Driver])

  if err != nil {
    log.Fatal(tools.Red(c.Driver+" connect error :"), err)
  } else {
    log.Info(tools.Green(c.Driver + " connect success !"))
  }

  runtime.App.SetDb(host, db)

  e := casbin.Setup(db, "")
  runtime.App.SetCasbin(host, e)
}
