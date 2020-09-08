package spider

type DataInfo struct {
	Cid          string `json:"cid"`
	BlockHeight  int    `json:"block_height"`
	TimesTamp    int    `json:"timestamp"`
	TimesTampStr string `json:"timestamp_str"`
	To           string `json:"to"`
}

type Info struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  DataI  `json:"data"`
}

type DataI struct {
	Data []DataInfo `json:"data"`
}
