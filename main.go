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

	exercicio1()
	exercicio2()
	exercicio3()
	exercicio5()
	exercicio6()
}

func exercicio1() {
	fmt.Println("\n--- EXERCÍCIO 1: Par ou Ímpar ---")
	numero := 4
	fmt.Println(verificaParOuImpar(numero))
	numero = 7
	fmt.Println(verificaParOuImpar(numero))

	// Exemplo com vários números (loop)
	fmt.Println("--- VERIFICANDO VÁRIOS NÚMEROS ---")
	numeros := []int{2, 3, 4, 5, 6}
	for _, n := range numeros {
		fmt.Println(verificaParOuImpar(n))
	}
	fmt.Println()

	// Para testar, é preciso declarar um número e chamar a função criada:
	numero = 4                              //alterei o valor da variável número para 4
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

// =============================================================================
// RESPOSTA DO EXERCÍCIO 2: Maior de Idade
// =============================================================================
func exercicio2() {
	fmt.Println("\nExercício 2: Maior de Idade")
	fmt.Println()

	//Primeiro, vamos criar as variáveis e chamar a função verificaMaiorIdade:
	idade := 20                            // criei a variável idade com valor 20 (cenário maior de idade)
	resultado := verificaMaiorIdade(idade) // criei a variável resultado que recebe o retorno da função verificaMaiorIdade, passando a variável idade como argumento
	fmt.Println(resultado)                 // imprime o resultado da verificação de idade

	idade = 15                            // criei a variável idade com valor 15 (cenário menor de idade)
	resultado = verificaMaiorIdade(idade) // chamei a função novamente, passando o novo valor da variável idade, e atribuindo o retorno à variável resultado
	fmt.Println(resultado)                // imprime o resultado da verificação de idade
}

func verificaMaiorIdade(idade int) string { // A idéia desse exercício é criar uma função que recebe a idade como um número inteiro e retorna uma string com a verificação se é maior ou menor de idade
	if idade >= 18 { // O if verifica se a idade é maior ou igual a 18
		return "Você é maior de idade" // Caso sim, retorna a string "Você é maior de idade"
	} else {
		return "Você é menor de idade" // Se não, retorna a string "Você é menor de idade"
	}
}

// =============================================================================
// RESPOSTA DO EXERCÍCIO 3: Classificação de Nota
// =============================================================================
func exercicio3() {
	fmt.Println("\nExercício 3: Classificação de Nota")
	fmt.Println()

	// Vamos criar as variáveis e chamar a função verificaClassificacaoNota:
	nota := 8.5                     // criei a variável nota com valor 8.5 (cenário de nota alta)
	resultado := verificaNota(nota) // criei a variável resultado que recebe o retorno da função verificaClassificacaoNota, passando a variável nota como argumento
	fmt.Println(resultado)          // imprime o resultado da verificação de nota

	nota = 5.0 // criei a variável nota com valor 5.0 (cenário de nota baixa)
	resultado = verificaNota(nota)
	fmt.Println(resultado)
}

func verificaNota(nota float64) string { // A idéia desse exercício é criar uma função que recebe uma nota, precisa ser do tipo float, e retorna uma string com a classificação da nota
	if nota >= 7.0 { // O primeiro if verifica se a nota é maior ou igual a 7.0
		return "Aprovado" // Caso sim, retorna a string "Aprovado"
	} else if nota >= 5.0 { // O segundo if verifica se a nota é maior ou igual a 5.0, como precisaremos verificar 3 cenarios, usei o else if aqui
		return "Recuperação" // Caso sim, retorna a string "Recuperação"
	} else { // Se a nota não for maior ou igual a 7.0, nem maior ou igual a 5.0, cai no else
		return "Reprovado" // Retorna a string "Reprovado"
	}
}

// =============================================================================
// RESPOSTA DO EXERCÍCIO 4: Desconto por Valor de Compra
// =============================================================================

func exercicio4() {
	fmt.Println("--- EXERCÍCIO 4: DESCONTO POR VALOR DE COMPRA ---")

	valores := []float64{600.0, 300.0, 150.0} // Para criar as variáveis, há algumas maneiras, uma delas é criar um slice com os valores de compra a serem testados, achei uma forma bacana de usar

	for _, valor := range valores { // aqui usei o for para percorrer o slice de valores, usando o _ para ignorar o índice e pegar apenas o valor
		desconto, valorFinal := calcularDesconto(valor)                            // chamei a função calcularDesconto, passando o valor da compra como argumento, e recebi dois retornos: desconto e valorFinal.
		fmt.Printf("Compra: R$ %.2f | Desconto: R$ %.2f | Valor Final: R$ %.2f\n", // aqui imprimo o resultado formatado com Printf, usando %.2f para formatar os valores como float com 2 casas decimais, o \n no final é para pular uma linha após a impressão
			valor, desconto, valorFinal) // dando continuidade à linha anterior, passo as variáveis valor, desconto e valorFinal para serem impressas na string formatada
	}

	fmt.Println() // linha em branco para separar a saída, deixa mais organizada a resposta
}

func calcularDesconto(valorCompra float64) (float64, float64) { // a função calcularDesconto recebe o valor da compra como float64 e retorna dois valores float64: o desconto e o valor final após o desconto
	var desconto float64 // criei a variável desconto do tipo float64 para armazenar o valor do desconto calculado

	if valorCompra >= 500 { // primeiro if verifica se o valor da compra é maior ou igual a 500
		desconto = valorCompra * 0.20 // se sim, calcula o desconto de 20% e atribui à variável desconto, usei o operador * para calcular a porcentagem, ou seja, multiplica o valor da compra por 0.20
	} else if valorCompra >= 200 { // segundo if verifica se o valor da compra é maior ou igual a 200
		desconto = valorCompra * 0.10 // se sim, calcula o desconto de 10% e atribui à variável desconto
	} else {
		desconto = 0 // se o valor da compra for menor que 200, não há desconto, então atribui 0 à variável desconto, ou seja, sem desconto
	}

	valorFinal := valorCompra - desconto //aqui calculo o valor final subtraindo o desconto do valor da compra, e atribuo à variável valorFinal

	return desconto, valorFinal // aqui o ponto importante: a vantagem da função ter dois retornos, posso retornar tanto o valor do desconto quanto o valor final após o desconto
}

// =============================================================================
// RESPOSTA DO EXERCÍCIO 5: Acesso ao Sistema
// =============================================================================

// exercicio5 agrupa a solução do Exercício 5 (Acesso ao Sistema com Operador &&)
func exercicio5() {
	fmt.Println("\n--- EXERCÍCIO 5: Acesso ao Sistema ---")

	// Teste 1: login correto e senha correta
	login := "admin"                           // criei a variável login com valor "admin"
	senha := "1234"                            // criei a variável senha com valor "1234"
	fmt.Println(verificarAcesso(login, senha)) // chamo a função verificarAcesso passando login e senha como argumentos

	// Teste 2: login correto mas senha incorreta
	login = "admin"                            // login continua correto
	senha = "0000"                             // mas agora a senha é incorreta
	fmt.Println(verificarAcesso(login, senha)) // chamo a função novamente com a senha errada

	// Teste 3: login incorreto e senha correta
	login = "user"                             // agora o login é incorreto
	senha = "1234"                             // mas a senha está correta
	fmt.Println(verificarAcesso(login, senha)) // função retorna acesso negado pois o login está errado

	// Teste 4: login incorreto e senha incorreta
	login = "user"                             // login incorreto
	senha = "0000"                             // senha também incorreta
	fmt.Println(verificarAcesso(login, senha)) // ambas incorretas, acesso negado
	fmt.Println()
}

// verificarAcesso recebe login e senha e retorna mensagem de acesso
func verificarAcesso(login string, senha string) string { // função recebe dois argumentos: login e senha, ambos strings, e retorna uma string
	if login == "admin" && senha == "1234" { // usa o operador && (E lógico) para verificar SE o login é "admin" E a senha é "1234"
		return "Acesso permitido" // SE ambas as condições forem verdadeiras, retorna "Acesso permitido"
	} else { // CASO CONTRÁRIO
		return "Acesso negado" // retorna "Acesso negado" se qualquer uma das condições for falsa
	}
}

// =============================================================================
// RESPOSTA DO EXERCÍCIO 6: Fim de Semana
// =============================================================================

// exercicio6 agrupa a solução do Exercício 6 (Fim de Semana com Operador ||)
func exercicio6() {
	fmt.Println("--- EXERCÍCIO 6: Fim de Semana ---")

	// Teste 1: verifica se é sábado
	dia := "sábado"                        // criei a variável dia com valor "sábado"
	fmt.Println(verificarFimDeSemana(dia)) // chamo a função verificarFimDeSemana passando "sábado"

	// Teste 2: verifica se é domingo
	dia = "domingo"                        // agora o dia é domingo
	fmt.Println(verificarFimDeSemana(dia)) // chamo a função com "domingo"

	// Teste 3: verifica se é segunda
	dia = "segunda"                        // agora o dia é segunda (dia útil)
	fmt.Println(verificarFimDeSemana(dia)) // função retorna que não é fim de semana

	// Teste 4: verifica se é quarta
	dia = "quarta"                         // quarta também é dia útil
	fmt.Println(verificarFimDeSemana(dia)) // função retorna que não é fim de semana

	// Teste 5: verifica se é sexta
	dia = "sexta"                          // sexta é dia útil (começa o fim de semana só no sábado)
	fmt.Println(verificarFimDeSemana(dia)) // função retorna que não é fim de semana
	fmt.Println()
}

// verificarFimDeSemana recebe um dia e verifica se é sábado ou domingo
func verificarFimDeSemana(dia string) string { // função recebe um argumento: dia (string), e retorna uma string
	if dia == "sábado" || dia == "domingo" { // usa o operador || (OU lógico) para verificar SE o dia é "sábado" OU "domingo"
		return "É fim de semana!" // SE qualquer uma das condições for verdadeira, retorna "É fim de semana!"
	} else { // CASO CONTRÁRIO
		return "Não é fim de semana" // retorna "Não é fim de semana" se nenhuma das condições for verdadeira
	}
}
