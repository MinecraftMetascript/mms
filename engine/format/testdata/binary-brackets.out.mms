// A left operand that breaks over lines keeps the operator after its
// closing bracket, and a lone object argument hugs the parentheses.
DensityFunction x = (Clamp(
  Abs(
    x_density_function *
      Spline({
        points = [SplinePoint(1, 1, inner_1)]
        coordinate = x_density_function
      }),
  ),
  0,
  1,
) - 1) * -1

DensityFunction long_right = Clamp(
  Abs(
    Spline({
      coordinate = x
      points = [SplinePoint(1, 1, 2)]
    }),
  ),
  0,
  1,
) * some_very_long_density_function_name_that_does_not_fit_on_the_line

DensityFunction commented = Spline({
  // inside the object
  coordinate = x
  points = [SplinePoint(1, 1, 2)]
})

DensityFunction trailing = Spline(
  {
    coordinate = x
    points = [SplinePoint(1, 1, 2)]
  }, // after the object
)

DensityFunction flat = (Abs(a) - 1) * -1

DensityFunction chain = Min(
  Interpolated(
    Min(cheese, starter_cavern) + Noise(wall_roughness, 1, 1) * 0.06,
    4,
    8,
  ),
  tunnels,
) + roof + floor
