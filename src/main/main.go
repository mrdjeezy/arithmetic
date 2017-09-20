package main

import (
	"arith"
	"fmt"
)

func main(){
	TestAdd()
	TestSub()
	TestMult()
	TestDiv()
	TestSqrt()
	TestExp()
	TestConv()
}

func TestConv(){
	binRes := arith.NToN("2", 10, 2)

	if binRes != "10" {
		fmt.Println("int conv: broken")
	}

	octRes := arith.NToN("10", 10, 8)

	if octRes != "12" {
		fmt.Println("bin conv: broken")
	}

	hexRes := arith.NToN("10", 10, 16)

	if hexRes != "a" {
		fmt.Println("oct conv: broken")
	}

	hexToOctRes := arith.NToN("a", 16, 8)

	if hexToOctRes != "12" {
		fmt.Println("oct conv: broken")
	}
}

func TestExp(){
	floatRes := arith.ExpFloat(2.0, 2.0)

	if floatRes != 4.0 {
		fmt.Println("float exp: broken")
	}

	intRes := arith.ExpN("2", "2", 10)

	if intRes != "4" {
		fmt.Println("int exp: broken")
	}

	binRes := arith.ExpN("10", "10", 2)

	if binRes != "100" {
		fmt.Println("bin exp: broken")
	}

	octRes := arith.ExpN("2", "2", 8)

	if octRes != "4" {
		fmt.Println("oct exp: broken")
	}

	hexRes := arith.ExpN("2", "2",16)

	if hexRes != "4"{
		fmt.Println("hex exp: broken")
	}
}

func TestSqrt(){
	floatRes := arith.SqrtFloat(4.0)

	if floatRes != 2.0 {
		fmt.Println("float sqrt: broken")
	}

	intRes := arith.SquareRootN("4", 10)

	if intRes != "2" {
		fmt.Println("int sqrt: broken")
	}

	binRes := arith.SquareRootN("100",2)

	if binRes != "10" {
		fmt.Println("bin sqrt: broken")
	}

	octRes := arith.SquareRootN("4", 8)

	if octRes != "2" {
		fmt.Println("oct sqrt: broken")
	}

	hexRes := arith.SquareRootN("4",16)

	if hexRes != "2"{
		fmt.Println("hex sqrt: broken")
	}
}

func TestDiv(){
	floatRes := arith.DivideFloat(10.0, 5.0)

	if floatRes != 2.0 {
		fmt.Println("float div: broken")
	}

	intRes := arith.DivideN("10","5", 10)

	if intRes != "2" {
		fmt.Println("int div: broken")
	}

	binRes := arith.DivideN("10", "01", 2)

	if(binRes != "10"){
		fmt.Println("bin div: broken")
	}

	octRes := arith.DivideN("12", "5",8)

	if(octRes != "2"){
		fmt.Println("oct div: broken")
	}

	hexRes := arith.DivideN("a", "5",16)

	if(hexRes != "2"){
		fmt.Println("hex div: broken")
	}
}

func TestMult(){
	floatRes := arith.MultiplyFloat(2.0, 5.0)

	if floatRes != 10.0 {
		fmt.Println("float mult: broken")
	}

	intRes := arith.MultiplyN("2","5", 10)

	if intRes != "10" {
		fmt.Println("int mult: broken")
	}

	binRes := arith.MultiplyN("10", "01", 2)

	if(binRes != "10"){
		fmt.Println("bin mult: broken")
	}

	octRes := arith.MultiplyN("2", "5",8)

	if(octRes != "12"){
		fmt.Println("oct mult: broken")
	}

	hexRes := arith.MultiplyN("2", "5",16)

	if(hexRes != "a"){
		fmt.Println("hex mult: broken")
	}
}

func TestSub(){
	floatRes := arith.SubtractFloat(10.0, 5.0)

	if floatRes != 5.0 {
		fmt.Println("float sub: broken")
	}

	intRes := arith.SubtractN("10","5", 10)

	if intRes != "5" {
		fmt.Println("int sub: broken")
	}

	binRes := arith.SubtractN("10", "01", 2)

	if(binRes != "1"){
		fmt.Println("bin sub: broken")
	}

	octRes := arith.SubtractN("12", "5",8)

	if(octRes != "5"){
		fmt.Println("oct sub: broken")
	}

	hexRes := arith.SubtractN("a", "5",16)

	if(hexRes != "5"){
		fmt.Println("hex sub: broken")
	}
}

func TestAdd(){
	floatRes := arith.AddFloat(5.0, 5.0)

	if floatRes != 10.0 {
		fmt.Println("float add: broken")
	}

	intRes := arith.AddN("5","5", 10)

	if intRes != "10" {
		fmt.Println("int add: broken")
	}

	binRes := arith.AddN("01", "01", 2)

	if(binRes != "10"){
		fmt.Println("bin add: broken")
	}

	octRes := arith.AddN("5", "5",8)

	if(octRes != "12"){
		fmt.Println("oct add: broken")
	}

	hexRes := arith.AddN("5", "5",16)

	if(hexRes != "a"){
		fmt.Println("hex add: broken")
	}
}