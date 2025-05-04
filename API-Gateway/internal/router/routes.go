package router

import (
	"github.com/gin-gonic/gin"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/user_service"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/stock_trading_server"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/account_service"
  "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/transaction_service"
)

func SetupRoutes(r *gin.Engine) {
	// Rutas de usuario
	r.POST   ("/register", user_service.RegisterUser)
	r.GET    ("/users",    user_service.ListUsers)
	r.PUT    ("/user",     user_service.UpdateUser)
	r.DELETE ("/user/:id", user_service.DeleteUser)

	// Rutas de stock
	r.GET    ("/stock/:symbol", stock_trading_server.GetStockPrice)
	r.POST   ("/stock",         stock_trading_server.CreateStock)
	r.PUT    ("/stock",         stock_trading_server.UpdateStockPrice)

	// Rutas de cuentas
	r.GET    ("/account/:id", account_service.GetAccountBalance)
	r.POST   ("/account",     account_service.CreateAccount)
	r.PUT    ("/account",     account_service.UpdateAccountBalance)

  // Rutas de transacciones
  r.POST("/transaction", transaction_service.CreateTransaction)

}
