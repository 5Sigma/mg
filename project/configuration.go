package project

//
// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"os"
// )
//
// type (
// 	Configuration struct {
// 		Driver           string `json:"driver"`
// 		MigrationPath    string `json:"migrationPath"`
// 		DatabaseHost     string `json:"databaseHost"`
// 		DatabasePort     string `json:"databasePort"`
// 		DatabaseName     string `json:"databaseName"`
// 		DatabasePath     string `json:"databasePath"`
// 		DatabaseUser     string `json:"databaseUser"`
// 		DatabasePassword string `json:"databasePassword"`
// 	}
// )
//
// func (c *Configuration) Json() []byte {
// 	b, err := json.Marshal(c)
// 	if err != nil {
// 		return []byte{}
// 	}
// 	var out bytes.Buffer
// 	json.Indent(&out, b, "", "    ")
// 	return out.Bytes()
// }
//
// func ReadConfiguration(filePath string) (*Configuration, error) {
// 	buf := bytes.NewBuffer(nil)
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	io.Copy(buf, f)
// 	f.Close()
// 	var c *Configuration
// 	json.Unmarshal(buf.Bytes(), &c)
// 	return c, nil
// }
//
// func (c *Configuration) Write(filePath string) {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	file.Write(c.Json())
// }
