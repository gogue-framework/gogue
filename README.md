# Gogue - Roguelike toolkit for Go

Gogue aims to create a simple to use toolkit for creating Roguelike games in the Go language. It uses BearLibTerminal for rendering, so that will be required to use the toolkit.
This is by no means a complete toolkit, but its got (and will have) a bunch of things I've found handy for building roguelikes. Hopefully someone else finds them handy as well.

Development is on-going, so use at your own risk.

## Features

This feature list is incomplete, as various pieces are still in development.

- Terminal creation and management (using BearlibTerminal)
    - Colors
    - Easy font management
    - Glyph rendering
    - Input handling
- Dynamic input registration system
- Lightweight Entity/Component/System implementation
- JSON data loading
- Dynamic entity generation from JSON data
- Map generation
- Scrolling camera
- Field of View (only raycasting at the moment, but more to come)
- UI
    - Logging
    - Screen Management
- Pathfinding (A* for now, but also Djikstra Maps)
- Djikstra Maps implementation (http://www.roguebasin.com/index.php?title=The_Incredible_Power_of_Dijkstra_Maps)
    - Single entity maps
    - multi-entity maps
    - combined map decisions
- Random number generation
    - Uniform, Normal Distribution, Ranges, Weighted choices
    - Dice rolls (normal and open ended)
- Random name generation (WIP)

... and whatever else I deem useful

## Getting Started

Standard Go package install - `go get github.com/jcerise/gogue`

### Prerequisites

BearLibTerminal is required to use this package. I'll put specific platform install instructions here, as it varies from platform to platform, and is not well documented (I feel).
