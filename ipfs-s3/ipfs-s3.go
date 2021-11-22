package ipfs_s3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type s3x struct {
	Url  string
	Auth string
}

func (s *s3x) GetBucketInfo(name string) {

	requestUrl := fmt.Sprintf("%s/info?bucket=%s", s.Url, name)
	response, _ := requestBody("GET", requestUrl, "", nil)

	fmt.Println(response)
}

func (s *s3x) ListBuckets() {
	requestUrl := fmt.Sprintf("%s/listBuckets", s.Url)

	headerMap := make(map[string]string)
	headerMap["Authorization"] = s.Auth

	response, _ := requestBody("GET", requestUrl, "", headerMap)
	fmt.Println(response)

}

func (s *s3x) DeleteBucket(bucket string) {

	requestUrl := fmt.Sprintf("%s/deleteBucket?bucket=%s", s.Url, bucket)
	response, _ := requestBody("DELETE", requestUrl, "", nil)

	fmt.Println(response)
}

func (s *s3x) PutBucket(bucket string) {
	requestUrl := fmt.Sprintf("%s/putBucket?bukcet=%s", s.Url, bucket)

	headerMap := make(map[string]string)
	headerMap["Authorization"] = s.Auth

	response, _ := requestBody("PUT", requestUrl, "", headerMap)
	fmt.Println(response)
}

func (s *s3x) PutObject(bucket, Object string) {

}

func requestBody(method, url, body string, headerMap map[string]string) (string, error) {
	client := http.Client{
		//超时时间为5s
		Timeout: 9 * time.Second,
	}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if string(body) == "" {
			return resp.Status, nil
		} else {
			return string(body), nil
		}
	}
}
