<?php

function isValidCNPJ(string $cnpj): bool
{
    // Remove caracteres não numéricos
    $cnpj = preg_replace('/\D/', '', $cnpj);

    // Verifica se o CNPJ tem 14 dígitos e não é uma sequência de dígitos iguais
    // Exemplo: 11111111111111, 22222222222222, etc.
    if (strlen($cnpj) !== 14 || preg_match('/^(\d)\1{13}$/', $cnpj)) {
        return false;
    }

    // Verifica se o CNPJ é válido
    // O CNPJ é composto por 12 dígitos + 2 dígitos verificadores
    // O primeiro dígito verificador é calculado a partir dos 12 primeiros dígitos
    // O segundo dígito verificador é calculado a partir dos 12 primeiros dígitos e do primeiro dígito verificador
    $calculateCheckDigit = function (array $weights) use ($cnpj) {
        $sum = 0;
        foreach ($weights as $i => $weight) {
            $sum += intval($cnpj[$i]) * $weight;
        }

        $remainder = $sum % 11;
        return $remainder < 2 ? '0' : strval(11 - $remainder);
    };

    // Os pesos para o primeiro e segundo dígitos verificadores
    $firstWeights = [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];
    $secondWeights = [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];

    // Calcula os dígitos verificadores
    $digit1 = $calculateCheckDigit($firstWeights);
    $digit2 = $calculateCheckDigit($secondWeights);

    // Verifica se os dígitos verificadores calculados são iguais aos dígitos verificadores do CNPJ
    // Ex.: XX.XXX.XXX/XXXX-95 -> dígitos verificadores são 9 e 5
    return substr($cnpj, -2) === $digit1 . $digit2;
}

// Exemplos de uso

// CNPJs válidos
$validCNPJs = [
    "12.345.678/0001-95", // true
    "12345678000195", // true
];
// CNPJs inválidos 
$invalidCNPJs = [
    "12.345.678/0001-96", // false
    "12345678000196", // false
    "12345678901234", // false
    "1234567890123", // false
    "123456789012345", // false
    "1234567890123a", // false
    "1234567890123-4", // false
    "1234567890123.4", // false
    "11.111.111/1111-11", // false
    "11111111111111", // false
    "00.000.000/0000-00", // false
    "00000000000000", // false
];
// Testando a função com os CNPJs válidos
// Exibindo os resultados
foreach ($validCNPJs as $cnpj) {
    echo (isValidCNPJ($cnpj) ? "true" : "false") . "\n";
}
// Testando a função com os CNPJs inválidos
// Exibindo os resultados
foreach ($invalidCNPJs as $cnpj) {
    echo (isValidCNPJ($cnpj) ? "true" : "false") . "\n";
}
