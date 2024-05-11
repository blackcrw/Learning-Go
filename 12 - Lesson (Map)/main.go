package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapCourses := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	courses := ""

	for courses != "q" {
		var workload int

		fmt.Print("Digite um curso ou digite 'q' para sair: ")
		scanner.Scan()
		courses = scanner.Text()

		if courses != "q" {
			fmt.Print("Digite a carga horária do curso: ")
			fmt.Scanf("%d", &workload)
			mapCourses[courses] = workload
		}
	}
	fmt.Println("Cursos Registrados: ")

	for nameCourse, workload := range mapCourses {
		fmt.Printf(" — %s: %dh \n", nameCourse, workload)
	}
}
