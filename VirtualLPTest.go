package main

import (
	fr "SuperPlasm/Super_LiqPool_Farm"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"log"
	"os"
	"time"
)

func SnapshooterPrinterSuper() []fr.Super {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER Amounts ...")
	Start := time.Now()
	SuperChain := fr.CreateSuperChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER Amounts to Chain_SUPER.txt")
	OutputFile, err := os.Create("Chain_SUPER.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SuperChain)
	fmt.Println("DONE Outputting SUPER Amounts to SUPER-chain.txt")
	fmt.Println("")

	return SuperChain
}

func SnapshooterPrinterSuperLP() []fr.SuperLP {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER-LP Amounts ...")
	Start := time.Now()
	LPChain := fr.CreateSuperLPChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER-LP Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-LP Amounts to Chain_SUPER-LP.txt")
	OutputFile, err := os.Create("Chain_SUPER-LP.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, LPChain)
	fmt.Println("DONE Outputting SUPER Amounts to SUPER-chain.txt")
	fmt.Println("")

	return LPChain
}

func SnapshooterPrinterSuperCamel() []fr.CamelAmount {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER-Camel Amounts ...")
	Start := time.Now()
	CamelChain := fr.CreateCamelChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER-Camel Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-Camel Amounts to Chain_SUPER-Camel.txt")
	OutputFile, err := os.Create("Chain_SUPER-Camel.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, CamelChain)
	fmt.Println("DONE Outputting SUPER Amounts to SUPER-chain.txt")
	fmt.Println("")

	return CamelChain
}

func SnapshooterPrinterSuperVirtualLP(Chain1 []fr.SuperLP, Chain2 []fr.CamelAmount) []fr.SuperVLP {
	fmt.Println("")
	fmt.Println("Computing  SUPER-VLP Amounts ...")
	Start := time.Now()
	VLPChain := fr.CreateVLPChain(Chain1, Chain2)
	Elapsed := time.Since(Start)
	fmt.Println("Done computing SUPER-VLP Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-VLP Amounts to Chain_SUPER-VLP.txt")
	OutputFile, err := os.Create("Chain_SUPER-VLP.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, VLPChain)
	fmt.Println("DONE Outputting SUPER-VLP Amounts to Chain_SUPER-VLP.txt")
	fmt.Println("")

	return VLPChain
}

func SnapshooterPrinterSuperVirtualLPRewards(Chain1 []fr.SuperVLP, Reward int64) []fr.SuperFarmReward {
	fmt.Println("")
	fmt.Println("Computing SUPER-VLP Rewards considering ", Reward, " per day")
	Start := time.Now()
	RewardChain := fr.SuperRewardComputer(Chain1, p.NFI(Reward))
	Elapsed := time.Since(Start)
	fmt.Println("Done computing SUPER-VLP Rewards, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-VLP-Rewards to Chain_SUPER-VLP-Rewards.txt")
	OutputFile, err := os.Create("Chain_SUPER-VLP-Rewards.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, RewardChain)
	fmt.Println("DONE Outputting SUPER-VLP Amounts to Chain_SUPER-VLP.txt")
	fmt.Println("")

	return RewardChain
}

func SnapshooterPrinterSuperPower(Chain1 []fr.Super, Chain2 []fr.SuperLP) []fr.SuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing  SUPER-Power Chain ...")
	Start1 := time.Now()
	SuperPowerChain := fr.CreateSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing  SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	SuperPowerPercentChain := fr.SuperPowerPercentComputer(SuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting  SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedSuperPowerPercentChain := fr.SortSuperPowerPercent(SuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting  SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted SUPER-Power-Percent-Chain to Chain_SUPER-Power-Percent-sorted.txt")
	OutputFile, err := os.Create("Chain_SUPER-Power-Percent-sorted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SortedSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted SUPER-Power-Percent-Chain to Chain_SUPER-Power-Percent-sorted.txt")
	fmt.Println("")

	return SortedSuperPowerPercentChain

}

func SnapshooterPrinterKosonicSuperPower(Chain1 []fr.Super, Chain2 []fr.SuperLP) []fr.SuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing  Kosonic SUPER-Power Chain ...")
	Start1 := time.Now()
	KosonicSuperPowerChain := fr.CreateKosonicSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  Kosonic SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing  Kosonic SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	KosonicSuperPowerPercentChain := fr.SuperPowerPercentComputer(KosonicSuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  Kosonic SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting  Kosonic SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedKosonicSuperPowerPercentChain := fr.SortSuperPowerPercent(KosonicSuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting Kosonic SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted Kosonic SUPER-Power-Percent-Chain to Chain_Kosonic-SUPER-Power-Percent-sorted.txt")
	OutputFile, err := os.Create("Chain_Kosonic-SUPER-Power-Percent-sorted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SortedKosonicSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted Kosonic SUPER-Power-Percent-Chain to Chain_Kosonic-SUPER-Power-Percent-sorted.txt")
	fmt.Println("")

	return SortedKosonicSuperPowerPercentChain

}

func SnapshooterPrinterMetaKosonicSuperPower(Chain1 []fr.Super, Chain2 []fr.SuperLP) []fr.MKSuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing  Meta-Kosonic SUPER-Power Chain ...")
	Start1 := time.Now()
	MetaKosonicSuperPowerChain := fr.CreateMetaKosonicSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  Meta-Kosonic-SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing  Meta-Kosonic SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	MetaKosonicSuperPowerPercentChain := fr.MKSuperPowerPercentComputer(MetaKosonicSuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  Meta-Kosonic SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting  Meta-Kosonic SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedMetaKosonicSuperPowerPercentChain := fr.SortMKSuperPowerPercent(MetaKosonicSuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting Meta-Kosonic SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted MEta-Kosonic SUPER-Power-Percent-Chain to Chain_MetaKosonic-SUPER-Power-Percent-sorted.txt")
	OutputFile, err := os.Create("Chain_MetaKosonic-SUPER-Power-Percent-sorted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SortedMetaKosonicSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted Meta-Kosonic SUPER-Power-Percent-Chain to Chain_MetaKosonic-SUPER-Power-Percent-sorted.txt")
	fmt.Println("")

	return SortedMetaKosonicSuperPowerPercentChain

}

func main() {
	SuperFarmRewardAmount := int64(1665)

	SuperChain := SnapshooterPrinterSuper()
	SuperLPChain := SnapshooterPrinterSuperLP()
	CamelChain := SnapshooterPrinterSuperCamel()
	VLPChain := SnapshooterPrinterSuperVirtualLP(SuperLPChain, CamelChain)
	SnapshooterPrinterSuperVirtualLPRewards(VLPChain, SuperFarmRewardAmount)

	SortedSuperPowerChain := SnapshooterPrinterSuperPower(SuperChain, SuperLPChain)
	SnapshooterPrinterKosonicSuperPower(SuperChain, SuperLPChain)

	//SortedMKSuperPowerChain := SnapshooterPrinterMetaKosonicSuperPower(SuperChain, SuperLPChain)
	SnapshooterPrinterMetaKosonicSuperPower(SuperChain, SuperLPChain)

	fmt.Println("")
	fmt.Println("======RESULTS======")
	fmt.Println("There are ", len(SuperLPChain), "addresses that have LP")
	fmt.Println("There are ", len(CamelChain), "addresses that have Camels")
	fmt.Println("There are only ", len(VLPChain), "addresses that are eligible for LP Rewards")
	fmt.Println("There are only ", len(SortedSuperPowerChain), "addresses that have SuperPower")

}
