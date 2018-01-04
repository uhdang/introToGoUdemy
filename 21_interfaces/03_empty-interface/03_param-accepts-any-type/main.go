package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

type elephant struct {
	animal
	big bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}

	momo := elephant{animal{"booo"}, true}
	fofo := elephant{animal{"mooo"}, false}

	specs(fido)
	specs(fifi)
	specs(momo)
	specs(fofo)

}
