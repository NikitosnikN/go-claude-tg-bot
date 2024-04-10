package db

import "gorm.io/gorm"

type NewDBTx func() *gorm.DB
