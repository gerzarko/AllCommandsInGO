package commands

import (
	"fmt"
	"os"
	// "reflect"
)


// type *os.unixDirent 

func Ls() {

	var f, err = os.Open("./")

	if err != nil {
		fmt.Println(f)

	}

	var files, error = f.ReadDir(0)

	if error != nil {
		fmt.Println(error)
		return

	}

	for n, v := range files {
		fmt.Println(n+1, "position", v.Name(), v.IsDir())

	}

}

func Grep(wordTosearch string){

    var f,err = os.Open("./")

    if err != nil{
        fmt.Print(f)
        return
    }

    var result,errResult = os.Create("report.txt")

    if errResult != nil {
        fmt.Print(errResult)
    }

    var _ , errorWriting = result.WriteString("report")

    if errorWriting != nil{
        fmt.Print(errorWriting)
    }

    var files, error = f.ReadDir(0)

    if error != nil {
        fmt.Print(error)
        return
    }
    fmt.Print(files)
    


}

func LsTree(){


    var f,err = os.Open("./")
    if err != nil{
        fmt.Print(f)
        return
    }

    // var result,errResult = os.Create("report.txt")
    // if errResult != nil {
    //     fmt.Print(errResult)
    // }


    // var totalResults := "hola"
    // var _ , errorWriting = result.WriteString(totalResults)

    // if errorWriting != nil{
    //     fmt.Print(errorWriting)
    // }

    var files, error = f.ReadDir(0)

    if error != nil {
        fmt.Print(error)
        return
    }

    for _,v := range files{

        if v.IsDir() == false {
         

        }
    }


}

func recursionPrint(){
    
}






    




