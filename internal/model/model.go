package model

import (
	"fmt"
	"github.com/wow-unbelievable/blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID         uint32                `gorm:"primary_key" json:"id"`
	CreatedBy  string                `json:"created_by"`
	ModifiedBy string                `json:"modified_by"`
	CreatedOn  uint32                `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn uint32                `gorm:"autoUpdateTime" json:"modified_on"`
	DeletedOn  uint32                `json:"deleted_on"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedOn,DeletedAtFieldUnit" json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
