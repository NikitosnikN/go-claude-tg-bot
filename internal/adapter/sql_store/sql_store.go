package sql_store

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type SQLStore struct {
	db *gorm.DB
}

func NewSQLStore(uri string) (*SQLStore, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open(uri), &gorm.Config{})

	//if strings.HasPrefix("sqlite", uri) {
	//	db, err = gorm.Open(sqlite.Open(uri), &gorm.Config{})
	//} else {
	//	err = errors.New("unsupported database type")
	//}

	if err != nil {
		return nil, err
	}

	return &SQLStore{
		db: db,
	}, nil
}

func (s *SQLStore) AutomigrateAll() (err error) {
	log.Println("Running automigrations...")

	err = s.db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}

	err = s.db.AutoMigrate(&dialog.Dialog{})
	if err != nil {
		return err
	}
	err = s.db.AutoMigrate(&message.Message{})

	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) DB() *gorm.DB {
	return s.db
}

func (s *SQLStore) NewTx() *gorm.DB {
	return s.db.Begin()
}
