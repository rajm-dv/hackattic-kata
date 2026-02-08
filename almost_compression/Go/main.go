package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	cnt := 1
	finalText := ""
	for i := 1; i < len(text); i++ {
		if text[i] == text[i-1] {
			cnt += 1
		} else {
			if cnt > 2 {
				finalText = finalText + strconv.Itoa(cnt) + string(text[i-1])
			} else if cnt == 2 {
				finalText = finalText + strings.Repeat(string(text[i-1]), 2)
			}	else {
				finalText = finalText + string(text[i-1])
			}
			cnt = 1
		}
	}

	fmt.Println(finalText)
}
