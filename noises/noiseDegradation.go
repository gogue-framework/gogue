package noises

import "github.com/jcerise/gogue/gamemap"

// DegradeNoises iterates over every tile on the map, and reduces the amount of noise on each tile by a set amount.
// If a noise reaches an intensity of 0, the noise is removed from the tile. This is intended to be run each frame.
func DegradeNoises(mapSurface *gamemap.Map, degradationRate float64) {
	for x := 0; x < mapSurface.Width; x++ {
		for y := 0; y < mapSurface.Height; y++ {
			if mapSurface.HasNoises(x, y) {
				updatedNoises := make(map[int]float64)
				for entity, noise := range mapSurface.Tiles[x][y].Noises {
					updatedNoise := noise - degradationRate

					if updatedNoise > 0 {
						updatedNoises[entity] = updatedNoise
					}
				}

				mapSurface.Tiles[x][y].Noises = updatedNoises
			}
		}
	}
}
