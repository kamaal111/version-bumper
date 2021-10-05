package main

import (
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

	buildNumberString := os.Getenv("BUILD_NUMBER")
	if buildNumberString == "" {
		log.Fatalln("now build number provided")
	}

	buildNumber, err := strconv.Atoi(buildNumberString)
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
