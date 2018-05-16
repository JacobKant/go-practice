package main

import "fmt"

type Person struct {
	Name string
	inn  string
}

type SecretAgent struct {
	Person        // встраивание структуры (типо вместо наследования)
	LicenseToKill bool
}

func (s SecretAgent) GetName() string {
	return "CLASSIFIED"
}

func main() {
	sa := SecretAgent{Person: Person{"James", "12312321321"}, LicenseToKill: true}
	fmt.Println("secret inn", sa.GetName())
}
