package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

var (
	clicks        = [3]int{}
	currentColumn = 0
	people        People
)

type Person struct {
	name   string
	age    int
	gender string
}

func NewPerson(name string, age int, gender string) *Person {
	person := &Person{name: name, age: age, gender: gender}
	return person
}

func click(col int) {
	clicks[col]++
	idx := func(clicks [3]int) int {
		idx := 0
		for i := range clicks {
			if clicks[idx] < clicks[i] {
				idx = i
			}
		}
		return idx
	}(clicks)
	currentColumn = idx
	fmt.Printf("Sorting with column %d\n", idx)
	sort.Sort(people)
}

type People []*Person

func (p People) Len() int      { return len(p) }
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p People) Less(i, j int) bool {
	if currentColumn == 0 {
		return p[i].name < p[j].name
	}
	if currentColumn == 1 {
		return p[i].age < p[j].age
	}
	if currentColumn == 2 {
		return p[i].gender < p[j].gender
	}
	return false
}

func printPeople(p *People) {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 2, 2, ' ', 0)
	const format = "%v\t%v\t%v\t\n"
	fmt.Fprintf(tw, format, "name", "age", "gender")
	fmt.Fprintf(tw, format, "----", "---", "------")
	for _, person := range people {
		fmt.Fprintf(tw, format, person.name, person.age, person.gender)
	}
	tw.Flush()
}

func main() {
	people = People{
		NewPerson("mohamed", 25, "male"),
		NewPerson("mohamed", 24, "male"),
		NewPerson("alaa", 26, "male"),
		NewPerson("alaa", 26, "female"),
	}
	click(1)
	click(1)
	click(2)
	click(2)
	printPeople(&people)
}
