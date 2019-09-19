package main

import (
	"fmt"
	"strings"
)

//Programme permettant de changer le mot par un autres mot
func processLine(line, old, new string) (found bool, res string, occ int) {
	oldlower := strings.ToLower(old)
	newlower := strings.ToLower(new)
	oldupper := strings.ToUpper(old)
	newupper := strings.ToUpper(new)
	res = line
	occ = 0
	if strings.Contains(line, old) || strings.Contains(line, oldlower) {
		found = true
		occ += strings.Count(line, old)
		res = strings.Replace(line, old, new, -1)
		occ += strings.Count(line, oldlower)
		res = strings.Replace(res, oldlower, newlower, -1)
	} else {
		found = true
		occ += strings.Count(line, oldupper)
		res = strings.Replace(line, oldupper, newupper, -1)

	}
	return found, res, occ

}

func main() {
	old := "go"
	new := "Python"
	old = old + " "
	new = new + " "
	a, b, c := processLine("Ce code Go se comporte comme go annoncé, mais les []bytepoints renvoyés Go sont placés dans un go tableau contenant le fichier entier Go", old, new)
	fmt.Println(a, c)
	fmt.Println("__________________________________________")
	fmt.Println(b)
}
