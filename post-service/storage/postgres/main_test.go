package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/baxromumarov/my-services/post-service/config"
	"github.com/baxromumarov/my-services/post-service/pkg/db"
	"github.com/baxromumarov/my-services/post-service/pkg/logger"
)

var repo *postRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres", logger.Error(err))
	}
	repo = NewPostRepo(connDB)
	os.Exit(m.Run())
}
