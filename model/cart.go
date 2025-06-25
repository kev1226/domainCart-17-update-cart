package model

import (
	"encoding/json"
	"fmt"
	"update-cart/config"
	"update-cart/entity"
)

func UpdateCartItem(userID string, item entity.CartItem) error {
	key := fmt.Sprintf("cart:%s", userID)

	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}

	// Guardar el item en Redis (actualiza si ya existe)
	return config.RedisClient.HSet(config.Ctx, key, item.ProductID, itemJSON).Err()
}
