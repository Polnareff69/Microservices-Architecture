package transaction_service

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/transaction_service/proto"
)

var TransactionClient proto.TransactionServiceClient

// Esta función se llama desde main.go para inicializar el cliente gRPC
func init() {
	// Establecer la conexión con el microservicio de Transaction
	conn, err := grpc.Dial(
		"localhost:9000", // Dirección del microservicio (ajusta el puerto si es necesario)
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar al microservicio gRPC: %v", err)
	}

	// Crear el cliente gRPC para el servicio TransactionService
	TransactionClient = proto.NewTransactionServiceClient(conn)
}
