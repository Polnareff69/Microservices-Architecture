package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	protobuf "Transaction/protobuf/transaction"
	accountpb "Transaction/protobuf/account/protobuf"

	"google.golang.org/grpc"
)

type TransactionService struct {
	protobuf.UnimplementedTransactionServiceServer
	db *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) Create(ctx context.Context, req *protobuf.CreateRequest) (*protobuf.CreateResponse, error) {
	// 1. Conectar al microservicio de Account
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Printf("❌ No se pudo conectar al microservicio de Account: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al conectar con el microservicio de Account",
		}, nil
	}
	defer conn.Close()

	client := accountpb.NewAccountClient(conn)

	// 2. Obtener balance del remitente (From)
	accReq := &accountpb.AccountRequest{
		IdAccount: req.From,
	}

	accCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accRes, err := client.GetAccountBalance(accCtx, accReq)
	if err != nil {
		log.Printf("❌ Error al obtener balance del emisor: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al obtener información del emisor",
		}, nil
	}

	log.Printf("📘 Balance del emisor %s: %d créditos", accRes.IdAccount, accRes.Credits)

	if accRes.Credits < req.Amount.Amount || req.Amount.Amount <= 0 {
		log.Printf("❌ Créditos insuficientes o cantidad inválida")
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Créditos insuficientes o cantidad inválida",
		}, nil
	}

	// 3. Guardar transacción en la base de datos
	_, err = s.db.Exec(`
		INSERT INTO transactions (from_acc, to_acc, amount, currency)
		VALUES ($1, $2, $3, $4)
	`, req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	if err != nil {
		log.Println("❌ Error al insertar en la base de datos:", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al guardar la transacción",
		}, nil
	}

	// 4. Descontar créditos al emisor
	newCreditsFrom := accRes.Credits - req.Amount.Amount
	updateReq := &accountpb.UpdateAccountRequest{
		IdAccount: req.From,
		Credits:   newCreditsFrom,
	}

	updateRes, err := client.UpdateAccountBalance(ctx, updateReq)
	if err != nil || !updateRes.Success {
		log.Printf("❌ Error al actualizar créditos del emisor: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Transacción guardada pero error al debitar al emisor",
		}, nil
	}

	log.Printf("✅ Créditos actualizados del emisor: %s", updateRes.Message)

	// 5. Obtener balance actual del receptor (To)
	toAccRes, err := client.GetAccountBalance(ctx, &accountpb.AccountRequest{
		IdAccount: req.To,
	})
	if err != nil {
		log.Printf("❌ Error al obtener balance del receptor: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al obtener información del receptor",
		}, nil
	}

	// 6. Acreditar al receptor
	newCreditsTo := toAccRes.Credits + req.Amount.Amount
	updateToReq := &accountpb.UpdateAccountRequest{
		IdAccount: req.To,
		Credits:   newCreditsTo,
	}

	updateToRes, err := client.UpdateAccountBalance(ctx, updateToReq)
	if err != nil || !updateToRes.Success {
		log.Printf("❌ Error al acreditar al receptor: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Transacción guardada pero error al acreditar al receptor",
		}, nil
	}

	log.Printf("✅ Créditos del receptor actualizados: %s", updateToRes.Message)

	// 7. Respuesta exitosa
	return &protobuf.CreateResponse{
		Success: true,
		Message: fmt.Sprintf("Transacción realizada exitosamente. Créditos restantes: %d", newCreditsFrom),
		Transaction: &protobuf.Transaction{
			From: req.From,
			To:   req.To,
			Amount: &protobuf.Amount{
				Amount:   req.Amount.Amount,
				Currency: req.Amount.Currency,
			},
		},
	}, nil
}
