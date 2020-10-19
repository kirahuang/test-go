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
var gresponseCookies []*http.Cookie

var gcookiesss string
var strArr []string

var sep =";"

var sep2 ="="

func  String(b []byte) string {    return *(*string)(unsafe.Pointer(&b))
}

func HttpTotal(url string){
	gcookiesss ="BIDUPSID=F2300C28CFF396DD0B725B58844A2167; PSTM=1601267805; BAIDUID=F2300C28CFF396DD4C01B20B23431C26:FG=1; BD_UPN=12314753; BDUSS=VaNm1na0MtbGcwU2x1MDJzZWlxcEpEdUp4RmNZSH52eW5uTGRvRTJIaFQ0N0JmSVFBQUFBJCQAAAAAAAAAAAEAAADTWC0OaHVhbmd5aW5oYWlhYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFNWiV9TVolfbl; BDUSS_BFESS=VaNm1na0MtbGcwU2x1MDJzZWlxcEpEdUp4RmNZSH52eW5uTGRvRTJIaFQ0N0JmSVFBQUFBJCQAAAAAAAAAAAEAAADTWC0OaHVhbmd5aW5oYWlhYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFNWiV9TVolfbl; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BD_HOME=1; delPer=0; BD_CK_SAM=1; PSINO=5; BDRCVFR[feWj1Vr5u3D]=mk3SLVN4HKm; BA_HECTOR=2k218k2h04a421bcig1foqd040l; H_PS_PSSID=7506_1460_32877_7566_31254_32723_32231_7517_7605_32115_31709_26350; sugstore=1"
	HttpRequest(url)

	HttpRequest(url)

}



func HttpRequest(url string)  {

	//func HttpRequest(url string) []byte {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("\n=11=====================系统出错啦==================")
	}


	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	fmt.Printf("\n=====================Requestcookie=%s", gcookiesss)
    //分割cookie
	strArr=strings.Split(gcookiesss,sep)
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
		fmt.Printf("\n=22=====================系统出错啦==================")
	}

	fmt.Printf("\nhttpResp.Header=%s", response.Header)
	fmt.Printf("\nhttpResp.Status=%s", response.Status)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("\n=33=====================系统出错啦==================")
	}

	//var coonew []*http.Cookie = nil
	gresponseCookies =  response.Cookies()
	//读取本地cookis
	//全局保存
	//gCurCookies = cookiejar.Cookies(request.URL)
	//dbgPrintCurCookies()

	handlenewreponsecookie()

	//打印html
	str := String(body)
	//bytes[0] = 'i'//注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
	fmt.Println(str)
	//return body
}


func handlenewreponsecookie() {
	var cookieNum int = len(gresponseCookies)
	var build strings.Builder
	//fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gresponseCookies[i]
		//fmt.Printf("\n------ Cookie [%d]------", i)

		for i := 0; i < len(strArr); i++ {
			//fmt.Println("i=", i, strArr[i])
			//分割某项cookie
			strArr3:=strings.Split(strArr[i],sep2)

			//cookie替换
			if strArr3[0]==curCk.Name{
				strArr3[1]=curCk.Value
				//替换父数组
				strArr[i]= strArr3[0]+sep2+ strArr3[1]
			}else{
				//新增
				build.WriteString(sep+curCk.Name +sep2+curCk.Value)

			}


		}

	}

	gcookiesss= strings.Join(strArr, sep)+build.String()
	fmt.Printf("\n=====================ReponseNewcookie=%s", gcookiesss)
}

func dbgPrintCurCookies() {
	var cookieNum int = len(gresponseCookies)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gresponseCookies[i]
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
