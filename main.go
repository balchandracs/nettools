package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	PROCNET_DEV_PATH string = "/proc/net/dev"
)

var version int

func main() {
	devfile, err := os.Open(PROCNET_DEV_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer devfile.Close()

	scanner := bufio.NewScanner(devfile)

	linecount := 0
	for scanner.Scan() {
		text := scanner.Text()
		linecount++

		if linecount == 2 {
			version = procnetversion(text)
		}
		//skip the headers containing names
		if strings.Contains(text, ":") {
			iname, _ := getInterface(text)
			addiface(&iface{name: iname})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	listiface()
}

func procnetversion(buffer string) int {
	if strings.Contains(buffer, "compressed") {
		return 3
	}
	if strings.Contains(buffer, "bytes") {
		return 2
	}
	return 1
}

func getInterface(buff string) (string, string) {

	splitstring := strings.Split(buff, ":")
	name := strings.ReplaceAll(splitstring[0], " ", "")
	if name == buff {
		return "", ""
	}
	return name, splitstring[1]
}

/*
func setInterfaceStat(buff string, inface *iface) {
	splitstring := strings.Split(buff, " ")

	if version == 3 {
		inface.stats.rx_bytes, _ = (uint64)strconv.Atoi(splitstring[0])
	}
	if version == 2 {

	}
	if version == 1 {

	}

} */
