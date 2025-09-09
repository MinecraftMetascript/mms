namespace xyz {
  DensityFn {
    simpleSpline = Spline(5)
        .Point(1,Spline(10),1)

    constant = YClampedGradient().Min(5).Max(2)
  }
}
