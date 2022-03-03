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

func WeightBonus(TB, CB *p.Decimal) *p.Decimal {
	P1 := mt.ADDxc(p.NFI(1), mt.DIVxc(TB, p.NFI(100)))
	P2 := mt.ADDxc(p.NFI(1), mt.DIVxc(CB, p.NFI(100)))
	TP := mt.MULxc(P1, P2)

	return TP
}

func VirtualLP(LpAmount, CamelAmount *p.Decimal) *p.Decimal {
	TierBonus := LpTierProcent(LpAmount)
	CamelBonus := CamelProcent(CamelAmount)
	Weight := WeightBonus(TierBonus, CamelBonus)

	Result := mt.MULxc(LpAmount, Weight)
	return mt.TruncateCustom(Result, 18)
}

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
