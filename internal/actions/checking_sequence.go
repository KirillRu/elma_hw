package actions

import "elma_hw/internal/models"

func GetDataCheckingSequence(taskName string) ([]models.CheckingSequence, []interface{}, error) {
	arr, question, err := GetDataCyclicRotation(taskName)
	if err != nil {
		return nil, nil, err
	}
	cr := []models.CheckingSequence{}
	for _, el := range arr {
		cr = append(cr, el.ArrayIn)
	}
	return cr, question, nil

}

func CsResponse(result []models.CheckingSequence, question []interface{}) Response {
	res := Response{
		UserName: "KirillRukhlyadev",
		Task:     "Проверка последовательности",
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
