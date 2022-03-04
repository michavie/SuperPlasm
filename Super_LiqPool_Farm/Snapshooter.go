package Super_LiqPool_Farm

import (
	mt "SuperPlasm/SuperMath"
	"encoding/json"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	SUPER       = "https://api.elrond.com/tokens/SUPER-507aa6/accounts?size=10000"
	SuperEgldLP = "https://api.elrond.com/tokens/SUPEREGLD-a793b9/accounts?size=10000"
	SuperCamel  = "https://api.elrond.com/nfts/SCYMETA-3104d5-01/owners?size=10000"
	ExA1        = ElrondAddress("erd1jd7gxdrv7qkghmm4afzk9hy6pw4qa5cfwt0nl7tmyhqujktc27rskzqmke")
	ExA2        = ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	ExA3        = ElrondAddress("erd1qqqqqqqqqqqqqpgqawkm2tlyyz6vtg02fcr5w02dyejp8yrw0y8qlucnj2")
)

type ElrondAddress string

type SuperFarmReward struct {
	Address ElrondAddress
	Reward  *p.Decimal
}

//1)Super Power Chain
//Made Chain Type, Balance is in Decimal
type SuperPower struct {
	Address    ElrondAddress
	SuperPower *p.Decimal
}

//Made Chain Type, Balance is in Decimal
type SuperPowerPercent struct {
	Main              SuperPower
	SuperPowerPercent *p.Decimal
}

//2)Super Power Chain
//Made Chain Type, Balance is in Decimal
type MKSuperPower struct {
	Address               ElrondAddress
	Super                 *p.Decimal
	MetaSuper             *p.Decimal
	MetaKosonicSuperPower *p.Decimal
}

//Made Chain Type, Balance is in Decimal
type MKSuperPowerPercent struct {
	Main                         MKSuperPower
	MetaKosonicSuperPowerPercent *p.Decimal
}

//3)Super Virtual LP Chain
//Made Chain Type, Balance is in Decimal
type SuperVLP struct {
	Address ElrondAddress
	VLP     *p.Decimal
}

//4)Snapshotted Chains
//Snapshot Chain, Balance is in string
type Super struct {
	Address ElrondAddress
	Balance string
}

//Snapshot Chain, Balance is in string
type SuperLP struct {
	Address ElrondAddress
	Balance string
}

//Snapshot Chain, Balance is in string
type CamelAmount struct {
	Address ElrondAddress
	Balance string
}

func OnPage(Link string) string {
	res, err := http.Get(Link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func CreateSuperChain() []Super {
	var OutputChain []Super
	SS := OnPage(SUPER)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func CreateSuperLPChain() []SuperLP {
	var OutputChain []SuperLP
	SS := OnPage(SuperEgldLP)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func CreateCamelChain() []CamelAmount {
	var OutputChain []CamelAmount
	SS := OnPage(SuperCamel)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func GetCamelAmount(Address ElrondAddress, Chain []CamelAmount) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

func GetSuperLPAmount(Address ElrondAddress, Chain []SuperLP) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

func CreateSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []SuperPower {
	var FinalChain []SuperPower
	for i := 0; i < len(Chain1); i++ {

		BaseStringPoint := "Iteration"
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, " ", i)

		if Chain1[i].Address == ExA1 || Chain1[i].Address == ExA2 || Chain1[i].Address == ExA3 {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			SuperAmount := ConvertAtomicUnits(Chain1[i].Balance)
			LPAmount := ConvertAtomicUnits(GetSuperLPAmount(Chain1[i].Address, Chain2))
			SP := SuperPowerComputer(SuperAmount, LPAmount)
			if mt.DecimalGreaterThan(SP, p.NFS("0")) == true {
				Unit := SuperPower{Chain1[i].Address, SP}
				FinalChain = append(FinalChain, Unit)
			}
		}
	}
	fmt.Println("")
	return FinalChain
}

func CreateKosonicSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []SuperPower {
	var FinalChain []SuperPower
	for i := 0; i < len(Chain1); i++ {

		BaseStringPoint := "Iteration"
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, " ", i)

		if Chain1[i].Address == ExA1 || Chain1[i].Address == ExA2 || Chain1[i].Address == ExA3 {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			SuperAmount := ConvertAtomicUnits(Chain1[i].Balance)
			LPAmount := ConvertAtomicUnits(GetSuperLPAmount(Chain1[i].Address, Chain2))
			KSP := KosonicSuperPowerComputer(SuperAmount, LPAmount)
			if mt.DecimalGreaterThan(KSP, p.NFS("0")) == true {
				Unit := SuperPower{Chain1[i].Address, KSP}
				FinalChain = append(FinalChain, Unit)
			}
		}
	}
	fmt.Println("")
	return FinalChain
}

func CreateMetaKosonicSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []MKSuperPower {
	var (
		FinalChain          []MKSuperPower
		SuperAmount         = new(p.Decimal)
		LPAmount            = new(p.Decimal)
		MetaSuperAmount     = new(p.Decimal)
		MetaKosonicPromille = new(p.Decimal)
		MKSP                = new(p.Decimal)
	)
	CamelChainz := CreateCamelChain()

	IzMeta := func(Addy ElrondAddress, Chain []CamelAmount) bool {
		var (
			IzzMeta bool
			IzCamel bool
		)

		CamelValue := GetCamelAmount(Addy, Chain)
		if mt.DecimalGreaterThanOrEqual(p.NFS(CamelValue), p.NFS("1")) == true {
			IzCamel = true
		} else {
			IzCamel = false
		}

		if IzCamel == true {
			IzzMeta = true
		}
		return IzzMeta
	}

	for i := 0; i < len(Chain1); i++ {

		BaseStringPoint := "Iteration"
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, " ", i)

		if Chain1[i].Address == ExA1 || Chain1[i].Address == ExA2 || Chain1[i].Address == ExA3 {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			//Super Amount is taken as integer, therefore truncated.
			SuperAmount = mt.TruncateCustom(ConvertAtomicUnits(Chain1[i].Balance), 0)
			LPAmount = ConvertAtomicUnits(GetSuperLPAmount(Chain1[i].Address, Chain2))

			IzAddressMeta := IzMeta(Chain1[i].Address, CamelChainz)

			if IzAddressMeta == true {
				MetaKosonicPromille = MetaKosonicSuperPowerPromille(SuperAmount)
				MetaSuperAmount = mt.TruncateCustom(mt.MULx(20, SuperAmount, mt.ADDx(20, p.NFI(1), mt.DIVx(20, MetaKosonicPromille, p.NFI(1000)))), 0)

			} else {
				MetaSuperAmount = SuperAmount
			}
			MKSP = KosonicSuperPowerComputer(MetaSuperAmount, LPAmount)

			Unit := MKSuperPower{Chain1[i].Address, SuperAmount, MetaSuperAmount, MKSP}
			FinalChain = append(FinalChain, Unit)
		}
	}
	fmt.Println("")
	return FinalChain
}

func CreateVLPChain(Chain1 []SuperLP, Chain2 []CamelAmount) []SuperVLP {
	var FinalChain []SuperVLP
	for i := 0; i < len(Chain1); i++ {
		if Chain1[i].Address == ExA1 || mt.DecimalLessThan(ConvertAtomicUnits(Chain1[i].Balance), p.NFS("0.5")) == true {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			LPAmount := ConvertAtomicUnits(Chain1[i].Balance)
			Camels := p.NFS(GetCamelAmount(Chain1[i].Address, Chain2))
			VLP := VirtualLP(LPAmount, Camels)
			Unit := SuperVLP{Chain1[i].Address, VLP}
			FinalChain = append(FinalChain, Unit)
		}
	}
	return FinalChain
}
