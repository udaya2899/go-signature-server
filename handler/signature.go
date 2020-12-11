package handler

import (
	"fmt"
	"net/http"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/logger"
	"challenge.summitto.com/udaya2899/challenge_result/signature"
	"github.com/gin-gonic/gin"
)

func setSignatureRoutes(router *gin.Engine) {
	router.GET("/public_key", getPublicKey)
	router.PUT("/transaction", putTransaction)
	router.POST("/signature", postSignature)
}

func getPublicKey(c *gin.Context) {
	if config.Env.PublicKey == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Public Key not found. Something went wrong. Try running 'make keygen' on the server if you have access",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key": config.Env.PublicKey,
	})
}

func putTransaction(c *gin.Context) {
	var transaction struct {
		Txn string `json:"txn" binding:"required"`
	}

	if c.Bind(&transaction) == nil {
		signatureService := signature.NewService(logger.Logger)

		id, err := signatureService.PutTransactionBlob(transaction.Txn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Errorf("Cannot save transaction blob, err: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
		return

	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Input is not of type transaction",
	})
}

func postSignature(c *gin.Context) {
	var transactions struct {
		IDs []string `json:"ids" binding:"required"`
	}

	var signedMessage *signature.SignedMessage

	if c.Bind(&transactions) == nil {
		signatureService := signature.NewService(logger.Logger)
		var err error

		signedMessage, err = signatureService.PostSignature(transactions.IDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Cannot sign for given transaction IDs, err: %v", err),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   signedMessage.MessageContent,
		"signature": signedMessage.Sign,
	})
}
