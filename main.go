package main

import "fmt"

const (
	tableBeginRow       string = "23"
	tableEndRow         string = "76"
	jobCostColumn       string = "E"
	employeeBeginColumn string = "F"
)

var excludeRows = []string{
	"52",
	"53",
	"76",
}

func main() {
	fmt.Println("RitaulPay запущен!")
}
