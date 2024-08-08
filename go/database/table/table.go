package table

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/javiorfo/nvim-tabula/go/database/table/border"
)

type Header struct {
	Name   string
	Length int
}

type Tabula struct {
	DestFolder  string
	BorderStyle int
	Headers     map[int]Header
	Rows        [][]string
}

func (t Tabula) Generate() {
	b := border.GetBorder(border.BorderOption(t.BorderStyle))

	headerUp := b.CornerUpLeft
	headerMid := b.Vertical
	headerBottom := b.VerticalLeft

	headers := t.Headers
	headersLength := len(headers)
	for key := 1; key < headersLength+1; key++ {
		length := headers[key].Length
		headerUp += strings.Repeat(b.Horizontal, length)
		headerBottom += strings.Repeat(b.Horizontal, length)
		headerMid += addSpaces(headers[key].Name, length)
		headerMid += b.Vertical

		if key < headersLength {
			headerUp += b.DivisionUp
			headerBottom += b.Intersection
		} else {
			headerUp += b.CornerUpRight
			headerBottom += b.VerticalRight
		}
	}

	rows := t.Rows
	table := make([]string, 3, (len(rows)*2)+3)
	table[0] = fmt.Sprintf("%s\n", headerUp)
	table[1] = fmt.Sprintf("%s\n", headerMid)
	table[2] = fmt.Sprintf("%s\n", headerBottom)

	rowsLength := len(rows) - 1
	rowFieldsLength := len(rows[1]) - 1
	for i, row := range rows {
		value := b.Vertical
		var line string

		if i < rowsLength {
			line += b.VerticalLeft
		} else {
			line += b.CornerBottomLeft
		}

		for j, field := range row {
			value += addSpaces(field, headers[j+1].Length)
			value += b.Vertical

			line += strings.Repeat(b.Horizontal, headers[j+1].Length)
			if i < rowsLength {
				if j < rowFieldsLength {
					line += b.Intersection
				} else {
					line += b.VerticalRight
				}
			} else if j < rowFieldsLength {
				line += b.DivisionBottom
			} else {
				line += b.CornerBottomRight
			}
		}
		table = append(table, fmt.Sprintf("%s\n", value), fmt.Sprintf("%s\n", line))
	}

	for _, v := range table {
		fmt.Println(v)
	}

	WriteToFile(t.DestFolder, "tabula", table...)
}

func addSpaces(inputString string, length int) string {
	result := inputString

	if length > len(inputString) {
		diff := length - len(inputString)
		result += strings.Repeat(" ", diff)
	}

	return result
}

func WriteToFile(destFolder, filename string, values ...string) {
	fmt.Println(fmt.Sprintf("%s/%s", destFolder, filename))
	file, err := os.Create(fmt.Sprintf("%s/%s", destFolder, filename))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, v := range values {
		_, err := writer.WriteString(v)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
