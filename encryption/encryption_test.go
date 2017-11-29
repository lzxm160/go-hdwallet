package encryption

import (
	"fmt"
	"testing"
)

// func Test1(t *testing.T) {
// 	ts_data := []struct {
// 		Key     string
// 		MasterKey       string
// 	}{
// 		{
// 			Key:     "123456",
// 			MasterKey:       "xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH",
// 		},
// 		{
// 			Key:     "1234565555",
// 			MasterKey:       "xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH",
// 		},
// 	}
// 		wa,err:= Encrypt("123456", "xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH")
// 		if err!=nil{
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("len(Encrypt):%d\n",len(wa))
// 		fmt.Println("Encrypt:",wa)
// 		fmt.Println("MasterKey:xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH")
// 	for _, v := range ts_data {
// 		fmt.Println("-------------------------------------------------")
// 		ret:=Validate(v.Key,wa)
// 		if ret{
// 			fmt.Println("password ok")
// 		}else{
// 			fmt.Println("password fail")
// 		}
// 		fmt.Println("######################################################")
// 	}
// }

func Test2(t *testing.T) {
	fmt.Println("test2-------------------------------------------------")
	wa,err:= Encrypt("1234567890", "xprv9s21ZrQH143K3M9e2Baq8wVXSMDbRhtEudJBMBB1y7EtRmYyxdHqnfGeFEUyn7CPZx82pTVE7HDTQtgW4MBE4hX6Qu1pzksZ9YyLkse7W4T")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Printf("len(Encrypt):%d\n",len(wa))
		fmt.Println("Encrypt:",wa)
}
func Test3(t *testing.T) {
	fmt.Println("test3-------------------------------------------------")
	wa,err:= Encrypt("123456789", "xprv9s21ZrQH143K3M9e2Baq8wVXSMDbRhtEudJBMBB1y7EtRmYyxdHqnfGeFEUyn7CPZx82pTVE7HDTQtgW4MBE4hX6Qu1pzksZ9YyLkse7W4T")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Printf("len(Encrypt):%d\n",len(wa))
		fmt.Println("Encrypt:",wa)
}