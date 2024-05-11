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

		fmt.Print("Enter a course or type 'q' to quit: ")
		scanner.Scan()
		courses = scanner.Text()

		if courses != "q" {
			fmt.Print("Enter the course workload: ")
			fmt.Scanf("%d", &workload)
			mapCourses[courses] = workload
		}
	}
	fmt.Println("Registered Courses: ")

	for nameCourse, workload := range mapCourses {
		fmt.Printf(" — %s: %dh \n", nameCourse, workload)
	}
}
