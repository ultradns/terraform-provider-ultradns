package record

import "strings"

func getMatchedRecordData(state []interface{}, server []string) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range state {
		dataMap[val.(string)] = true
	}

	for _, val := range server {
		if dataMap[val] {
			data = append(data, val)
		}
	}

	return data
}

func getUnMatchedRecordData(state []interface{}, server []string) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range state {
		dataMap[val.(string)] = true
	}

	for _, val := range server {
		if !dataMap[val] {
			data = append(data, val)
		}
	}

	return data
}

func getDiffRecordData(first []interface{}, second []interface{}) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range first {
		dataMap[val.(string)] = true
	}

	for _, val := range second {
		if !dataMap[val.(string)] {
			data = append(data, val.(string))
		}
	}

	return data
}

func rmRecordData(data, target []string) []string {
	dataMap := make(map[string]bool)

	for _, val := range data {
		dataMap[val] = true
	}

	for i, val := range target {
		if dataMap[val] {
			target[i] = target[len(target)-1]
			target[len(target)-1] = ""
			target = target[:len(target)-1]
		}
	}

	return target
}

func addRecordData(data, target []string) []string {
	target = append(target, data...)

	return target
}

func escapeSOAEmail(email string) string {
	index1 := strings.Index(email, "@")
	index2 := strings.LastIndex(email[:index1], ".")

	if index2 == -1 {
		return email[:index1] + "." + email[index1+1:]
	}

	return strings.Replace(email[:index2]+"\\"+email[index2:], "@", ".", 1)
}

func formatSOAEmail(email string) string {
	index := strings.Index(email, "\\.")

	if index == -1 {
		return strings.Replace(email, ".", "@", 1)
	}

	return email[:index] + "." + strings.Replace(email[index+2:], ".", "@", 1)
}
