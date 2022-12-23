package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "database/sql"
)

type trainRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewTrainRepo(db *sqlx.DB) *trainRepo {
	return &trainRepo{db: db}
}

//create table users

func (t *trainRepo) GetJson() {
	fmt.Println("dir1")

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("error while getting directory", err)
	// 	return &models.JsonFile{}, nil

	// }
	fmt.Println("dir")
	// return models.JsonFile{}, nil

}
