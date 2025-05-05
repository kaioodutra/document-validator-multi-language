function isValidCNPJ(cnpj) {
  // Verifica se o CNPJ é uma string
  if (typeof cnpj !== "string") return false;

  // Remove caracteres não numéricos
  cnpj = cnpj.replace(/[^\d]+/g, "");

  // Verifica se o CNPJ tem 14 dígitos e não é uma sequência de dígitos iguais
  // Exemplo: 11111111111111, 22222222222222, etc.
  if (cnpj.length !== 14 || /^(\d)\1+$/.test(cnpj)) return false;

  // Verifica se o CNPJ é válido
  // O CNPJ é composto por 12 dígitos + 2 dígitos verificadores
  // O primeiro dígito verificador é calculado a partir dos 12 primeiros dígitos
  // O segundo dígito verificador é calculado a partir dos 12 primeiros dígitos e do primeiro dígito verificador
  const calcVerifier = (base, weights) => {
    const sum = base
      .split("")
      .reduce((acc, num, idx) => acc + Number(num) * weights[idx], 0);
    const rest = sum % 11;
    return rest < 2 ? "0" : String(11 - rest);
  };

  // Pega os 12 primeiros dígitos do CNPJ
  const base = cnpj.slice(0, 12);

  // Define os pesos para o cálculo dos dígitos verificadores
  // Calcula os dígitos verificadores
  const firstVerifier = calcVerifier(
    base,
    [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]
  );
  const secondVerifier = calcVerifier(
    base + firstVerifier,
    [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]
  );

  // Verifica se os dígitos verificadores calculados são iguais aos dígitos verificadores do CNPJ
  // Ex.: XX.XXX.XXX/XXXX-95 -> dígitos verificadores são 9 e 5
  return cnpj.endsWith(firstVerifier + secondVerifier);
}

// Exemplos de uso

// CNPJs válidos
console.log(isValidCNPJ("12.345.678/0001-95")); // true
console.log(isValidCNPJ("12345678000195")); // true

// CNPJs inválidos
console.log(isValidCNPJ("12.345.678/0001-96")); // false
console.log(isValidCNPJ("12345678000196")); // false
console.log(isValidCNPJ("12345678901234")); // false
console.log(isValidCNPJ("1234567890123")); // false
console.log(isValidCNPJ("123456789012345")); // false
console.log(isValidCNPJ("1234567890123a")); // false
console.log(isValidCNPJ("1234567890123-4")); // false
console.log(isValidCNPJ("1234567890123.4")); // false
console.log(isValidCNPJ("11.111.111/1111-11")); // false
console.log(isValidCNPJ("11111111111111")); // false
console.log(isValidCNPJ("00.000.000/0000-00")); // false
console.log(isValidCNPJ("00000000000000")); // false
