package fabric

import (
	"log"
	"strconv"
	"strings"
)

type Claim struct {
	ID    int
	X     int
	Y     int
	Wide  int
	Tall  int
	Clean bool
}

func NewClaim(claim string) *Claim {
	claim = formatClaim(claim)
	fields := strings.Fields(claim)
	return &Claim{
		ID:   stringToInt(fields[0]),
		X:    stringToInt(fields[1]),
		Y:    stringToInt(fields[2]),
		Wide: stringToInt(fields[3]),
		Tall: stringToInt(fields[4]),
	}
}

func stringToInt(claim string) int {
	field, err := strconv.Atoi(claim)
	if err != nil {
		log.Fatalf("Could not convert string to int: %s", err)
	}
	return field
}

func formatClaim(claim string) string {
	claim = strings.Replace(claim, "#", "", -1)
	claim = strings.Replace(claim, "@ ", "", -1)
	claim = strings.Replace(claim, ",", " ", -1)
	claim = strings.Replace(claim, ":", "", -1)
	return strings.Replace(claim, "x", " ", -1)
}
