package presenter

import (
	"update-cart/entity"
	"update-cart/model"
)

func UpdateCart(userID string, item entity.CartItem) error {
	return model.UpdateCartItem(userID, item)
}
