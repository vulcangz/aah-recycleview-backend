package repository

import (
	"database/sql"
)

var Repo Repository

// // GormRepository 基于GORM的存储库封装
// type GormRepository struct {
// 	db *gorm.DB
// }

// // NewGormRepository 初始化并生成存储库的实例
// func NewGormRepository(db *gorm.DB) (Repository, error) {
// 	Repo := &GormRepository{
// 		db: db,
// 	}
// 	return Repo, nil
// }

// // get repository instance
// func GetRepo() Repository {
// 	return Repo
// }

// SormRepository 基于SQLBoiler的存储库封装
type SormRepository struct {
	db *sql.DB
}

// NewSormRepository 初始化并生成存储库的实例
func NewSormRepository(db *sql.DB) (Repository, error) {
	Repo := &SormRepository{
		db: db,
	}
	return Repo, nil
}
