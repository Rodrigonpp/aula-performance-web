/*
FALHA A: Escalabilidade
Justificativa: O código foi feito pensando em uma escala menor de usuários simultâneos,
com o aumento drástico de usuários, o código se tornou menos eficiente, gerando carregamentos excessivos.

FALHA B: Confiabilidade
Justificativa: O código deixa de funcionar corretamente quando está diante de adversidades, como é o caso
dos servidores superlotados na black friday, o que pode gerar falhas e vazamento não intencional de dados.

FALHA C: Manutenabilidade
Justificativa: O código não é construído pensando em manutenções futuras, qualquer alteração que precisa ser 
feita acaba demorando horas.
*/

package main

import "fmt"

func main() {
	fmt.Println("Ambiente configurado por: Rodrigo Nunes")
	fmt.Println("Desafio analisado!")
	fmt.Scanln()
}
