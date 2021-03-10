package responses

type ResponseBool struct {
	Data bool `json:"data" xml:"data"`
}

type ResponseCommonSingle struct {
	Data interface{} `json:"data" xml:"data"`
}

type ResponseCommonArray struct {
	Data       interface{} `json:"data" xml:"data"`
	TotalCount int         `json:"total_count" xml:"total_count"`
}
