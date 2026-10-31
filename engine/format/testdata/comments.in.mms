// header comment

/** doc */
Recipe r { type = "minecraft:crafting_shaped" // the type
  pattern = [
    "RRR", // top row
    "R R"
    // bottom row is empty
  ]
  key = { R = ruby /* inline */ }
  // dangling at end of block
}

DensityFunction d = Add(
  // first arg
  1,
  2 // second
)

Function g {
  if (a) { Say(a) } // after if
  else {
    /say b
  }
  /* before let */ let y = 2
}
// eof comment
