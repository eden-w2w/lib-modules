package gaode

type CommonRequest struct {
	Key string `in:"query" name:"key"`
}

type DistrictRequest struct {
	CommonRequest
	Keywords    string `in:"query" name:"keywords" default:""`
	SubDistrict int    `in:"query" name:"subdistrict" default:"1"`
	Page        int    `in:"query" name:"page" default:"1"`
	Offset      int    `in:"query" name:"offset" default:"20"`
	Extensions  string `in:"query" name:"extensions" default:"base"`
	Filter      string `in:"query" name:"filter" default:""`
	Callback    string `in:"query" name:"callback" default:""`
	Output      string `in:"query" name:"output" default:"JSON"`
}

type DistrictItem struct {
	CityCode  interface{}    `json:"citycode" default:""`
	ADCode    string         `json:"adcode" default:""`
	Name      string         `json:"name"`
	Polyline  string         `json:"polyline" default:""`
	Center    string         `json:"center"`
	Level     string         `json:"level"`
	Districts []DistrictItem `json:"districts"`
}

type DistrictResponse struct {
	Status     int    `json:"status,string"`
	Info       string `json:"info"`
	InfoCode   string `json:"infocode"`
	Suggestion struct {
		Keywords []string `json:"keywords"`
		Cities   []string `json:"cities"`
	} `json:"suggestion"`
	Districts []DistrictItem `json:"districts"`
}
