package db

import (
	orm "opencodes/pkg/database"
)

type Order struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	GoodName   string  `json:"good_name"`
	GoodPrice  float64 `json:"good_price"`
	GoodCounts int     `json:"good_counts"`
	CreatedAt  int64   `json:"created_at"`
	UpdatedAt  int64   `json:"updated_at"`
}

func CreateOrder(order *Order) error {
	db := orm.Eloquent.Table("o_order").Create(order)
	if db.Error != nil {

	}
	return nil
}
