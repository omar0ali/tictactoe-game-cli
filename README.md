# TicTacToe CLI Game

A simple terminal-based implementation of the classic **TicTacToe** game written in **Go**, using the [`tcell`](https://github.com/gdamore/tcell) library for handling terminal graphics and input.

## TODOS
- [x] Implement a window Object (tcell screen), ready to be drawn on the screen.
    - Uses `ticker` Update() iterate over frames. Currently shows FPS top left the screen.
- [X] Single Box Object that can be clicked via the mouse button to draw on `X` or `O`
- [ ] Set 9 boxes in the middle of the screen
    - Ensure all boxes are centered

## Getting Started

Clone repository

```bash
git clone https://github.com/omar0ali/tictactoe-game-cli.git
```

Run the game

```bash
go run .
```

