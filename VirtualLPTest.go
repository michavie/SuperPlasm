package main

import (
	fr "SuperPlasm/Super_LiqPool_Farm"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"log"
	"os"
)

func main() {

	//Snapshoting LP and Camel Amounts and Creating the VLP Chain

	LPChain := fr.CreateLPChain()
	CamelChain := fr.CreateCamelChain()
	VLPChain := fr.CreateVLPChain(LPChain, CamelChain)
	Reward := fr.SuperRewardComputer(VLPChain, p.NFI(fr.SuperFarmRewardAmount))

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

	OutputFile4, err := os.Create("Reward-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile4.Close()
	_, _ = fmt.Fprintln(OutputFile4, Reward)

	fmt.Println("There are ", len(LPChain), "addresses that have LP")
	fmt.Println("There are ", len(CamelChain), "addresses that have Camels")
	fmt.Println("There are only ", len(VLPChain), "addresses that are eligible for rewards")
}
