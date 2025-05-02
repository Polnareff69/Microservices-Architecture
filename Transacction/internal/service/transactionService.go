package service

import (
	"context"
	"database/sql"
	"log"

	"Transaction/protobuf/invoicer"
)

type InvoiceService struct {
	invoicer.UnimplementedInvoicerServer
	db *sql.DB
}

func NewInvoiceService(db *sql.DB) *InvoiceService {
	return &InvoiceService{db: db}
}

func (s *InvoiceService) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	_, err := s.db.Exec(`
		INSERT INTO transactions (from_user, to_user, amount, currency)
		VALUES ($1, $2, $3, $4)
	`, req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	if err != nil {
		log.Println("❌ Error al insertar en la base de datos:", err)
		return &invoicer.CreateResponse{
			Success: false,
			Message: "Error al guardar la transacción",
		}, nil
	}

	return &invoicer.CreateResponse{
		Success: true,
		Message: "Transacción guardada y creada correctamente",
		Transaction: &invoicer.Transaction{
			From: req.From,
			To:   req.To,
			Amount: &invoicer.Amount{
				Amount:   req.Amount.Amount,
				Currency: req.Amount.Currency,
			},
		},
	}, nil
}
