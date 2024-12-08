package main

import (
	"fmt"
	"time"
)

type Human struct {
	Age      int
	Name     string
	Height   int
	passport string //Анонимное поле
}

type Action struct {
	Human      //Встраиваемая структура без имени, что позволяет нам использовать синтаксические сокращения
	RunSpeed   int
	sleepHours time.Duration
}

func (a *Human) Sayage() {
	fmt.Println(a.Age)
}

func (a *Action) ActionSleep(d int) {
	a.sleepHours = time.Duration(d) * time.Hour
	fmt.Println("suggested time to sleep is:", a.sleepHours.Hours(), "hours")
}

func main() {
	var man = Human{22, "Глеб", 185, "3329:661322"}
	man.Sayage()

	var Gleb Action
	Gleb.Sayage() //Встраивание метода от родительской структуры

	Gleb.Age = 33
	Gleb.ActionSleep(8)
	Gleb.Sayage()
	fmt.Println(Gleb.Age)

}
