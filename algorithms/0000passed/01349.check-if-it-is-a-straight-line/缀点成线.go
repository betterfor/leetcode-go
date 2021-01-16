package code

func checkStraightLine(coordinates [][]int) bool {
	x0, y0, x1, y1 := coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (y1-y0)*(x-x1) != (y-y1)*(x1-x0) {
			return false
		}
	}
	return true
}
