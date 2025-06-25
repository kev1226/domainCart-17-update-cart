package routes

import (
	"update-cart/view"

	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func SetupRoutes(r *gin.Engine) {
	cart := r.Group("/cart", jwt.AuthGuard("user"))
	{
		cart.PUT("", view.UpdateCart)
	}
}
