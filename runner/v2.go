package runner

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type V2 struct {
}

// Use a custom split function for our usecase which also leads to
// lesser allocations due to not having to convert a byte array to string
// Although the gen_data script generates keys with a length of 10, using
// this assumption to improve the speed would feel like cheating
func (v V2) Run(inp string) map[string][]float64 {
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

		builder := strings.Builder{}
		idx := -1

		for i, b := range line {
			if b == ',' {
				idx = i
				break
			}
			builder.WriteByte(b)
		}

		name := builder.String()

		val := BufToFloat(line, idx)
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
