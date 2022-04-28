package docker

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func CheckDocker() error {
	_, err := exec.LookPath("docker")
	return err
}

func DockerAuth(registry string) error {
	fmt.Printf("Authorization in %s registry\n", registry)
	cmd := exec.Command("docker", "login", registry)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}

func UpdateImage(wg *sync.WaitGroup, o, n string) {
	if err := exec.Command("docker", "pull", o).Run(); err != nil {
		fmt.Printf("error running docker pull on %s, %+v\n", o, err)
	}
	if err := exec.Command("docker", "tag", o, n).Run(); err != nil {
		fmt.Printf("error running docker tag on %s, %+v\n", o, err)
	}
	if err := exec.Command("docker", "push", n).Run(); err != nil {
		fmt.Printf("error running docker push on %s, %+v\n", o, err)
	} else {
		fmt.Printf("Image: %s has successfully pushed.\n", n)
	}
	wg.Done()
}
