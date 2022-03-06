package Super_LiqPool_Farm

import (
	mt "SuperPlasm/SuperMath"
	"encoding/json"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

var (
	SUPER       = "https://api.elrond.com/tokens/SUPER-507aa6/accounts?size=10000"
	SuperEgldLP = "https://api.elrond.com/tokens/SUPEREGLD-a793b9/accounts?size=10000"
	SuperCamel  = "https://api.elrond.com/nfts/SCYMETA-3104d5-01/owners?size=10000"
	ExA1        = ElrondAddress("erd1jd7gxdrv7qkghmm4afzk9hy6pw4qa5cfwt0nl7tmyhqujktc27rskzqmke") //Community Funds
	ExA2        = ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp") //Maiar Super-EGLD-LP Pool
	ExA3        = ElrondAddress("erd1qqqqqqqqqqqqqpgqawkm2tlyyz6vtg02fcr5w02dyejp8yrw0y8qlucnj2") //Jexchange Smart Contract
)

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//01.General Elrond Types
//======================================================================================================================
//
//01_01 - Elrond Address Type
//
type ElrondAddress string

//======================================================================================================================
//
//01_02 - Super Type
//
type Super struct {
	Address ElrondAddress
	Balance string
}

//======================================================================================================================
//
//01_03 - Super-EGLD-LP Type
//
type SuperLP struct {
	Address ElrondAddress
	Balance string
}

//======================================================================================================================
//
//01_04 - Super SFT Camel Type
//
type CamelAmount struct {
	Address ElrondAddress
	Balance string
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//02.Artificial Elrond Types
//======================================================================================================================
//
//02_01 - Virtual (Super-EGLD-LP) Type
//
type SuperVLP struct {
	Address ElrondAddress
	VLP     *p.Decimal
}

//======================================================================================================================
//
//02_02 - Super Liquidity Farming Reward Type
//
type SuperFarmReward struct {
	Address ElrondAddress
	Reward  *p.Decimal
}

//======================================================================================================================
//
//02_03a - SuperPower Type (used for Super-Power and Kosonic Super-Power)
//
type SuperPower struct {
	Address    ElrondAddress
	SuperPower *p.Decimal
}

//======================================================================================================================
//
//02_03b - SuperPowerPercent Type (used for Super-Power and Kosonic Super-Power to display Percentages)
//
type SuperPowerPercent struct {
	Main              SuperPower
	SuperPowerPercent *p.Decimal
}

//======================================================================================================================
//
//02_04a - MKSuperPower Type (used for Meta-Kosonic Super-Power)
//
type MKSuperPower struct {
	Address    ElrondAddress
	Super      *p.Decimal
	MetaSuper  *p.Decimal
	SuperPower *p.Decimal
}

//======================================================================================================================
//
//02_03b - MKSuperPowerPercent Type (used for Meta-Kosonic Super-Power to display Percentages)
//
type MKSuperPowerPercent struct {
	Main                         MKSuperPower
	MetaKosonicSuperPowerPercent *p.Decimal
}

//
//
//
//
//
//======================================================================================================================
//======================================================================================================================
//
//
//[A] - Snapshooter Functions
//
//
//[A]00 - Main Snapshooter Function - OnPage
//Snapshots given Token based on Input Link
//
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

//======================================================================================================================
//
//[A]01 - Super Snapshooter Function; Creates a Chain of Super Values
//
func CreateSuperChain() []Super {
	var OutputChain []Super
	SS := OnPage(SUPER)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//======================================================================================================================
//
//[A]02 - Super-EGLD-LP Snapshooter Function; Creates a Chain of Super-EGLD-Values
//
func CreateSuperLPChain() []SuperLP {
	var OutputChain []SuperLP
	SS := OnPage(SuperEgldLP)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//======================================================================================================================
//
//[A]03 - Camel Snapshooter Function; Creates a Chain of Camel Values
//
func CreateCamelChain() []CamelAmount {
	var OutputChain []CamelAmount
	SS := OnPage(SuperCamel)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//
//
//
//
//
//======================================================================================================================
//======================================================================================================================
//
//
//[B] - LookUp Functions
//
//
//
//[B]01 - SUPER LookUP Function - OnPage
//Returns the SUPER Amount of a given Elrond Address using a given Super Chain
//
func GetSuperAmount(Address ElrondAddress, Chain []Super) *p.Decimal {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}

	//Converting ReadString to Decimal
	EndResult := ConvertAU(Result)
	return EndResult
}

//======================================================================================================================
//
//[B]02 - SUPER-EGLD-LP LookUP Function - OnPage
//Returns the SUPER-EGLD-LP Amount of a given Elrond Address using a given SuperLP Chain
//
func GetSuperLPAmount(Address ElrondAddress, Chain []SuperLP) *p.Decimal {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	//Converting ReadString to Decimal
	EndResult := ConvertAU(Result)
	return EndResult
}

//======================================================================================================================
//[B]03 - Camel LookUP Function - OnPage
//Returns the Camel Amount of a given Elrond Address using a given Camel Chain
//
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

//
//
//
//
//
//======================================================================================================================
//======================================================================================================================
//
//
//[C] - Artificial Chain Functions
//
//
//[C]01 - CreateVLPChain
//Creates a Chain of Virtual (Super-EGLD-LP)
//Virtual (Super-EGLD-LP) is a weighted Super-EGLD-LP by its Amount based on Tiers and Camel Bonus if existent.
//
func CreateVLPChain(Chain1 []SuperLP, Chain2 []CamelAmount) []SuperVLP {
	var FinalChain []SuperVLP
	for i := 0; i < len(Chain1); i++ {
		if Chain1[i].Address == ExA1 || mt.DecimalLessThan(ConvertAU(Chain1[i].Balance), p.NFS("0.5")) == true {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			LPAmount := ConvertAU(Chain1[i].Balance)
			Camels := p.NFS(GetCamelAmount(Chain1[i].Address, Chain2))
			VLP := VirtualLP(LPAmount, Camels)
			Unit := SuperVLP{Chain1[i].Address, VLP}
			FinalChain = append(FinalChain, Unit)
		}
	}
	return FinalChain
}

//======================================================================================================================
//
//[C]02a - CreateSuperPowerChainCore Function
//The Core Function that is used to create a Chain of SuperPower values.
//SuperPower values are 1)SuperPower 2)Kosonic SuperPower 3)Meta-Kosonic SuperPower
//Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
func CreateSuperPowerChainCore(Chain1 []Super, Chain2 []SuperLP, SuperPowerFunction func(*p.Decimal, *p.Decimal) *p.Decimal) []MKSuperPower {
	var (
		FinalChain      []MKSuperPower
		GetMeta         = false
		SuperPowerValue = new(p.Decimal)
	)

	//MetaCheck Snapshots - used for Meta-Kosonic Super-Power
	//Multiple Chains can be added if multiple SFTs must be checked
	//Remember to add Checks in the IzMeta Function as well
	SFT1Chain := CreateCamelChain()

	//Getting a Function name.
	GetFunctionName := func(temp interface{}) string {
		Value := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
		return Value[len(Value)-1]
	}

	for i := 0; i < len(Chain1); i++ {
		var MetaSuperAmount = new(p.Decimal)

		BaseStringPoint := "Iteration"
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, " ", i)

		if Chain1[i].Address == ExA1 || Chain1[i].Address == ExA2 || Chain1[i].Address == ExA3 {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			//0)Address is Chain1[i].Address

			//1)Getting the Super Value
			SuperAmount := ConvertAU(Chain1[i].Balance)
			TruncatedSuperAmount := mt.TruncateCustom(SuperAmount, 0)
			//Integers the non-integer Super

			//2)Getting the LP Amount
			LPAmount := GetSuperLPAmount(Chain1[i].Address, Chain2)

			//3)Computing MetaSuper and SuperPower
			if GetFunctionName(SuperPowerFunction) == "MetaKosonicSuperPowerComputer" {
				GetMeta = IzMeta(Chain1[i].Address, SFT1Chain)
				if GetMeta == true {
					//if meta is true, SuperPower applies the input Function,
					//which is in this case "MetaKosonicSuperPowerComputer"
					//It has built in Super to meta-Super conversion
					//That is why it is used with SuperAmount
					MetaSuperAmount = ComputeMetaSuper(SuperAmount)
					SuperPowerValue = SuperPowerFunction(SuperAmount, LPAmount)
				} else {
					//if meta is false, SuperPower applies doesnt apply the input Function
					//but applies the KosonicSuperPowerComputer function
					//because this doesnt use meta-Super
					MetaSuperAmount = SuperAmount
					SuperPowerValue = KosonicSuperPowerComputer(SuperAmount, LPAmount)
				}
			} else {
				//case where non Meta-Kosonic Super-Power has to be calculated
				//namely the Super-Power and Kosonic Super-Power
				MetaSuperAmount = SuperAmount
				SuperPowerValue = SuperPowerFunction(SuperAmount, LPAmount)
			}

			//Truncating the meta-Super since it must be integer
			TruncatedMetaSuperAmount := mt.TruncateCustom(MetaSuperAmount, 0)

			//Creating the Chain element. Only SuperPower values greater than 0 are added to the chain.
			//Since the SuperPower computing Function sets the SuperPower Result to 0
			//if it is below the "SuperPowerExistenceThreshold" this code here
			//Incorporates in Chain only non zero-values.

			if mt.DecimalGreaterThan(SuperPowerValue, p.NFS("0")) == true {
				Unit := MKSuperPower{Chain1[i].Address, TruncatedSuperAmount, TruncatedMetaSuperAmount, SuperPowerValue}
				FinalChain = append(FinalChain, Unit)
			}
		}
	}
	fmt.Println("")
	return FinalChain
}

//======================================================================================================================
//
//[C]02b - CreateSuperPowerChain Function
//Creates a chain of SuperPower values
//Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
func CreateSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []MKSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, SuperPowerComputer)
	return Result
} //======================================================================================================================
//
//[C]02c - CreateKosonicSuperPowerChain Function
//Creates a chain of Kosonic SuperPower values
//Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
func CreateKosonicSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []MKSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, KosonicSuperPowerComputer)
	return Result
}

//
//[C]02d - CreateMetaKosonicSuperPowerChain Function
//Creates a chain of Meta-Kosonic SuperPower values
//Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
func CreateMetaKosonicSuperPowerChain(Chain1 []Super, Chain2 []SuperLP) []MKSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, MetaKosonicSuperPowerComputer)
	return Result
}

//======================================================================================================================
//======================================================================================================================
//
//
//[D] - Checking Function
//
//
//[D]01 - IzMeta
//Checks if an address is Superciety Meta
//Checking is done by checking 1 or multiple SFTs/NFTs
//
func IzMeta(Addy ElrondAddress, Chain1 []CamelAmount) bool {
	var (
		MetaResult bool //Total boolean Value
		IzCamel    bool //1st boolean value to check
	)

	//1st SFT Check
	CamelValue := GetCamelAmount(Addy, Chain1)
	if mt.DecimalGreaterThanOrEqual(p.NFS(CamelValue), p.NFS("1")) == true {
		IzCamel = true
	} else {
		IzCamel = false
	}

	//If all SFT Checks are true, IzMeta is true.
	if IzCamel == true {
		MetaResult = true
	}
	return MetaResult
}
