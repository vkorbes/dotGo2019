package main

import (
	"encoding/base64"
	"log"
	"os"
	"os/exec"
)

func svgtopng(input, output string) {

	file, err := os.Create(output)
	if err != nil {
		log.Fatalln(err)
	}

	if input == "" {
		empty, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z/C/HgAGgwJ/lK3Q6wAAAABJRU5ErkJggg==")
		_, err = file.Write(empty)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	cmd := exec.Command("inkscape", "-z", "-e", output, input, "--export-area-drawing", "--export-height=1000")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
