def is_valid_cnpj(cnpj: str) -> bool:
    # Remove qualquer caracteres não numéricos
    cnpj = ''.join(filter(str.isdigit, cnpj))
    
    # Verifica se o CNPJ tem 14 dígitos e não é uma sequência de dígitos iguais
    # Exemplo: 11111111111111, 22222222222222, etc.
    if len(cnpj) != 14 or cnpj == cnpj[0] * 14:
        return False

    # Verifica se o CNPJ é válido
    # O CNPJ é composto por 12 dígitos + 2 dígitos verificadores
    # O primeiro dígito verificador é calculado a partir dos 12 primeiros dígitos
    # O segundo dígito verificador é calculado a partir dos 12 primeiros dígitos e do primeiro dígito verificador
    def calculate_check_digit(weights):
        total = sum(int(cnpj[i]) * weights[i] for i in range(len(weights)))
        remainder = total % 11
        return '0' if remainder < 2 else str(11 - remainder)

    # Define os pesos para o cálculo dos dígitos verificadores
    first_weights = [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]
    second_weights = [6] + first_weights

    # Calcula os dígitos verificadores
    check_digit1 = calculate_check_digit(first_weights)
    check_digit2 = calculate_check_digit(second_weights)

    # Verifica se os dígitos verificadores calculados são iguais aos dígitos verificadores do CNPJ
    # Ex.: XX.XXX.XXX/XXXX-95 -> dígitos verificadores são 9 e 5
    return cnpj[-2:] == check_digit1 + check_digit2

# Exemplos de uso

# CNPJs válidos
print(is_valid_cnpj("12.345.678/0001-95"))  # True
print(is_valid_cnpj("12345678000195"))      # True

# CNPJs inválidos
print(is_valid_cnpj("12.345.678/0001-96"))  # False
print(is_valid_cnpj("12345678000196"))      # False
print(is_valid_cnpj("12345678901234"))      # False
print(is_valid_cnpj("1234567890123"))       # False
print(is_valid_cnpj("123456789012345"))     # False
print(is_valid_cnpj("1234567890123a"))      # False
print(is_valid_cnpj("1234567890123-4"))     # False
print(is_valid_cnpj("1234567890123.4"))     # False
print(is_valid_cnpj("11.111.111/1111-11"))  # False
print(is_valid_cnpj("11111111111111"))      # False
print(is_valid_cnpj("00.000.000/0000-00"))  # False
print(is_valid_cnpj("00000000000000"))      # False