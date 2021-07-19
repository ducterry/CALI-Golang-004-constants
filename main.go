package main

import "fmt"

/*
	- Ngay: 19.07.2021
	- By: DucND
*/
func main() {
	// Cach 1: Khai bao Constants
	const myName = "Duc Nguyen Dang"
	const age = 30

	// Cach 2: Khai bao tren 1 dong
	const employeeId, code = "Backend123", 15000

	// Cach 3: Khai bao tap chung
	const (
		company string  = "ABC"
		salary  float64 = 1000000000
	)

	// Print
	fmt.Println(myName,age,employeeId,code,company,salary)
}
