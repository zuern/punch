/*
   Punch - The simple CLI timekeeper.
   Copyright © 2017 Kevin Zuern <kevin@kevinzuern.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"time"
)

// initialize will check the user's home directory for the presence of a `.punch` folder, creating it if not already present
func initialize(punchPath string) {
	if fileExists(punchPath) == false {
		err := os.Mkdir(punchPath, 0755)
		check(err)
	}
}

// Returns true if the file at the specified path exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil && !os.IsExist(err)
}

func check(err error) {
	checkWithMessage(err, err.Error(), true)
}

// Gracefully exit the program if an error has occured
func checkWithMessage(err error, errorMessage string, exitOnError bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", errorMessage)
		if exitOnError {
			os.Exit(1)
		}
	}
}

// Returns the current system user's name as a string
func getUserName() string {
	user, err := user.Current()
	check(err)
	return user.Username
}

// Returns the current time as a string
func getCurrentTimeString() string {
	now := time.Now()
	return now.Format(time.UnixDate)
}

// Create an entry in the correct punchCard, using the verb specified.
// Verbs should be either "IN" or "OUT"
func punch(punchPath string, projectName string, verb string) {
	punchCardPath := punchPath + "/" + projectName + ".punch"

	// Open punchCard in append mode, or if it doesn't exist, create it then open
	punchCard, err := os.OpenFile(punchCardPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	check(err)

	// Set up the punch message
	punchMessage := fmt.Sprintf("%v\t%v\t%v\n", getUserName(), verb, getCurrentTimeString())

	punchCard.Seek(0, 0)
	punchCard.WriteString(punchMessage)
	punchCard.Close()

	fmt.Print(punchMessage)

}

// List all the punch cards in the directory specified by punchPath
func listPunchCards(punchPath string) {
	files, err := ioutil.ReadDir(punchPath)
	check(err)

	listedFiles := false

	for _, file := range files {
		listedFiles = true
		fileName := file.Name()
		extension := path.Ext(fileName)
		if extension == ".punch" {
			fmt.Println(fileName[0 : len(fileName)-len(extension)])
		}
	}

	if listedFiles == false {
		fmt.Println("No punch cards yet.")
	}
}

// Outputs the contents of the punchCard.
func printPunchCard(punchPath string, punchCard string) {
	punchCardPath := punchPath + "/" + punchCard + ".punch"

	bytes, err := ioutil.ReadFile(punchCardPath)
	checkWithMessage(err, "The punch card \""+punchCard+"\" could not be found.", true)

	text := string(bytes)

	fmt.Print(text)

}

func main() {
	usage := `Punch - The simple CLI timekeeper.
Usage:
  punch (-h | -v)
  punch (in|out|log) [<project_name>]
  punch list
Examples:
  punch in
  punch in "ACME Website Development"
  punch out "ACME Website Development"
  punch log "ACME Website Development"
Options:
  -h, --help
  -v, --version
`
	// Validate the program args, Parse the program args into a dict, and handle -h and -v
	arguments, _ := docopt.Parse(usage, nil, true, "Punch 0.0.0", false)

	homeDir, _ := homedir.Dir()
	punchPath := homeDir + "/.punch"

	initialize(punchPath)

	// Get the project name, or the default option
	projectName := "default"
	if arguments["<project_name>"] != nil {
		projectName = arguments["<project_name>"].(string)
	}

	// Execute the correct command
	if arguments["list"].(bool) {
		listPunchCards(punchPath)
	} else if arguments["in"].(bool) {
		punch(punchPath, projectName, "IN")
	} else if arguments["out"].(bool) {
		punch(punchPath, projectName, "OUT")
	} else {
		printPunchCard(punchPath, projectName)
	}
}
