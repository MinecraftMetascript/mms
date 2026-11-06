/** One structure file as a weighted pool entry. */
Template Piece(
  path: string,
  weight: int = 1,
  processors: ProcessorList = minecraft:empty,
): WeightedPoolElement = WeightedElement(
  weight,
  SinglePoolElement(
    rigid,
    StructureFile("structures/${path}.nbt", path),
    processors,
  ),
)

TemplatePool halls {
  fallback = caps
  elements = [Piece("halls/hall", 3), Piece("halls/short_hall")]
}
