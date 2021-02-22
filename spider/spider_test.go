package spider

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
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

func TestGetFil(t *testing.T) {
	mienrList := []string{"f023985", "f023984", "f023982", "f023462", "f023534", "f023983", "f02529", "f02613", "f03002"}

	for _, minerId := range mienrList {
		////*[@id="__layout"]/div/div[1]/div[1]/div/div[3]/div[2]/div[1]/div/div[2]/p[5]
		node := HttpFetchDoc("https://filfox.info/zh/address/"+minerId, "GET")
		fmt.Print(minerId)
		sector := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[3]/div[2]/div[1]/div/div[2]/p[5]/text()")
		fmt.Print(sector.Data)
		data := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[3]/div[2]/div[1]/div/div[2]/p[6]/text()")
		fmt.Print(data.Data)
		owner := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[6]/div/div[2]/div[2]/a/@href")
		fmt.Print(GetFil(owner.FirstChild.Data))
		worker := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[6]/div/div[2]/div[3]/a/@href")
		fmt.Println(GetFil(worker.FirstChild.Data))
	}

}

func TestOwner(t *testing.T) {
	minerId := "f023985"
	node := HttpFetchDoc("https://filfox.info/zh/address/"+minerId, "GET")
	fmt.Print(minerId)
	owner := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[6]/div/div[2]/div[2]/a/@href")
	fmt.Print(GetFil(owner.FirstChild.Data))
	worker := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div/div[6]/div/div[2]/div[3]/a/@href")
	fmt.Println(GetFil(worker.FirstChild.Data))
}

func GetFil(address string) string {
	node := HttpFetchDoc("https://filfox.info"+address, "GET")
	data := htmlquery.FindOne(node, "//*[@id=\"__layout\"]/div/div[1]/div[1]/div[2]/dl[4]/dd/text()")
	return data.Data
}
