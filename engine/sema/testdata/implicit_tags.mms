// Fields marked tags="implicit" hold a tag written without '#'. Both
// spellings compile to the bare id.
PlacedFeature pond = {
  feature = "minecraft:lake_lava"
  placement = [
    BlockPredicateFilter(predicate = MatchingBlockTag(tag = #air))
    BlockPredicateFilter(predicate = MatchingBlockTag(tag = air))
    BlockPredicateFilter(predicate = MatchingBlockTag(tag = "#minecraft:features_cannot_replace"))
  ]
}
