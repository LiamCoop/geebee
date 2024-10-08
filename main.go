package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
    // run git branch, get all branches
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

    // format output
	branches := strings.Split(out.String(), "\n")
    var cleanBranches []string

    // clean data
	for _, branch := range branches {
		trimmedBranch := strings.TrimSpace(branch)
		if trimmedBranch != "" {
			// '*' indicates the current branch, remove it if needed
            cleanBranches = append(cleanBranches, strings.TrimPrefix(trimmedBranch, "* "))
		}
	}
    

    
	// show checkbox, allow user to select
	var selectedBranches []string
	prompt := &survey.MultiSelect{
		Message: "Select branches to delete:",
		Options: cleanBranches,
	}
	err = survey.AskOne(prompt, &selectedBranches)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}


    // delete selected branches
    for _, branch := range selectedBranches {
        delCmd := exec.Command("git", "branch", "-D", branch)
        err := delCmd.Run()
        if err != nil {
            fmt.Printf("Error deleting branch %s: %v\n", branch, err)
        } else {
            fmt.Printf("Deleted branch: %s\n", branch)
        }
	} 
}
