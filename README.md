# A module generator written in go

This is a project/module generator written in go. It is intended to generate standard project/module layouts in the targetted languages.

### How to Use
 * Install the Modulenator package
 > go install github.com/cybermind-nick/modulenator
 * This should automatically add it to your PATH.
 If it doesn't, find it in your GOPATH and manually add it to your
 PATH

 * Create the folder in the desired location. The folder name is the project name 
 > mkdir my_project
 > cd my_project

* Run Modulenator . Specificing the github flag is a must for go projects
> modulenator go --github "githubUsername"




## This project currently only supports go, but more languages will be added as I improve on it when I can take time from my other committments. Contributions are welcome