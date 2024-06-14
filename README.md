# quake-log-reader
![quake-log-reader](https://img.shields.io/badge/quake--log--reader-gray?logo=go)
![technology Go 1.22](https://img.shields.io/badge/technology-go%201.22-blue.svg)
[![Go Coverage](https://github.com/FabsHC/quake-log-reader/wiki/coverage.svg)](https://raw.githack.com/wiki/FabsHC/quake-log-reader/coverage.html)

## Overview
This project reads a log file that was generated by a Quake 3 Arena server and extract some infos like:

- Total Kills per match
- Players
- Kills by player
- Kills by means

Example: 
```json
{
    "total_kills": 131,
    "players": [
      "Dono da Bola",
      "Zeh",
      "Mal",
      "Isgalamido",
      "Assasinu Credi",
      "Oootsimo"
    ],
    "kills": {
      "Assasinu Credi": 19,
      "Dono da Bola": 14,
      "Isgalamido": 17,
      "Mal": 6,
      "Oootsimo": 22,
      "Zeh": 19
    },
    "kills_by_means": {
      "MOD_FALLING": 3,
      "MOD_MACHINEGUN": 4,
      "MOD_RAILGUN": 9,
      "MOD_ROCKET": 37,
      "MOD_ROCKET_SPLASH": 60,
      "MOD_SHOTGUN": 4,
      "MOD_TRIGGER_HURT": 14
    }
  }
```

When a player kills another player, gains 1 kill point. Example log:
```
 22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH
 22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH
```

But when <world> kill a player, that player loses -1 kill score. Example log:
```
 21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
```

When a player kills itself, the killing score is not affected.
```
 22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH
```

## Project Organization
```
├── bin..................: Contain the shell script to generate mockgen files
├── cmd..................: Contains the main file to run the application.
├── cross-cut............: contain files that are used in cmd and internal package as well
├── internal.............: all core implementation its here.
│   ├── model............: All application structures (DTO).
│   ├── usecase..........: Application core validations.
│   └── util.............: General stuff.
├── mock.................: contains mock files generated by mockgen for unit tests

```

## How to run
The application expects data in STDIN format and will return json in STDOUT format. 
So you need to provide a log file that was generated by a Quake 3 Arena server.
Examples below:

### Using Go Run
You can run the commands below to run the application.
```shell
go run cmd/main.go < resources/qgames.log 
```

### Using Docker
Just run the docker commands below.
``` shell
docker run -i --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22.0 go run cmd/main.go < resources/qgames.log
```

## How to run the tests
Run the command below in the terminal to run the application tests.
### Using docker:
```shell
make -f Makefile test-docker
```
### Using go:
```shell
make -f Makefile test-go
```

