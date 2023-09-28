package scannertools

import (
	"strconv"
	"strings"
)

func StringToIntPorts(ports *string) []int {
	// Switch case to find 3 types of ports:
	// comma ","
	// dash "-"
	// Or just a single port.
	switch {
	case strings.Contains(*ports, "-"):
		portsslice := ConvertArrayPortsToIntDash(*ports)
		return portsslice
	case strings.Contains(*ports, ","):
		portsslice := ConvertArrayPortsToIntComma(*ports)
		return portsslice
	default:
		var ints []int
		portConvertToInt, err := strconv.Atoi(*ports)
		if err != nil {
			panic(err)
		}
		ints = append(ints, portConvertToInt)
		return ints
	}
}

func ConvertArrayPortsToIntDash(ports string) []int {
	// Accept a string of ports, example: 20-30, then return a slice of ports
	// Example: [20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30]
	splitByDash := strings.Split(ports, "-")
	first := splitByDash[0]
	firstInt, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}
	second := splitByDash[1]
	secondInt, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}
	var rangeslice []int
	for index := firstInt; index <= secondInt; index++ {
		rangeslice = append(rangeslice, index)
	}

	return rangeslice
}

func ConvertArrayPortsToIntComma(ports string) []int {
	// Accept a string of ports, example: 20,30 then return a slice of ports
	// Example: [20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30]
	splitByComma := strings.Split(ports, ",")
	ints := make([]int, len(splitByComma))
	for i, s := range splitByComma {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
