package httpUtil

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)
//var gCurCookieJar *cookiejar.Jar
var gCurCookies []*http.Cookie

func  String(b []byte) string {    return *(*string)(unsafe.Pointer(&b))
}

func HttpRequest(url string,cook string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}
	}


	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

    //分割cookie
	sep:=";"
	sep2:="="
	strArr:=strings.Split(cook,sep)
	//fmt.Println("arr:",strArr)

	for i := 0; i < len(strArr); i++ {
		//fmt.Println("i=", i, strArr[i])
		//分割某项cookie
		strArr2:=strings.Split(strArr[i],sep2)
		//cookie := &http.Cookie{Name:strArr2[0],Value:strArr2[1]};
		//加载cookie
		request.AddCookie(&http.Cookie{Name:strArr2[0],Value:strArr2[1]})
	}


	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}


	//client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}
	}

	fmt.Printf("httpResp.Header=%s", response.Header)
	fmt.Printf("httpResp.Status=%s", response.Status)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}

	//var coonew []*http.Cookie = nil
	gCurCookies =  response.Cookies()
	//读取本地cookis
	//全局保存
	//gCurCookies = cookiejar.Cookies(request.URL)
	dbgPrintCurCookies()


	return body
}


func dbgPrintCurCookies() {
	var cookieNum int = len(gCurCookies)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		fmt.Printf("\n------ Cookie [%d]------", i)
		fmt.Printf("\tName=%s", curCk.Name)
		fmt.Printf("\tValue=%s", curCk.Value)
		fmt.Printf("\tPath=%s", curCk.Path)
		fmt.Printf("\tDomain=%s", curCk.Domain)
		fmt.Printf("\tExpires=%s", curCk.Expires)
		fmt.Printf("\tRawExpires=%s", curCk.RawExpires)
		fmt.Printf("\tMaxAge=%d", curCk.MaxAge)
		fmt.Printf("\tSecure=%t", curCk.Secure)
		fmt.Printf("\tHttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("\tRaw=%s", curCk.Raw)
		fmt.Printf("\tUnparsed=%s", curCk.Unparsed)
	}
}
