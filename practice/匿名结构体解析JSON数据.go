package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //水平和垂直分辨率
}

type Battery struct {
	Capacity int //电池容量
}

func getJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchID: true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}
func main() {
	jsonData := getJsonData()
	fmt.Printf(string(jsonData))
	fmt.Println("--------------------")
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v", screenAndTouch)
}

/*
从结构体中解析出json
package main

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
)

var publicKey=[]byte("")
var privateKey=[]byte("")

func RsaEncrypto(data[]byte)  ([]byte,error){
	block, _ := pem.Decode(publicKey)
	pkixPublicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pkixPublicKey.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader,pub,data)

}

func RsaDecrypto(data[]byte)([]byte,error)  {
	block,_ := pem.Decode(privateKey)
	pkixPriKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return rsa.DecryptPKCS1v15(rand.Reader,pkixPriKey,data)

}

type Product struct {
	Name string
	ProductId int64
	Number int
	Price float64
	IsOnSale bool
}

//不用匿名结构体体的话，很麻烦，代码长？感觉用匿名结构体也没简单哪去

func main()  {
	p := Product{}
	p.Name="apple"
	p.ProductId=1
	p.Number=100
	p.Price=3.45
	p.IsOnSale=false
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	//json中解析出结构体
	var p1 Product
	json.Unmarshal(data, &p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.ProductId)
	fmt.Println(p1.Price)

}

*/
