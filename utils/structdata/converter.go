package structdata

func MapIStringToArray(data map[string]bool) []string {
	res := []string{}
	for i := range data {
		res = append(res, i)
	}
	return res
}

func ArrayToMapIString(data []string) map[string]bool {
	res := map[string]bool{}
	for _, i := range data {
		res[i] = true
	}
	return res
}