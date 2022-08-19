package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/baxromumarov/my-services/user-service/config"
	"github.com/baxromumarov/my-services/user-service/pkg/db"
	"github.com/baxromumarov/my-services/user-service/pkg/logger"
)

var repo *userRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres", logger.Error(err))
	}
	repo = NewUserRepo(connDB)
	os.Exit(m.Run())
}
