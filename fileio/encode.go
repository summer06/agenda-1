package fileio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func logsome(s string) {
	fileName := "l.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Print("ERROR: ", err.Error)
	}
	logger := log.New(file, "[Fileio]", log.Llongfile)
	logger.Println("file io log:")
	logger.Println(s)
}

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
	logsome(filename)
}

func ReadFile(filename string) ([]map[string]interface{}, error) {
	logsome(filename)
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
		file, _ := os.Create(filename)
		defer file.Close()
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
