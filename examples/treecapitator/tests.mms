// Tests of the tree check. Run them with `mms test`, in the web editor's
// Tests tab, or in VS Code's Testing view. They are never written to the
// pack.

Test logs_with_leaves_on_four_sides_are_a_tree {
  setup {
    nether = false
    east = 1
    west = 1
    south = 1
    north = 1
    top = 0
  }
  expect is_tree()
}

Test a_cabin_with_three_leafy_sides_stands {
  setup {
    nether = false
    east = 1
    west = 0
    south = 1
    north = 0
    top = 1
  }
  expect !is_tree()
}

Test a_fungus_needs_one_cap_block {
  setup {
    nether = true
    east = 0
    west = 0
    south = 0
    north = 0
    top = 1
  }
  expect is_tree()
}

Test a_bare_stem_is_not_a_fungus {
  setup {
    nether = true
    east = 0
    west = 0
    south = 0
    north = 0
    top = 0
  }
  expect !is_tree()
}
