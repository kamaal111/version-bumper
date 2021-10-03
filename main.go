package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	xCodeProjBytes, err := ioutil.ReadFile("../../Password-Generator.xcodeproj/project.pbxproj")
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
			newString := fmt.Sprintf("%sCURRENT_PROJECT_VERSION = %d;", tabsToAdd, 5)
			if lines[lineNumber] != newString {
				lines[lineNumber] = newString
				hasChanges = true
			}
		}
	}

	if hasChanges {
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile("../../Password-Generator.xcodeproj/project.pbxproj", []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
}
