package main

import (
	"net/http"
	"net/url"
)

func createFivePatients() error{
	if exist := fivePatientsExist(); exist {
		return nil
	} else {
		patients := []map[string]string {
			{"Id": "1", "Name": "小明"},
			{"Id": "2", "Name": "小蔡"},
			{"Id": "3", "Name": "小楊"},
			{"Id": "4", "Name": "小李"},
			{"Id": "5", "Name": "小高"},
		}

		for _, patient := range patients {
			_, err := http.PostForm("http://localhost:8888/patients/create", url.Values{"Id": {patient["Id"]}, "Name": {patient["Name"]}})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func fivePatientsExist() bool {
	resp, _ := http.Get("http://localhost:8888/patients")
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func main() {
	err := createFivePatients()
		if err != nil {
			panic("failed to create five patients")
		}
}
