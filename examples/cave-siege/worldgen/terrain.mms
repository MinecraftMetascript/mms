// The shape of the cave world: solid rock from y -64 to y 191, carved into
// chambers and the tunnels between them. Where the final density is above 0
// the generator places stone, and where it is below 0 it leaves air.
//
// Three kinds of open space are combined with Min, so a block is open when
// any of them opens it:
// - cheese caves: big blobs where a 3D noise is high. In "chamber" regions,
//   a few hundred blocks apart, the threshold drops and the blobs grow into
//   halls that make good base sites.
// - tunnels: level passages in five layers, 36 blocks apart, winding
//   between the chambers.
// - the starter cavern: an open hall around x 0, z 0, y 40, where new
//   players arrive.
//
// Gradients near the top and bottom keep the roof and floor solid, so there
// is never sky and the bedrock layers are never breached.

/** Cheese caves. About 130 blocks across, flattened in y. */
Noise cave_cheese = {
  base_octave = -7
  amplitude_modifiers = [1, 0.6, 0.3, 0.15]
}

/** Where the big chambers are: a flat 2D field, about 500 blocks across. */
Noise chamber_field = {
  base_octave = -9
  amplitude_modifiers = [1, 0.5]
}

/** Where tunnels run: along the lines where this noise is 0. */
Noise tunnel_path = {
  base_octave = -7
  amplitude_modifiers = [1, 0.5]
}

/** How far a tunnel's floor rises and falls along its length. */
Noise tunnel_height = {
  base_octave = -7
  amplitude_modifiers = [1]
}

/** Small bumps on every wall, so caves don't look smooth. */
Noise wall_roughness = {
  base_octave = -4
  amplitude_modifiers = [1, 0.5]
}

/** 0 in most of the world, rising to 1 in the middle of a chamber region. */
DensityFunction chamber_bias = Clamp(Noise(chamber_field, 1, 0) * 2 - 0.4, 0, 1)

/**
 * Below 0 inside a cheese cave. The noise is squashed by 2 in y, so the
 * blobs are about twice as wide as they are tall: 8 to 40 blocks.
 */
DensityFunction cheese = 0.32 - chamber_bias * 0.28 - Noise(cave_cheese, 1, 2)

/** The height, as a density function: y itself. */
DensityFunction height = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = -64,
  to_value = 192,
)

// Tunnels come in five layers, 36 blocks apart. In each layer a tunnel
// runs along the lines where tunnel_path is 0. It is about 6 blocks wide
// and 5 tall, and its centre wanders up to 12 blocks above or below the
// layer's height, rising at most about half a block per block, so
// enemies can walk it. The noises are flat (2D); each layer shifts them
// by a different amount, so it gets its own map of tunnels.

DensityFunction tunnel_centre_1 = -24 + Noise(tunnel_height, 1, 0, shift_x = 4096) * 12

DensityFunction tunnels_1 = Max(
  Abs(Noise(tunnel_path, 1, 0, shift_x = 4096)) * 16,
  Abs(height - tunnel_centre_1) / 2.5,
) -
  1

DensityFunction tunnel_centre_2 = 12 + Noise(tunnel_height, 1, 0, shift_x = 8192) * 12

DensityFunction tunnels_2 = Max(
  Abs(Noise(tunnel_path, 1, 0, shift_x = 8192)) * 16,
  Abs(height - tunnel_centre_2) / 2.5,
) -
  1

DensityFunction tunnel_centre_3 = 48 + Noise(tunnel_height, 1, 0, shift_x = 12288) * 12

DensityFunction tunnels_3 = Max(
  Abs(Noise(tunnel_path, 1, 0, shift_x = 12288)) * 16,
  Abs(height - tunnel_centre_3) / 2.5,
) -
  1

DensityFunction tunnel_centre_4 = 84 + Noise(tunnel_height, 1, 0, shift_x = 16384) * 12

DensityFunction tunnels_4 = Max(
  Abs(Noise(tunnel_path, 1, 0, shift_x = 16384)) * 16,
  Abs(height - tunnel_centre_4) / 2.5,
) -
  1

DensityFunction tunnel_centre_5 = 120 + Noise(tunnel_height, 1, 0, shift_x = 20480) * 12

DensityFunction tunnels_5 = Max(
  Abs(Noise(tunnel_path, 1, 0, shift_x = 20480)) * 16,
  Abs(height - tunnel_centre_5) / 2.5,
) -
  1

/** Below 0 inside any tunnel. */
DensityFunction tunnels = Min(Min(Min(tunnels_1, tunnels_2), Min(tunnels_3, tunnels_4)), tunnels_5)

/** -1 to 1 across the starter cavern, along each axis. */
DensityFunction starter_x = Gradient(
  axis = x,
  from_coordinate = -36,
  to_coordinate = 36,
  from_value = -1,
  to_value = 1,
)

DensityFunction starter_z = Gradient(
  axis = z,
  from_coordinate = -36,
  to_coordinate = 36,
  from_value = -1,
  to_value = 1,
)

DensityFunction starter_y = Gradient(
  axis = y,
  from_coordinate = 26,
  to_coordinate = 54,
  from_value = -1,
  to_value = 1,
)

/**
 * Below 0 inside an ellipsoid about 64 wide and 26 tall around (0, 40, 0).
 * The gradients stop at 1 past their ends, so the 0.8 keeps everything
 * outside solid.
 */
DensityFunction starter_cavern = Square(starter_x) + Square(starter_z) + Square(starter_y) - 0.8

/** Solid from y 172 up, so the top 20 blocks are always roof. */
DensityFunction roof = Gradient(
  axis = y,
  from_coordinate = 150,
  to_coordinate = 172,
  from_value = 0,
  to_value = 2,
)

/** Solid below y -56, so the bedrock floor is never open. */
DensityFunction floor = Gradient(
  axis = y,
  from_coordinate = -60,
  to_coordinate = -44,
  from_value = 2,
  to_value = 0,
)

/**
 * The terrain. Chambers are interpolated over 4 by 8 block cells, as
 * vanilla does, to keep them fast and smooth. Tunnels are not: they are
 * only a few blocks tall, and interpolation would smear them away.
 */
DensityFunction caves = Min(
  Interpolated(Min(cheese, starter_cavern) + Noise(wall_roughness, 1, 1) * 0.06, 4, 8),
  tunnels,
) +
  roof +
  floor
