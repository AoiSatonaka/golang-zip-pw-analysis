package main

import (
	"flag"
	"fmt"
	"github.com/yeka/zip"
	"io/ioutil"
	"log"
	"math"
)

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	rangeFlag:=flag.Bool("r",false,"false=>\"0-9\",5digit,true=>\"0-9\"+\"a-z,A-Z\",8digit")

	flag.Parse()

	//sourceDir := getInput("対象ファイルのパスを入力してください")
	sourceDir := flag.Arg(0)

	log.Print(sourceDir)

	r,err :=zip.OpenReader(sourceDir)

	if err !=nil {
		panic(err)
	}

	for _,file := range r.File {
		if file.IsEncrypted(){
			passwordFormat := "%05d"
			var maxCombiCount int
			if *rangeFlag {
				maxCombiCount = int(math.Pow(float64(len(letters)),float64(8)))
				passwordFormat="%08d"
			}else{
				maxCombiCount = 100000
			}
			//startTime := time.Now()
			for i:=0; i<maxCombiCount; i++ {
				//if i!=0 && i%10000 == 0 {
				//	fmt.Printf("解析終了まで最大約%d分\n",time.Now().Sub(startTime)*time.Duration(int64(maxCombiCount/i))/time.Minute)
				//}
				password := fmt.Sprintf(passwordFormat,i)
				file.SetPassword(password)
				rc,errOpen:= file.Open()
				if errOpen!=nil {
					panic(errOpen.Error())
				}
				_,errRead := ioutil.ReadAll(rc)
				rc.Close()
				fmt.Printf("%t\t%s\n",errRead==nil,password)
				if errRead != nil {
					continue
				}else{
					return
				}
			}
		}else{
			panic("this file is not Encrypted")
		}
	}
}

//func getInput(hint string) string {
//	stdin := bufio.NewScanner(os.Stdin)
//	fmt.Print(hint,":")
//	if stdin.Scan(){
//		return stdin.Text()
//	}else{
//		panic("Standard Input Error")
//	}
//}
