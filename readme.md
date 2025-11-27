# runlog

A very simple service to track your runs.

The main goal of this project is to practise Golang. The backend is built with focus on minimal amount of AI and dependencies.

The idea is to create a simple frontend to go with this, and then self-host and use it as a form of run journaling. Like strava,
but a little more personal.

So far, the experience of using `http` and `sql` directly as opposed to using `chi` or `sqlx` hasn't been great, and I see why
they're so popular. Just wanted to try to do it myself to see how it was. Also because people seem to love the go stdlib.

## Development

Run postgres using the docker compose file; `docker compose up -d postgres`. 

### Migration

Create a file called `{version}_{name}.{up|down}.sql` and add your SQL.

Run `task migrate:up`. This uses [taskfile.dev](https://taskfile.dev). 
If you don't feel like installing that, you can just copy out the relevant code and run it yourself.
