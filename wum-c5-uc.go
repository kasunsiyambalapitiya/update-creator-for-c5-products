package main

import (
	"os"
	"archive/zip"
	"sync"
	"net/http"
	"fmt"
	"path/filepath"
	"io"
)

//var sync sync.WaitGroup`

func main() {
	/*command := "diff --brief -r /home/kasun/Documents/Delete-Them/Updated-Products/C4-Products/Old/ " +
		"/home/kasun/Documents/Delete-Them/Updated-Products/C4-Products/New/"

	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)*/
	//test()
	Unzip("/home/kasun/Documents/Delete-Them/Updated-Products/C5-Products/c5-custom-product-5.2.0-update1.zip",
		"/tmp/c5-custom-product-5.2.0-update1")

}
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		//return error and exit
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		// this  is annonymous class, hence no name is there

		rc, err := f.Open()
		//rc has files contents
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc) // since rc has files content, we are copying it to the created empty
			// file
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		fmt.Println(f.Name+" created")
		err := extractAndWriteFile(f)
		if err != nil {
			// return and exit
			return err
		}
	}

	return nil
}

func getDiff() {

}

func test() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

