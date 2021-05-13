package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetCacheDir() string {
	if cacheDir, err := os.UserCacheDir(); err == nil {
		return path.Join(cacheDir, "jumpo")
	}
	return ""
}

func createIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModePerm)
	}
	return errors.New("Directory already exists")
}

func CreateConfig(dir string) string {
	configPath := ConfigFileLocation(dir)
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Println("Creating config file")
		file, err := os.Create(configPath)
		defer file.Close()
		if err != nil {
			log.Fatalf("Error in creating file %v\n", err)
		}
	}
	return configPath
}

func ConfigFileLocation(dir string) string {
	configPath := path.Join(dir, "config.json")
	return configPath
}

func checkConflict(table map[string]string, key string) bool {
	if _, ok := table[key]; ok {
		return true
	}
	return false
}

func AddKeyToStore(config string, obj Jumpo) {
	data := extractData(config)
	if !checkConflict(data, obj.Prefix) {
		data[obj.Prefix] = obj.Location
		writeData(config, data)
	} else {
		log.Fatalln("Prefix already present in cache")
	}
}

func RemoveKey(config string, pref string) {
	data := extractData(config)
	if _, ok := data[pref]; ok {
		delete(data, pref)
	}
	writeData(config, data)
}

func ListData(config string) {
	data := extractData(config)
	for k, v := range data {
		println(k, v)
	}
}

func FetchValue(config string, pref string) string {
	data := extractData(config)
	if val, ok := data[pref]; ok {
		return val
	}
	log.Fatalln("Prefix not found")
	return ""
}

func writeData(config string, data map[string]string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("Error in indenting data %v\n", err)
	}
	err = ioutil.WriteFile(config, file, os.ModePerm)
	if err != nil {
		log.Fatalf("Error in writing file: %v\n", err)
	}
}

func extractData(config string) map[string]string {
	jsonFile, err := os.OpenFile(config, os.O_RDWR, os.ModePerm)
	defer jsonFile.Close()
	if err != nil {
		log.Fatalf("Error in reading json file %v\n", err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error in fetching bytes %v\n", err)
	}
	var data map[string]string
	json.Unmarshal(byteValue, &data)
	return data
}
