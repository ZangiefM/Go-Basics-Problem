package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	} else if order == "chips" {
		return 10
	} else if order == "nuggets" {
		return 12
	} else if order == "burger" {
		return 37
	} else {
		return 404
	}
}
