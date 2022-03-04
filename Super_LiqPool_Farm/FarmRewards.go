package Super_LiqPool_Farm

import (
	mt "SuperPlasm/SuperMath"
	p "github.com/Crypt0plasm/Firefly-APD"
)

var (
	SuperEgldLPDecimals = int64(18)
)

func ConvertAtomicUnits(Number string) *p.Decimal {
	Value := p.NFS(Number)
	Result := mt.DIVxc(Value, mt.POWxc(p.NFI(10), p.NFI(SuperEgldLPDecimals)))
	return Result
}

//Returns the MetaKosonicSuperPowerPromille used to Compute the Meta-Kosonic-Super-Power
func MetaKosonicSuperPowerPromille(SuperAmount *p.Decimal) *p.Decimal {
	var (
		ProMille = new(p.Decimal)

		T00 = p.NFI(0)
		T01 = p.NFI(10)
		T02 = p.NFI(25)
		T03 = p.NFI(50)
		T04 = p.NFI(75)
		T05 = p.NFI(100)
		T06 = p.NFI(250)
		T07 = p.NFI(500)
		T08 = p.NFI(750)
		T09 = p.NFI(1000)
		T10 = p.NFI(2500)
		T11 = p.NFI(5000)
		T12 = p.NFI(7500)
		T13 = p.NFI(10000)
		T14 = p.NFI(25000)
		T15 = p.NFI(50000)
		T16 = p.NFI(75000)
		T17 = p.NFI(100000)
		T18 = p.NFI(250000)
		T19 = p.NFI(500000)
		T20 = p.NFI(750000)
		T21 = p.NFI(1000000)
		T22 = p.NFI(2500000)
		T23 = p.NFI(5000000)
		T24 = p.NFI(7500000)
		STS = p.NFI(10000000) //SuperTotalSupply.

		TP01 = p.NFI(10000)
		TP02 = p.NFI(8000)
		TP03 = p.NFI(7000)
		TP04 = p.NFI(6000)
		TP05 = p.NFI(5000)
		TP06 = p.NFI(4000)
		TP07 = p.NFI(3000)
		TP08 = p.NFI(2000)
		TP09 = p.NFI(1000)

		TP10 = mt.DIVxc(TP02, p.NFI(10)) //800
		TP11 = mt.DIVxc(TP03, p.NFI(10)) //700
		TP12 = mt.DIVxc(TP04, p.NFI(10)) //600
		TP13 = mt.DIVxc(TP05, p.NFI(10)) //500
		TP14 = mt.DIVxc(TP06, p.NFI(10)) //400
		TP15 = mt.DIVxc(TP07, p.NFI(10)) //300
		TP16 = mt.DIVxc(TP08, p.NFI(10)) //200
		TP17 = mt.DIVxc(TP09, p.NFI(10)) //100

		TP18 = mt.DIVxc(TP10, p.NFI(10)) //80
		TP19 = mt.DIVxc(TP11, p.NFI(10)) //70
		TP20 = mt.DIVxc(TP12, p.NFI(10)) //60
		TP21 = mt.DIVxc(TP13, p.NFI(10)) //50
		TP22 = mt.DIVxc(TP14, p.NFI(10)) //40
		TP23 = mt.DIVxc(TP15, p.NFI(10)) //30
		TP24 = mt.DIVxc(TP16, p.NFI(10)) //20
		TP25 = mt.DIVxc(TP17, p.NFI(10)) //10

		TPW0 = p.NFI(2000)
		TPW1 = p.NFI(1000)
		TPW2 = mt.DIVxc(TPW0, p.NFI(10)) //200
		TPW3 = mt.DIVxc(TPW1, p.NFI(10)) //100
		TPW4 = mt.DIVxc(TPW2, p.NFI(10)) //20
		TPW5 = mt.DIVxc(TPW3, p.NFI(10)) //10
	)

	//Rounding SuperAmount to IntegerValue
	SA := mt.TruncateCustom(SuperAmount, 0)

	ProMilleFunc := func(StartThreshold, EndThreshold, TopPromille, WidthPromille, Value *p.Decimal) *p.Decimal {
		var (
			V1, V2, V3, V4, PM = new(p.Decimal), new(p.Decimal), new(p.Decimal), new(p.Decimal), new(p.Decimal)
		)
		V1 = mt.SUBxc(Value, StartThreshold)
		V2 = mt.SUBxc(EndThreshold, StartThreshold)
		V3 = mt.MULxc(V1, WidthPromille)
		V4 = mt.DIVxc(V3, V2)
		PM = mt.SUBxc(TopPromille, V4)
		return PM
	}

	if mt.DecimalGreaterThan(SA, T00) == true && mt.DecimalLessThanOrEqual(SA, T01) == true {
		ProMille = ProMilleFunc(T00, T01, TP01, TPW0, SA)
	} else if mt.DecimalGreaterThan(SA, T01) == true && mt.DecimalLessThanOrEqual(SA, T02) == true {
		ProMille = ProMilleFunc(T01, T02, TP02, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T02) == true && mt.DecimalLessThanOrEqual(SA, T03) == true {
		ProMille = ProMilleFunc(T02, T03, TP03, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T03) == true && mt.DecimalLessThanOrEqual(SA, T04) == true {
		ProMille = ProMilleFunc(T03, T04, TP04, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T04) == true && mt.DecimalLessThanOrEqual(SA, T05) == true {
		ProMille = ProMilleFunc(T04, T05, TP05, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T05) == true && mt.DecimalLessThanOrEqual(SA, T06) == true {
		ProMille = ProMilleFunc(T05, T06, TP06, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T06) == true && mt.DecimalLessThanOrEqual(SA, T07) == true {
		ProMille = ProMilleFunc(T06, T07, TP07, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T07) == true && mt.DecimalLessThanOrEqual(SA, T08) == true {
		ProMille = ProMilleFunc(T07, T08, TP08, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T08) == true && mt.DecimalLessThanOrEqual(SA, T09) == true {
		ProMille = ProMilleFunc(T08, T09, TP09, TPW2, SA)
	} else if mt.DecimalGreaterThan(SA, T09) == true && mt.DecimalLessThanOrEqual(SA, T10) == true {
		ProMille = ProMilleFunc(T09, T10, TP10, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T10) == true && mt.DecimalLessThanOrEqual(SA, T11) == true {
		ProMille = ProMilleFunc(T10, T11, TP11, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T11) == true && mt.DecimalLessThanOrEqual(SA, T12) == true {
		ProMille = ProMilleFunc(T11, T12, TP12, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T12) == true && mt.DecimalLessThanOrEqual(SA, T13) == true {
		ProMille = ProMilleFunc(T12, T13, TP13, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T13) == true && mt.DecimalLessThanOrEqual(SA, T14) == true {
		ProMille = ProMilleFunc(T13, T14, TP14, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T14) == true && mt.DecimalLessThanOrEqual(SA, T15) == true {
		ProMille = ProMilleFunc(T14, T15, TP15, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T15) == true && mt.DecimalLessThanOrEqual(SA, T16) == true {
		ProMille = ProMilleFunc(T15, T16, TP16, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T16) == true && mt.DecimalLessThanOrEqual(SA, T17) == true {
		ProMille = ProMilleFunc(T16, T17, TP17, TPW4, SA)
	} else if mt.DecimalGreaterThan(SA, T17) == true && mt.DecimalLessThanOrEqual(SA, T18) == true {
		ProMille = ProMilleFunc(T17, T18, TP18, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T18) == true && mt.DecimalLessThanOrEqual(SA, T19) == true {
		ProMille = ProMilleFunc(T18, T19, TP19, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T19) == true && mt.DecimalLessThanOrEqual(SA, T20) == true {
		ProMille = ProMilleFunc(T19, T20, TP20, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T20) == true && mt.DecimalLessThanOrEqual(SA, T21) == true {
		ProMille = ProMilleFunc(T20, T21, TP21, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T21) == true && mt.DecimalLessThanOrEqual(SA, T22) == true {
		ProMille = ProMilleFunc(T21, T22, TP22, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T22) == true && mt.DecimalLessThanOrEqual(SA, T23) == true {
		ProMille = ProMilleFunc(T22, T23, TP23, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T23) == true && mt.DecimalLessThanOrEqual(SA, T24) == true {
		ProMille = ProMilleFunc(T23, T24, TP24, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T24) == true && mt.DecimalLessThanOrEqual(SA, STS) == true {
		ProMille = ProMilleFunc(T24, STS, TP25, TPW5, SA)
	}

	PM := mt.TruncateCustom(ProMille, 0)
	return PM
}

//Returns Bonus Given By LP Tier
func LpTierProcent(LP *p.Decimal) *p.Decimal {
	var (
		TierBonus = new(p.Decimal)
		T1        = p.NFI(1)
		T2        = p.NFI(2)
		T3        = p.NFI(5)
		T4        = p.NFI(10)
		T5        = p.NFI(20)
		T6        = p.NFI(50)
		T7        = p.NFI(100)

		NoBonus = p.NFI(0)
		BonusT1 = p.NFI(1)
		BonusT2 = p.NFI(2)
		BonusT3 = p.NFI(3)
		BonusT4 = p.NFI(4)
		BonusT5 = p.NFI(5)
		BonusT6 = p.NFI(6)
		BonusT7 = p.NFI(7)
	)

	if mt.DecimalLessThan(LP, T1) == true {
		TierBonus = NoBonus
	} else if mt.DecimalGreaterThanOrEqual(LP, T1) == true && mt.DecimalLessThan(LP, T2) == true {
		TierBonus = BonusT1
	} else if mt.DecimalGreaterThanOrEqual(LP, T2) == true && mt.DecimalLessThan(LP, T3) == true {
		TierBonus = BonusT2
	} else if mt.DecimalGreaterThanOrEqual(LP, T3) == true && mt.DecimalLessThan(LP, T4) == true {
		TierBonus = BonusT3
	} else if mt.DecimalGreaterThanOrEqual(LP, T4) == true && mt.DecimalLessThan(LP, T5) == true {
		TierBonus = BonusT4
	} else if mt.DecimalGreaterThanOrEqual(LP, T5) == true && mt.DecimalLessThan(LP, T6) == true {
		TierBonus = BonusT5
	} else if mt.DecimalGreaterThanOrEqual(LP, T6) == true && mt.DecimalLessThan(LP, T7) == true {
		TierBonus = BonusT6
	} else {
		TierBonus = BonusT7
	}

	return TierBonus
}

// Returns Bonus Given by Camel Ownership
func CamelProcent(Camel *p.Decimal) *p.Decimal {
	var (
		CamelTotalBonus = new(p.Decimal)
		Bonus           = p.NFI(10)
	)

	if mt.DecimalGreaterThanOrEqual(Camel, p.NFI(1)) == true {
		CamelTotalBonus = Bonus
	} else {
		CamelTotalBonus = p.NFI(1)
	}
	//CamelTotalBonus = mt.MULxc(Camel, Bonus)
	return CamelTotalBonus

}

//Returns Weighted Bonus Given By Camel Bonus and SuperLP Bonus
func WeightBonus(TB, CB *p.Decimal) *p.Decimal {
	P1 := mt.ADDxc(p.NFI(1), mt.DIVxc(TB, p.NFI(100)))
	P2 := mt.ADDxc(p.NFI(1), mt.DIVxc(CB, p.NFI(100)))
	TP := mt.MULxc(P1, P2)

	return TP
}

//Returns the Virtual LP Amount, which is the Super LP Amount weighted by the Weight Bonus
func VirtualLP(LpAmount, CamelAmount *p.Decimal) *p.Decimal {
	TierBonus := LpTierProcent(LpAmount)
	CamelBonus := CamelProcent(CamelAmount)
	Weight := WeightBonus(TierBonus, CamelBonus)

	Result := mt.MULxc(LpAmount, Weight)
	return mt.TruncateCustom(Result, 18)
}

// Computes Super Power : Super * SuperLP
func SuperPowerComputer(SuperAmount, LPAmount *p.Decimal) *p.Decimal {
	var SP = new(p.Decimal)
	if mt.DecimalLessThanOrEqual(LPAmount, p.NFS("1")) == true {
		if mt.DecimalLessThan(SuperAmount, p.NFS("1")) == true {
			SP = p.NFS("0")
		} else {
			SP = mt.TruncateCustom(SuperAmount, 0)
		}
	} else {
		SP1 := mt.MULxc(SuperAmount, LPAmount)
		if mt.DecimalGreaterThan(SP1, p.NFS("1")) == true {
			SP = SP1
		} else {
			SP = p.NFS("0")
		}
	}

	return mt.TruncateCustom(SP, 0)
}

// Computes Kosonic Super Power : Super * log(2,SuperLP)
func KosonicSuperPowerComputer(SuperAmount, LPAmount *p.Decimal) *p.Decimal {
	var KSP = new(p.Decimal)
	LBase := p.NFS("2")
	if mt.DecimalLessThanOrEqual(LPAmount, LBase) == true {
		if mt.DecimalLessThan(SuperAmount, p.NFS("1")) == true {
			KSP = p.NFS("0")
		} else {
			KSP = SuperAmount
		}
	} else {
		KosonicLP := mt.TruncateCustom(mt.Logarithm(LBase, LPAmount), 18)
		SP1 := mt.MULxc(SuperAmount, KosonicLP)
		if mt.DecimalGreaterThan(SP1, p.NFS("1")) == true {
			KSP = SP1
		} else {
			KSP = p.NFS("0")
		}
	}

	return mt.TruncateCustom(KSP, 0)
}

//Computes Rewards earned by VirtualLP using a given Amount of Reward per Day
func SuperRewardComputer(Chain1 []SuperVLP, RewardAmount *p.Decimal) []SuperFarmReward {
	var (
		VLPSum     = new(p.Decimal)
		FinalChain []SuperFarmReward
	)
	for i := 0; i < len(Chain1); i++ {
		VLPSum = mt.ADDxc(VLPSum, Chain1[i].VLP)
	}
	for i := 0; i < len(Chain1); i++ {
		Reward := mt.TruncateCustom(mt.DIVxc(mt.MULxc(Chain1[i].VLP, RewardAmount), VLPSum), 18)
		Unit := SuperFarmReward{Chain1[i].Address, Reward}
		FinalChain = append(FinalChain, Unit)
	}
	return FinalChain
}

//Kosonic SUPER POWER Percent Computer and Percent Sorter
//Returns the individual Percents of Super Power
func SuperPowerPercentComputer(Chain []SuperPower) []SuperPowerPercent {
	var (
		SPSum      = new(p.Decimal)
		FinalChain []SuperPowerPercent
	)
	for i := 0; i < len(Chain); i++ {
		SPSum = mt.ADDxc(SPSum, Chain[i].SuperPower)
	}
	for i := 0; i < len(Chain); i++ {
		Percent := mt.TruncateCustom(mt.DIVxc(mt.MULxc(Chain[i].SuperPower, p.NFS("100")), SPSum), 18)
		Unit := SuperPowerPercent{SuperPower{Chain[i].Address, Chain[i].SuperPower}, Percent}
		FinalChain = append(FinalChain, Unit)
	}
	return FinalChain
}

//Sorts SuperPowerPercent Chain based on Percent
func SortSuperPowerPercent(Chain []SuperPowerPercent) []SuperPowerPercent {
	var (
		SortedChain []SuperPowerPercent
	)
	GetMaxElement := func(Chain []SuperPowerPercent) int {
		Max := 0
		for i := 0; i < len(Chain)-2; i++ {
			if mt.DecimalGreaterThanOrEqual(Chain[i].SuperPowerPercent, Chain[Max].SuperPowerPercent) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := SuperPowerPercent{Chain2Sort[Biggest].Main, Chain2Sort[Biggest].SuperPowerPercent}
		SortedChain = append(SortedChain, Unit)

		//Removing biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)

	}
	return SortedChain
}

//Meta Kosonic SUPER POWER Percent Computer and Percent Sorter
//Returns the individual Percents of Super Power
func MKSuperPowerPercentComputer(Chain []MKSuperPower) []MKSuperPowerPercent {
	var (
		SPSum      = new(p.Decimal)
		FinalChain []MKSuperPowerPercent
	)
	for i := 0; i < len(Chain); i++ {
		SPSum = mt.ADDxc(SPSum, Chain[i].MetaKosonicSuperPower)
	}
	for i := 0; i < len(Chain); i++ {
		Percent := mt.TruncateCustom(mt.DIVxc(mt.MULxc(Chain[i].MetaKosonicSuperPower, p.NFS("100")), SPSum), 18)
		Unit := MKSuperPowerPercent{MKSuperPower{Chain[i].Address, Chain[i].Super, Chain[i].MetaSuper, Chain[i].MetaKosonicSuperPower}, Percent}
		FinalChain = append(FinalChain, Unit)
	}
	return FinalChain
}

//Sorts SuperPowerPercent Chain based on Percent
func SortMKSuperPowerPercent(Chain []MKSuperPowerPercent) []MKSuperPowerPercent {
	var (
		SortedChain []MKSuperPowerPercent
	)
	GetMaxElement := func(Chain []MKSuperPowerPercent) int {
		Max := 0
		for i := 0; i < len(Chain)-2; i++ {
			if mt.DecimalGreaterThanOrEqual(Chain[i].MetaKosonicSuperPowerPercent, Chain[Max].MetaKosonicSuperPowerPercent) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := MKSuperPowerPercent{Chain2Sort[Biggest].Main, Chain2Sort[Biggest].MetaKosonicSuperPowerPercent}
		SortedChain = append(SortedChain, Unit)

		//Removing biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)

	}
	return SortedChain
}
