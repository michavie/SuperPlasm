package main

import (
	fr "SuperPlasm/Super_LiqPool_Farm"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"log"
	"os"
	"time"
)

func main() {

	//Snapshotting SUPER Amounts
	fmt.Println("")
	fmt.Println("Snapshotting SUPER Amounts ...")
	Start0 := time.Now()
	SuperChain := fr.CreateSuperChain()
	Elapsed0 := time.Since(Start0)
	fmt.Println("Done snapshotting Super Amounts, time required", Elapsed0)
	//Printing Snapshot
	fmt.Println("Outputting SUPER Amounts to SUPER-chain.txt")
	OutputFile0, err := os.Create("SUPER-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile0.Close()
	_, _ = fmt.Fprintln(OutputFile0, SuperChain)
	fmt.Println("DONE Outputting SUPER Amounts to SUPER-chain.txt")
	fmt.Println("")

	//Snapshotting LP Amounts
	fmt.Println("")
	fmt.Println("Snapshotting LP Amounts ...")
	Start1 := time.Now()
	LPChain := fr.CreateLPChain()
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done snapshotting LP Amounts, time required", Elapsed1)
	//Printing Snapshot
	fmt.Println("Outputting LP Amounts to LP-chain.txt")
	OutputFile1, err := os.Create("LP-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile1.Close()
	_, _ = fmt.Fprintln(OutputFile1, LPChain)
	fmt.Println("DONE Outputting LP Amounts to LP-chain.txt")
	fmt.Println("")

	//Snapshotting Camel Amount
	fmt.Println("Snapshotting Camel Amounts ...")
	Start2 := time.Now()
	CamelChain := fr.CreateCamelChain()
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done snapshotting Camel Amounts, time required", Elapsed2)
	//Printing Snapshot
	fmt.Println("Outputting Camel Amounts to Camel-Chain.txt")
	OutputFile2, err := os.Create("Camel-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile2.Close()
	_, _ = fmt.Fprintln(OutputFile2, CamelChain)
	fmt.Println("DONE Outputting Camel Amounts to Camel-Chain.txt")
	fmt.Println("")

	//Computing  Virtual LP Chain
	fmt.Println("Start creating the Virtual LP Chain")
	Start3 := time.Now()
	VLPChain := fr.CreateVLPChain(LPChain, CamelChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done creating the Virtual LP Chain, time required", Elapsed3)
	//Printing Virtual LP Chain
	fmt.Println("Outputting Virtual LP Chain to VLP-Chain.txt")
	OutputFile3, err := os.Create("VLP-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile3.Close()
	_, _ = fmt.Fprintln(OutputFile3, VLPChain)
	fmt.Println("DONE Outputting Virtual LP Chain to VLP-Chain.txt")
	fmt.Println("")

	//Computing Rewards
	fmt.Println("Starting computing Rewards considering", fr.SuperFarmRewardAmount, "per day.")
	Start4 := time.Now()
	Reward := fr.SuperRewardComputer(VLPChain, p.NFI(fr.SuperFarmRewardAmount))
	Elapsed4 := time.Since(Start4)
	fmt.Println("Done computing Rewards, time required", Elapsed4)
	//Outputting Rewards
	fmt.Println("Outputting Rewards Chain to Reward-Chain.txt")
	OutputFile4, err := os.Create("Reward-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile4.Close()
	_, _ = fmt.Fprintln(OutputFile4, Reward)
	fmt.Println("DONE Outputting Rewards Chain to Reward-Chain.txt")
	fmt.Println("")

	//Computing SUPER Power
	fmt.Println("Start creating the Super-Power Chain")
	Start5 := time.Now()
	SuperPowerChain := fr.CreateSuperPowerChain(SuperChain, LPChain)
	Elapsed5 := time.Since(Start5)
	fmt.Println("Done creating the SUPER Power Chain, time required", Elapsed5)
	//Printing SUPER Power Chain
	fmt.Println("Outputting Super Power Chain to SUPER-Power-Chain.txt")
	OutputFile5, err := os.Create("SUPER-Power-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile5.Close()
	_, _ = fmt.Fprintln(OutputFile5, SuperPowerChain)
	fmt.Println("DONE Outputting Super Power Chain Chain to SUPER-Power-Chain.txt")
	fmt.Println("")

	//Computing SUPER Power percent
	fmt.Println("Start creating the Super-Power-Percent Chain")
	Start6 := time.Now()
	SuperPowerPercentChain := fr.SuperPowerPercentComputer(SuperPowerChain)
	Elapsed6 := time.Since(Start6)
	fmt.Println("Done creating the SUPER Power Percent Chain, time required", Elapsed6)
	//Printing SUPER Power Percent Chain
	fmt.Println("Outputting Super Power Percent Chain to SUPER-Power-Percent-Chain.txt")
	OutputFile6, err := os.Create("SUPER-Power-Percent-Chain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile6.Close()
	_, _ = fmt.Fprintln(OutputFile6, SuperPowerPercentChain)
	fmt.Println("DONE Outputting Super Power Chain Percent Chain to SUPER-Power-Percent-Chain.txt")
	fmt.Println("")

	fmt.Println("There are ", len(LPChain), "addresses that have LP")
	fmt.Println("There are ", len(CamelChain), "addresses that have Camels")
	fmt.Println("There are only ", len(VLPChain), "addresses that are eligible for rewards")
	fmt.Println("There are only ", len(SuperPowerChain), "addresses that have SuperPower")
}
