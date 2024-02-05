package commands

// func Cat()  {
//
//     start := time.Now()
//
//     defer func(){
//         fmt.Println(time.Since(start))
//
//     }()
//
//     f,err := os.Open("./")
//
//     if err != nil{
//         log.Fatal(err)
//     }
//
//     files,error := f.ReadDir(0)
//
//     if error != nil{
//         log.Fatal(error)
//     }
//
//     fmt.Println(files)
//
//     fileToCat := os.Args[1]
//
//     channel := make(chan string,len(files))
//    //
//    //
//
//
//     for _,v := range files{
//         fmt.Println(reflect.TypeOf(v))
//         go func(v os.DirEntry){
//             if v.Name() == fileToCat{
//             channel <- mostrarContenido(v.Name())
//             }
//         }(v)
//     }
//
// }
//
// func mostrarContenido(fileName string){
//
//
//     f,err := os.Open(fileName)
//
//     if err != nil{
//         log.Fatal(err)
//     }
//
//     defer f.Close()
//
//     buf := make([]byte,8)
//
//
//     for i := 0;i<4;i++{
//         n, err := f.Read(buf)
//         if err == io.EOF {
//             break
//         }
//         if err != nil {
//             fmt.Println(err)
//             continue
//         }
//         if n > 0 {
//             fmt.Println(string(buf[:n]))
//         }
//     }
//
// }
