ItemTag planks_and_logs = [
  ...[for w in woods: "${w}_planks"],
  ...[for w in woods if w != mangrove: "${w}_log"],
]
