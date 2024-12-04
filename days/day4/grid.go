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

func (m *grid) numRows() int {
	return len(m.source)
}

func (m *grid) numColumns() int {
	if m.numRows() == 0 {
		return 0
	}

	return len(m.source[0])
}

func (m *grid) coordinatesAreValid(row, col int) bool {
	return row >= 0 && row < m.numRows() && col >= 0 && col < m.numColumns()
}

func (m *grid) get(row, col int) byte {
	return m.source[row][col]
}

func (m *grid) orientedSlice(
	startRow, startCol, rowOrientation, colOrientation, length int,
) string {
	endRow := startRow + rowOrientation*(length-1)
	endCol := startCol + colOrientation*(length-1)

	if !m.coordinatesAreValid(endRow, endCol) {
		return ""
	}

	slice := make([]byte, 0, length)

	for currentOffset := 0; currentOffset < length; currentOffset++ {
		currentRow := startRow + rowOrientation*currentOffset
		curentCol := startCol + colOrientation*currentOffset

		slice = append(
			slice,
			m.get(currentRow, curentCol),
		)
	}

	return string(slice)
}
