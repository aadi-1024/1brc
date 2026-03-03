package runner

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Baseline approach, simply read the file line by line sequentially
type V1 struct{}

func (v V1) Run(inp string) map[string][]float64 {
	res := make(map[string][]float64)

	file, err := os.Open(inp)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer file.Close()

	rdr := bufio.NewReader(file)

	for {
		line, _, err := rdr.ReadLine()
		if err != nil {
			break
		}

		x := strings.Split(string(line), ",")
		name := x[0]
		val, err := strconv.ParseFloat(x[1], 32)

		v, ok := res[name]
		if ok {
			v[0] += 1
			v[1] += val
		} else {
			arr := make([]float64, 3)
			arr[0] = 1
			arr[1] = val
			res[name] = arr
		}
	}

	return res
}
