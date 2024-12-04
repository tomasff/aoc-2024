package day4

type grid struct {
	source []string
}

var allOrientations = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func (grid *grid) numRows() int {
	return len(grid.source)
}

func (grid *grid) numColumns() int {
	if grid.numRows() == 0 {
		return 0
	}

	return len(grid.source[0])
}

func (grid *grid) coordinatesAreValid(row, col int) bool {
	return row >= 0 && row < grid.numRows() && col >= 0 && col < grid.numColumns()
}

func (grid *grid) get(row, col int) byte {
	return grid.source[row][col]
}

func (grid *grid) orientedSlice(
	startRow, startCol, rowOrientation, colOrientation, length int,
) string {
	endRow := startRow + rowOrientation*(length-1)
	endCol := startCol + colOrientation*(length-1)

	if !grid.coordinatesAreValid(endRow, endCol) {
		return ""
	}

	slice := make([]byte, 0, length)

	for currentOffset := 0; currentOffset < length; currentOffset++ {
		currentRow := startRow + rowOrientation*currentOffset
		curentCol := startCol + colOrientation*currentOffset

		slice = append(
			slice,
			grid.get(currentRow, curentCol),
		)
	}

	return string(slice)
}
