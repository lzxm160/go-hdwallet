package encryption

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	ts_data := []struct {
		Key     string
		MasterKey       string
	}{
		{
			Key:     "123456",
			MasterKey:       "xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH",
		},
		{
			Key:     "1234565555",
			MasterKey:       "xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH",
		},
	}
		wa:= Encrypt([]byte("123456"), []byte("xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH"))
		fmt.Printf("len(Encrypt):%d\n",len(wa))
		fmt.Println("Encrypt:",wa)
		fmt.Println("MasterKey:xprv9s21ZrQH143K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWSfFPHhKar6nqGYxukp5BjrvFqCjoLTxQ9izmBmes4sSR7KH")
	for _, v := range ts_data {
		fmt.Println("-------------------------------------------------")
		ret:=Validate([]byte(v.Key),wa)
		if ret{
			fmt.Println("password ok")
		}else{
			fmt.Println("password fail")
		}
		fmt.Println("######################################################")
	}
}

