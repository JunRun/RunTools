package tron

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/JunRun/RunTools/spider"
	"io/ioutil"
	"net/http"
	"testing"
)

type Account struct {
}
type GenerateAddress struct {
	PrivateKey string
	Address    string
	HexAddress string
}

const Tron_Url = "https://api.shasta.trongrid.io"

func TestPostAddress(t *testing.T) {
	//生成地址
	result, _ := http.Get(Tron_Url + "/wallet/generateaddress")
	var generateAddress GenerateAddress
	s, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(s))

	json.Unmarshal(s, &generateAddress)

	fmt.Println(generateAddress.Address)

}

func CreateAddress() {
	res, _ := spider.PostBody(Tron_Url+"/wallet/createaddress", "{\"value\":\"20202\"}")
	fmt.Println(res)
}

//创建地址
func TestCreateAddress(t *testing.T) {
	CreateAddress()

	//结果
	//{"base58checkAddress":"TAQX8qgRB9LmbQv7zBLVCsDjG9FrqiSMZD","value":"4104c98963a7f5a0cee9d7dabdf1841a583425cfc7"}
	//{"base58checkAddress":"TDWX5pz8uBR5RgLEvPbVYdhVvAe5B2b9fJ","value":"4126d45dcfa21092993127adca1f19dab1c66de508"}
}

//校验地址
func TestValidateAddress(t *testing.T) {
	oldAddress := &GenerateAddress{
		PrivateKey: "bf91697e3e638db3515863176e6f18d68a5019d93fe52d5e5d81e09fb269ef2d",
		Address:    "TWRq3Q4isUcpBnfa157n5d5dsqrER19hPw",
		HexAddress: "4126d45dcfa21092993127adca1f19dab1c66de508",
	}

	//js,_:=json.Marshal(oldAddress)
	//fmt.Println(string(js))
	//校验地址
	res, err := spider.PostBody(Tron_Url+"/wallet/validateaddress", "{'address':"+"'"+oldAddress.HexAddress+"'}")
	//res, err:= http.Post(Tron_Url+"/wallet/validateaddress", "application/json", strings.NewReader("'address':"+"'"+oldAddress.HexAddress+"'"))
	//res,err:=http.PostForm(Tron_Url+"/wallet/validateaddress",url.Values{"address":{oldAddress.HexAddress}})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func EasyTransferByPrivate(privateKey, toAddress, amount string) {
	v := make(map[string]string)

	byteData := []byte(privateKey)
	hexS := hex.EncodeToString(byteData)
	byteAddress := []byte(toAddress)
	address := hex.EncodeToString(byteAddress)

	fmt.Println(hexS)
	fmt.Println(address)
	v["privateKey"] = hexS
	v["toAddress"] = address
	v["amount"] = amount

	s, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, _ := spider.PostBody(Tron_Url+"/wallet/easytransferbyprivate", string(s))
	fmt.Println(res)
}

func TestEasy(t *testing.T) {
	EasyTransferByPrivate("25ba9c10421f6089742300c23505befb2abd9f54121f5fceb13df77522172e5d", "41150a23eec262f4a59d7a10d778481fee83742df7", "100")
}
