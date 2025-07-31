# TicTacToe CLI Game

A simple terminal-based implementation of the classic **TicTacToe** game written in **Go**, using the [`tcell`](https://github.com/gdamore/tcell) library for handling terminal graphics and input.

## TODOS
- [x] Implement a window Object (tcell screen), ready to be drawn on the screen.
    - Uses `ticker` Update() iterate over frames. Currently shows FPS top left the screen.
- [X] Single Box Object that can be clicked via the mouse button to draw on `X` or `O`.
- [X] Set 9 boxes in the middle of the screen.
    - [X] Ensure all boxes are centered.
    - [X] GridView created and can draw list of boxes.
- [ ] Player Scores when win pattern correct and reset.
    - [X] The game show who wins the game `X` or `O` at the end of every match.
- [ ] Clear or Restart game when pressing i.e `r` key.
- [ ] Show logs and current Turn.

## Getting Started

Clone repository

```bash
git clone https://github.com/omar0ali/tictactoe-game-cli.git
```

Run the game

```bash
go run .
```

