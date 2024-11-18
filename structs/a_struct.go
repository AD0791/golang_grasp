package main

import (
	"fmt"
)

// enum pattern
type role int

const (
	messenger role = iota
	receiver  role = iota
)

type Person struct {
	name string
	role
	money int
}

type Message struct {
	emitter Person
	content string
}

func (r role) roleString() string {
	switch r {
	case messenger:
		return "messenger"
	case receiver:
		return "receiver"
	default:
		return "Not a Role"
	}
}

// the pointer allows to mutate the struct
// the original struct will not be uptated
func (p *Person) UpdateMoney(mnt int) {
	p.money = mnt
}

// here:
// the original struct will be updated. no copy of the struct is done
func (p *Person) StrictUpdateMoney(mnt int) {
	(*p).money = mnt
}

func main() {
	mathieu := Person{name: "Mathieu Monroe", role: messenger, money: 0}
	originalPerson := Person{name: "Zeus son of cronus", role: receiver, money: 0}
	message := Message{emitter: mathieu, content: "free the republic"}
	// now we can have access to orginalPerson
	pointOriginalPerson := &originalPerson
	// will copy mathieu in memory
	mathieu.UpdateMoney(100)
	// we wont have a copy in memory
	originalPerson.StrictUpdateMoney(500)

	fmt.Println("The message from " + mathieu.name + " is clear: " + message.content + " !")
	fmt.Println("The emitter was from "+message.emitter.name+" is role was", message.emitter.role.roleString())

	fmt.Println("somme de Mathieu: ", mathieu.money)
	fmt.Printf("%+v \n", mathieu)

	fmt.Println("the original person (no copy in memory) money is: ", originalPerson.money)
	fmt.Printf("%+v \n", originalPerson)

	fmt.Println(pointOriginalPerson)
	fmt.Println(pointOriginalPerson.roleString())

	pointOriginalPerson.money = 1000
	fmt.Println(pointOriginalPerson.money)

	// we will see the exact changes
	fmt.Println(originalPerson)
	fmt.Println(originalPerson.money)
}
