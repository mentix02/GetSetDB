// common utilities and functions
// used again and again in the code
package main

import (
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// commonly used error both
// for spaceCommands as well
// as for databaseCommands
func commandError(command string) error {
	return errors.New("command `" + command + "` not recognized")
}

// commonly used error
// for spaceCommands to
// inform for lack of
// database name in the
// command query
func databaseNameError(command string) error {
	return errors.New("database name not specified for command `" + command + "`")
}

func removeExtension(fileName string) string {
	return fileName[0:len(fileName) - len(filepath.Ext(fileName))]
}

// simply panic if
// error is found
// QUITS the server
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// checks for a
// character within
// a string
func charInString(text, del string) bool {
	for _, char := range text {
		if del == string(char) {
			return true
		}
	}
	return false
}

// regex function to
// check whether argument
// provided is a uuid or not
func isUUID(uuid string) bool {

	// pattern copied again shamelessly
	// from https://bit.ly/2Ruz66f
	pattern := "\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b"
	matched, _ := regexp.MatchString(pattern, uuid) // check exactly for pattern matching

	return matched
}

// function to split strings
// according to a delimiter
func splitString(command string, del string) []string {
	return strings.Split(command, del)
}

// a linear search function for
// iterating over a list of strings
// usually containing database names
// to compare with the database
// entered by the user to give
// commands on the specified database
func stringInSlice(command string, list []string) bool {
	for _, b := range list {
		if b == command {
			return true
		}
	}
	return false
}

// return full path of the datbase
// along with the file extension
func path(database string) string {
	return "/tmp/gsdb/" + database + ".gsdb"
}

// simply return the first string
// of the command string which
// would specify the database name
func extractFirstTerm(command string) string {
	return strings.Split(command, " ")[0]
}

// return array of strings of
// the command string splitted
// with delemiter of space except
// the first which is the database
// name on which the command is to
// be executed upon
func extractCommandFromDatabaseCommand(command string) []string {
	return strings.Split(command, " ")[1:]
}

// optimized line counter copied
// shamelessly from https://bit.ly/2TkIEOy
func lineCounter(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count + 1

		case err != nil:
			return count + 1
		}
	}
}

// checks if string
// provided is a float
// or an integer
func isNumeric(num string) bool {
	_, err := strconv.ParseFloat(num ,64)
	return err == nil
}

// infers the datatype from
// string provided as a value
// of a key pair on a single line
func dataInferer(value string) string {

	if charInString(value, " ") { 					  // check if there's a space in the value
		return "list"
	} else if isNumeric(value) { 						  // check if the value can be converted to a number
		return "number"
	} else if isUUID(value) {							  // check if the value matches a UUID regexp
		return "uuid"
	} else { 											  // return String as default value
		return "string"
	}

}
