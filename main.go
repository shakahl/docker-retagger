package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"docker-retagger/pkg/docker"
	"docker-retagger/pkg/images"
)

var origin, newRegistry, file string
var skiplogin bool

func init() {
	flag.StringVar(&origin, "image", "", "the original image to pull")
	flag.StringVar(&newRegistry, "newregistry", "", "the registry to set the image to")
	flag.StringVar(&file, "file", "", "use an input file rather than one-off flags")
	flag.BoolVar(&skiplogin, "skiplogin", false, "if you want continue without login in original registry")
	flag.Parse()

	if err := docker.CheckDocker(); err != nil {
		log.Fatalf("docker does not appear to be installed, %+v", err)
	}
}

func belongsToList(list []string, lookup string) bool {
	for _, val := range list {
		if val == lookup {
			return true
		}
	}
	return false
}

func main() {
	var img *images.Image
	var originalRegistry string
	var allRegistries []string
	d := make(map[*images.Image]string)

	if file != "" {
		if newRegistry == "" {
			log.Fatal("--newregistry flag must not be \"\"")
		}
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), "\n")
			img, err = images.ParseImage(split[0])
			if err != nil {
				log.Fatal(err)
			}
			d[img] = split[0]
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		switch {
		case origin == "":
			log.Fatal("--image flag must not be \"\"")
		case newRegistry == "":
			log.Fatal("--newregistry flag must not be \"\"")
		default:
			var err error
			img, err = images.ParseImage(origin)
			if err != nil {
				log.Fatal(err)
			}
			d[img] = newRegistry
		}
	}

	if err := docker.DockerAuth(newRegistry); err != nil {
		log.Fatalf("Failed to authorization in %s", newRegistry)
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
		originalRegistry = k.Registry
		k.Registry = newRegistry
		k.User = ""
		newImage, err := k.Marshal()
		if err != nil {
			log.Fatal(err)
		}
		if !skiplogin && !belongsToList(allRegistries, originalRegistry) {
			allRegistries = append(allRegistries, originalRegistry)
			if err := docker.DockerAuth(originalRegistry); err != nil {
				log.Fatalf("Failed to authorization in %s", originalRegistry)
			}
		}

		go docker.UpdateImage(wg, originalImage, newImage)
	}
	wg.Wait()
}
