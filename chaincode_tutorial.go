package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strings"
	"fmt"
	"os"
)

// 啟用Logger
var logger = shim.NewLogger("lbh_tuna_demo")

// 定義智能合約
type SampleChaincodeByLBH struct {

}

// 為此智能合約定義方法
func (sampleCC *SampleChaincodeByLBH ) putData( stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) <1 {
		return shim.Error("incorrect parameter length.")
	}

	logger.Debugf("Prepare to insert data, value is %s", strings.Join(args,"|$%#|"))
	logger.Info("Prepare to insert data, key is %s",args[0])

	putErr := stub.PutState(args[0],[]byte(args[1]))

	if putErr != nil {
		logger.Errorf("can not insert data, key is %s, value is %s ",args[0],args[1])
	}

	logger.Infof("key %s insert success", args[0])

	return shim.Success([]byte(fmt.Sprintf("This key is %s, value is %s, insert success.", args[0],args[1])))

}

// ComposeKey PutData
// 使用邏輯說明
// 第一個參數為 composekeyIndex
// 第二個參數為 其他相關key
// 第三個參數為 其他相關key
// 第四個參數為 值
// 將index與key組合作composekey，而後將值與其塞入
func (sampleCC *SampleChaincodeByLBH) putComposeKeyData( stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) <3 {
		return shim.Error("incorrect parameter length.")
	}

	logger.Debugf("Prepare to get value from argument, value is %s ", strings.Join(args,"|$%#|"))

	composeKeyIndex := args[0]
	txId := stub.GetTxID()

	// 用主鍵搭配其他附帶值，組合出ComposeKey
	composeKey,_ := stub.CreateCompositeKey(composeKeyIndex,[]string{args[1],args[2],txId})

	putErr := stub.PutState(composeKey,[]byte(args[3]))

	if putErr != nil {

		shim.Error("can not insert data")

	}

	return shim.Success([]byte(composeKey))

}


// ComposeKey GetData
// 思想邏輯
// 用戶輸入 composeKeyIndex
// 我方依據 composeKeyIndex 找出相關 keyvalue
// 再把個別key 拆解，取出用來跟Index結合的key，加總
// 並把value 加總
// 傳回
func (sampleCC *SampleChaincodeByLBH) getComposeKeyData( stub shim.ChaincodeStubInterface, args []string) sc.Response{
	// 取得與object相關的key value
	composeKeyIndex := args[0]

	// 得到一系列的(key,value)
	deltaResultIter, _  := stub.GetStateByPartialCompositeKey(composeKeyIndex,[]string{args[1]})


	var finalKeyPart1 string
	var finalKeyPart2 string
	var finalKeyPart3 string
	var finalVal string

	for deltaResultIter.HasNext() {
		eachRecordKV, _ := deltaResultIter.Next()

		fmt.Print(eachRecordKV)

		// 解析ComposeKey, 取出用來作ComposerKey的個別key
		_ ,keyParts, _ := stub.SplitCompositeKey(eachRecordKV.Key)

		elementOne := keyParts[0]
		elementTwo := keyParts[1]
		elementTxId := keyParts[2]

		logger.Debugf("key is %s, keyPart1 is %s, keyPart2 is %s", eachRecordKV.Key,elementOne,elementTwo)

		finalVal += string(eachRecordKV.Value)
		finalKeyPart1 += elementOne
		finalKeyPart2 += elementTwo
		finalKeyPart3 += elementTxId
	}

	returnValue := finalKeyPart1 + "||" + finalKeyPart2 +"||"  + "||"+ finalVal

	return shim.Success([]byte(returnValue))
}


// 記得解釋
func (sampleCC *SampleChaincodeByLBH) Init(stub shim.ChaincodeStubInterface) sc.Response{

	return shim.Success(nil)
}

// 調用列表
func (sampleCC *SampleChaincodeByLBH) Invoke(stub shim.ChaincodeStubInterface)sc.Response{

	fn, args := stub.GetFunctionAndParameters()

	switch fn {
		case "putData" :
			return sampleCC.putData(stub, args)
		case "putCompose" :
			return sampleCC.putComposeKeyData(stub,args)
		case "getCompose":
			return sampleCC.getComposeKeyData(stub,args)
		default:
			return shim.Error("Invalid Smart Contract function name.")
	}
	return shim.Error("Invalid Smart Contract function name.")

}


func main() {

	os.Setenv("CORE_CHAINCODE_LOGGING_LEVEL","DEBUG")
	shim.SetupChaincodeLogging()
	//shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(SampleChaincodeByLBH))
	if err != nil {
		fmt.Println("Could not start SampleChaincode")
	} else {
		fmt.Println("SampleChaincode successfully started")
	}
	
}
