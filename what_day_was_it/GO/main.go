package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countDay(input string) {
	atoi, _ := strconv.Atoi(input)
	days := [7]string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	idx := ((atoi % 7) + 7) % 7
	fmt.Println(days[idx])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		countDay(line)
	}
}
