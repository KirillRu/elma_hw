package actions

import (
	"elma_hw/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CrResponse(result []models.CyclicRotation, question []interface{}) Response {
	res := Response{
		UserName: "KirillRukhlyadev",
		Task:     "Циклическая ротация",
		Results: Results{
			Payload: question,
		},
	}
	r := make([]interface{}, len(result))
	for i, el := range result {
		r[i] = el.Solution()
	}

	res.Results.Results = r
	return res
}

func GetDataCyclicRotation(taskName string) ([]models.CyclicRotation, []interface{}, error) {
	var netClient = http.Client{
		Timeout: ForceDisconnectAfter,
	}
	if taskName == "" {
		taskName = "%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F"
	}
	res, err := netClient.Get(URL_TASK + "/tasks/" + taskName)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var answer []interface{}
	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		return nil, nil, errors.New(fmt.Sprintf("No response in json format. Err: %s. %s", jsonErr.Error(), string(body)))
	}
	cr := []models.CyclicRotation{}
	for _, el := range answer {
		row := el.([]interface{})
		a := []int{}
		for _, i := range row[0].([]interface{}) {
			a = append(a, int(i.(float64)))
		}
		rotate := 0
		if len(row) > 1 {
			rotate = int(row[1].(float64))
		}
		cr = append(cr, models.CyclicRotation{ArrayIn: a, RotationCounts: rotate})
	}
	return cr, answer, nil

}
