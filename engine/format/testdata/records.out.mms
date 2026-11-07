Type Pos {
  x: int
  y: int = 64
}

Type S {
  mode: "x" | "y"
  blah: {
    x: int
    y: int = 2
    /** The way it faces. */
    facing: "north" | "south"
  }
  inline: { a: int, b: int }
  list: [{ a: int }]
  home: Pos
}
