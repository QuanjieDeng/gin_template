package   main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
}
func genAppIdKey()(string,string){

	//TODO 生成appid
	appid := CreateCaptcha()
	//TODO 生成appkey
	str := appid +"gvp"
	sum := sha256.Sum256([]byte(str))

	appkey:= fmt.Sprintf("%x",sum)

	return   appid,appkey
}


func  main(){
	a,b :=  genAppIdKey()
	fmt.Printf("a[%s]\n",a);
	fmt.Printf("b[%s]\n",b);
}
