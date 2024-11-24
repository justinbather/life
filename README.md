# Life - Track workouts, meals, and macros, all from the terminal

## Install the CLI
`go install github.com/justinbather/life/life@latest`

## Getting around
- `life-server/`
> Contains the server
> Currently privately hosted until auth is baked into the cli.

- `life/`
> Contains the CLI.
> See readme in there for usage.

## Lazy?
Fork the repo, deploy anywhere (or host locally)
All you need to do is configure the `API_URL` at `~/.life.yaml` (if deploying), and setup postgres

## Make it not suck
- Build in auth, with global config
- Generate a monthly or weekly report of everything
- Reminders to do this stuff
