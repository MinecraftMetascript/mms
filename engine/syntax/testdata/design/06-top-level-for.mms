// woods is the const declared above: [oak, spruce, birch, ...]
for w in woods {
  /** A crate made of ${w}. */
  Recipe ${w}_crate = Shaped(["PPP", "P P", "PPP"], { P = "${w}_planks" }, barrel)

  LootTable blocks/${w}_crate {
    pools = [Pool(1, "${w}_crate")]
  }
}
