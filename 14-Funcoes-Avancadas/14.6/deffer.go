package main

import "fmt"

/*
Casos de uso	Descrição
Fechando arquivos =>	Garantir que o arquivo seja fechado assim que a função retornar, mesmo se a função retornar devido a um erro.
Desalocando memória	=>	Garantir que a memória seja desalocada assim que a função retornar, independentemente de como a função retorna.
Fechando conexões de rede	=>	Garantir que a conexão seja fechada assim que a função retornar, independentemente de como a função retorna.
Fechando recursos de banco de dados =>		Garantir que a conexão seja fechada assim que a função retornar, independentemente de como a função retorna.
Fechando semáforos	=>	 Garantir que o semáforo seja liberado assim que a função retornar, independentemente de como a função retorna.
*/
func vaiSerExecutadoPorUltimo() {
	fmt.Println("Defer Fn é sempre a última a ser executada!")
}

func main() {
	defer vaiSerExecutadoPorUltimo()

	fmt.Println("Sou uma Fn normal, e serei executada primeiro embora declarada posteriormente")
}
