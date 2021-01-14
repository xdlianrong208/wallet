package main

import "fmt"

// 解锁账户需要的参数
var(
	ethaccount = "0xd7ff4f9fd04811ad63509979676c772db815827c"
	ethkey     = ""
)
// 发送转账交易需要的参数
var(
	spk        = "234b7f8dcdec50b47127a9ba7f03d629bd751b571ff07ac8879c4ca0a91b146205e72bd1ac5e39bcf34cbbbcf48a13edc865f862a85ce69866be24e078a3942a33333f914834ced561c145797d9b5782719dbd1b43a668d4b01151f9c0e67d9f1569899100a4ce41de3c549b649ff72d5d7c9fe8983c244cc28f2ce84b2a758c"
	rpk        = "234b7f8dcdec50b47127a9ba7f03d629bd751b571ff07ac8879c4ca0a91b146205e72bd1ac5e39bcf34cbbbcf48a13edc865f862a85ce69866be24e078a3942a33333f914834ced561c145797d9b5782719dbd1b43a668d4b01151f9c0e67d9f1569899100a4ce41de3c549b649ff72d5d7c9fe8983c244cc28f2ce84b2a758c"
	r          = "0x9"
	s          = "0x2"
	vor        = "0x0c21ccfaaa23f4562094fa71c16bbfeb1db461c2f96dc72c3a70b8cd266bd37c"
	cmo        = "0x145efb9d48584450198d2fb30a1ba7e9396eb08e0b5c662dd9414d9d8fa1abe4"
)
// 查看交易内容需要的参数
var(
	txhash     = "0xf4c2854732f0eb772456232aa2666dbb7f6311875d2f6a14ba3180168f2dd993"
)

func main(){
	//TestUnlock()
	//TestSendTransaction()
	TestGetTransaction()
}

func TestUnlock(){
	if(UnlockAccount(ethaccount, ethkey) == true){
		fmt.Println("unlock account " + ethaccount + " right")
	}else{
		fmt.Println("unlock account " + ethaccount + " erro")
	}
}

func TestSendTransaction(){
	if(SendTransaction(spk, rpk, s, r, vor, cmo) == true){
		fmt.Println("sendtx right")
	}else{
		fmt.Println("sendtx erro")
	}
}

func TestGetTransaction() {
	if(GetTransaction(txhash) == true){
		fmt.Println("gettx "+ txhash)
	}else{
		fmt.Println("gettx erro")
	}
}
