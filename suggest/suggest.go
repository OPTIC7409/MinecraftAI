package suggest

func SuggestAction(health int) string {
	if health < 20 {
		return "Find food or healing items!"
	}
	return "Keep exploring."
}
