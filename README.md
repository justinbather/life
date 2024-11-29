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
    - Auth is partially implemented in the server, would like to make this a bit better
    - Sidequest of a small lib to make the middleware stuff a bit nicer
    - Then will need to implement this in the client
        - Thinking of a config helper with username and pass, or just put that in .life.yaml
        - then can store the jwt in here too, reading that in for each request, if its expired, re-log in, if server sends a reauthenticate, re log in
            - logging in would consist of hitting the login endpoint, and storing the jwt and expiry

- Generate a monthly or weekly report that gets emailed or something (recap?)
- Would be interesting to build a reverse proxy for this service too to add in rate limiting
