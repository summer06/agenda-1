package fileio

import (
	"encodng/json"
	"fmt"
	"io/ioutil"
)

func StructToJson(s interface{}) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return b
	}
}

func ReadFile(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return xxx, nil
}
