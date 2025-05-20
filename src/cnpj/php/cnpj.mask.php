<?php

function maskCNPJ($cnpj) {
  // Verifica se o CNPJ é uma string e não está vazio
  if (!is_string($cnpj) || trim($cnpj) === '') {
    return false;
  }

  // Remove espaços em branco
  $cnpj = trim($cnpj);

  // Se já estiver no formato com máscara
  if (preg_match('/^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$/', $cnpj)) {
    return $cnpj;
  }

  // Se for exatamente 14 dígitos numéricos
  if (preg_match('/^\d{14}$/', $cnpj)) {
    return sprintf(
      '%s.%s.%s/%s-%s',
      substr($cnpj, 0, 2),
      substr($cnpj, 2, 3),
      substr($cnpj, 5, 3),
      substr($cnpj, 8, 4),
      substr($cnpj, 12, 2)
    );
  }

  // Qualquer outro caso
  return false;
}

// Exemplos de uso

// CNPJs válidos
$validCNPJs = [
  "12345678000195", // true
  "12.345.678/0001-95", // true
  "    12345678000195 " //true
];

// CNPJs inválidos 
$invalidCNPJs = [
  "123.456.780/001-95", // false
  "1a2.3b4.5c6/7d8-9e", // false
  "abc12345678000195", // false
  "12345678", // false
  "" // false
];

// Testando a função com os CNPJs válidos
echo "// CNPJs válidos\n";
foreach ($validCNPJs as $cnpj) {
    $result = maskCNPJ($cnpj);
    echo $cnpj . ' → ' . ($result !== false ? $result : 'false') . "\n";
}

// Testando a função com os CNPJs inválidos
echo "\n// CNPJs inválidos\n";
foreach ($invalidCNPJs as $cnpj) {
    $result = maskCNPJ($cnpj);
    echo $cnpj . ' → ' . ($result !== false ? $result : 'false') . "\n";
}