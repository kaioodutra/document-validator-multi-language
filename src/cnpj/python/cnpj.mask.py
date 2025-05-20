import re

def mask_cnpj(cnpj):
    # Verifica se o CNPJ é uma string e não está vazio
    if not isinstance(cnpj, str) or cnpj.strip() == "":
        return False

    # Remove espaços em branco
    cnpj = cnpj.strip()

    # Se já estiver no formato com máscara
    masked_regex = re.compile(r'^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$')
    if masked_regex.fullmatch(cnpj):
        return cnpj

    # Se for exatamente 14 dígitos numéricos
    unmasked_regex = re.compile(r'^\d{14}$')
    if unmasked_regex.fullmatch(cnpj):
        return f"{cnpj[:2]}.{cnpj[2:5]}.{cnpj[5:8]}/{cnpj[8:12]}-{cnpj[12:]}"

    # Qualquer outro caso
    return False

# Exemplos de uso

# CNPJs válidos
print(mask_cnpj("12345678000195"))         # "12.345.678/0001-95"
print(mask_cnpj("12.345.678/0001-95"))     # "12.345.678/0001-95"
print(mask_cnpj("    12345678000195 "))    # "12.345.678/0001-95"

# CNPJs inválidos
print(mask_cnpj("123.456.780/001-95"))     # False
print(mask_cnpj("1a2.3b4.5c6/7d8-9e"))      # False
print(mask_cnpj("12345678"))               # False
print(mask_cnpj("abc12345678000195"))      # False
print(mask_cnpj(12345678000195))           # False
print(mask_cnpj(True))                     # False
print(mask_cnpj(False))                    # False
print(mask_cnpj(None))                     # False