package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número será sorteado. Tente acertar. O número é um inteiro entre 1 e 100.")

	x := rand.Int64N(101) // Numero aleatorio entre 0 e 100
	x = 10
	scanner := bufio.NewScanner((os.Stdin)) // Lendo o que o usuário escreve
	chutes := [10]int64{}                   // Criando um array de 10 posições para guardar os chutes do usuário

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()          // passando o que eu escrevi para variavel que criamos chute em texto
		chute = strings.TrimSpace(chute) // Removendo os espaços em branco

		chuteInt, err := strconv.ParseInt(chute, 10, 64) // Estamos convertendo o chute que é uma string para um inteiro de 64 bits

		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! Você acertou! O número era: %d\n"+
					"Você acertou em %d tentativas\n"+
					"Essas foram as suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt // Guardando o chute do usuário no array
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram as suas tentativas: %v\n",
		x, chutes,
	)
}
