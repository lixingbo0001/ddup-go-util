package where

import (
	"github.com/jinzhu/gorm"
)

func ForGorm(wr *Where, db *gorm.DB) *gorm.DB {
	
	switch wr.Name {
	case "like":
		return db.Where("`"+wr.Name+"` like '%?%'", wr.Val)
	}
	
	return db.Where("`"+wr.Name+"`"+wr.Con+"?", wr.Val)
}
