# runlog

A very simple service to track your runs.

The main goal of this project is to practise Golang. The backend is built with focus on minimal amount of AI and dependencies.

## Development

### Migration

Create a file called `{version}_{name}.{up|down}.sql` and add your SQL.

Run `task migrate:up`. This uses [taskfile.dev](https://taskfile.dev). 
If you don't feel like installing that, you can just copy out the relevant code and run it yourself.
