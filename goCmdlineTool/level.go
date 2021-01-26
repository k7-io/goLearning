package main

func Level(point int) string {
	switch {
	case point < 60:
		return "not pass"
	case point < 70:
		return "pass"
	case point < 80:
		return "good"
	case point < 100:
		return "excellent"
	default:
		return "error data"
	}
}
