package VCoin

import (
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Price() (CurrentPrice, UpdatedTime string) {
	resp, err := http.Get("https://pool-open-data.mcfallout.net/pool-status-java.log")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	output := strings.Split(string(body), "\n")
	/*
			   APi Raw Data Struct:
		        VCoin              Emerald 	       Transaction Time
			   625282.9712234671,37439455239.0504,2023-10-17 05:13:31

			   Simple Calculation:
		       Emerald / VCoin	which return **float64 value

	*/
	TransactionDetails := strings.Split(output[len(output)-2], ",")

	/*
	   Check If Last Transaction Exists
	*/
	if len(TransactionDetails) >= 1 {
		coin, _ := decimal.NewFromString(TransactionDetails[0])
		emer, _ := decimal.NewFromString(TransactionDetails[1])                          //Result
		return strconv.Itoa(int(emer.IntPart() / coin.IntPart())), TransactionDetails[2] //Convert the result into a string value , Change it if needed.
	}
	return "", ""
}
