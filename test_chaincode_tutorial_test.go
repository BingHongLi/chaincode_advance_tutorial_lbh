package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	sc "github.com/hyperledger/fabric/protos/peer"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte ) {
	res := stub.MockInit("1",nil)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

// 檢驗調用，並傳回結果
func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte)  sc.Response {
	// uuid 為交易編號
	res := stub.MockInvoke("1", args)

	// 此處只檢驗回傳狀態，之後可以考慮回傳結果，或回傳結果寫在其他地方
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}

	return res
}

// 測試是否正常啟用
func TestInit(t *testing.T) {

	scc := new(SampleChaincodeByLBH)
	stub := shim.NewMockStub("lbh_tuna_demo", scc)

	// Init
	checkInit(t, stub,nil)

}

// 實作測試業務功能
func TestPutData (t *testing.T) {

	scc := new(SampleChaincodeByLBH)
	stub := shim.NewMockStub("SampleChaincodeByLBH", scc)

	checkInit(t, stub,nil)

	putRes := checkInvoke(t, stub, [][]byte{[]byte("putData"),[]byte("binghongli"),[]byte("123")})

	// 測狀態是否正常
	if putRes.Status != shim.OK {
		fmt.Println("Invoke failed" + string(putRes.Message))
		t.FailNow()
	}

	// 測內容是否正確
	if string(putRes.Payload) != "This key is binghongli, value is 123, insert success." {
		t.FailNow()
	}

	t.Log("success")

}