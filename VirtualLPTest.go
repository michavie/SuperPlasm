package main

import (
	fr "SuperPlasm/Super_LiqPool_Farm"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"log"
	"os"
)

func main() {
	var (
		LpAmount string
		CA       string
	)

	//Virtual LP Computer

	fmt.Println("Enter LP amount")
	_, _ = fmt.Scanln(&LpAmount)
	LP := p.NFS(LpAmount)

	fmt.Println("Camel Amount")
	_, _ = fmt.Scanln(&CA)
	CAmount := p.NFS(CA)

	VLP := fr.VirtualLP(LP, CAmount)
	fmt.Println("Your Base    LP is:", LP)
	fmt.Println("Your Virtual LP is:", VLP)

	//Snapshoting LP and Camel Amounts and Creating the VLP Chain

	LPChain := fr.CreateLPChain()
	CamelChain := fr.CreateCamelChain()
	VLPChain := fr.CreateVLPChain(LPChain, CamelChain)

	OutputFile1, err := os.Create("LP-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile1.Close()
	_, _ = fmt.Fprintln(OutputFile1, LPChain)

	OutputFile2, err := os.Create("Camel-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile2.Close()
	_, _ = fmt.Fprintln(OutputFile2, CamelChain)

	OutputFile3, err := os.Create("VLP-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile3.Close()
	_, _ = fmt.Fprintln(OutputFile3, VLPChain)

	fmt.Println("There are only ", len(VLPChain), "addresses that are eligible for rewards")
}
