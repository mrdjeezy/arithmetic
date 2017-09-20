package arith

import (
	"strconv"
	"math"
)

func AddFloat(a float64, b float64) (ret float64){
	var res = a+b
	return res
}

func SubtractFloat(a float64, b float64) (ret float64){
	var res = a-b
	return res
}

func DivideFloat(a float64, b float64) (ret float64){
	var res = a/b
	return res
}

func MultiplyFloat(a float64, b float64) (ret float64){
	var res = a*b
	return res
}

func SqrtFloat(a float64) (ret float64){
	return math.Sqrt(a)
}

func ExpFloat(a float64, b float64) (ret float64){
	return math.Pow(a, b)
}

func AddN(nA string, nB string, base int) (ret string){
	aRes, _:=strconv.ParseUint(nA, base, 64)
	bRes, _:=strconv.ParseUint(nB, base, 64)
	res := aRes + bRes

	return strconv.FormatUint(res, base)
}

func SubtractN(nA string, nB string, base int) (ret string){
	aRes, _:=strconv.ParseUint(nA, base, 64)
	bRes, _:=strconv.ParseUint(nB, base, 64)
	res := aRes - bRes

	return strconv.FormatUint(res, base)
}

func MultiplyN(nA string, nB string, base int) (ret string){
	aRes, _:=strconv.ParseUint(nA, base, 64)
	bRes, _:=strconv.ParseUint(nB, base, 64)
	res := aRes * bRes

	return strconv.FormatUint(res, base)
}

func DivideN(nA string, nB string, base int) (ret string){
	aRes, _:=strconv.ParseUint(nA, base, 64)
	bRes, _:=strconv.ParseUint(nB, base, 64)
	res := aRes / bRes

	return strconv.FormatUint(res, base)
}

func SquareRootN(nA string, base int) (ret string) {
	aRes, _ := strconv.ParseUint(nA, base, 64)
	res := math.Sqrt(float64(aRes))

	return strconv.FormatUint(uint64(res), base)
}

func ExpN(nA string, nB string, base int) (ret string){
	aRes, _:=strconv.ParseUint(nA, base, 64)
	bRes, _:=strconv.ParseUint(nB, base, 64)
	res := math.Pow(float64(aRes), float64(bRes))

	return strconv.FormatUint(uint64(res), base)
}

func NToN(nA string, from int, to int) (ret string){
	aRes,_ := strconv.ParseUint(nA, from, 64);
	return strconv.FormatUint(aRes, to);
}
