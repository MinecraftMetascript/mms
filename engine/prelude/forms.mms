// The forms of expression the language builds in, and what each one
// stands for where a type is expected. The parser gives each form its
// parts, named by the fields below; an `as T = value` is used where a T
// is expected and the value written does not already fit, with `self`
// holding the parts. A part that is left out, such as the low end of
// `..3`, drops the field it would fill.

/** `lo..hi`, a range. Either end can be left out. */
extern Type Range {
  /** The low end. */
  lo: any?
  /** The high end. */
  hi: any?
  /** `1..3` where a number provider goes: a uniform number from 1 to 3. */
  as IntNumberProviderRef = { type = uniform, min = self.lo, max = self.hi }
  /** `0.5..2` where a number provider goes: a uniform number from 0.5 to 2. */
  as FloatNumberProviderRef = { type = uniform, min = self.lo, max = self.hi }
  /** `1..3`, `1..` or `..3` where bounds go. */
  as IntRange = { min = self.lo, max = self.hi }
  /** `0.5..2`, `0.5..` or `..2` where bounds go. */
  as FloatRange = { min = self.lo, max = self.hi }
  /** `1..3`, `1..` or `..3` where bounds go. */
  as MinMaxBounds = { min = self.lo, max = self.hi }
}

/** `!x`. */
extern Type Not {
  operand: any
  /** `!a` on a loot condition: true when a is false. */
  as LootCondition = { type = inverted, term = self.operand }
  /** `!a` on a surface condition. */
  as MaterialCondition = { type = not, invert = self.operand }
}

/** `a && b && ...`. A chain of `&&` is one form, with every term; parentheses keep a group. */
extern Type And {
  terms: [any]
  /** `a && b` on loot conditions: true when every term is. */
  as LootCondition = { type = all_of, terms = self.terms }
}

/** `a || b || ...`. A chain of `||` is one form, with every term; parentheses keep a group. */
extern Type Or {
  terms: [any]
  /** `a || b` on loot conditions: true when any term is. */
  as LootCondition = { type = any_of, terms = self.terms }
}

/** `a + b`. */
extern Type Sum {
  left: any
  right: any
  /** `a + b` on density functions. */
  as DensityFunction = { type = add, left = self.left, right = self.right }
}

/** `a - b`. */
extern Type Difference {
  left: any
  right: any
  /** `a - b` on density functions. */
  as DensityFunction = { type = sub, left = self.left, right = self.right }
}

/** `a * b`. */
extern Type Product {
  left: any
  right: any
  /** `a * b` on density functions. */
  as DensityFunction = { type = mul, left = self.left, right = self.right }
  /** `diamond * 4`: a stack of four diamonds. */
  as ItemStackTemplate = { id = self.left, count = self.right }
}

/** `a / b`. */
extern Type Quotient {
  left: any
  right: any
  /** `a / b` on density functions. */
  as DensityFunction = { type = div, left = self.left, right = self.right }
}

/** `-x`, where x is not a number. */
extern Type Negation {
  operand: any
  /** `-a` on a density function. */
  as DensityFunction = { type = mul, left = -1, right = self.operand }
}

/** A bare id, such as `diamond` or `minecraft:diamond`. */
extern Type BareId {
  id: any
  /** `diamond` in a pool's entries: an item entry. */
  as LootPoolEntry = { type = item, name = self.id }
}

/** A tag, such as `#logs`. */
extern Type TagRef {
  tag: any
  /** `#logs` in a pool's entries: one entry for each item of the tag, so each is equally likely. */
  as LootPoolEntry = { type = tag, items = self.tag, expand = true }
}

/** A number. */
extern Type NumberLiteral {
  value: any
  /** `64` where a height goes: the absolute height 64. */
  as VerticalAnchor = { absolute = self.value }
}

/** A list. Lists inside it are flattened into it. */
extern Type ListLiteral {
  items: [any]
  /** `[a, b]` where a surface rule goes: the first rule that applies. */
  as MaterialRule = { type = sequence, sequence = self.items }
}
