package primitive

func clamp(minVal int, val int, maxVal int) int {
	return max(minVal, min(maxVal, val))
}
