# rock-paper-scissors
A very very basic implementation of Rock Paper Scissors in Go

## Get Started
Install Go 1.18+

### Commands

Run `go build` to generate a binary in local repo. 

### Usage
Pass in two strings, Player 1's move and Player 2's move. The program will output which move wins.
```sh
local-machine % ./rock-paper-scissors Rock Scissors
Rock Wins!
```

If both strings are the same, it's a draw: 
```sh
local-machine % ./rock-paper-scissors paper paper
Its a Draw!
```

Thats it.