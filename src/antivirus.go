package main

import (
	"crypto/md5"
        "encoding/hex" 
	"fmt"
	"io"
	"log"
	"os"
)
//FIXME: I don't know go
func get_hash(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
func scan_file(file string,list []string) {
    for item:=0;item < len(list);item++{
        if get_hash(file) == list[item]{
		    fmt.Println("malware")
            break;
	    } else {
		    if item == len(list)-1{
                fmt.Println("not malware")
            }
        }
	}
}
