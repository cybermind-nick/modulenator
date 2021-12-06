package golang

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cybermind-nick/modulenator/utils"
)

// behaviours in language folders are inherited from base.

// go project object
type GoProject struct {
	ProjectPath string
	ProjectName string
	Github      string
}

func NewProject(github string) error {
	dir, err := utils.DirFolder()
	if err != nil {
		log.Fatal(err)
	}

	module := "github.com/" + github + "/" + dir
	fmt.Println(module)
	cmd := exec.Command("go", "mod", "init", module)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Could not create project. An error occured, or a project may already exist")
		return err
	}

	cmd = exec.Command("git", "init")
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to initialize repo. One may already exist in this dirctory")
		return err
	}

	return nil
}

func NewProjectWithPath(github, path string) error {
	os.MkdirAll(path, 0777)
	dir := utils.DirPathFolder(path)

	cd := exec.Command("cd", path)
	_, err := cd.CombinedOutput()
	if err != nil {
		log.Fatal("Could not change directories")
	}

	module := "github.com/" + github + "/" + dir
	cmd := exec.Command("go", path, "mod", "init", module)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Could not create project, check your path: ", path, err)
		return err
	}

	return nil
}
