package main

import "bytes"

func stringSum(n1, n2 string) string {

	i, j := len(n1)-1, len(n2)-1
	var carry int = 0
	var result []byte

	for i >= 0 || j >= 0 || carry > 0 {
		leftDigit := 0
		if i >= 0 {
			leftDigit = int(n1[i] - '0')
		}

		rightDigit := 0
		if j >= 0 {
			rightDigit = int(n2[j] - '0')
		}
		sum := leftDigit + rightDigit + carry
		carry = sum / 10
		sum = sum % 10
		result = append([]byte{'0' + byte(sum)}, result...)
		i--
		j--
	}

	for k := 0; k < len(result); k++ {
		if result[k] != '0' {
			return string(result[k:])
		}
	}

	return "0"

}

func solution(num1, num2 string) string {
	// TODO: Implement the solution
	var placeValue int = 0
	var productCarry int = 0
	var result []byte

	i := len(num1) - 1

	for i >= 0 {
		j := len(num2) - 1
		productRow := []byte{}

		for j >= 0 || productCarry > 0 {
			firstDigit := 0
			if i >= 0 {
				firstDigit = int(num1[i] - '0')
			}

			secondDigit := 0
			if j >= 0 {
				secondDigit = int(num2[j] - '0')
			}
			product := firstDigit*secondDigit + productCarry
			productCarry = product / 10
			productDigit := product % 10

			productRow = append([]byte{'0' + byte(productDigit)}, productRow...)

			j--
		}

		if placeValue > 0 {
			productRow = append(productRow, bytes.Repeat([]byte{'0'}, placeValue)...)
		}
		placeValue++
		result = []byte(stringSum(string(productRow), string(result)))
		i--
	}

	return string(result)
}
