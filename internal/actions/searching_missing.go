package actions

import "elma_hw/internal/models"

func GetDataSearchingMissing(taskName string) ([]models.SearchingMissing, []interface{}, error) {
	arr, question, err := GetDataCyclicRotation(taskName)
	if err != nil {
		return nil, nil, err
	}
	cr := []models.SearchingMissing{}
	for _, el := range arr {
		cr = append(cr, el.ArrayIn)
	}
	return cr, question, nil

}

func SmResponse(result []models.SearchingMissing, question []interface{}) Response {
	res := Response{
		UserName: "KirillRukhlyadev",
		Task:     "Поиск отсутствующего элемента",
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
