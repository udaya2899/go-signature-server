package handler

import (
	"crypto/ed25519"
	"log"
	"net/http"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setSignatureRoutes(router *gin.Engine) {
	router.GET("/public_key", getPublicKey)
	router.PUT("/transaction", putTransaction)
	router.POST("/signature", postSignature)
}

func getPublicKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"public_key": config.Env.PublicKey,
	})
}

func putTransaction(c *gin.Context) {
	var transaction struct {
		Txn string `json:"txn" binding:"required"`
	}

	var id string

	if c.Bind(&transaction) == nil {
		id = uuid.New().String()

		err := db.PutTransaction(id, transaction.Txn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input is not of type transaction",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func postSignature(c *gin.Context) {
	var transactions struct {
		IDs []string `json:"ids" binding:"required"`
	}

	var values []byte

	if c.Bind(&transactions) == nil {
		for _, id := range transactions.IDs {

			value, err := db.GetTransactionByID(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			values = append(values, value...)
		}
	}

	log.Printf("Value obtained from ID: %+v", values)

	signed := ed25519.Sign(config.Env.PrivateKey, values)

	if ok := ed25519.Verify(config.Env.PublicKey, values, signed); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot verify self-signed signature, something wrong",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   string(values),
		"signature": signed,
	})
}
