package awsw

import (
	"fmt"
	"os"
	"strings"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")

	d, err := LoadYamlFile()
	if err != nil {
		panic(err)
	}

	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	if len(args) < 2 {
		fmt.Println("Please select envioment.")
		return
	} else if !(args[1] == "development" || args[1] == "production") {
		fmt.Println("Please select from [development / production]")
		return
	}
	project := Project{}
	env := Environment{}
	color := "99BCE3"

	for _, p := range d.Projects {
		if p.Name == args[0] {
			project = p
			break
		}
	}
	for _, e := range project.Envs {
		if e.Env == args[1] {
			env = e
			break
		}
	}
	if len(args) == 3 {
		switch args[2] {
		case "red":
			color = "F2B0A9"
		case "orange":
			color = "FBBF93"
		case "yellow":
			color = "FAD791"
		case "green":
			color = "B7CA9D"
		case "blue":
			color = "99BCE3"
		default:
			color = "99BCE3"
		}
	}
	awsBaseURL := "https://signin.aws.amazon.com/switchrole"
	fmt.Println("--------------------")
	fmt.Println("ProjectName: " + project.Name)
	fmt.Println("AccountID: " + env.ID)
	fmt.Println("Environment: " + env.Env)
	fmt.Println("RoleName: " + project.RoleName)
	str := fmt.Sprintf("%s?roleName=%s&account=%s&displayName=%s&color=%s",
		awsBaseURL,
		project.RoleName,
		env.ID,
		env.Name,
		color)
	fmt.Println("")
	fmt.Println(str)
	fmt.Println("--------------------")
	os.Exit(0)
}
