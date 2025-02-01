new ver
## Installation

To clone the project, create a folder and use the git clone command. Then please read the [makefile](makefile) file to learn how to install all the tooling and docker images.

```
$ git clone https://github.com/Klimentin0/courses-service 
$ cd courses-service
```



Now you have a copy with your own module name. Now all you need to do is initialize the project for git.

## Running The Project

To run the project use the following commands.

```
# Install Tooling
$ make dev-gotooling
$ make dev-docker

# Run Tests
$ make test

# Shutdown Tests
$ make test-down

# Run Project
$ make dev-up
$ make dev-update-apply
$ make token
$ export TOKEN=<COPY TOKEN>
$ make users

# Run Load
$ make load

# Run Tooling
$ make grafana
$ make statsviz

# Shut Project
$ make dev-down
```