package runner

func BufToFloat(line []byte, idx int) float64 {
	val := 0
	isNeg := false

	if line[idx+1] == '-' {
		isNeg = true
		idx++
	}

	//will go on till .
	for i := idx + 1; i < len(line); i++ {
		b := line[i]
		if b == '.' {
			idx = i
			break
		}

		val *= 10
		val += int(b - 48)
	}

	//fractional part
	val2 := 0
	num := 1

	//dont need to count \n
	for i := idx + 1; i < len(line); i++ {
		b := line[i]
		val2 *= 10
		val2 += int(b - 48)
		num *= 10
	}

	final := float64(val*num+val2) / float64(num)
	if isNeg {
		final *= -1
	}
	return final
}
