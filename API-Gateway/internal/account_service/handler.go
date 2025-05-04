package account_service

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JeroZp/gRPC-MOM/API-Gateway/internal/account_service/proto"
)


// GET /account/:id - obtener saldo actual de la cuenta
func GetAccountBalance(c *gin.Context) {
	if AccountClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AccountClient no está inicializado"})
		return
	}

	id := c.Param("id")

	grpcReq := &proto.AccountRequest{IdAccount: id}

	resp, err := AccountClient.GetAccountBalance(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_account":    resp.GetIdAccount(),
		"credits":       resp.GetCredits(),
		"creation_date": resp.GetCreationDate(),
	})
}

// POST /account - crear una nueva cuenta
func CreateAccount(c *gin.Context) {
	var req struct {
		IdAccount string  `json:"id_userForAccount" binding:"required"`
		Credits   float64 `json:"credits" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &proto.CreateAccountRequest{
		IdAccount: req.IdAccount,
		Credits:   req.Credits,
	}

	resp, err := AccountClient.CreateAccount(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.GetSuccess(),
		"message": resp.GetMessage(),
	})
}

// PUT /account - actualizar créditos de una cuenta
func UpdateAccountBalance(c *gin.Context) {
	var req struct {
		IdAccount string `json:"id_account" binding:"required"`
		Credits   int64  `json:"credits" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &proto.UpdateAccountRequest{
		IdAccount: req.IdAccount,
		Credits:   req.Credits,
	}

	resp, err := AccountClient.UpdateAccountBalance(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.GetSuccess(),
		"message": resp.GetMessage(),
	})
}
