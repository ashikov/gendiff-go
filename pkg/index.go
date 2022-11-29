package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func isFileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	 } else {
		message := fmt.Sprintf("file %s does not exist", path)
		return false, errors.New(message)
	 }
}

func parse(data []byte, s *interface{}) {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()

	err := d.Decode(&s)
	if err != nil {
		log.Fatal("file contains invalid JSON")
	}
}

func GenDiff(path1, path2, formatter string) string {
	absPath1, err := filepath.Abs(path1)
		if err != nil {
			log.Fatal(err)
		}

		absPath2, err := filepath.Abs(path2)
		if err != nil {
			log.Fatal(err)
		}

		_, err = isFileExists(absPath1)
		if err != nil {
			log.Fatal(err)
		}

		_, err = isFileExists(absPath2)
		if err != nil {
			log.Fatal(err)
		}

		data1, err := os.ReadFile(absPath1)
		if err != nil {
			log.Fatal(err)
		}

		data2, err := os.ReadFile(absPath2)
		if err != nil {
			log.Fatal(err)
		}

		var s1, s2 interface{}

		parse(data1, &s1)
		parse(data2, &s2)

		return ""
}
