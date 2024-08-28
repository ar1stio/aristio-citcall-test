package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// time execution 1.9385434s
	testA()
	// time execution 1.0463004s
	testB()
	callBebek()

	// Calculate the elapsed time
	elapsed := time.Since(start)

	// Print the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)

	// Example 1
	input1 := "II + II = HIU"
	fmt.Printf("Input  : %s\nOutput : %s\n\n", input1, solveCryptarithmExpression(input1))

	// Example 2
	input2 := "ABD - AD = DKL"
	fmt.Printf("Input  : %s\nOutput : %s\n", input2, solveCryptarithmExpression(input2))
}

// Check if a digit is used in the current mapping
func isUsed(mapping map[rune]int, digit int) bool {
	for _, v := range mapping {
		if v == digit {
			return true
		}
	}
	return false
}

// Check if the character is at the start of a number
func isLeadingZero(ch rune, expression string, mapping map[rune]int) bool {
	parts := strings.FieldsFunc(expression, func(r rune) bool { return r == '+' || r == '-' || r == '=' || r == ' ' })
	for _, part := range parts {
		if len(part) > 1 && rune(part[0]) == ch {
			return true
		}
	}
	return false
}

// Validate if the expression with the current mapping is correct
func validateExpression(expression string, mapping map[rune]int) (string, bool) {
	replaced := ""
	for _, ch := range expression {
		if num, exists := mapping[ch]; exists {
			replaced += strconv.Itoa(num)
		} else {
			replaced += string(ch)
		}
	}

	parts := strings.Split(replaced, "=")
	if len(parts) != 2 {
		return "", false
	}

	leftSide := strings.TrimSpace(parts[0])
	rightSide := strings.TrimSpace(parts[1])

	leftParts := strings.Fields(leftSide)
	if len(leftParts) != 3 {
		return "", false
	}

	leftOperand1, _ := strconv.Atoi(leftParts[0])
	leftOperand2, _ := strconv.Atoi(leftParts[2])
	rightOperand, _ := strconv.Atoi(rightSide)
	operator := leftParts[1]

	var calculatedValue int
	if operator == "+" {
		calculatedValue = leftOperand1 + leftOperand2
	} else if operator == "-" {
		calculatedValue = leftOperand1 - leftOperand2
	}

	return replaced, calculatedValue == rightOperand
}

// Solve the cryptarithm by trying all possible mappings
func solveCryptarithm(expression string, uniqueChars []rune, mapping map[rune]int, usedDigits map[int]bool, index int) (bool, string) {
	if index == len(uniqueChars) {
		result, isValid := validateExpression(expression, mapping)
		return isValid, result
	}

	for digit := 0; digit <= 9; digit++ {
		if isUsed(mapping, digit) || (digit == 0 && isLeadingZero(uniqueChars[index], expression, mapping)) {
			continue
		}

		mapping[uniqueChars[index]] = digit
		usedDigits[digit] = true

		if solved, result := solveCryptarithm(expression, uniqueChars, mapping, usedDigits, index+1); solved {
			return true, result
		}

		delete(mapping, uniqueChars[index])
		usedDigits[digit] = false
	}

	return false, ""
}

// Get all unique characters from the expression
func getUniqueChars(expression string) []rune {
	charSet := make(map[rune]bool)
	for _, ch := range expression {
		if ch != ' ' && ch != '+' && ch != '-' && ch != '=' {
			charSet[ch] = true
		}
	}
	uniqueChars := make([]rune, 0, len(charSet))
	for ch := range charSet {
		uniqueChars = append(uniqueChars, ch)
	}
	return uniqueChars
}

// Main function to solve the cryptarithm problem
func solveCryptarithmExpression(expression string) string {
	uniqueChars := getUniqueChars(expression)
	mapping := make(map[rune]int)
	usedDigits := make(map[int]bool)

	if solved, result := solveCryptarithm(expression, uniqueChars, mapping, usedDigits, 0); solved {
		return result
	}
	return "No solution found"
}

func longRunningTask() {
	// Simulate a long-running task
	time.Sleep(2 * time.Second)
}

var x int
var i uint64
var n uint64
var tmp []byte

func testB() {
	fmt.Println("test B: ")
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}
	func() {
		// The Test
		for i := 0; i < len(tmp); i++ {
			fmt.Println("test B: ", i)
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}

func testA() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}
	func() {
		// The test
		size := len(tmp)
		for i := 0; i < size; i++ {
			fmt.Println("test A: ", i)
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

// Method untuk mengubah status hidup bebek menjadi mati
func (b *Bebek) Mati() {
	b.hidup = false
}

// Method untuk melakukan aksi terbang
func (b *Bebek) Terbang() {
	if b.energi > 0 && b.hidup && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			b.Mati()
		}
	}
}

// Method untuk menambah energi bebek
func (b *Bebek) Makan() {
	if b.hidup {
		b.energi += 1
	}
}

func callBebek() {
	// Contoh penggunaan struct Bebek dan method-methodnya
	bebek := Bebek{
		energi:       5,
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "Kwek kwek",
	}

	// Aksi terbang
	bebek.Terbang()
	fmt.Printf("Energi setelah terbang: %d\n", bebek.energi)

	// Aksi makan
	bebek.Makan()
	fmt.Printf("Energi setelah makan: %d\n", bebek.energi)

	// Aksi terbang lagi hingga energi habis
	for bebek.energi > 0 {
		bebek.Terbang()
		fmt.Printf("Energi setelah terbang: %d\n", bebek.energi)
	}

	// Cek status hidup bebek
	fmt.Printf("Apakah bebek masih hidup? %v\n", bebek.hidup)
}
