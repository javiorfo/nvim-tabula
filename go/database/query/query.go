package query

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Header struct {
	Name   string
	Length int
}

type Tabula struct {
	Headers map[int]Header
// 	Rows   map[int][]string
	Rows   [][]string
}

func (t Tabula) Generate() {
	const corner_up_left = "┏"
	const corner_up_right = "┓"
	const corner_bottom_left = "┗"
	const corner_bottom_right = "┛"
	const div_up = "┳"
	const div_bottom = "┻"
	const hor = "━"
	const vert = "┃"
	const intersection = "╋"
	const vert_left = "┣"
	const vert_right = "┫"

	header_up := corner_up_left
	header_mid := vert
	header_bottom := vert_left

	headers := t.Headers
	headersLength := len(headers)
	for key := 1; key < headersLength+1; key++ {
		length := headers[key].Length
		header_up += strings.Repeat(hor, length)
		header_bottom += strings.Repeat(hor, length)
		header_mid += addSpaces(headers[key].Name, length)
		header_mid += vert

		if key < headersLength {
			header_up += div_up
			header_bottom += intersection
		} else {
			header_up += corner_up_right
			header_bottom += vert_right
		}
	}

	rows := t.Rows
	table := make([]string, 3, (len(rows)*2)+3)
    table[0] = header_up
    table[1] = header_mid
    table[2] = header_bottom

	rowsLength := len(rows) - 1
	rowFieldsLength := len(rows[1]) - 1
	for i, row := range rows {
		value := vert
		var line string

		if i < rowsLength {
			line += vert_left
		} else {
			line += corner_bottom_left
		}

		for j, field := range row {
			value += addSpaces(field, headers[j+1].Length)
			value += vert

			line += strings.Repeat(hor, headers[j+1].Length)
			if i < rowsLength {
				if j < rowFieldsLength {
					line += intersection
				} else {
					line += vert_right
				}
			} else if j < rowFieldsLength {
				line += div_bottom
			} else {
				line += corner_bottom_right
			}
		}
		table = append(table, value, line)
	}

	for _, v := range table {
		fmt.Println(v)
	}

//     writeTable(table)
}

func addSpaces(inputString string, length int) string {
	result := inputString

	if length > len(inputString) {
		diff := length - len(inputString)
		result += strings.Repeat(" ", diff)
	}

	return result
}

func writeTable(values []string) {
	file, err := os.Create("1.tabula")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

    for _, v := range values {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", v)) // agregar \n en la función principal
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
