package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

var origin, newOrigin, inFile string

func init() {
	flag.StringVar(&origin, "image", "", "the original image to pull")
	flag.StringVar(&newOrigin, "new-origin", "", "the origin to set the image to")
	flag.StringVar(&inFile, "infile", "", "use an input file rather than one-off flags")
	flag.Parse()
}

func main() {
	var img *Image
	d := make(map[*Image]string)

	if inFile != "" {
		f, err := os.Open(inFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), " ")
			img, err = parseImage(split[0])
			if err != nil {
				log.Fatal(err)
			}
			d[img] = split[1]
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		switch {
		case origin == "":
			log.Fatal("--image flag must not be \"\"")
		case newOrigin == "":
			log.Fatal("--new-origin flag must not be \"\"")
		default:
			var err error
			img, err = parseImage(origin)
			if err != nil {
				log.Fatal(err)
			}
			d[img] = newOrigin
		}
	}

	if err := checkDocker(); err != nil {
		log.Fatalf("docker does not appear to be installed, %+v", err)
	}

	wg := &sync.WaitGroup{}
	for k, v := range d {
		wg.Add(1)
		originalImage, err := k.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		_, err = url.Parse(v)
		if err != nil {
			log.Fatalf("%s is not a valid origin url", v)
		}

		k.Registry = v
		newImage, err := k.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		go updateImage(wg, originalImage, newImage)
	}
	wg.Wait()
}
