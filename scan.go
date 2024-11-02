package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

// appends all the found git repos and encloses them in a folder .gogitlocalstats
func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

// Opens file located at filepath, creates it and accesses it if missing.
func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			// other error
			panic(err)
		}
	}
	return f
}

// opens the file from the selected file path
// Gets all content in each line, parses it, and slices them into seperate strings

func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	return lines
}

// checks if slice contains some thing -> returns true, else returns false
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// joins slices if the existing []string does not contain the things in the new []string, returns existing
func joinSlice(new []string, existing []string) []string {

	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// dumps (overwrites current data or appends to existing the strings found from joinSlice into
func dumpStrsToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

// combines existing repos and with new repos and dumps them to the file, thus overwriting the existing content
func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	completeRepos := joinSlice(existingRepos, newRepos)
	dumpStrsToFile(completeRepos, filePath)
}

// recursive scan function to call back the scanGitFolders to call in main func (scan)
func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// essential function that scans selected dir / folders for any .git files to add the folder and path to the

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, error := os.Open(folder)
	if error != nil {
		log.Fatal(error)
	}
	files, error := f.Readdir(-1)
	f.Close()
	if error != nil {
		log.Fatal(error)
	}
	var path string

	for _, file := range files {

		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "node_modules" || file.Name() == "vendor" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}

	}
	return folders
}

// main scan func which starts the recursive search of git repositories
// living in the `folder` subtree
// i.e. scans a new folder for git repos

func scan(folder string) {
	fmt.Printf("\nFolders found:\n\n")
	repos := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repos)
	fmt.Printf("\n\nSuccessfully added.\n")
}
