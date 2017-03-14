package jsonutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "   ")
	log.Println(string(b))
}

func PrettyWrite(v interface{}, filePath string) {
	b, _ := json.MarshalIndent(v, "", "   ")
	ioutil.WriteFile(filePath, b, 0644)
}

func RawWrite(v interface{}, filePath string) {
	b, _ := json.Marshal(v)
	ioutil.WriteFile(filePath, b, 0644)
}

func RawPrint(v interface{}) {
	b, _ := json.Marshal(v)
	log.Println(string(b))
}
