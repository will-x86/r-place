This is the code for an r-place clone that can only be used via code:)

Elements:
cmd/ -- entrypoint for golang
pkg/ -- packages for golang
ui/ -- React userinterface for requests
curl/ -- Basic curl commands for interfacing
.air.toml -- hot reloading

# Cache:

- K:V are stored in appendonly valkey db
- On startup all key:pair's are fetched up to xmax and ymax ( see .env.example )
- Two caches are stored, both the fully marshalled

# Development:

## Enviroment variables

Copy .env.example to .env and fillout values

## Nix

- Use nix package manager and run direnv allow

```
(cd ui && npm install && npm run build) && docker compose up -d && air
``

## non-nix
- Things required:
    - Air
    - Golang
    - docker compose
    - npm
```

(cd ui && npm install && npm run build) && docker compose up -d && air
``

## Windows

uhhhhhhhhhhh todo later

# Common errors

- Panics on startup
  - most likely missing an env variable
- Valkey keeps restarting constantly ( most likely
  - Most likely missed enviroment variables, ` docker compose down -v` then restart
