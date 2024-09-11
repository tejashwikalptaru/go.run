# Go.Run

**Current Version: 0.0.1 (Alpha)**

Go.Run is a 2D side-scrolling game developed using Go and Ebitengine. In this game, the player must navigate through various obstacles, complete jumps, and progress through increasingly difficult levels. With immersive background music and dynamic obstacles, Go.Run offers an exciting gameplay experience for players of all skill levels.

## 🚀 Features

- **Dynamic Obstacles**: Face off against a variety of animated enemies like snakes, hyenas, vultures, and more.
- **Jump Mechanics**: Time your jumps to avoid obstacles and clear levels.
- **Level Progression**: Complete jumps to advance to higher levels, with increasing difficulty.
- **Seamless Background Transitions**: Enjoy scrolling backgrounds that change as you progress.
- **Sound Effects**: Background music and sound effects for jumps and collisions add to the immersion.
- **Cross-Platform**: Build and play the game on macOS, Linux, and Windows.

## 🎮 Gameplay

The objective of Go.Run is simple: avoid obstacles, jump over them, and complete each level by clearing a certain number of jumps. As you advance through the levels, the obstacles get faster, and the difficulty increases.

## 🛠 Installation

### Prerequisites

- **Go (Golang)**: Version 1.23+ is recommended.
- **Ebitengine**: The game is built using [Ebitengine](https://ebiten.org/), a lightweight game engine for Go.

### Cloning the Repository

Clone the repository using the following command:

```bash
git clone https://github.com/tejashwikalptaru/go.run.git
cd go.run
```

### Building the Game

You can build the game for macOS, Linux, and Windows using the included Makefile.

For macOS:

```bash
make build-darwin
```

For Linux:

```bash
make build-linux
```

For Windows:

```bash
make build-windows
```

To run the game locally:

```bash
make run
```

To clean up build artifacts:

```bash
make clean
```

## 🕹 Controls

- **Spacebar**: Jump
- **Arrow Keys**: Navigate through menus (future implementation)
- **Escape**: Pause/Exit (future implementation)

## 📦 Dependencies

The game uses the following dependencies:

- [Ebitengine](https://ebiten.org/): Game engine
- [Golangci-lint](https://github.com/golangci/golangci-lint): Linter for Go
- [Go mod](https://golang.org/ref/mod): For dependency management

To install necessary dependencies, run:

```bash
make dep
```

## 📝 Contributing

Contributions are welcome! Please feel free to submit a pull request, report issues, or suggest new features.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

## 🎉 Acknowledgements

- Thanks to the [Ebitengine](https://ebiten.org/) community for their great work on the engine.
- 1001fonts.com for JungleAdventurer font
- craftpix.net for background images
- pixabay.com authors for game sounds
---