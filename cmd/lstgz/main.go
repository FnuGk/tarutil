package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fnugk/tarutil"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Must provide atleast one tar.gz file")
	}

	for _, path := range args {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		tarReader, err := tarutil.GZipToTar(f)
		if err != nil {
			log.Fatal(err)
		}

		err = tarutil.DoEachFile(tarReader, func(name string, r io.Reader) error {
			fmt.Println(name)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
