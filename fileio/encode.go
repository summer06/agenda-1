package fileio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"log"
	"os"
)

func StructToJson(s interface{}) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return b
}

func WriteFile(filename string, data interface{}) {
	var b []byte
	b = StructToJson(data)
	if b != nil {
		ioutil.WriteFile(filename, b, os.ModeAppend)
	}
}

func ReadFile(filename string) ([]map[string]interface{}, error) {
	if checkFileIsExist(filename) {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("ReadFile: ", err.Error())
			return nil, err
		}
		var xxx []map[string]interface{}
		if err := json.Unmarshal(bytes, &xxx); err != nil {
			fmt.Println("Unmarshal: ", err.Error())
			return nil, err
		}
		return xxx, nil
	} else {
		os.Create(filename)
		return nil, nil
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
