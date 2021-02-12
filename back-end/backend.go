package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())    //seed: 認証番号のために、現時点の時間をrandomな数字の生成のseedにする
	verCodeMap = make(map[string]int)   //学生: id、認証番号
	verCodeMapTc = make(map[int]string) //先生: 認証番号、id
	info = make(map[string]*person)     //Info: id、{name, phoneNumber...}
	arrInfo()                           //information(DB)を用意

	//gin: framework
	router := gin.Default()
	router.GET("/login/:id", login)
	router.GET("/view/:id", view)
	router.RunTLS(":8081", "cert.pem", "key.pem") //port, certFile, keyFile
}

var verCode int
var verCodeMap map[string]int
var verCodeMapTc map[int]string
var info map[string]*person

type person struct {
	name     string
	phoNum   string
	subArray []string
	scoArray []int
}

func arrInfo() {
	info["tc1"] = &person{name: "神田", phoNum: "00000000001"}
	info["tc2"] = &person{name: "上田", phoNum: "00000000002"}
	info["st001"] = &person{"上村", "00000000111", []string{"数学", "物理", "国語", "英語"}, []int{100, 100, 100, 100}}
	info["st002"] = &person{"中村", "00000000222", []string{"数学", "物理", "国語"}, []int{90, 90, 90}}
	info["st003"] = &person{"下村", "00000000333", []string{"物理", "英語"}, []int{85, 85}}
}

//1【login/:id】-> 認証番号、学生/先生: 名前
func login(c *gin.Context) {
	id := c.Param("id")
	perInfo, ok := info[id]

	if !ok {
		c.String(http.StatusOK, "id does not exist.")
		return
	}

	verCode = rand.Intn(9000) + 1000
	if id[:2] == "tc" { // 先生
		verCodeMapTc[verCode] = id
	} else { // 学生
		verCodeMap[id] = verCode
	}
	c.String(http.StatusOK, perInfo.name)
	fmt.Println("verification code: ", verCode)
}

// 2【/view/:id】
func view(c *gin.Context) {
	id := c.Param("id")
	// inVerCodeStr := c.Param("vercode")
	inVerCodeStr := c.Query("vercode") //【/view/:id(?vercode=xxx)】なら

	inVerCode, _ := strconv.Atoi(inVerCodeStr)
	verCodeSt, okSt := verCodeMap[id]     // 学生: id -> 認証番号
	tcId, okTc := verCodeMapTc[inVerCode] // 先生: 認証番号 -> id, ok

	fmt.Println("0", id, inVerCodeStr, inVerCode, verCodeSt, okTc) //test
	if (inVerCode != verCodeSt && !okTc) || verCode == 0 || inVerCode == 0 || inVerCodeStr == "" {
		fmt.Println("error: please log in.")
		c.String(http.StatusOK, "please log in.")
		return
	}

	// 【1】/view/:tcId/:tcVercode -> 学生list: id, name
	if okTc && tcId == id {
		fmt.Println("1", id, inVerCodeStr, inVerCode, verCodeSt, okTc) //test
		stList := ""
		for id, perInfo := range info {
			if id[:2] == "st" {
				stList += id + "    " + perInfo.name + "\n"
			}
		}
		c.String(http.StatusOK, stList)
		return
	}
	// 【2】/view/:stId/:tcVercode -> 成績list: 学科名 点数
	//     /view/:stId/:stVercode
	if (okTc && tcId != id) || (okSt && verCodeSt == inVerCode) {
		perInfo, _ := info[id]
		scoList := "学科名 点数\n"
		for i := 0; i < len(perInfo.subArray); i++ {
			scoList += perInfo.subArray[i] + "  " + strconv.Itoa(perInfo.scoArray[i]) + "\n"
		}
		c.String(http.StatusOK, scoList)
		return
	}
}
