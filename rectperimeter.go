package piscine

func RectPerimeter(w, h int) int {
	perimeter := 0
	if w < 0 || h < 0 {
		return -1
	} else {
		perimeter = (w + h) * 2
	}
	return perimeter
}
