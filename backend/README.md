# Project structure and organization

The project structure looks something like:

- bzone/backend
    - **root** -> any config/non Go specific code 
    - cmd
        - web
        - another_application_which_does_not_exist_yet

    - internal
        - models, etc
    


The `cmd` directory will contain the application-specific code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the `cmd/web` directory.

The `internal` directory will contain the ancillary non-application-specific code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the database/Bumbal API models for the project.

## Why?

There are two big benefits:

1. It gives a clean separation between Go and non-Go assets. All the Go code we write will live exclusively under the `cmd` and `internal` directories, leaving the project root free to hold non-Go assets like UI files, makefiles and module definitions (including `go.mod` file). This can make things easier to manage when it comes to building and deploying the app later.

2. It scales really nicely if we want to add another executable application to our project. For example, we might want to add a CLI to automate some administrative tasks in the future. With this structure, we could create this CLI application under `cmd/cli` and it will be able to import and reuse all the code you’ve written under the internal directory.


