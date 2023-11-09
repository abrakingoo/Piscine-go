package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {

	chips := food{
		preptime: 10,
	}
	burger := food{
		preptime: 15,
	}
	nuggets := food{
		preptime: 12,
	}

	switch order {
	case "chips":
		return chips.preptime
	case "burger":
		return burger.preptime
	case "nuggets":
		return nuggets.preptime
	}
	return 404

}
