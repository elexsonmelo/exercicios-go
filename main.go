package main

import (
	"fmt"
	"math/rand"
)

/*
=============================================================================
EXERCÍCIOS DE LÓGICA EM GO
=============================================================================

Tópicos abordados:
- Operadores Lógicos (&&, ||, !)
- Estruturas Condicionais (if, else if, else)
- Loops de Repetição (for)

Instruções:
- Resolva os exercícios na ordem
- Use o Debug (F5) para ver o código executando passo a passo
- Comente os exercícios anteriores para testar o próximo
=============================================================================
*/

// =============================================================================
// BLOCO 1: ESTRUTURAS CONDICIONAIS SIMPLES (IF/ELSE)
// =============================================================================

/*
EXERCÍCIO 1: Par ou Ímpar
Descrição: Crie uma variável com um número inteiro e verifique se ele é par ou ímpar.
Dica: Use o operador % (resto da divisão). Se numero % 2 == 0, é par.
Exemplo:
  - Se numero = 4 → "O número 4 é par"
  - Se numero = 7 → "O número 7 é ímpar"
*/

/*
EXERCÍCIO 2: Maior de Idade
Descrição: Crie uma variável "idade" e verifique se a pessoa é maior de idade (18 anos ou mais).
Exemplo:
  - Se idade = 20 → "Você é maior de idade"
  - Se idade = 15 → "Você é menor de idade"
*/

/*
EXERCÍCIO 3: Classificação de Nota
Descrição: Crie uma variável "nota" (0 a 10) e classifique:
  - nota >= 7: "Aprovado"
  - nota >= 5 e < 7: "Recuperação"
  - nota < 5: "Reprovado"
Dica: Use if, else if, else
*/

/*
EXERCÍCIO 4: Desconto por Valor de Compra
Descrição: Uma loja dá desconto baseado no valor da compra:
  - Compra >= R$ 500: 20% de desconto
  - Compra >= R$ 200: 10% de desconto
  - Compra < R$ 200: sem desconto
Calcule o valor final após o desconto.
Exemplo: Se valor = 600, desconto = 120, valorFinal = 480
*/

// =============================================================================
// BLOCO 2: OPERADORES LÓGICOS (&&, ||, !)
// =============================================================================

/*
EXERCÍCIO 5: Acesso ao Sistema
Descrição: Um usuário só pode acessar o sistema se:
  - Tiver login correto (login == "admin")
  - E tiver senha correta (senha == "1234")
Use o operador && (E lógico)
Exemplo:
  - login = "admin", senha = "1234" → "Acesso permitido"
  - login = "admin", senha = "0000" → "Acesso negado"
*/

/*
EXERCÍCIO 6: Fim de Semana
Descrição: Verifique se um dia da semana é fim de semana.
Um dia é fim de semana se for sábado OU domingo.
Use o operador || (OU lógico)
Exemplo:
  - dia = "sábado" → "É fim de semana!"
  - dia = "segunda" → "Não é fim de semana"
*/

/*
EXERCÍCIO 7: Entrada no Cinema
Descrição: Uma pessoa pode entrar no cinema se:
  - Tiver idade >= 18 anos
  - OU tiver idade >= 12 anos E estar acompanhada (acompanhado == true)
Use && e ||
Exemplo:
  - idade = 20, acompanhado = false → "Pode entrar"
  - idade = 15, acompanhado = true → "Pode entrar"
  - idade = 10, acompanhado = false → "Não pode entrar"
*/

/*
EXERCÍCIO 8: Negação de Condição
Descrição: Verifique se um número NÃO é negativo.
Use o operador ! (NÃO lógico)
Exemplo:
  - numero = 5 → "O número não é negativo"
  - numero = -3 → "O número é negativo"
Dica: if !(numero < 0) ou if numero >= 0
*/

// =============================================================================
// BLOCO 3: LOOPS - FOR SIMPLES
// =============================================================================

/*
EXERCÍCIO 9: Contar de 1 até 10
Descrição: Use um loop for para imprimir os números de 1 até 10.
Dica: for i := 1; i <= 10; i++ { ... }
Saída esperada: 1 2 3 4 5 6 7 8 9 10
*/

/*
EXERCÍCIO 10: Números Pares de 0 a 20
Descrição: Use um loop for para imprimir apenas os números pares de 0 a 20.
Dica: Use i := 0; i <= 20; i = i + 2 OU verifique se i % 2 == 0
Saída esperada: 0 2 4 6 8 10 12 14 16 18 20
*/

/*
EXERCÍCIO 11: Tabuada do 5
Descrição: Imprima a tabuada do 5 (de 5x1 até 5x10).
Saída esperada:
  5 x 1 = 5
  5 x 2 = 10
  5 x 3 = 15
  ...
  5 x 10 = 50
*/

/*
EXERCÍCIO 12: Contagem Regressiva
Descrição: Faça uma contagem regressiva de 10 até 1.
Dica: for i := 10; i >= 1; i-- { ... }
Saída esperada: 10 9 8 7 6 5 4 3 2 1
*/

// =============================================================================
// BLOCO 4: LOOPS COM CONDICIONAIS
// =============================================================================

/*
EXERCÍCIO 13: Números Ímpares de 1 a 20
Descrição: Use um loop for e um if para imprimir apenas números ímpares de 1 a 20.
Saída esperada: 1 3 5 7 9 11 13 15 17 19
*/

/*
EXERCÍCIO 14: Soma dos Números de 1 a 100
Descrição: Use um loop for para calcular a soma de todos os números de 1 a 100.
Dica: Crie uma variável soma := 0 antes do loop
Saída esperada: A soma é: 5050
*/

/*
EXERCÍCIO 15: Contador de Números Positivos
Descrição: Dado um array de números (crie você mesmo),
conte quantos números são positivos (maiores que 0).
Exemplo: numeros := []int{5, -3, 0, 8, -1, 12, 7}
Saída esperada: Existem 4 números positivos
Dica: for i := 0; i < len(numeros); i++ { ... }
*/

/*
EXERCÍCIO 16: FizzBuzz Simplificado
Descrição: Para números de 1 a 15, imprima:
  - "Fizz" se o número for divisível por 3
  - "Buzz" se o número for divisível por 5
  - "FizzBuzz" se for divisível por 3 E por 5
  - O próprio número se não for divisível por 3 nem por 5
Saída esperada:
  1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
*/

// =============================================================================
// BLOCO 5: DESAFIOS (COMBINANDO TUDO)
// =============================================================================

/*
EXERCÍCIO 17: Validador de Senha
Descrição: Crie uma validação de senha que verifica se:
  - A senha tem pelo menos 6 caracteres (use len(senha) >= 6)
  - A senha NÃO é igual a "123456" (senha muito fraca)
Use if, else e operadores lógicos
Exemplo:
  - senha = "abc123" → "Senha válida"
  - senha = "123456" → "Senha muito fraca, escolha outra"
  - senha = "abc" → "Senha muito curta (mínimo 6 caracteres)"
*/

/*
EXERCÍCIO 18: Calculadora de IMC com Classificação
Descrição: Calcule o IMC (Índice de Massa Corporal) e classifique:
  IMC = peso / (altura * altura)

  Classificação:
  - IMC < 18.5: "Abaixo do peso"
  - IMC >= 18.5 e < 25: "Peso normal"
  - IMC >= 25 e < 30: "Sobrepeso"
  - IMC >= 30: "Obesidade"

Exemplo: peso = 70.0, altura = 1.75
  IMC = 70 / (1.75 * 1.75) = 22.86 → "Peso normal"
*/

/*
EXERCÍCIO 19: Jogo de Adivinhação com Tentativas Limitadas
Descrição: Crie um jogo onde:
  - O computador "pensa" em um número (defina você, ex: 7)
  - O usuário tem 3 tentativas para adivinhar
  - A cada tentativa errada, diga se o palpite foi maior ou menor
  - Se acertar, parabenize
  - Se errar as 3 tentativas, mostre o número correto

Dica: Use um for com contador de tentativas
*/

/*
EXERCÍCIO 20: Múltiplos de 3 e 5 até 50
Descrição: Imprima todos os números de 1 a 50 que são:
  - Múltiplos de 3 OU múltiplos de 5
  - Mas não imprima os que são múltiplos de 3 E 5 ao mesmo tempo

Exemplo de saída: 3 5 6 9 10 12 18 20 21 24 25 27 33 35 36 39 40 42 48 50
(Note que 15, 30, 45 NÃO aparecem pois são múltiplos de ambos)

Dica: Use %, && e ||
*/

// =============================================================================
// RESPOSTA DO EXERCÍCIO 1
// =============================================================================

// Função que verifica se um número é par ou ímpar
func verificaParOuImpar(numero int) string { // a função recebe um número inteiro e retorna uma string
	if numero%2 == 0 { // verifica se o resto da divisão do número por 2 é igual a 0, usando o operador % (módulo, ou resto da divisão) e o operador de igualdade ==
		return fmt.Sprintf("O número %d é par", numero) // primeiro retorno da função, para imprimir a variável da string, usei Sprintf, que formata a string e retorna o valor
	} else {
		return fmt.Sprintf("O número %d é ímpar", numero) // segundo retorno da função, caso o número não seja par, retorna que o número é ímpar
	}
}

func main() {
	fmt.Println("=== EXERCÍCIOS DE LÓGICA EM GO ===")
	fmt.Println("Exercício 1: Par ou Ímpar")
	fmt.Println()

	// Para testar, é preciso declarar um número e chamar a função criada:
	numero := 4                             //criei uma variável número com valor inteiro 4, em go há várias maneiras de declaracao de variáveis, aqui usei := que é por inferencia de tipo
	resultado := verificaParOuImpar(numero) // criei uma variável resultado, que recebe o retorno da função, chamando a função e passando a variável número como argumento
	fmt.Println(resultado)                  // imprime o resultado

	// Teste com outro número:
	numero = 7                             // alterei o valor da variável número para 7
	resultado = verificaParOuImpar(numero) // aqui usei a variável resultado para receber o novo retorno da função, atribuindo o novo valor à variável resultado com =
	fmt.Println(resultado)                 // imprime o novo resultado

	// Exercicio 1: Usando números aleatórios (Sugestao do Rafa)
	fmt.Println("\n--- NÚMEROS ALEATÓRIOS ---") // título para resposta do exercício usando números aleatórios

	// Gera e verifica os números aleatórios
	for numero := 0; numero < 5; numero++ { // usei o for para criar um loop que vai executar 5 vezes, ou seja, executar a instrução de código dentro do for 5 vezes
		numeroAleatorio := rand.Intn(10)                // criei a variável numeroAleatorio, que receberá a função rand.Intn, que gera números aleatórios conforme estabelecido, neste caso, de 0 a 10
		resultado = verificaParOuImpar(numeroAleatorio) //aqui chamo a função verificaParOuImpar, passando o número aleatório gerado como argumento, e atribuindo o retorno à variável resultado
		fmt.Println(resultado)                          // imprime o resultado usando o Println, assim cada resultado aparece em uma linha abaixo da outra
	}
}
