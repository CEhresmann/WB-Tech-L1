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

func (a *Action) ActionSleep(d int) string {
	a.sleepHours = time.Duration(d) * time.Hour
	return fmt.Sprintf("suggested time to sleep is: %.0f hours", a.sleepHours.Hours())
}

func (a *Action) AgedSleep(d int) string {
	a.sleepHours = time.Duration(d) * time.Hour
	return fmt.Sprintf("suggested time to sleep is: %.0f hours for people %d years old", a.sleepHours.Hours(), a.Age)
}

func main() {
	var man = Human{22, "Глеб", 185, "3329:661322"}
	man.Sayage()

	var Gleb Action
	Gleb.Sayage() //Использование метода от родительской структуры

	Gleb.Age = 33
	fmt.Println(Gleb.ActionSleep(8))
	Gleb.Sayage()
	fmt.Println(Gleb.Age)
	fmt.Println(Gleb.AgedSleep(6)) //Встраивание метода от родительской структуры

}
