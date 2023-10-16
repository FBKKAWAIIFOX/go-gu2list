package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gu2list/database/database"
	"gu2list/ent"
	"gu2list/ent/schema"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	SqlClient *ent.Client
)

func main() {

	var err error

	//Choose Database on your own
	//Here is Example Using SqlClient Database.
	//Testing env Using XAMPP , SqlClient
	//Connect To Database

	SqlClient, err = database.CreateClient("root", "", "test", "localhost", 3306)
	if err != nil {
		panic(err)
	}

	err = database.UpdateUser(858610110002888724, &schema.Roles{1153040242}, false)
	if err != nil {
		panic(err)
	}

	//err = database.UpdateUserLogs(858610110002888724, "FBK", true, time.Now().Unix())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(database.IsManager(858610110002888724))
}

/*
  Usage: LatestVCoinPrice() --> return Latest Price and UpdatedTime
*/

func LatestVCoinPrice() (CurrentPrice, UpdatedTime string) {
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
