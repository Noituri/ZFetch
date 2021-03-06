package main

import (
	"bufio"
	"fmt"
	sigar "github.com/cloudfoundry/gosigar"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type OsInfo struct {
	OS		 	 string
	Kernel	 	 string
	CPU		 	 string
	Cores	 	 string
	GPU		 	 string
	Shell	 	 string
	Terminal 	 string
	Hostname 	 string
	MaxRam	 	 string
	UsedRAM  	 string
	//MaxStorage	 string
	//UsedStorage	 string
	Username	 string
	Uptime		 string
}

func format(val uint64) uint64 {
	return val / 1024 / 1024
}

func (oi *OsInfo) GetInfo() {
	// Hostname
	hn, _ := os.Hostname()
	oi.Hostname = hn

	// OS name
	file, err := os.Open("/etc/os-release")

	if err != nil {
		log.Fatalf("Could not open '/etc/os-release'")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			osName := strings.TrimLeft(scanner.Text(), "PRETTY_NAME=")
			osName = strings.Replace(osName, `"`, "", -1)
			oi.OS = osName
			break
		}
	}

	_ = file.Close()

	// Kernel
	res, err := exec.Command("uname", "-r").Output()

	var kernel string

	if err != nil {
		kernel = ""
	} else {
		kernel = string(res)
	}

	oi.Kernel = kernel

	// Shell
	shell := strings.Split(os.Getenv("SHELL"), "/")
	oi.Shell = shell[len(shell) - 1]

	// User
	oi.Username = os.Getenv("USER")

	// Terminal
	oi.Terminal = os.Getenv("TERM")

	// CPU & CpuCores
	file, err = os.Open("/proc/cpuinfo")

	if err != nil {
		log.Fatalf("Could not open '/proc/cpuinfo'")
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "model name") {
			cpuModel := strings.Replace(scanner.Text(), "model name", "", -1)
			cpuModel = strings.Replace(cpuModel, `:`, "", -1)
			cpuModel = strings.TrimSpace(cpuModel)
			oi.CPU = cpuModel
		} else if strings.Contains(scanner.Text(), "cpu cores") {
			cpuCores := strings.Replace(scanner.Text(), "cpu cores", "", -1)
			cpuCores = strings.Replace(cpuCores, `:`, "", -1)
			cpuCores = strings.TrimSpace(cpuCores)
			oi.Cores = cpuCores
		}
	}

	_ = file.Close()

	// GPU
	gpuRegexp, _ := regexp.Compile(`(GeForce)|(3D)|(VGA)|(Display)|(ATI)\w+`)
	quotesRegexp, _ := regexp.Compile(`"(.*?)"`)

	res, err = exec.Command("lspci", "-mm").Output()
	lspciInfo := strings.Split(string(res), "\n")

	var gpu string
	if err != nil {
		gpu = ""
	} else {
		foundGPUs := []string{}

		for _, v := range lspciInfo {
			if gpuRegexp.MatchString(v) {
				foundGPUs = append(foundGPUs, v)
			}
		}

		gpu = strings.Replace(quotesRegexp.FindAllString(foundGPUs[len(foundGPUs) - 1], -1)[2], `"`, "", -1)
	}

	oi.GPU = gpu

	// Max RAM & Used RAM
	mem := sigar.Mem{}
	_ = mem.Get()

	oi.MaxRam = strconv.FormatInt(int64(format(mem.Total)), 10)
	oi.UsedRAM = strconv.FormatInt(int64(format(mem.ActualUsed)), 10)

	// Uptime
	uptime := sigar.Uptime{}
	_ = uptime.Get()

	oi.Uptime = uptime.Format()
}

func GetDefaultResponse() (string, error) {
	oi := OsInfo{}
	oi.GetInfo()

	finalResponse := fmt.Sprintf("OS: %s\n", oi.OS)
	finalResponse += fmt.Sprintf("Kernel: %s", oi.Kernel)
	finalResponse += fmt.Sprintf("Shell: %s\n", oi.Shell)
	finalResponse += fmt.Sprintf("User: %s\n", oi.Username)
	finalResponse += fmt.Sprintf("Hostname: %s\n", oi.Hostname)
	finalResponse += fmt.Sprintf("Terminal: %s\n", oi.Terminal)
	finalResponse += fmt.Sprintf("CPU: %s\n", oi.CPU)
	finalResponse += fmt.Sprintf("Cores: %s\n", oi.Cores)
	finalResponse += fmt.Sprintf("GPU: %s\n", oi.GPU)
	finalResponse += fmt.Sprintf("RAM: %s/%s\n", oi.UsedRAM, oi.MaxRam)
	finalResponse += fmt.Sprintf("Uptime: %s\n", oi.Uptime)

	return finalResponse, nil
}