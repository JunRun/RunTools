package spider

import (
	"encoding/json"
	"fmt"
	"testing"
)

//运用map 将 ok:=map[string] 触发 就进行时间的计算
func TestSpider(t *testing.T) {
	node := HttpFetchDoc("https://filscoutv3api.ipfsunion.cn/message/list?address=t01287&page=1&method=PreCommitSector&page_size=1000", "GET")
	sp := node.FirstChild.LastChild.FirstChild

	info := &Info{}
	err := json.Unmarshal([]byte(sp.Data), info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sp.Data)
}
