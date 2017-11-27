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
	}

	for _, v := range ts_data {
		wa:= Encrypt([]byte(v.Key), []byte(v.MasterKey))
		fmt.Printf("len(Encrypt):%d\n",len(wa))
		fmt.Println("Encrypt:",wa)
		fmt.Println("-------------------------------------------------")
		ret:=DecryptAndValidate([]byte(v.Key),wa)
		if ret{
			fmt.Println("true")
		}
	}
}

