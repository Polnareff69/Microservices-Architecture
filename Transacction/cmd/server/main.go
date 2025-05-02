package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"Transaction/internal/config"
	"Transaction/internal/service"
	"Transaction/protobuf/transaction"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Cannot create listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	transactionService := service.NewTransactionService(db)
	invoicer.RegistertransactionServer(grpcServer, transactionService)

	log.Println("Servidor escuchando en puerto 9000...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
