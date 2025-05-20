<h1 align="center">
  Validador de Documentos Multilinguagens
</h1>

---

<p align="center">
  Sabemos que existem inúmeras libs (bibliotecas, extensões etc.) para validação de documentos. No entanto, dependendo do tamanho do projeto e da necessidade de validações documentais, talvez não seja necessário implementar uma dependência externa. Uma simples função em sua árvore de arquivos, como por exemplo <code>src/utils/validators/cnpj.validator.ts</code>, já pode ser o suficiente.<br/>
  Pensando nisso, criamos este repositório com funções avulsas que lidam com documentos, escritas em diferentes linguagens — simples, direto e reutilizável.
  <br/>
  <br/>
  <img src="https://img.shields.io/badge/--00ADD8?style=flat&logo=go&logoColor=white" alt="Go" style="width:32px; height:auto"/>
  <img src="https://img.shields.io/badge/--F7DF1E?style=flat&logo=javascript&logoColor=000000" alt="JavaScript" style="width:32px; height:auto"/>
  <img src="https://img.shields.io/badge/--777BB4?style=flat&logo=php&logoColor=white" alt="PHP" style="width:32px; height:auto"/>
  <img src="https://img.shields.io/badge/--3776AB?style=flat&logo=python&logoColor=white" alt="Python" style="width:32px; height:auto"/>
  <img src="https://img.shields.io/badge/--3178C6?style=flat&logo=typescript&logoColor=white" alt="TypeScript" style="width:32px; height:auto"/>
</p>

---

## Features highlights

> Funções organizadas por categoria e implementadas em múltiplas linguagens.  
> Todas seguem o mesmo padrão de entrada, validação e retorno, com testes e comentários didáticos.

| Funções       | ![Go](https://img.shields.io/badge/--00ADD8?style=flat&logo=go&logoColor=ffffff) | ![JavaScript](https://img.shields.io/badge/--F7DF1E?style=flat&logo=javascript&logoColor=000000) | ![PHP](https://img.shields.io/badge/--777BB4?style=flat&logo=php&logoColor=ffffff) | ![Python](https://img.shields.io/badge/--3776AB?style=flat&logo=python&logoColor=ffffff) | ![TypeScript](https://img.shields.io/badge/--3178C6?style=flat&logo=typescript&logoColor=ffffff) |                Retorno                |
| ------------- | :------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------: | :-----------------------------------: |
| `maskCNPJ`    |                                        ✓                                         |                                                ✓                                                 |                                         ✓                                          |                                            ✓                                             |                                                ✓                                                 | string com máscara ou `false` ou `""` |
| `isValidCNPJ` |                                        ✓                                         |                                                ✓                                                 |                                         ✓                                          |                                            ✓                                             |                                                ✓                                                 |           `true` ou `false`           |
| `unMaskCNPJ`  |                                        ✘                                         |                                                ✘                                                 |                                         ✘                                          |                                            ✘                                             |                                                ✘                                                 |                   -                   |

---

## Como usar

A implementação ocorre de forma simples em qualquer linguagem chamando apenas a função, exemplo em **TypeScript**:

- **isValidCNPJ** - Verifica se um CNPJ é Válido

```ts
import { isValidCNPJ } from "./cnpj.validator.ts";

const cnpj = isValidCNPJ("12.345.678/0001-95");
console.log(cnpj); //true
```

O repositório hoje conta com funções para: **Go**, **JavaScript**, **PHP**, **Python** e **TypeScript**. Você poderá verificar quais funções existem hoje a cima em **Features highlights**.

---

## Contribuição

##### INSTRUÇÕES

- Clone o repositório.

```bash
git clone git@github.com:kaioodutra/multi-language-document-validator.git
```

- Crie uma nova branch.

```bash
git checkout -b feature/nome-da-sua-feature
```

- Crie a funçao na pasta da linguagem ou crie a pasta da nova linguagem.

```
src
└── cnpj
    ├── go
    │   └── cnpj.mask.go
    │   └── cnpj.validator.go
    └── java  👈 //nova pasta
    │   └── cnpj.mask.java  👈 //nova função mask
    │   └── cnpj.unmask.java  👈 //nova função unMask
    │   └── cnpj.validator.java  👈 //nova função validator
    ├── javaScript
    │   └── cnpj.mask.js
    │   └── cnpj.validator.js
    ├── php
    │   └── cnpj.mask.php
    │   └── cnpj.validator.php
    ├── python
    │   └── cnpj.mask.py
    │   └── cnpj.validator.py
    └── typeScript
        └── cnpj.mask.ts
        └── cnpj.validator.ts
```

<br/>

##### COMMIT

- Use `feat:` para nova função.

```bash
git commit -m "feat: add cnpj validator and mask java"
```

- Use `fix:` para correção de bugs.

```bash
git commit -m "fix: update cnpj validator typeScript"
```

- Use `refactor:` para refatoração de código.

```bash
git commit -m "refactor: update cnpj validator python"
```

- Use `perf:` para refatoração de performance de código.

```bash
git commit -m "perf: update cnpj validator php"
```

<br/>

##### PUSH

- Envie para seu fork

```bash
git push origin feature/nome-da-sua-feature
```

<br/>

##### PULL REQUEST

- Abra um Pull Request explicando o que foi feito.
  1- Vá até seu repositório no GitHub (ex: github.com/seu-usuario/seu-fork)
  2- O GitHub vai mostrar um botão "Compare & Pull Request". Clique nele.
  3- Escreva um título e descrição explicando o que foi feito.
  4- Clique em "Create Pull Request".

---

### Licença

Este é um utilitário **aberto e gratuito**, livre para uso, modificação e distribuição.

<br/>

_Criado com 💙 e muito ☕️ por desenvolvedores apaixonados por código limpo e padrões brasileiros._

###### "Simplicidade é a sofisticação máxima."
