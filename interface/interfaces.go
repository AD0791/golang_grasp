package main

import (
	"errors"
	"fmt"
)

type ip int

var process_ip ip

type english string
type french string
type creole string

type talk interface {
	english() english
	french() french
	creole() creole
}

type person struct {
	name string
	cob  string
}

func (*person) english() english {
	return "hi"
}

func (*person) french() french {
	return "Bonjour"
}

func (*person) creole() creole {
	return "pale"
}

func comm(c talk, langue string) (string, error) {
	switch langue {
	case "english":
		return string(c.english()), nil
	case "french":
		return string(c.french()), nil
	case "creole":
		return string(c.creole()), nil
	default:
		return "", errors.New("Invalid language")
	}
}

func main() {
	// work with pointer
	process_ip = 3333
	test_point := &process_ip

	fmt.Println(process_ip.MutateProcessIP())
	fmt.Println(test_point)
	fmt.Println(*test_point)

	*test_point = 50
	fmt.Println(process_ip.MutateProcessIP())

	// "a" will not change process_ip
	a := process_ip
	a = 10
	fmt.Println(a)
	fmt.Println(process_ip.MutateProcessIP())

	//work in interfaces
	pablo := person{name: "Pablo Hernandez", cob: "Mexico"}

	/*
		no need
		fmt.Println(pablo)
		fmt.Println(pablo.creole())
	*/
	message, err_message := comm(&pablo, "english")
	if err_message != nil {
		fmt.Println("Error:", err_message)
	} else {
		fmt.Println("Message:", message)
	}

	mess, err_mess := comm(&pablo, "french")
	if err_mess != nil {
		fmt.Println("Error:", err_mess)
	} else {
		fmt.Println("Mess:", mess)
	}

	mes, err_mes := comm(&pablo, "creole")
	if err_mes != nil {
		fmt.Println("Error:", err_mes)
	} else {
		fmt.Println("Mes:", mes)
	}

}

func (pip ip) MutateProcessIP() ip {
	return pip * 100
}
