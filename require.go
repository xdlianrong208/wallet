package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//url
var(
	Ethurl  = "http://localhost:8545"

	EthFrom = "0xd7ff4f9fd04811ad63509979676c772db815827c"
	EthTo   = "0xc8703ca0ba2db1e9dda37a7d020addf776c5e21a"
)

// unlock publisher eth_account struct
type toETH struct {
	Jsonrpc	 string			`json:"jsonrpc"`
	Method	 string	    	`json:"method"`
	Params   []interface{}  `json:"params"`
	Id       int			`json:"id"`
}

// get result from unlock to ethereum
type unlockget struct {
	Jsonrpc	 string			`json:"jsonrpc"`
	Id       int			`json:"id"`
	Result   bool           `json:"result"`
}

type SendTx struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Gas      string  `json:"gas"`
	GasPrice string  `json:"gasPrice"`
	Value    string  `json:"value"`
	ID       string  `json:"id"`
	Data     string  `json:"data"`
	Spk      string  `json:"spk"`
	Rpk      string  `json:"rpk"`
	S        string  `json:"s"`
	R        string  `json:"r"`
	Vor      string  `json:"vor"`
	Cmo      string  `json:"cmo"`
}

// 解锁区块链上的账户
func UnlockAccount(ethaccount string, ethkey string) bool{
	paramsul := make([]interface{}, 3)
	paramsul[0] = ethaccount
	paramsul[1] = ethkey
	paramsul[2] = 30000

	data := toETH{"2.0", "personal_unlockAccount", paramsul,67}

	datapost, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req, err := http.NewRequest("POST", Ethurl, bytes.NewBuffer(datapost))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyC))
	var s unlockget;
	json.Unmarshal([]byte(bodyC), &s)
	if(s.Result == true){
		return true
	}else{
		return false
	}
}

// 发送转账交易到区块链
func SendTransaction(spk string, rpk string, s string, r string, vor string, cmo string) bool {
	paramstx   := make([]interface{}, 1)

	//epkrc1 = strings.TrimLeft(epkrc1, "0x")
	//fmt.Println(hex.DecodeString(epkrc1))
	paramstx[0] = SendTx{EthFrom, EthTo, "0x0", "0x0", "0x0", "0x0", "0x00", spk, rpk, s, r,vor,cmo}
	data := toETH{"2.0", "eth_sendTransaction", paramstx,67}
	fmt.Println(data)
	datapost, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req, err := http.NewRequest("POST", Ethurl, bytes.NewBuffer(datapost))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	bodyC, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyC))
	return true
}

// 获取交易信息
func GetTransaction(txhash string) bool{
	paramgtx   := make([]interface{}, 1)
	paramgtx[0] = txhash
	data := toETH{"2.0", "eth_getTransactionByHash", paramgtx,67}
	datapost, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req, err := http.NewRequest("POST", Ethurl, bytes.NewBuffer(datapost))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	bodyC, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyC))
	return true
}