package go_koans

import "io/ioutil"
import "strings"
import "fmt"

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	// assert("func aboutFiles() {" == "func aboutFiles() {") // handling files is too trivial
	// assert(1 == 2)
	for i, _ := range lines {
		fmt.Println("line", i, ":", lines[i])
	}

	fmt.Println(lines[6])
	fmt.Println("func aboutFiles() {")
	for k, _ := range lines[6] {
		v1 := lines[6][k]
		v2 := "func aboutFiles() {\r"[k]
		if v1 == v2 {
			fmt.Printf("match! (%s, %s)\n", v1, v2)
		} else {
			fmt.Printf("no match! (%s, %s)\n", v1, v2)
		}

	}

	fmt.Println(lines[6] == "func aboutFiles() {")
	fmt.Println("func aboutFiles() {" == lines[6])
	assert(lines[6] == "func aboutFiles() {\r") // handling files is too trivial
}
