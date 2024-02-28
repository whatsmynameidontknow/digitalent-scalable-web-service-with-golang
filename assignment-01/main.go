package main

import (
	"assignment-01/generator"
	"fmt"
	"os"
	"strconv"
)

const (
	start int = 1
	n     int = 200
)

func init() {
	absenPerson = make(map[noAbsen]person, n)
	for i := range n {
		absenPerson[noAbsen(i+start)] = person{
			name:       generator.GenerateRandomString(15+generator.R.IntN(16), 1+generator.R.IntN(3)),    // 15 - 30 chars, 1 - 3 words
			address:    generator.GenerateRandomString(25+generator.R.IntN(26), 4+generator.R.IntN(5)),    // 25- 50 chars,  4 - 8 words
			occupation: generator.GenerateRandomString(10+generator.R.IntN(11), 1+generator.R.IntN(2)),    // 10 - 20 chars,  1 - 2 words
			reason:     generator.GenerateRandomString(150+generator.R.IntN(101), 10+generator.R.IntN(6)), // 150 - 250 chars,  10 - 15 words
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("masukin nomor absen, pls (assignment-01_XXX <no-absen>). jgn ditambah/kurangin, y.")
		os.Exit(1)
	}
	noAbsenUint64, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("nomor absen cuma bisa angka positif, y.")
		os.Exit(1)
	}
	noAbsen := noAbsen(noAbsenUint64)
	if person, ok := absenPerson[noAbsen]; !ok {
		fmt.Printf("gaada orang yg nomor absennya %d (no absen cuma dari %d - %d).\n", noAbsen, start, start+n-1)
	} else {
		fmt.Println(person)
	}
}

type noAbsen uint64

var absenPerson map[noAbsen]person

type person struct {
	name       string
	address    string
	occupation string
	reason     string
}

// implement Stringer interface
func (p person) String() string {
	return fmt.Sprintf("Nama\t\t\t\t: %s\nAlamat\t\t\t\t: %s\nPekerjaan\t\t\t: %s\nAlasan memilih kelas Golang\t: %s", p.name, p.address, p.occupation, p.reason)
}
