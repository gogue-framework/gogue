package noises

import "github.com/gogue-framework/gogue/gamemap"

// DegradeNoises iterates over every tile on the map, and reduces the amount of noise on each tile by a set amount.
// If a noise reaches an intensity of 0, the noise is removed from the tile. This is intended to be run each frame.
// This simulates sounds being intially made, and maybe echoing around for a brief time, but then eventually disappearing
// An example would be if an entity is tracking another entity by the sound its generating, if the tracked entity stops
// generating sounds, eventually, the tracking entity will no longer be able to follow the trail of sound, as the sound
// no longer exists.
func DegradeNoises(mapSurface *gamemap.GameMap, degradationRate float64) {
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
