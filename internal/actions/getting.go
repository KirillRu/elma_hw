package actions

import (
	"elma_hw/internal/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ForceDisconnectAfter = 60 * time.Second
	URL_CYCLICROTATION   = "https://kuvaev-ituniversity.vps.elewise.com/tasks/%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F"
)

func GetDataCyclicRotation() ([]models.CyclicRotation, error) {
	var netClient = http.Client{
		Timeout: ForceDisconnectAfter,
	}
	res, err := netClient.Get(URL_CYCLICROTATION)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var answer []interface{}
	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		return nil, errors.New("No response in json format")
	}
	cr := []models.CyclicRotation{}
	for _, el := range answer {
		row := el.([]interface{})
		a := []int16{}
		for _, i := range row[0].([]interface{}) {
			a = append(a, int16(i.(float64)))
		}
		cr = append(cr, models.CyclicRotation{ArrayIn: a, RotationCounts: uint8(row[1].(float64))})
	}
	return cr, nil

}
