package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		reversedMenu[len(menu)-i-1] = menu[i]
	}
	return reversedMenu
}
