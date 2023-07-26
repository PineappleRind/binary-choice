package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Data []string

func readPath() (data []byte, valid bool) {
	if len(os.Args) < 2 {
		return []byte{}, false
	}
	path := os.Args[1]

	contents, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("Specified data path does not exist, using default data")
		time.Sleep(time.Second)
		return []byte{}, false
	}

	return contents, true
}

func getData() Data {
	contents, valid := readPath()
	var data []string
	if valid == false {
		return []string{"Apples", "Oranges", "Bananas"}
	}
	if err := json.Unmarshal(contents, &data); err != nil {
		panic(err)
	}
	return data
}
