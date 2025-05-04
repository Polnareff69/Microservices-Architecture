package account_service


import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/account_service/proto"
)

var AccountClient proto.AccountClient

// Esta función se llama desde main.go para inicializar el cliente gRPC
func init() {
	// Establecer la conexión con el microservicio de Account
	conn, err := grpc.Dial(
		"localhost:9090", // Dirección del microservicio
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar al microservicio gRPC: %v", err)
	}

	// Crear el cliente gRPC para el servicio Account
	AccountClient = proto.NewAccountClient(conn)
}
