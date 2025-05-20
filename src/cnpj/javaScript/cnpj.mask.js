function maskCNPJ(cnpj) {
  // Verifica se o CNPJ é uma string e não está vazio
  if (typeof cnpj !== "string" || cnpj === "") return false;

  // Remove espaços em branco
  cnpj = cnpj.trim();
  if (cnpj === "") return false;

  // Se já estiver no formato com máscara
  const maskedRegex = /^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$/;
  if (maskedRegex.test(cnpj)) return cnpj;

  // Se for exatamente 14 dígitos numéricos
  const unmaskedRegex = /^\d{14}$/;
  if (unmaskedRegex.test(cnpj)) {
    return (
      cnpj.slice(0, 2) +
      "." +
      cnpj.slice(2, 5) +
      "." +
      cnpj.slice(5, 8) +
      "/" +
      cnpj.slice(8, 12) +
      "-" +
      cnpj.slice(12)
    );
  }

  // Qualquer outro caso
  return false;
}

// Exemplos de uso

// Válidos
console.log(maskCNPJ("12345678000195")); // "12.345.678/0001-95"
console.log(maskCNPJ("12.345.678/0001-95")); // "12.345.678/0001-95"
console.log(maskCNPJ("    12345678000195   ")); // "12.345.678/0001-95"

// Inválidos
console.log(maskCNPJ("123.456.780/001-95")); // false
console.log(maskCNPJ("1a2.3b4.5c6/7d8-9e")); // false
console.log(maskCNPJ("abc12345678000195")); // false
console.log(maskCNPJ(12345678000195)); // false
console.log(maskCNPJ(true)); // false
console.log(maskCNPJ(false)); // false
console.log(maskCNPJ(null)); // false
console.log(maskCNPJ(undefined)); // false
