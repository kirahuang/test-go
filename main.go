package main

import (
	"fmt"
	httpUtil "github.com/kirahuang/test-go/src/http"
)

func main() {

	var countflag int32 = 1
	countflag = 3
	fmt.Println(countflag)

	httpUtil.HttpTotal("https://www.baidu.com")

	//chrome 拿cookies
	//coo :="BIDUPSID=F2300C28CFF396DD0B725B58844A2167; PSTM=1601267805; BAIDUID=F2300C28CFF396DD4C01B20B23431C26:FG=1; BD_UPN=12314753; BDUSS=VaNm1na0MtbGcwU2x1MDJzZWlxcEpEdUp4RmNZSH52eW5uTGRvRTJIaFQ0N0JmSVFBQUFBJCQAAAAAAAAAAAEAAADTWC0OaHVhbmd5aW5oYWlhYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFNWiV9TVolfbl; BDUSS_BFESS=VaNm1na0MtbGcwU2x1MDJzZWlxcEpEdUp4RmNZSH52eW5uTGRvRTJIaFQ0N0JmSVFBQUFBJCQAAAAAAAAAAAEAAADTWC0OaHVhbmd5aW5oYWlhYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFNWiV9TVolfbl; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BD_HOME=1; delPer=0; BD_CK_SAM=1; PSINO=5; BDRCVFR[feWj1Vr5u3D]=mk3SLVN4HKm; BA_HECTOR=2k218k2h04a421bcig1foqd040l; H_PS_PSSID=7506_1460_32877_7566_31254_32723_32231_7517_7605_32115_31709_26350; sugstore=1"
	//bytes:=httpUtil.HttpRequest("https://www.baidu.com")
	//str := httpUtil.String(bytes)
	////bytes[0] = 'i'//注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
	//fmt.Println(str)

	////需要先导入strings包
	//s1 := "字符串"
	//s2 := "拼接"
	////定义一个字符串数组包含上述的字符串
	//var str2 []string = []string{s1, s2}
	////调用Join函数
	//s3 := strings.Join(str2, ";")
	//fmt.Print(s3)



//	s:="BIDUPSID=F2300C28CFF396DD0B725B58844A2167; PSTM=1601267805; BAIDUID=F2300C28CFF396DD4C01B20B23431C26:FG=1; BD_UPN=12314753; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; delPer=0; BD_CK_SAM=1; PSINO=5; COOKIE_SESSION=1019_0_9_9_9_8_0_0_9_4_6_0_0_0_273_0_1602817342_0_1602818088%7C9%230_1_1601287320%7C1; H_PS_PSSID=7506_1460_31254_32723_32231_7517_7605_32115_31709_26350; sugstore=1; H_PS_645EC=78bbuf3Nbeqc2SOikN1lNlThMG87ya89YcVb7JuxUGzK0c63Ymi6s%2BYDl%2B2NFsQOdTw%2F; BA_HECTOR=a08g8g2l8g8l21a7iu1foilil0i"
//	sep:=";"
//	strArr:=strings.Split(s,sep)
//	fmt.Println("arr:",strArr)
//
//	for i := 0; i < len(strArr); i++ {
//		fmt.Println("i=", i, strArr[i])
//	}
}
