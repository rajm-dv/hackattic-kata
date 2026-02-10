package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compress(input string) string {
	compressedText := ""
	i := 0
	n := len(input)
	for i < n {
		cnt := 1
		j := i + 1
		isSame := true
		for j < n && isSame {
			if input[i] == input[j] {
				cnt += 1
				j += 1
			} else {
				isSame = false
			}
		}
		if cnt > 2 {
			compressedText = compressedText + strconv.Itoa(cnt)
			compressedText += string(input[i])
		} else if cnt == 2 {
			compressedText += string(input[i])
			compressedText += string(input[i])
		} else {
			compressedText += string(input[i])
		}
		i = j
	}
	return compressedText
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		compressedText := compress(line)
		fmt.Println(compressedText)
	}
}
