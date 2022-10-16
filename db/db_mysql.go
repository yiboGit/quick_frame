/**
* @Author: yibo_LastDay
* @Date: 2022/10/15 16:21
 */

package db

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"quick_frame/const_data"
)

var (
	MysqlDB *gorm.DB
	err     error
)

func InitMySql() error {
	if const_data.Mysql == "" {
		return errors.New("MySql DNS is nil")
	}
	MysqlDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       const_data.Mysql, // DSN data source name
		DefaultStringSize:         256,              // string 类型字段的默认长度
		DisableDatetimePrecision:  true,             // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,             // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,             // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,            // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return nil
}
