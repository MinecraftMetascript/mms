// 3. The terrain: islands where 3D noise is high, fading out above and
// below y = 128.
Noise island_noise = {
  base_octave = -6
  amplitude_modifiers = [1, 1, 0.5]
}

// 0 at y = 128, rising to 1 at y = 32 and y = 224.
DensityFunction height = Gradient(
  axis = y,
  from_coordinate = 32,
  to_coordinate = 224,
  from_value = -1,
  to_value = 1,
)

DensityFunction islands = Noise(island_noise, 1, 2) - height * height - 0.3

// 4. What the surface is made of.
SurfaceRule surface = [
  If(StoneDepth(Floor)) { Block(grass_block) },
  If(StoneDepth(Floor).Offset(3)) { Block(dirt) },
]

NoiseSettings islands = {
  sea_level = 0
  default_block = stone
  default_fluid = air
  disable_mob_generation = false
  legacy_random_source = false
  noise = {
    min_y = 0
    height = 256
  }
  material_rule = surface
  spawn_target = []
  noise_router = {
    final_density = islands
    temperature = 0
    vegetation = 0
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
  }
}
