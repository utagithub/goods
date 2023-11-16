/*
 * @Author: 尤_Ta
 * @Date:  14:23
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  14:23
 */

package models

import (
	"time"
)

type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

func (Migration) TableName() string {
	return "sys_migration"
}
