package go_decimal_converter

import (
	"fmt"
	"strconv"
)

func DecimalToBiner(biner int64) {
	fmt.Println("Enter Decimal Number :")
	output := strconv.FormatInt(biner, 2)
	fmt.Println("Output Biner:", output)
}

func DecimalToOctal(octal int64) {
	output := strconv.FormatInt(octal, 8)
	fmt.Println("Output Octal:", output)
}

func DecimalToHexa(hexa int64) {
	output := strconv.FormatInt(hexa, 16)
	fmt.Println("Output Hexa:", output)
}
