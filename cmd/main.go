package main

import (
	"log"
	"os"

	"github.com/Nishad4140/proto_files/pb"
	"github.com/Nishad4140/user_service/db"
	"github.com/Nishad4140/user_service/initializer"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err.Error())
	}

	addr := os.Getenv("DATABASE_ADDR")

	DB, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	services := initializer.Initialize(DB)

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, services)
}
