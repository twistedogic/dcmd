package docker

import (
	"os/exec"
	"regexp"
	"strconv"
)

const docker = "docker"

func dockerCmd(args ...string) *exec.Cmd {
	return exec.Command(docker, args...)
}

func CreateCmd(image string, vol, args []string) *exec.Cmd {
	dockerArgs := []string{"run", "-it", "--rm"}
	dockerArgs = append(dockerArgs, vol...)
	dockerArgs = append(dockerArgs, image)
	dockerArgs = append(dockerArgs, args...)
	return dockerCmd(dockerArgs...)
}

func HasEntrypoint(image string) bool {
	args := []string{"image", "inspect", "-f", "'{{.Config.Entrypoint}}'", image}
	out, err := dockerCmd(args...).CombinedOutput()
	if err != nil {
		return false
	}
	return string(out) == "[]"
}

func GetExposePorts(image string) []int {
	args := []string{"image", "inspect", "-f", "'{{.Config.ExposedPorts}}'", image}
	out, err := dockerCmd(args...).CombinedOutput()
	if err != nil {
		return make([]int, 0)
	}
	pattern := regexp.MustCompile("([0-9]+)")
	portString := pattern.FindAllString(string(out), -1)
	ports := make([]int, len(portString))
	for i, v := range portString {
		ports[i], _ = strconv.Atoi(v)
	}
	return ports
}
