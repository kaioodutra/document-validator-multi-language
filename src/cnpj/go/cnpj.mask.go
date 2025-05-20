package main

import (
	"fmt"
	"regexp"
	"strings"
)

func maskCNPJ(cnpj string) string {
	// Remove espaços em branco
	cnpj = strings.TrimSpace(cnpj)

	// Se já estiver no formato com máscara
	maskedPattern := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	if maskedPattern.MatchString(cnpj) {
		return cnpj
	}

	// Se for exatamente 14 dígitos numéricos
	unmaskedPattern := regexp.MustCompile(`^\d{14}$`)
	if unmaskedPattern.MatchString(cnpj) {
		return fmt.Sprintf("%s.%s.%s/%s-%s",
			cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
	}

	// Qualquer outro caso
	return ""
}

func main() {
	// Exemplo de uso

	// Válidos
	fmt.Println(maskCNPJ("12345678000195")); // "12.345.678/0001-95"
	fmt.Println(maskCNPJ("12.345.678/0001-95")); // "12.345.678/0001-95"
	fmt.Println(maskCNPJ("    12345678000195   ")); // "12.345.678/0001-95"

	// Inválidos
	fmt.Println(maskCNPJ("123.456.780/001-95")); // ""
	fmt.Println(maskCNPJ("1a2.3b4.5c6/7d8-9e")); // ""
	fmt.Println(maskCNPJ("abc12345678000195")); // ""
	fmt.Println(maskCNPJ("12345678")); // ""
	fmt.Println(maskCNPJ("")); // ""
}