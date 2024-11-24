# Life CLI

## Install
`go install github.com/justinbather/life/life@latest`

## Setup
> Note: The CLI wont work unless running the server locally, until auth is released.

## Usage

`life <action> <resource> [flags]`

### Available actions
`new` : Creates a new resource. See `life new --help` for help
`get` : Gets the resource. See `life get --help` for help.

### Available resources
`workout`/`workouts`: A workout.
`meal`/`meals`: A meal:
`macros`: An aggregate view of workouts and meals within a timeframe

### Examples
`life create meal --calories 1000 --carbs 100 --protein 120 --fat 32 --type Breakfast --desc "went out for breakfast"`
> This creates a meal, with 1000 cals, 100g carbsm 120g protein and 32g of fat. Its a breakfast, and has a description

`life create workout --calories 933 --type Legs --load 8 --duration 43 --desc "Hit a hard leg day"`
> This creates a workout with 933 cals burned, 8/10 load, 43 mins long, type is legs, and has a description

`life get meals` or `life get workouts` or `life get macros` - gets the specified resource history from the past week
can specify timeframe `--timeframe today|week|month|year`. 
