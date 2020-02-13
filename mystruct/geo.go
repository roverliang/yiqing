package mystruct

type CityData struct {
	Message string `json:message`
	Data    Data   `json:data`
}

type Data struct {
	Mso      Mso      `json:mso`
	Ip       string   `json:ip`
	Position Position `json:position`
}

type Mso struct {
	Y string `json:y`
	X string `json:x`
}
type Position struct {
	Province    string `json:province`
	Isp         string `json:isp`
	Adcode      string `json:adcode`
	Area        string `json:area`
	Address     string `json:address`
	City        string `json:city`
	City_code   string `json:city_code`
	Country     string `json:country`
	Street      string `json:street`
	Weathercode string `json:weathercode`
}