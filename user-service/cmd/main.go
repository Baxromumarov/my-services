package main

import (
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/baxromumarov/my-services/user-service/config"
	"github.com/baxromumarov/my-services/user-service/events"
	pb "github.com/baxromumarov/my-services/user-service/genproto"
	"github.com/baxromumarov/my-services/user-service/pkg/db"
	"github.com/baxromumarov/my-services/user-service/pkg/logger"
	"github.com/baxromumarov/my-services/user-service/pkg/messagebroker"
	"github.com/baxromumarov/my-services/user-service/service"
	grpcClient "github.com/baxromumarov/my-services/user-service/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	
	grpcC, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("grpc client error", logger.Error(err))	
		return
	}
	//Kafka
	publisherMap := make(map[string]messagebroker.Publisher)

	userTopicPublisher := events.NewKafkaProducerBroker(cfg,log,"user.user")
	defer func(){
		err:=userTopicPublisher.Stop()
		if err!=nil{
			log.Fatal("failed to stop kafka user",logger.Error(err))
		}
	}()

	publisherMap["user"] = userTopicPublisher
	//Kafka End


	userService := service.NewUserService(connDB, log, grpcC, publisherMap)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
