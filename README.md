# TicTacToe CLI Game

A simple terminal-based implementation of the classic **TicTacToe** game written in **Go**, using the [`tcell`](https://github.com/gdamore/tcell) library for handling terminal graphics and input.

![TicTacToe Game](https://github.com/omar0ali/tictactoe-game-cli/blob/main/screenshots/tictactoe-game-cli.png)

## TODOS
- [x] Implement a window Object (tcell screen), ready to be drawn on the screen.
    - Uses `ticker` Update() iterate over frames. Currently shows FPS top left the screen.
- [X] Single Box Object that can be clicked via the mouse button to draw on `X` or `O`.
- [X] Set 9 boxes in the middle of the screen.
    - [X] Ensure all boxes are centered.
    - [X] GridView created and can draw list of boxes.
- [ ] Player Scores when win pattern correct and reset.
    - [X] The game show who wins the game `X` or `O` at the end of every match.
    - [ ] Color the Winner.
- [ ] Show logs and current Turn. (Should be shown on the bottom-right of screen).
- [X] After each end of the game, should should a dialog, showing the winner 
- [ ] Ask player to restart the game or exit.
    - [X] Dialog Box added

## Getting Started

Clone repository

```bash
git clone https://github.com/omar0ali/tictactoe-game-cli.git
```

Run the game

```bash
go run .
```

