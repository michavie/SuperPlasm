package main

import (
	fr "SuperPlasm/Super_LiqPool_Farm"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

func main() {
	var (
		LpAmount string
		CA       string
	)

	fmt.Println("Enter LP amount")
	_, _ = fmt.Scanln(&LpAmount)
	LP := p.NFS(LpAmount)

	fmt.Println("Camel Amount")
	_, _ = fmt.Scanln(&CA)
	CAmount := p.NFS(CA)

	VLP := fr.VirtualLP(LP, CAmount)
	fmt.Println("Your Base    LP is:", LP)
	fmt.Println("Your Virtual LP is:", VLP)
}
