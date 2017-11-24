package main
/* 
#cgo  CFLAGS:  -I  /root/bip44cxx 
#cgo  LDFLAGS:  -L /root/bip44cxx -lbip44wallet
#include <stdlib.h>
#include "interface.h"
*/  
import "C"    
import (
	"flag"
	"fmt"
	"os"
	// "bytes"
	// "crypto/ecdsa"
	// "crypto/elliptic"
	// "crypto/hmac"
	// "crypto/sha256"
	// "errors"
	// "fmt"
	// "hash"
	// "math/big"
	"unsafe"
	"./wallet"
	"github.com/btcsuite/btcd/btcec"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/common"
	"encoding/hex"
	// "github.com/btcsuite/btcd/btcjson"
	// "github.com/btcsuite/btcd"
	// "github.com/btcsuite/btcwallet/rpc/legacyrpc"
)
type GoWallet struct {
     cxxwallet C.voidstar;
}
type Prefixes struct {
    bip44_code uint32
	HDprivate uint32
	HDpublic uint32
	P2KH uint32
	P2SH uint32
}
func (p *Prefixes)setPrefixes(coin string){
	// Prefixes BTC =  {0x80000000, 0x0488ADE4, 0x0488B21E, 0x00, 0x05};
	// Prefixes tBTC = {0x80000001, 0x04358394, 0x043587CF, 0x6f, 0xC4};
	// Prefixes LTC =  {0x80000002, 0x0488ADE4, 0x0488B21E, 0x30, 0x32};
	// Prefixes tLTC = {0x80000001, 0x04358394, 0x04358394, 0x6f, 0xC0};
	// Prefixes POT =  {0x80000081, 0x0488ADE4, 0x0488B21E, 55, 0x05};
	if(coin == "BTC"){
		p.bip44_code=0x80000000
		p.HDprivate=0x0488ADE4
		p.HDpublic=0x0488B21E
		p.P2KH=0x00
		p.P2SH=0x05
	}else if(coin == "tBTC"){
		p.bip44_code=0x80000001
		p.HDprivate=0x04358394
		p.HDpublic=0x043587CF
		p.P2KH=0x6f
		p.P2SH=0xC4
	}else if(coin == "LTC"){
		p.bip44_code=0x80000002
		p.HDprivate=0x0488ADE4
		p.HDpublic=0x0488B21E
		p.P2KH=0x30
		p.P2SH=0x32
	} else if(coin == "tLTC"){
		p.bip44_code=0x80000001
		p.HDprivate=0x04358394
		p.HDpublic=0x04358394
		p.P2KH=0x6f
		p.P2SH=0xC0
	} else if(coin=="POT"){
		p.bip44_code=0x80000081
		p.HDprivate=0x0488ADE4
		p.HDpublic=0x0488B21E
		p.P2KH=55
		p.P2SH=0x05
	}else{
		p.bip44_code=0x80000000
		p.HDprivate=0x0488ADE4
		p.HDpublic=0x0488B21E
		p.P2KH=0x00
		p.P2SH=0x05
	}
}
func New()(GoWallet){
     var ret GoWallet;
     ret.cxxwallet = C.walletInit();
     return ret;
}
func New1(mnemonic string)(GoWallet){
     var ret GoWallet;
     ret.cxxwallet = C.walletInitFromMnemonic(C.CString(mnemonic));
     return ret;
}
func New2(coin_code Prefixes)(GoWallet){
     var ret GoWallet;
     CPrefixes := *(*C.struct_Prefixes)(unsafe.Pointer(&coin_code))
     ret.cxxwallet = C.walletInitFromCointype(CPrefixes);
     return ret;
}
func New3(mnemonic string,coin_code Prefixes)(GoWallet){
     var ret GoWallet;
     CPrefixes := *(*C.struct_Prefixes)(unsafe.Pointer(&coin_code))
     ret.cxxwallet = C.walletInitFromCointypeAndMnemonic(C.CString(mnemonic),CPrefixes);
     return ret;
}
// voidstar walletInitFromCointype(Prefixes coin_code);
// voidstar walletInitFromCointypeAndMnemonic(const char* mnemonicSeed, Prefixes coin_code);
func (f GoWallet)Free(){
     C.walletFree(f.cxxwallet)
}
func (f GoWallet)getMnemonic()string{
	csRet:=C.getMnemonic(f.cxxwallet)
	// fmt.Printf("fmt: %s\n", C.GoString(csRet))
	// defer C.free(unsafe.Pointer(cstr))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)getMasterKey()string{
	csRet:=C.getMasterKey(f.cxxwallet)
	// fmt.Printf("fmt: %s\n", C.GoString(csRet))
	// defer C.free(unsafe.Pointer(cstr))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)getChildKeyPath()string{
	csRet:=C.getChildKeyPath(f.cxxwallet)
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)getChildSecretKey(index int)string{
	csRet:=C.getChildSecretKey(f.cxxwallet,C.int(index))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)getChildPublicKey(index int)string{
	csRet:=C.getChildPublicKey(f.cxxwallet,C.int(index))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)getChildAddress(index int)string{
	csRet:=C.getChildAddress(f.cxxwallet,C.int(index))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func (f GoWallet)set_account(index int){
	C.set_account(f.cxxwallet,C.int(index))
}
func (f GoWallet)getCurrentAccount()int{
	csRet:=C.getCurrentAccount(f.cxxwallet)
	// defer C.free(unsafe.Pointer(csRet))
    return int(csRet)
}
// func (f GoWallet)createrawtransaction(reqjson string)string{
// 	csRet:=C.createrawtransaction(f.cxxwallet,C.CString(reqjson))
// 	// defer C.free(unsafe.Pointer(csRet))
//     return C.GoString(csRet)
// }
func FromMnemonicToMasterKey(mnemonic string)string {	
	csRet:=C.FromMnemonicToMasterKey(C.CString(mnemonic))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}
func fromHex(s string) []byte {
	r, err := hex.DecodeString(s)
	if err != nil {
		panic("invalid hex in source file: " + s)
	}
	return r
}
func test2() {
	p := new(Prefixes)
	p.setPrefixes("BTC")
	// fmt.Println(p.bip44_code)
	wallet := New3("label stick flat innocent brother frost rebel aim creek six baby copper need side cannon student announce alpha",*p);
    ret:=wallet.getMnemonic()
    fmt.Println(ret)
    wallet.set_account(2)

	getmasterkey:=wallet.getMasterKey()
    fmt.Println(getmasterkey)
   	
   	getChildKeyPath:=wallet.getChildKeyPath()
    fmt.Println(getChildKeyPath)
	
	getChildSecretKey:=wallet.getChildSecretKey(0)
	fmt.Println(getChildSecretKey)

	getChildPublicKey:=wallet.getChildPublicKey(0)
	fmt.Println(getChildPublicKey)
	
	getChildAddress:=wallet.getChildAddress(0)
	fmt.Println(getChildAddress)

	fmt.Println("-------------------------------------")

{
	//use btcd
	// txInputs := []btcjson.TransactionInput{
	// 				{Txid: "6c3f611cbd624e8a094f08b10f849b765d3548c13ace1704de050a44f504caff", Vout: 0},
	// 			}
	// 			amounts := map[string]float64{"mxu9tvJsuZq1rxiaevcUJkuu6mv2LFhpSr": 0.1}
	// cmd:=btcjson.NewCreateRawTransactionCmd(txInputs, amounts, nil)
	// // func HandleCreateRawTransaction(s *rpcServer, cmd interface{}, closeChan <-chan struct{}) (interface{}, error)
	// closeChan:=make(chan,1)
	// ret,err:=HandleCreateRawTransaction(nil,cmd,closeChan)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(ret)

	// func SignRawTransaction(icmd interface{}, w *wallet.Wallet, chainClient *chain.RPCClient) (interface{}, error)
	// `[{"txid":"6c3f611cbd624e8a094f08b10f849b765d3548c13ace1704de050a44f504caff","vout":n,"scriptpubkey":"value","redeemscript":"value"}] ["privkey"] flags="ALL"`
	// cmd:=`"signrawtransaction 0100000001ffca04f5440a05de0417ce3ac148355d769b840fb1084f098a4e62bd1c613f6c0000000000ffffffff0180969800000000001976a914beacf93f739b48324e79d5c3314c8a434d18d2ba88ac00000000 ["privkey"]"`
	// txInputs := []btcjson.RawTxInput{} 				
	// privKeys := []string{"cVJiFesQn1duqM6RThR3N8oXL6xkYFo1r5h4PtCaXV3qXkxd3DBT"} 				
	// cmd:=btcjson.NewSignRawTransactionCmd("0100000001ffca04f5440a05de0417ce3ac148355d769b840fb1084f098a4e62bd1c613f6c0000000000ffffffff0180969800000000001976a914beacf93f739b48324e79d5c3314c8a434d18d2ba88ac00000000", &txInputs, &privKeys, nil)

	// retsign,err:=legacyrpc.SignRawTransaction(cmd,nil)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(retsign)

	// reqjson:=`"[{"txid":"6c3f611cbd624e8a094f08b10f849b765d3548c13ace1704de050a44f504caff","vout":0}]" "{"mxu9tvJsuZq1rxiaevcUJkuu6mv2LFhpSr":0.1}"`
	// ret:=wallet.createrawtransaction(reqjson)
	// fmt.Println(ret)
	//0100000001c0f97438287f944d1ed73b5d1fe3349440cd470584847faaba109639e272e48d0000000000ffffffff0140420f00000000001976a914d3a4a0e66f494a95942e45b26561c07f81bacfd788ac00000000
	//test sign
	// test:=`[{"txid" : "8de472e2399610baaa7f84840547cd409434e31f5d3bd71e4d947f283874f9c0","vout":0}]" "{"mzp267vBXdD5Q79Gnx26LsH2r2e7uYDMyt":0.01}`
	// pkBytes:=[]byte(getChildSecretKey)
	// privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)
	// // key, _ := crypto.HexToECDSA()
	// msgHash := fromHex("0100000001c0f97438287f944d1ed73b5d1fe3349440cd470584847faaba109639e272e48d0000000000ffffffff0140420f00000000001976a914d3a4a0e66f494a95942e45b26561c07f81bacfd788ac00000000")
	// // msgHash :=[]byte(test)
	// // signRFC6979(privateKey *PrivateKey, hash []byte) (*Signature, error)
	// sig,_:=privKey.Sign(msgHash)
	// // sig := Signature{
	// // 	R: fromHex("fef45d2892953aa5bbcdb057b5e98b208f1617a7498af7eb765574e29b5d9c2c"),
	// // 	S: fromHex("d47563f52aac6b04b55de236b7c515eb9311757db01e02cff079c3ca6efb063f"),
	// // }
	// out:=sig.Serialize()
	// fmt.Println(hex.EncodeToString(out))
	// fmt.Println(len(hex.EncodeToString(out)))
	// if !sig.Verify(msgHash, pubKey) {
	// 	fmt.Println("Signature failed to verify")
	// }
}
	{
		// kh := getChildSecretKey
		// k0, _ := crypto.HexToECDSA(kh)

		// msg0 := crypto.Keccak256([]byte("foo"))
		// sig0, _ := Sign(msg0, k0)

		// // msg1 := common.FromHex("0000000000000000000000000000000122311")
		// // sig1, _ := Sign(msg1, k0)

		// fmt.Printf("msg: %x\nprivkey: %s\nsig: %x\n", msg0, k0, sig0)
		// fmt.Printf("msg: %x\nprivkey: %s\nsig: %x\n", msg1, k0, sig1)
	}

		fmt.Println("********************************************")
	
	{
		// fmt.Println(getChildSecretKey)
		// fmt.Println(getChildAddress)
		// key, _ := crypto.HexToECDSA(getChildSecretKey)
		// addr := common.HexToAddress(getChildAddress)

		// fmt.Printf("\n\naddr: %s\n\n", addr.Hex())

		// msg := crypto.Keccak256([]byte("foo"))
		// sig, err := Sign(msg, key)
		// if err != nil {
		// 	fmt.Printf("Sign error: %s\n", err)
		// }
		// fmt.Printf("msg: %x\nprivkey: %s\n\nsig: %x\n", msg, key, sig)

		// fmt.Println("********************************************")
		// recoveredPub, err := crypto.Ecrecover(msg, sig)
		// if err != nil {
		// 	fmt.Printf("ECRecover error: %s\n", err)
		// }
		// pubKey := crypto.ToECDSAPub(recoveredPub)
		// // genAddr := PubkeyToAddress(key.PublicKey)
		// fmt.Printf("pubkey:%s\n\n",pubKey)
		// recoveredAddr := common.BytesToAddress(crypto.FromECDSAPub(pubKey))
		// if addr != recoveredAddr {
		// 	fmt.Printf("Address mismatch: want: %x have: %x\n", addr, recoveredAddr)
		// }
		// fmt.Println("********************************************")
		// // should be equal to SigToPub
		// recoveredPub2, err := SigToPub(msg, sig)
		// if err != nil {
		// 	fmt.Printf("ECRecover error: %s\n", err)
		// }
		// fmt.Printf("pubkey:%s\n",recoveredPub2)
		// // recoveredAddr2 := crypto.PubkeyToAddress(recoveredPub2.PublicKey)
		// recoveredAddr2 := common.BytesToAddress(crypto.FromECDSAPub(recoveredPub2))
		// if addr != recoveredAddr2 {
		// 	fmt.Printf("Address mismatch: want: %s have: %x\n", addr.Hex(), recoveredAddr2)
		// }
	}

    wallet.Free();
}
func test(){

    {
		wallet := New();
	    ret:=wallet.getMnemonic()
	    fmt.Println(ret)
	    getmasterkey:=wallet.getMasterKey()
	    fmt.Println(getmasterkey)

	    masterkey:=FromMnemonicToMasterKey(ret)
	    fmt.Println(masterkey)

	    wallet.Free();
	    fmt.Println("/////////////////////////////////////")
	}
	{
		mnemonic:="label stick flat innocent brother frost rebel aim creek six baby copper need side cannon student announce alpha"
		wallet := New1(mnemonic)
	    
	    ret:=wallet.getMnemonic()
	    fmt.Println(ret)

		getmasterkey:=wallet.getMasterKey()
	    fmt.Println(getmasterkey)

	    masterkey:=FromMnemonicToMasterKey(mnemonic)
	    fmt.Println(masterkey)
	    wallet.Free();
	    fmt.Println("/////////////////////////////////////")
	}
	{
		p := new(Prefixes)
		p.setPrefixes("LTC")
		wallet := New2(*p);
	    ret:=wallet.getMnemonic()
	    fmt.Println(ret)
		getmasterkey:=wallet.getMasterKey()
	    fmt.Println(getmasterkey)
	    masterkey:=FromMnemonicToMasterKey(ret)
	    fmt.Println(masterkey)

	    wallet.Free();
	    fmt.Println("/////////////////////////////////////")
	}
	{
		p := new(Prefixes)
		p.setPrefixes("BTC")
		wallet := New3("label stick flat innocent brother frost rebel aim creek six baby copper need side cannon student announce alpha",*p);
	    ret:=wallet.getMnemonic()
	    fmt.Println(ret)

		getmasterkey:=wallet.getMasterKey()
	    fmt.Println(getmasterkey)
	    masterkey:=FromMnemonicToMasterKey(ret)
	    fmt.Println(masterkey)

	    wallet.Free();
	}
}

func main() {
	number, vanity, export := parseParams()
	if number > 0 {
		err := generateWallets(uint32(number), vanity, export)
		if err != nil {
			println(err.Error())
			return
		}
	} else {
		test2()



		// view.ShowSplashView(view.SplashStartView)

		// var ws []*wallet.Wallet
		// if !wallet.IsFileExists() {
		// 	var err error
		// 	ws, err = createWallets(1, 10)
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		return
		// 	}
		// 	// save wallets
		// 	wf := wallet.NewWalletFile(ws)
		// 	wf.Save()
		// } else {
		// 	wf, err := wallet.LoadWalletFile()
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		return
		// 	}
		// 	ws = wf.Wallets
		// }

		// showUI(ws)
	}
}

// func showUI(ws []*wallet.Wallet) {

// 	accountView := view.NewAccountView(ws)
// 	accountView.Show()

// 	for accountView.Data != nil {
// 		cmd := accountView.Data.(string)
// 		if cmd == "quit" {
// 			break
// 		}
// 		tipView := view.NewTipView(cmd)
// 		if tipView != nil {
// 			tipView.Show()
// 		}
// 		accountView.Show()
// 	}
// }


// create wallets by secret and salt
// func createWallets(start, count uint32) (ws []*wallet.Wallet, err error) {
// 	// view.ShowSplashView(view.SplashCreateView)

// 	// create wallets
// 	wp, err := wallet.InputNewParameters(3)
// 	if err != nil {
// 		return
// 	}
// 	// wp = wallet.WalletParam{Secret:"https://github.com/aiportal", Salt:"gowallet"}

// 	wa, err := wallet.NewWalletAccount(wp.SecretBytes(), wp.SaltBytes())
// 	if err != nil {
// 		return
// 	}
// 	ws, err = wa.GenerateWallets(start, count)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

//Parse command line parameters
func parseParams() (number uint, vanity, export string) {

	flag.UintVar(&number, "number", 0, "Number of wallets to generate.")
	flag.UintVar(&number, "n", 0, "Number of wallets to generate.")

	flag.StringVar(&vanity, "vanity", "", "Find vanity wallet address matching. (prefix)")
	flag.StringVar(&vanity, "v", "", "Find vanity wallet address matching. (prefix)")

	flag.StringVar(&export, "export", "", "Export wallets in WIF format.")
	flag.StringVar(&export, "e", "", "Export wallets in WIF format.")

	flag.Parse()
	return
}

func generateWallets(number uint32, vanity, export string) (err error) {

	// view.ShowSplashView(view.SplashStartView)
	// view.ShowSplashView(view.SplashCreateView)
	// wp, err := view.InputNewParameters(3)
	wp, err := wallet.InputNewParameters(3)
	if err != nil {
		return
	}
	wa, err := wallet.NewWalletAccount(wp.SecretBytes(), wp.SaltBytes())
	//NewWalletAccount里面产生masterkey,然后产生account key包括公私钥
	fmt.Println("pubkey:"+wa.PublicKey+"\nprivatekey:"+wa.PrivateKey)
	if err != nil {
		return
	}
	var ws []*wallet.Wallet
	if vanity == "" {
		ws, err = wa.GenerateWallets(0, uint32(number))
		if err != nil {
			return
		}
	} else {
		var patterns []string
		patterns, err = wa.NormalizeVanities([]string{vanity})
		if err != nil {
			return
		}
		ws, err = wa.FindVanities(patterns, func(i, c, n uint32) bool {
			fmt.Printf("progress: %d, %d, %d\n", i, c, n)
			return (n >= number)
		})
	}
	if export == "" {
		for _, w := range ws {
			fmt.Printf("wallet (%d): \n", w.No)
			fmt.Println("  " + w.Private)
			fmt.Println("  " + w.Address)
		}
	} else {
		var f *os.File
		f, err = os.Create(export)
		if err != nil {
			return
		}
		defer f.Close()
		for _, w := range ws {
			f.WriteString(fmt.Sprintf("wallet(%d): \r\n", w.No))
			f.WriteString(fmt.Sprintf("   private: %s\r\n", w.Private))
			f.WriteString(fmt.Sprintf("   address: %s\r\n", w.Address))
		}
	}

	return
}
