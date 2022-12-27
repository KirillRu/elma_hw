package actions

import (
	"elma_hw/internal/models"
)

func GetDataWonderfulOccurrences(taskName string) ([]models.WonderfulOccurrences, []interface{}, error) {
	arr, question, err := GetDataCyclicRotation(taskName)
	if err != nil {
		return nil, nil, err
	}
	cr := []models.WonderfulOccurrences{}
	for _, el := range arr {
		cr = append(cr, el.ArrayIn)
	}
	return cr, question, nil

}

func WoResponse(result []models.WonderfulOccurrences, question []interface{}) Response {
	res := Response{
		UserName: "KirillRukhlyadev",
		Task:     "Чудные вхождения в массив",
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
