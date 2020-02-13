package mystruct

type DataSource struct {
	ReqId     int64    `json:reqId`
	Code      int      `json:code`
	Msg       string   `json:msg`
	Data      ResultData 	`json:data`
	Timestamp int64    `json:timestamp`
}

type ResultData struct {
	ChinaTotal ChinaTotal `json:chinaTotal`
	ChinaDayList []ChinaDayList `json:chinaDayList`
	LastUpdateTime string `json:lastUpdateTime`
	AreaTree []AreaCountry `json:areaTree`
}

type ChinaTotal struct {
	Today TotalStruct `json:today`
	Total  TotalStruct `json:total`
}

type TotalStruct struct {
	Confirm int64 `json:confirm`
	Suspect int64 `json:suspect`
	Heal int64 `json:heal`
	Dead int64 `json:dead`
}

type ChinaDayList struct {
	Date string `json:date`
	Today TotalStruct `json:today`
	Total  TotalStruct `json:total`
}

//中国
type AreaCountry struct {
	Today TotalStruct `json:today`
	Total  TotalStruct `json:total`
	Name string `json:name`
	Id string `json:id`
	Children []AreaProvince `json:children`
}

//省区
type AreaProvince struct {
	Today TotalStruct `json:today`
	Total  TotalStruct `json:total`
	Name string `json:name`
	Id string `json:id`
	Children []struct{
		Today TotalStruct `json:today`
		Total  TotalStruct `json:total`
		Name string `json:name`
		Id string `json:id`
		Children []AreaCity `json:children`
	}`json:children`
}

//城市
type AreaCity struct {
	Today TotalStruct `json:today`
	Total  TotalStruct `json:total`
	Name string `json:name`
	Id string `json:id`
	Children struct{} `json:children`
}



