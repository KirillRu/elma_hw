package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	ForceDisconnectAfter = 60 * time.Second
	URL_TASK             = "https://kuvaev-ituniversity.vps.elewise.com"
)

type Results struct {
	Payload []interface{} `json:"payload"`
	Results []interface{} `json:"results"`
}

type CheckResult struct {
	Percent int           `json:"percent"`
	Fails   []interface{} `json:"fails"`
}

type Response struct {
	UserName string  `json:"user_name"`
	Task     string  `json:"task"`
	Results  Results `json:"results"`
}

func SendData(result Response) (CheckResult, error) {
	var netClient = http.Client{
		Timeout: ForceDisconnectAfter,
	}
	responseContent, err := json.Marshal(result)
	if err != nil {
		return CheckResult{}, err
	}

	res, err := netClient.Post(URL_TASK+"/tasks/solution", "application/json", bytes.NewReader(responseContent))
	if err != nil {
		return CheckResult{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CheckResult{}, err
	}

	var answer CheckResult
	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		return CheckResult{}, errors.New(fmt.Sprintf("No response in json format. Err: %s. %s", jsonErr.Error(), string(body)))
	}
	return answer, nil
}

var wg sync.WaitGroup

func GetAll() []CheckResult {
	allResult := []CheckResult{}

	r := make(chan CheckResult, 1)
	go func(res *[]CheckResult) {
		saveCheck(r, res)
	}(&allResult)

	go func(r chan CheckResult) {
		wg.Add(1)
		arr, question, err := GetDataCyclicRotation("%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F")
		if err == nil {
			check, err := SendData(CrResponse(arr, question))
			if err == nil {
				r <- check
			}
		}
		wg.Done()
	}(r)
	go func(r chan CheckResult) {
		wg.Add(1)
		arr, question, err := GetDataWonderfulOccurrences("%D0%A7%D1%83%D0%B4%D0%BD%D1%8B%D0%B5%20%D0%B2%D1%85%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%B8%D1%8F%20%D0%B2%20%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2")
		if err == nil {
			check, err := SendData(WoResponse(arr, question))
			if err == nil {
				r <- check
			}
		}
		wg.Done()
	}(r)
	go func(r chan CheckResult) {
		wg.Add(1)
		arr, question, err := GetDataCheckingSequence("%D0%9F%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%BA%D0%B0%20%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8")
		if err == nil {
			check, err := SendData(CsResponse(arr, question))
			if err == nil {
				r <- check
			}
		}
		wg.Done()
	}(r)
	go func(r chan CheckResult) {
		wg.Add(1)
		arr, question, err := GetDataSearchingMissing("%D0%9F%D0%BE%D0%B8%D1%81%D0%BA%20%D0%BE%D1%82%D1%81%D1%83%D1%82%D1%81%D1%82%D0%B2%D1%83%D1%8E%D1%89%D0%B5%D0%B3%D0%BE%20%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%B0")
		if err == nil {
			check, err := SendData(SmResponse(arr, question))
			if err == nil {
				r <- check
			}
		}
		wg.Done()
	}(r)
	wg.Wait()
	i := 0
	for len(allResult) < 4 && i < 50 {
		i++
		//I give up
		time.Sleep(time.Second)
	}
	return allResult
}

func saveCheck(r chan CheckResult, res *[]CheckResult) {
	for {
		*res = append(*res, <-r)
	}
}
