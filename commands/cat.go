package commands

import (
	"fmt"
	"io"
	"log"
	"os"
)


func Cat()  {


    f,err := os.Open("./")

    if err != nil{
        log.Fatal(err) 
    }

    files,error := f.ReadDir(0)

    if error != nil{
        log.Fatal(error) 
    }

    fmt.Println(files)

    fileToCat := os.Args[1]
    


    for _,v := range files{
        if v.Name() == fileToCat{
            mostrarContenido(v.Name())            
        }     
    }
    
}

func mostrarContenido(fileName string){


    f,err := os.Open(fileName)

    if err != nil{
        log.Fatal(err) 
    }

    defer f.Close()

    buf := make([]byte,1024)

    for {
        n, err := f.Read(buf)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(err)
            continue
        }
        if n > 0 {
            fmt.Println(string(buf[:n]))
        }
    }

}
