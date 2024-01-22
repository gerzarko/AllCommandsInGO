package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Cd() string {

	f, err := os.Open("./")

	if err != nil {
		fmt.Println(err)

	}

	files, error := f.ReadDir(0)

	if error != nil {
		fmt.Println(error)
	}

	for _, v := range files {
		if v.Type().IsDir() {
			fmt.Println(v)
		}
	}

	response1 := ""

	_, errorr := fmt.Scan(&response1)

	if errorr != nil {
		log.Fatal(errorr)
	}

	locationDir := -1

	for i, v := range files {
		if v.Name() == response1 {
			locationDir = i
		}
	}


    if locationDir != 0{
        pwd,error := os.Getwd()
        if error != nil{
            log.Fatal(error)
            os.Exit(-1)
        }
        err := os.Chdir(filepath.Join(pwd,response1))
        if err != nil{
            log.Fatal(err)
        }


    }
    
        pwd,error := os.Getwd()
        if error != nil{
            log.Fatal(error)
            os.Exit(-1)
        }

	return pwd

}
