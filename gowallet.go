package main
/* 
#cgo  CFLAGS:  -I  /root/bip44cxx 
#cgo  LDFLAGS:  -L /root/bip44cxx  -lbip44wallet -lbitcoin -lbitcoin-client
#include <stdlib.h>
#include "interface.h" 
*/  
import "C"    
import (
	"flag"
	"fmt"
	"os"
	// "./view"
	"unsafe"
	"./wallet"
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


func FromMnemonicToMasterKey(mnemonic string)string {	
	csRet:=C.FromMnemonicToMasterKey(C.CString(mnemonic))
	defer C.free(unsafe.Pointer(csRet))
    return C.GoString(csRet)
}

func test2() {
	p := new(Prefixes)
	p.setPrefixes("LTC")
	// fmt.Println(p.bip44_code)
	wallet := New3("label stick flat innocent brother frost rebel aim creek six baby copper need side cannon student announce alpha",*p);
    ret:=wallet.getMnemonic()
    fmt.Println(ret)

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
