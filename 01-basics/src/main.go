package main

import (
	"fmt"
)

func main() {
	// Buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Unbuffered channel
	ch2 := make(chan int)
	go func() {
		ch2 <- 3
	}()
	y := <-ch2
	fmt.Println(y)
}

// // Inicializar un proyecto
// go mod init path_del_proyecto

// // Verificar que el código externo no esté corrupto
// go mod verify

// // Reemplazar fuente del código
// go mod edit -replace path_del_repo_online=path_del_repo_en_local

// // Quitar el replace
// go mod edit -dropreplace path_del_repo_online

// // Empaquetar todo el código de terceros que usa nuestro código
// go mod vendor

// // Eliminar todos los paquetes externos que no estemos usando
// go mod tidy

// // Aprender más de go modules
// go help mod
