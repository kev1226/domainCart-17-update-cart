package view

import (
	"fmt"
	"net/http"
	"update-cart/dto"
	"update-cart/entity"
	"update-cart/presenter"

	"github.com/gin-gonic/gin"
)

func UpdateCart(c *gin.Context) {
	var body dto.UpdateCartDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener el usuario"})
		return
	}
	userID := fmt.Sprintf("%v", userIDInterface)

	item := entity.CartItem{
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
	}

	if err := presenter.UpdateCart(userID, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}
