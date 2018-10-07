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
