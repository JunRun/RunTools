/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-03-17 13:12
 */
package test

import (
	"fmt"
	"github.com/JunRun/RunTools/rfile"
	"testing"
)

func TestFile(t *testing.T) {

	rfile.FileRead("/Volumes/videos/crunchyroll_video", "")
	fmt.Println(len(rfile.VideoList))
	for _, v := range rfile.VideoList {
		fmt.Println(v.Name, v.Url)
	}

}
