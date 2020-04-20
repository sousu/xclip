package main

import (
    "os"
    "fmt"
    "time"
    "io/ioutil"
    "regexp"
    "./lib"
)
func main() {
    if len(os.Args) < 2 {
        panic("usage: xclip [log_dir]")
    }
    path := os.Args[1]
    fmt.Print("Log dir path [",path,"]\n")
    clipboard.EachText(func(txt string){
        t := time.Now().Format("20060102_030405")
        x := GetXml(txt)
        if x != "" {
            ioutil.WriteFile(path+"\\"+t+".xml",[]byte(x),os.ModePerm)
            time.Sleep(1 * time.Second)
            ioutil.WriteFile(path+"\\"+"latest.xml",[]byte(x),os.ModePerm)
            fmt.Printf("[%s] Save\n",t)
        }else{
            fmt.Printf("[%s] Not Match\n",t)
        }
    })
}
func GetXml(txt string) string{
    //fmt.Print(txt)
    r := regexp.MustCompile(`^(<\?xml.+?\?>[\s\S]+)$`) //[\s\S]=>改行等全て
    if r.MatchString(txt){
        return r.ReplaceAllString(txt,"$1")
    }else{
        return ""
    }
}

