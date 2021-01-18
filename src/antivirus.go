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
func scan_file(file string,data []string) {
    for n:=0;n < len(data);n++{
        if get_hash(file) == data[n]{
		    fmt.Println("malware")
            break;
	    } else {
		    if n == len(data)-1{
                fmt.Println("not malware")
            }
        }
	}
}
