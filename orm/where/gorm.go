package where

import (
	"github.com/jinzhu/gorm"
)

func ForGorm(wr *Where, db *gorm.DB) *gorm.DB {
	return db.Where("`"+wr.Name+"`"+wr.Con+"?", wr.Val)
}
