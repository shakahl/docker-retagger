package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func checkDocker() error {
	_, err := exec.LookPath("docker")
	return err
}

func updateImage(wg *sync.WaitGroup, o, n string) {
	if err := exec.Command("docker", "pull", o).Run(); err != nil {
		fmt.Printf("error running docker pull on %s, %+v\n", o, err)
	}
	if err := exec.Command("docker", "tag", o, n).Run(); err != nil {
		fmt.Printf("error running docker tag on %s, %+v\n", o, err)
	}
	if err := exec.Command("docker", "push", n).Run(); err != nil {
		fmt.Printf("error running docker push on %s, %+v\n", o, err)
	}
	fmt.Printf("Image: %s has successfully pushed.\n", n)
	wg.Done()
}
