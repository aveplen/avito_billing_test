package utils

func reverseByteSlice(x []byte) {
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}
}

func removeDotFromByteSlice(x []byte) []byte {
	res := make([]byte, 0, len(x))
	for _, val := range x {
		if val == '.' {
			continue
		}
		res = append(res, val)
	}
	return res
}

func calcFractLen(x []byte) int {
	res := 0
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == '.' {
			return res
		}
		res++
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func multReversedSliceByByte(str []byte, x byte) []byte {
	res := make([]byte, len(str), len(str)+1)
	var overflow int
	var digitSum int
	for i := 0; i < len(str); i++ {
		digitSum = getDigit(str, i)*int(x-'0') + overflow
		overflow = digitSum / 10
		res[i] = byte(digitSum%10 + '0')
	}
	if overflow != 0 {
		res = append(res, byte(overflow+'0'))
	}
	for len(res) >= 2 && res[len(res)-1] == '0' {
		res = res[:len(res)-1]
	}
	return res
}

func getDigit(str []byte, i int) int {
	if i >= len(str) || i < 0 {
		return 0
	}
	return int(str[i] - '0')
}

func addReversedByteSlices(a, b []byte) []byte {
	res := make([]byte, max(len(a), len(b)), max(len(a), len(b))+1)
	var overflow int
	var digitSum int
	for i := 0; i < max(len(a), len(b)); i++ {
		digitSum = getDigit(a, i) + getDigit(b, i) + overflow
		overflow = digitSum / 10
		res[i] = byte(digitSum%10 + '0')
	}
	if overflow != 0 {
		res = append(res, byte(overflow+'0'))
	}
	return res
}

func constructResult(mult []byte, fract int) []byte {
	if fract == 0 {
		reverseByteSlice(mult)
		return mult
	}
	if fract >= len(mult) {
		for fract > len(mult) {
			mult = append(mult, '0')
		}
		mult = append(mult, '.')
		mult = append(mult, '0')
		reverseByteSlice(mult)
		mult = truncateTrailingZeros(mult)
		return mult
	}
	res := make([]byte, 0, len(mult)+1)
	for i := 0; i < len(mult); i++ {
		if i == fract {
			res = append(res, '.')
		}
		res = append(res, mult[i])
	}
	reverseByteSlice(res)
	res = truncateTrailingZeros(res)
	return res
}

func truncateTrailingZeros(slice []byte) []byte {
	fract := calcFractLen(slice)
	if fract == 0 {
		return slice
	}
	var fractZeros int
	for i := len(slice) - 1; i >= len(slice)-fract-1; i-- {
		if slice[i] != '0' {
			break
		}
		fractZeros++
	}
	if fractZeros == fract {
		return slice[:len(slice)-fract-1]
	}
	return slice[:len(slice)-fractZeros]
}

func StringMultiplication(a, b []byte) []byte {
	if len(a) == 0 || len(b) == 0 {
		return []byte("")
	}
	fractLen := calcFractLen(a) + calcFractLen(b)

	reverseByteSlice(a)
	reverseByteSlice(b)

	a = removeDotFromByteSlice(a)
	b = removeDotFromByteSlice(b)

	res := []byte{'0'}
	for i := 0; i < len(b); i++ {
		temp := make([]byte, 0, i)
		for j := 0; j < i; j++ {
			temp = append(temp, '0')
		}
		temp = append(temp, multReversedSliceByByte(a, b[i])...)
		res = addReversedByteSlices(res, temp)
	}
	return constructResult(res, fractLen)
}

func StringFloor(str []byte, fractLen int) []byte {
	res := make([]byte, 0, len(str))
	var fract bool
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			res = append(res, str[i])
			fract = true
			continue
		}
		if !fract {
			res = append(res, str[i])
			continue
		}
		if fractLen == 0 {
			return res
		}
		res = append(res, str[i])
		fractLen--
	}
	return res
}
