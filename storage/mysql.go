package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(mysqlURL string) (*MySQLStorage, error) {
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mysql: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	return &MySQLStorage{
		db: db,
	}, nil
}

func (s *MySQLStorage) ListCategory(ctx context.Context) ([]Category, error) {
	rows, err := s.db.QueryContext(ctx, "select name from category")
	if err != nil {
		return []Category{}, err
	}

	defer rows.Close()

	var categories []Category

	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Close(); err != nil {
		log.Fatal(err)
	}

	return categories, nil
}

func (s *MySQLStorage) WeaponParams(ctx context.Context, category string) ([]*Params, error) {

	return nil, nil
}

func (s *MySQLStorage) InsertWeapon(ctx context.Context, params *Params) error {

	return nil
}

func (s *MySQLStorage) UpdateWeapon(ctx context.Context, params *Params) error {

	return nil
}
