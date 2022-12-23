package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/baxromumarov/my-services/api-gateway/config"
	"github.com/baxromumarov/my-services/api-gateway/pkg/db"
)

var repo *userRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres", err)
	}
	repo = NewUserRepo(connDB)
	os.Exit(m.Run())
}
