package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"Transaction/protobuf"
	"Transaction/protobuf/stockpb"

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
	// 1. Llamar al microservicio de stock para obtener el balance
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Printf("❌ No se pudo conectar al microservicio de stock: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al conectar con el microservicio de stock",
		}, nil
	}
	defer conn.Close()

	client := stockpb.NewStockTradingServiceClient(conn)

	stockReq := &stockpb.StockRequest{
		StockSymbol: req.From,
	}

	stockCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stockRes, err := client.GetStockPrice(stockCtx, stockReq)
	if err != nil {
		log.Printf("❌ Error al obtener precio del stock: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al obtener información del stock",
		}, nil
	}

	log.Printf("📈 Precio de %s: %.2f (timestamp: %s)", stockRes.StockSymbol, stockRes.Price, stockRes.Timestamp)

	if stockRes.Price < float64(req.Amount.Amount) || req.Amount.Amount == 0 {
		log.Printf("❌ Saldo insuficiente: %.2f < %d", stockRes.Price, req.Amount.Amount)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Saldo insuficiente",
		}, nil
	}

	// 2. Insertar en la base de datos
	_, err = s.db.Exec(`
		INSERT INTO transactions (from_user, to_user, amount, currency)
		VALUES ($1, $2, $3, $4)
	`, req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	if err != nil {
		log.Println("❌ Error al insertar en la base de datos:", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al guardar la transacción",
		}, nil
	}

	// 3. Descontar valor al origen (From)
	newPrice := stockRes.Price - float64(req.Amount.Amount)
	updateReq := &stockpb.UpdateStockRequest{
		StockSymbol: req.From,
		NewPrice:    newPrice,
	}
	updateRes, err := client.UpdateStockPrice(ctx, updateReq)
	if err != nil || !updateRes.Success {
		log.Printf("❌ Error al actualizar el precio del stock origen: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Transacción guardada pero error al actualizar saldo del emisor",
		}, nil
	}

	log.Printf("✅ Precio actualizado del emisor: %s", updateRes.Message)

	// 4. Sumar valor al receptor (To)
	stockToReq := &stockpb.StockRequest{StockSymbol: req.To}
	stockToRes, err := client.GetStockPrice(ctx, stockToReq)
	if err != nil {
		log.Printf("❌ Error al obtener precio del stock destino: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Error al obtener información del receptor",
		}, nil
	}

	newToPrice := stockToRes.Price + float64(req.Amount.Amount)
	updateToReq := &stockpb.UpdateStockRequest{
		StockSymbol: req.To,
		NewPrice:    newToPrice,
	}
	updateToRes, err := client.UpdateStockPrice(ctx, updateToReq)
	if err != nil || !updateToRes.Success {
		log.Printf("❌ Error al actualizar saldo del receptor: %v", err)
		return &protobuf.CreateResponse{
			Success: false,
			Message: "Transacción guardada pero error al acreditar saldo al receptor",
		}, nil
	}

	log.Printf("✅ Saldo del receptor actualizado: %s", updateToRes.Message)

	// 5. Devolver respuesta final
	return &protobuf.CreateResponse{
		Success: true,
		Message: fmt.Sprintf("Transacción realizada correctamente. Saldo restante: %.2f", newPrice),
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
