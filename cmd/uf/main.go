package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iqer/algo4_go/algs4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	uf := algs4.NewUF(n)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		p, _ := strconv.Atoi(values[0])
		q, _ := strconv.Atoi(values[1])
		fmt.Println(p, q)
		uf.Union(p, q)
	}
}
