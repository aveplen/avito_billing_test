package utils

import "testing"

func byteSlicesAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestByteSlicesAreEqual(t *testing.T) {
	var a, b []byte
	var res bool

	a = []byte("")
	b = []byte("")
	res = byteSlicesAreEqual(a, b)
	if !res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' != '%s'", a, b)
	}

	a = []byte("abcd")
	b = []byte("abcd")
	res = byteSlicesAreEqual(a, b)
	if !res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' != '%s'", a, b)
	}

	a = []byte("")
	b = []byte("abcd")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}

	a = []byte("abcd")
	b = []byte("")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}

	a = []byte("abcd123")
	b = []byte("abcd")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}

	a = []byte("abcd")
	b = []byte("abcd123")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}

	a = []byte("123abcd")
	b = []byte("abcd")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}

	a = []byte("abcd")
	b = []byte("123abcd")
	res = byteSlicesAreEqual(a, b)
	if res {
		t.Errorf("byte slice equality assert was incorrect, got '%s' == '%s'", a, b)
	}
}

func TestReverseByteSlice(t *testing.T) {
	var a, b []byte

	a = []byte("")
	b = []byte("")
	reverseByteSlice(b)
	if !byteSlicesAreEqual(a, b) {
		t.Errorf("reversing byte slice failed, got '%s' != '%s'", a, b)
	}

	a = []byte("a")
	b = []byte("a")
	reverseByteSlice(b)
	if !byteSlicesAreEqual(a, b) {
		t.Errorf("reversing byte slice failed, got '%s' != '%s'", a, b)
	}

	a = []byte("ab")
	b = []byte("ba")
	reverseByteSlice(b)
	if !byteSlicesAreEqual(a, b) {
		t.Errorf("reversing byte slice failed, got '%s' != '%s'", a, b)
	}

	a = []byte("abc")
	b = []byte("cba")
	reverseByteSlice(b)
	if !byteSlicesAreEqual(a, b) {
		t.Errorf("reversing byte slice failed, got '%s' != '%s'", a, b)
	}

	a = []byte("abcdefghijklmnopqrstuvwxyz")
	b = []byte("zyxwvutsrqponmlkjihgfedcba")
	reverseByteSlice(b)
	if !byteSlicesAreEqual(a, b) {
		t.Errorf("reversing byte slice failed, got '%s' != '%s'", a, b)
	}
}

func TestRemoveDotFromByteSlice(t *testing.T) {
	var a, b, aWithoutDots []byte

	a = []byte("")
	b = []byte("")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte(".")
	b = []byte("")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte("....")
	b = []byte("")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte("123.456789")
	b = []byte("123456789")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte(".123456789")
	b = []byte("123456789")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte("123456789.")
	b = []byte("123456789")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte("....123....456....789...")
	b = []byte("123456789")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}

	a = []byte("....abc....def....ghi...")
	b = []byte("abcdefghi")
	aWithoutDots = removeDotFromByteSlice(a)
	if !byteSlicesAreEqual(aWithoutDots, b) {
		t.Errorf("removing dots from slice array failed, got '%s' != '%s'", aWithoutDots, b)
	}
}

func TestCalcFractLen(t *testing.T) {
	var a []byte
	var fractLen, expect int

	a = []byte("")
	expect = 0
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0.")
	expect = 0
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0.1")
	expect = 1
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0.12")
	expect = 2
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0.123")
	expect = 3
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0.123456789")
	expect = 9
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0..1")
	expect = 1
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("0..123")
	expect = 3
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("123.456.789")
	expect = 3
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("123.456.789.")
	expect = 0
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}

	a = []byte("123.456.789.abc")
	expect = 3
	fractLen = calcFractLen(a)
	if fractLen != expect {
		t.Errorf("calculating fract len failed, expected: %d, got: %d", expect, fractLen)
	}
}

func TestMax(t *testing.T) {
	var a, b, maximum, expect int

	a = 0
	b = 0
	maximum = max(a, b)
	expect = 0
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = -1
	b = 0
	maximum = max(a, b)
	expect = 0
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = 0
	b = -1
	maximum = max(a, b)
	expect = 0
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = 1
	b = 0
	maximum = max(a, b)
	expect = 1
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = 0
	b = 1
	maximum = max(a, b)
	expect = 1
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = -1
	b = 1
	maximum = max(a, b)
	expect = 1
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = 1
	b = -1
	maximum = max(a, b)
	expect = 1
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}

	a = (^0) & (1 << 32)
	b = ^0
	maximum = max(a, b)
	expect = (^0) & (1 << 32)
	if maximum != expect {
		t.Errorf("max failed, expected: %d, got: %d", expect, maximum)
	}
}

func TestGetDigit(t *testing.T) {
	var slice []byte
	var i int
	var got int
	var expected int

	slice = []byte("1")
	i = 0
	got = getDigit(slice, i)
	expected = 1
	if got != expected {
		t.Errorf("get digit failed, expected: %d, got: %d", expected, got)
	}

	slice = []byte("1")
	i = 1
	got = getDigit(slice, i)
	expected = 0
	if got != expected {
		t.Errorf("get digit failed, expected: %d, got: %d", expected, got)
	}

	slice = []byte("123456789")
	i = 8
	got = getDigit(slice, i)
	expected = 9
	if got != expected {
		t.Errorf("get digit failed, expected: %d, got: %d", expected, got)
	}

	slice = []byte("123456789")
	i = -1
	got = getDigit(slice, i)
	expected = 0
	if got != expected {
		t.Errorf("get digit failed, expected: %d, got: %d", expected, got)
	}

	slice = []byte("123456789")
	i = 9
	got = getDigit(slice, i)
	expected = 0
	if got != expected {
		t.Errorf("get digit failed, expected: %d, got: %d", expected, got)
	}
}

func TestMultRevesedSliceByByte(t *testing.T) {
	var slice, expected, got []byte
	var x byte

	slice = []byte("")
	x = '0'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("0")
	x = '0'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("321")
	x = '0'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("321000")
	x = '0'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("321")
	x = '1'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("321")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("3210000")
	x = '1'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("321")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("321")
	x = '3'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("963")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("321")
	x = '4'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("294")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("222")
	x = '5'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("0111")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("0000022200000")
	x = '5'
	got = multReversedSliceByByte(slice, x)
	expected = []byte("000000111")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("mult reversed slice by byte failed, expected: '%s', got: '%s'", expected, got)
	}
}

func TestAddReversedByteSlices(t *testing.T) {
	var a, b, expected, got []byte

	a = []byte("")
	b = []byte("")
	got = addReversedByteSlices(a, b)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0")
	b = []byte("")
	got = addReversedByteSlices(a, b)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("")
	b = []byte("0")
	got = addReversedByteSlices(a, b)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0")
	b = []byte("0")
	got = addReversedByteSlices(a, b)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("1")
	b = []byte("0")
	got = addReversedByteSlices(a, b)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0")
	b = []byte("1")
	got = addReversedByteSlices(a, b)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("2")
	b = []byte("3")
	got = addReversedByteSlices(a, b)
	expected = []byte("5")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("2")
	b = []byte("8")
	got = addReversedByteSlices(a, b)
	expected = []byte("01")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("9999")
	b = []byte("1")
	got = addReversedByteSlices(a, b)
	expected = []byte("00001")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("891")
	b = []byte("201")
	got = addReversedByteSlices(a, b)
	expected = []byte("003")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("add reversed slices failed, expected: '%s', got: '%s'", expected, got)
	}
}

func TestConstructResult(t *testing.T) {
	var slice, expected, got []byte
	var fractLen int

	slice = []byte("")
	fractLen = 0
	got = constructResult(slice, fractLen)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("0")
	fractLen = 0
	got = constructResult(slice, fractLen)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("0")
	fractLen = 1
	got = constructResult(slice, fractLen)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("0")
	fractLen = 2
	got = constructResult(slice, fractLen)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("654321")
	fractLen = 2
	got = constructResult(slice, fractLen)
	expected = []byte("1234.56")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("654321")
	fractLen = 8
	got = constructResult(slice, fractLen)
	expected = []byte("0.00123456")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("00654321")
	fractLen = 10
	got = constructResult(slice, fractLen)
	expected = []byte("0.00123456")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("construct result failed, expected: '%s', got: '%s'", expected, got)
	}
}

func TestStringMultiplication(t *testing.T) {
	var a, b, expected, got []byte

	a = []byte("")
	b = []byte("")
	got = StringMultiplication(a, b)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("1")
	b = []byte("")
	got = StringMultiplication(a, b)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("")
	b = []byte("1")
	got = StringMultiplication(a, b)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("1")
	b = []byte("0")
	got = StringMultiplication(a, b)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0")
	b = []byte("1")
	got = StringMultiplication(a, b)
	expected = []byte("0")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("20")
	b = []byte("5")
	got = StringMultiplication(a, b)
	expected = []byte("100")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0.00001")
	b = []byte("100000")
	got = StringMultiplication(a, b)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0.0000100000000000")
	b = []byte("100000.00000000000")
	got = StringMultiplication(a, b)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("100000.00000000000")
	b = []byte("0.0000100000000000")
	got = StringMultiplication(a, b)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("20.21123456")
	b = []byte("5.123")
	got = StringMultiplication(a, b)
	expected = []byte("103.54215465088")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("5.45")
	b = []byte("0.00123123123123123123123123123")
	got = StringMultiplication(a, b)
	expected = []byte("0.0067102102102102102102102102035")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0.0012")
	b = []byte("0.0012")
	got = StringMultiplication(a, b)
	expected = []byte("0.00000144")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}

	a = []byte("0.5555")
	b = []byte("2002")
	got = StringMultiplication(a, b)
	expected = []byte("1112.111")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string multiplication failed, expected: '%s', got: '%s'", expected, got)
	}
}

func TestStringFloor(t *testing.T) {
	var slice, expected, got []byte

	slice = []byte("")
	got = StringFloor(slice, 10)
	expected = []byte("")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("1")
	got = StringFloor(slice, 10)
	expected = []byte("1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("1.15")
	got = StringFloor(slice, 10)
	expected = []byte("1.15")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("1.123456")
	got = StringFloor(slice, 1)
	expected = []byte("1.1")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("1.123456")
	got = StringFloor(slice, 2)
	expected = []byte("1.12")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}

	slice = []byte("1.123456")
	got = StringFloor(slice, 5)
	expected = []byte("1.12345")
	if !byteSlicesAreEqual(got, expected) {
		t.Errorf("string floor failed, expected: '%s', got: '%s'", expected, got)
	}
}
