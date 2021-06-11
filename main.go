package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var data []uint8

var i int
var msg string

func main() {
	fmt.Println("Starting Parsing")

	i = 0
	data = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}
	code := "-[--->+<]>-.[---->+++++<]>-.+.++++++++++.+[---->+<]>+++.-[--->++<]>-.++++++++++.+[---->+<]>+++.[-->+++++++<]>.++.-------------.[--->+<]>---..+++++.-[---->+<]>++.+[->+++<]>.++++++++++++..---.[-->+<]>--------."
	msg = ""
	interpret(code)
	fmt.Println(msg)
}

func interpret(code string) {
	reader := bufio.NewReader(os.Stdin)

	v := strings.SplitAfter(code, "")

	for idx := 0; idx < len(v); idx++ {
		token := v[idx]

		switch token {
		case "<":
			i -= 1
			if i < 0 {
				i = 0
			}
			break
		case ">":
			if len(data)-1 <= i {
				data = append(data, 0)
			}
			i += 1
			break
		case "+":
			data[i] += 1
			break
		case "-":
			data[i] -= 1
			break
		case ".":
			msg += string(data[i])
			break
		case ",":
			text, _ := reader.ReadString('\n')
			data[i] = uint8(text[0])
			break
		case "[":
			fclosure := ""

			lc := 0
			rc := 0
			tr := true
			for tr == true {

				switch string(token) {
				case "[":
					lc += 1
				case "]":
					rc += 1
					if lc == rc {
						idx--
						tr = false
					}
				default:
					break
				}
				fclosure += token
				idx++
				if !(idx >= len(v)) {
					token = v[idx]
				} else {
					tr = false
				}
			}

			fclosure = fclosure[1 : len(fclosure)-1]

			for data[i] != 0 {
				interpret(fclosure)
			}
			break
		}
	}
}
