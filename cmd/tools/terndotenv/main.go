package main

import (
	"os/exec"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern", 
		"migrate", 
		"--migrations", 
		"./internal/store/pgstore/migrations", 
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// if err := cmd.Run(); err != nil {
	// 	panic(err)
	// }
	output, err := cmd.CombinedOutput()
	if err != nil {
			fmt.Println("Erro ao executar o comando:")
			fmt.Println(string(output))
			panic(err)
	}

	fmt.Println("Comando executado com sucesso:")
	fmt.Println(string(output))
}
