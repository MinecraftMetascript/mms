Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })
Global n: int = 0

Function is_sneaking(): bool {
  return sneaking
}

Function f {
  if (sneaking && n > 2) {
    /say hi
  }
  if (!sneaking) {
    /say no
  }
  let s: bool = siege:sneaking
  while (sneaking || s) {
    /say w
  }
}
