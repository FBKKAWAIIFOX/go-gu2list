package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gu2list/database/database"
	"gu2list/ent"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	Mysql *ent.Client
)

func main() {
	/*
		Connect To Databse
	*/
	var err error
	Mysql, err = database.CreateClient("root", "", "test", "localhost", 3306)
	if err != nil {
		panic(err)
	}
	//err = database.UpdateUser(858610110002888724, &schema.Roles{1153040240685109311}, true)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateUserLogs(858610110002888724, FBK", true, time.Now().Unix())
	//if err != nil {
	//	panic(err)
	//}
	database.IsManager(858610110002888724)
}

/*
  Usage: CheckCoinPrice() --> return Latest Price and UpdatedTime
*/

func LatestVCoinPrice() (CurrentPrice, UpdatedTime string) {
	resp, err := http.Get("https://pool-open-data.mcfallout.net/pool-status-java.log")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	output := strings.Split(string(body), "\n")
	TransactionDetails := strings.Split(output[len(output)-2], ",")
	if len(TransactionDetails) >= 1 {
		coin, _ := decimal.NewFromString(TransactionDetails[0])
		emer, _ := decimal.NewFromString(TransactionDetails[1])
		return strconv.Itoa(int(emer.IntPart() / coin.IntPart())), TransactionDetails[2]
	}
	return "", ""
}
