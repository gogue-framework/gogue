# Gogue - Roguelike toolkit for Go

![Go](https://github.com/gogue-framework/gogue/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogue-framework/gogue)](https://goreportcard.com/report/github.com/gogue-framework/gogue)

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
    - Menu system (primitive)
- Pathfinding (A* for now, but also Djikstra Maps)
- Djikstra Maps implementation (http://www.roguebasin.com/index.php?title=The_Incredible_Power_of_Dijkstra_Maps)
    - Single entity maps
    - multi-entity maps
    - combined maps
- Random number generation
    - Uniform, Normal Distribution, Ranges, Weighted choices
    - Dice rolls (normal and open ended)
- Random name generation (WIP)

... and whatever else I deem useful

## Getting Started

Standard Go package install - `go get github.com/gogue-framework/gogue`

Or if using Modules, simply include `github.com/gogue-framework/gogue` in your project imports and build.

### Prerequisites

BearLibTerminal is required to use this package. You can find install instructions for various operating systems here: https://github.com/gogue-framework/gogue/wiki
