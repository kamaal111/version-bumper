package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	buildNumberString := initializeFlag("build number", "", "build-number", "bn")
	infoPlistPath := initializeFlag("info plist path", "", "info-plist", "ip")
	flag.Parse()

	if *buildNumberString == "" {
		log.Fatalln("no build number provided, you can provide a build number by giving -build-number or -bn as a argument with a build number")
	}
	if *infoPlistPath == "" {
		log.Fatalln("no info plist path has been provided, you can provide a build number by giving -info-plist or -ip as a argument with a build number")
	}

	buildNumber, err := strconv.Atoi(*buildNumberString)
	if err != nil {
		log.Fatalln(err)
	}

	infoPlistFile, err := os.Open(*infoPlistPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer infoPlistFile.Close()

	_, err = ioutil.ReadAll(infoPlistFile)
	if err != nil {
		log.Fatalln(err)
	}

	projectFiles, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}

	var foundProjectFile string
	for _, projectFile := range projectFiles {
		if strings.Contains(projectFile.Name(), "xcodeproj") && projectFile.IsDir() {
			foundProjectFile = filepath.Join(".", projectFile.Name(), "project.pbxproj")
			break
		}
	}

	if foundProjectFile == "" {
		log.Fatalln("could not find project in directory")
	}

	xCodeProjBytes, err := ioutil.ReadFile(foundProjectFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(xCodeProjBytes), "\n")
	hasChanges := false
	for lineNumber, line := range lines {
		if strings.Contains(line, "CURRENT_PROJECT_VERSION = ") {
			oneTab := "	"
			amountOfTabs := strings.Count(line, oneTab)
			tabsToAdd := ""
			for i := 0; i < amountOfTabs; i++ {
				tabsToAdd += oneTab
			}
			newString := fmt.Sprintf("%sCURRENT_PROJECT_VERSION = %d;", tabsToAdd, buildNumber)
			if lines[lineNumber] != newString {
				lines[lineNumber] = newString
				hasChanges = true
			}
		}
	}

	if hasChanges {
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(foundProjectFile, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("done bumping build %s\n", elapsed)
}

func initializeFlag(usage string, flagDefault string, longVariable string, shortVariable string) *string {
	var value string
	flag.StringVar(&value, longVariable, flagDefault, usage)
	flag.StringVar(&value, shortVariable, flagDefault, usage)
	return &value
}
