package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isValidCNPJ(cnpj string) bool {
	
	// Remove caracteres não numéricos
	cnpj = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, cnpj)

	// Verifica se o CNPJ tem 14 dígitos e não é uma sequência de dígitos iguais
  // Exemplo: 11111111111111, 22222222222222, etc.
	if len(cnpj) != 14 || strings.Repeat(string(cnpj[0]), 14) == cnpj {
		return false
	}

	// Pesos utilizados para o cálculo dos dois dígitos verificadores
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	/// Verifica se o CNPJ é válido
  // O CNPJ é composto por 12 dígitos + 2 dígitos verificadores
  // O primeiro dígito verificador é calculado a partir dos 12 primeiros dígitos
  // O segundo dígito verificador é calculado a partir dos 12 primeiros dígitos e do primeiro dígito verificador
	calculateDigit := func(weights []int) int {
		sum := 0
		for i := 0; i < len(weights); i++ {
			num, _ := strconv.Atoi(string(cnpj[i]))
			sum += num * weights[i]
		}
		remainder := sum % 11
		if remainder < 2 {
			return 0
		}
		return 11 - remainder
	}

	// Calcula os dois dígitos verificadores
	d1 := calculateDigit(weights1)
	d2 := calculateDigit(weights2)

	// Verifica se os dígitos verificadores calculados são iguais aos dígitos verificadores do CNPJ
  // Ex.: XX.XXX.XXX/XXXX-95 -> dígitos verificadores são 9 e 5
	return strconv.Itoa(d1)+strconv.Itoa(d2) == cnpj[12:]
}

func main() {
	// Exemplos de uso

	// Exemplos de CNPJs válidos
	fmt.Println(isValidCNPJ("12.345.678/0001-95"))  // true
	fmt.Println(isValidCNPJ("12345678000195"))      // true

	// Exemplos de CNPJs inválidos
	fmt.Println(isValidCNPJ("12.345.678/0001-96"))  // false
	fmt.Println(isValidCNPJ("12345678000196"))      // false
	fmt.Println(isValidCNPJ("12345678901234"))      // false
	fmt.Println(isValidCNPJ("1234567890123"))       // false
	fmt.Println(isValidCNPJ("123456789012345"))     // false
	fmt.Println(isValidCNPJ("1234567890123a"))      // false
	fmt.Println(isValidCNPJ("1234567890123-4"))     // false
	fmt.Println(isValidCNPJ("1234567890123.4"))     // false
	fmt.Println(isValidCNPJ("11.111.111/1111-11"))  // false
	fmt.Println(isValidCNPJ("11111111111111"))      // false
	fmt.Println(isValidCNPJ("00.000.000/0000-00"))  // false
	fmt.Println(isValidCNPJ("00000000000000"))      // false
}