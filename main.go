package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"net/http"
	"strings"
	"yiqing/mystruct"
)

type CitySlice struct {
	Province string
	Confirm  int
}

var city string
var ret mystruct.DataSource
var cityret mystruct.CityData

func main() {
	cityBytes := httpget(mystruct.CityUrl)
	bytes := httpget(mystruct.ApiUrl)
	json.Unmarshal(bytes, &ret)
	json.Unmarshal(cityBytes, &cityret)
	
	city = cityret.Data.Position.Province
	city = strings.Trim(city, "市")
	city = strings.Trim(city, "省")

	color.Set(color.BgWhite)
	color.FgWhite.Println(getEmptyStr(57))
	color.Set(color.BgWhite)
	color.Danger.Println(getEmptyStr(6),"全国数据截止更新时间:", ret.Data.LastUpdateTime,getEmptyStr(8))
	color.Set(color.BgWhite)
	color.FgWhite.Println(getEmptyStr(57))
	color.Set(color.BgBlue)

	color.FgWhite.Println(getEmptyStr(20),"历史数据",getEmptyStr(27))
	color.FgRed.Print("确诊:", ret.Data.ChinaTotal.Total.Confirm, "\t")
	color.FgYellow.Print("疑似:", ret.Data.ChinaTotal.Total.Suspect, "\t")
	color.FgGreen.Print("治愈:", ret.Data.ChinaTotal.Total.Heal, "\t")
	color.FgLightMagenta.Print("死亡:", ret.Data.ChinaTotal.Total.Dead, "\n")

	fmt.Println()
	color.Set(color.BgBlue)

	color.FgWhite.Println(getEmptyStr(20),"今日数据",getEmptyStr(27))
	color.FgRed.Print("确诊:", ret.Data.ChinaTotal.Today.Confirm, "\t")
	color.FgYellow.Print("疑似:", ret.Data.ChinaTotal.Today.Suspect, "\t")
	color.FgGreen.Print("治愈:", ret.Data.ChinaTotal.Today.Heal, "\t")
	color.FgLightMagenta.Print("死亡:", ret.Data.ChinaTotal.Today.Dead, "\n")


	//chinaData := ret.Data.AreaTree[0].Children
	fmt.Println()
	color.Set(color.BgBlue)
	color.FgWhite.Println(
		"地区", getEmptyStr(11),
		"确诊",getEmptyStr(6),
		"疑似",getEmptyStr(3),
		"治愈",getEmptyStr(4),
		"死亡",getEmptyStr(4))
	fmt.Println("－－－－－－－－－－－－－－－－－－－－－－－－－－－－－")
	for _,chinaProvince := range ret.Data.AreaTree[0].Children {
		if (chinaProvince.Name == city) {
			for _,chinaCity := range chinaProvince.Children {
				fmt.Printf("|%-s%s", chinaCity.Name, getFormatStr(chinaCity.Name))
				fmt.Printf("|%-9d", chinaCity.Today.Confirm)
				fmt.Printf( "|%-10d",chinaCity.Today.Suspect)
				fmt.Printf("|%-10d",chinaCity.Today.Heal)
				fmt.Printf("|%-11d|\n",chinaCity.Today.Dead)
			}
		}
	}
	fmt.Println("－－－－－－－－－－－－－－－－－－－－－－－－－－－－－")

}

func getFormatStr(str string)string{
	setMaxLen := 18; //9个汉字
	needEmpty := (setMaxLen - len(str))/3
	return getEmptyZHStr(needEmpty)
}



func getEmptyStr(num int)string {
	emptyStr := " "
	var retStr string
	for i := 0; i<num; i++ {
		retStr = emptyStr + retStr;
	}
	return retStr
}


func getEmptyZHStr(num int)string {
	emptyStr := "　"
	var retStr string
	for i := 0; i<num; i++ {
		retStr = emptyStr + retStr;
	}
	return retStr
}

func httpget(url string) []byte {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	return body
}
