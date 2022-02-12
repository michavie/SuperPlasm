package Super_LiqPool_Farm

import (
	mt "SuperPlasm/SuperMath"
	"encoding/json"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	SuperEgldLP = "https://api.elrond.com/tokens/SUPEREGLD-a793b9/accounts?size=10000"
	SuperCamel  = "https://api.elrond.com/nfts/SCYMETA-3104d5-01/owners?size=10000"
	ExA1        = ElrondAddress("erd1jd7gxdrv7qkghmm4afzk9hy6pw4qa5cfwt0nl7tmyhqujktc27rskzqmke")
)

type ElrondAddress string

type SuperFarmReward struct {
	Address ElrondAddress
	Reward  *p.Decimal
}

type SuperVLP struct {
	Address ElrondAddress
	VLP     *p.Decimal
}

type SuperLP struct {
	Address ElrondAddress
	Balance string
}

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

func Snapshot(Link, OutputName string) {
	OutputFile, err := os.Create(OutputName)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()

	SnapshotContent := OnPage(Link)
	_, _ = fmt.Fprintln(OutputFile, SnapshotContent)
}

func CreateLPChain() []SuperLP {
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
