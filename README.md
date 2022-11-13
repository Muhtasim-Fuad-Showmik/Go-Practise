# Go Practise

## Initial Setup

Make sure you have go correctly installed on the system by running the command
"go version" and receiving a proper output displaying the version of Go language
installed on your system.
Run "go mod tidy" on the projects to install any dependencies if any to get
the required packages setup inside the project. This step is similar to the
"npm i" command for node projects.

## GO Current Working Directory Setup

If the project fails to run and looks for files for commands in the default "src" 
folder within the "C:/Program Files/Go" directory, then you need to setup your
GOPATH and GOROOT first.
Open your start menu and search "env". Open up the "Edit the system environment
variables" program and then click on "Environment Variables". Select the "GOPATH"
variable and define it to be your own project folder so that your files are
looked for in the correct working directory.
Finally open the terminal within your project directory and run:
"go env -w GO111MODULE=off"
At this point your projects should continue to run commands on the correct 
directory.

## Initializing new projects

You can run "go mod init "