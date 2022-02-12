package Super_LiqPool_Farm

import (
	mt "SuperPlasm/SuperMath"
	p "github.com/Crypt0plasm/Firefly-APD"
)

var (
	SuperEgldLPDecimals   = int64(18)
	SuperFarmRewardAmount = int64(1650)
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

	CamelTotalBonus = mt.MULxc(Camel, Bonus)
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
