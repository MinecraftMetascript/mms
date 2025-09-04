# File Structure:

Every file _must_ begin with a namespace declaration:
```
namespace my_namespace;
```

Files are then made up of declarations, which follow this format:

```
<Identifier> := <Resource Type> { <Resource Definition> }
```

To declare a simple noise:

```
MyNoise := Noise { -5 [ 1 3 ] }
```

# Language Reference:

## Syntax

Whitespace (spaces and tabs) is ignored.

New lines can appear in most places and are not significant.

Comments can be created using `//` for single-line comments, and `/* */` for multi-line comments.

## Basic Types

### Int vs. Number

`Int` is a _whole_ number, and cannot contain decimal points.

`Number` is either an `Int` or a `Float`, and _can_ contain decimal points.

### Strings

Strings are any text that is surrounded by double quotes. 
Many cases don't actually require strings unless explicitly stated (e.g., resource references do not need quotes) 

### Identifiers
Identifiers are case-sensitive and can contain letters, numbers, and underscores.
They must start with a letter.

Identifiers must be unique within a namespace but can be used in multiple namespaces.

### Resource References
Resource references can be _qualified_ or _unqualified_.

| Qualified | Unqualified |
| --------- | ----------- |
| `minecraft:stone` | `stone` |


## Surface Rules

### Bandlands:

No special arguments are required:
```
SurfaceRule { 
    Bandlands 
}
```

### Block:

Fundamental surface rule
```
SurfaceRule { 
    Block <ResourceReference> 
}
```

### Reference:
Use a [resource reference](#resource-references) in place of an explicit rule:

```
SurfaceRule { 
    <ResourceReference> 
}
```

### Conditional (If)

Construct made up of a [condition](#surface-conditions) and a [rule](#surface-rules):

```
SurfaceRule {
    If ( <SurfaceCondition> ) {
        <SurfaceRule>
    }
}

```

### Sequence

Construct made up of a list of [rules](#surface-rules):
```
SurfaceRule {
    Sequence [
        <SurfaceRule>
        <SurfaceRule...>
    ]
}
```

## Surface Conditions

### Simple Conditions:
These conditions require no additional arguments:

```
SurfaceCondition { AboveSurface }
SurfaceCondition { Freezing }
SurfaceCondition { Hole }
SurfaceCondition { Steep }
```


### AboveWater

Creates a `minecraft:water` surface condition:

```
SurfaceCondition {
    AboveWater <Int (Offset)> 
               <Number (Multiplier)> 
               <"Add" | "Sub" ("add_stone_depth")>
}
```

### Biome

Filters to within a set of biomes

```
SurfaceCondition {
    Biome [ 
        <ResourceReference> 
        <ResourceReference...> 
    ]
}
```

### Compound Conditions

#### Not

Inverts the result of a condition.

```
SurfaceCondition {
    Not ( <SurfaceCondition> )
}
```

#### And
Allows for multiple conditions to be combined with `and`. All conditions must be true
for the rule to be applied.

```
SurfaceCondition {
    And (
        <SurfaceCondition>
        <SurfaceCondition...>
    )
}
```

#### Or

Allows for multiple conditions to be combined with `or`. At least one condition must be true
for the rule to be applied.

```
SurfaceCondition {
    Or (
        <SurfaceCondition>
        <SurfaceCondition...>
    )
}
```

`Or` lets you combine multiple conditions that all lead to the same result in an easier to read fashion.
For example, these rules are equivalent:
```
MyRule := SurfaceRule {
    If (
        Or (
            AboveSurface
            Freezing
        )
    )
    Block stone
}

// -- is the same as --

MyRule := SurfaceRule {
    If (
        AboveSurface
    ) Block stone
    If (
        Freezing
    ) Block stone
}
```



### Noise

Uses a [noise](#noise) to determine whether the condition passes or fails.

```
SurfaceCondition {
    Noise 
        <ResourceReference> 
        [ <Number (min)>, <Number (max)> ]
}
```

### Reference

Use a [resource reference](#resource-references) in place of an explicit condition:

```
SurfaceCondition { 
    <ResourceReference> 
}
```

### StoneDepth

Uses if the position is within a specified distance from the surface.

```
SurfaceCondition {
    StoneDepth 
        <"Floor" | "Ceiling" (surface_type)> 
        <Int (vertical offset)> 
        <"Add" | "Sub" (add_surface_depth)> 
        <Int (secondary_depth_range)>
}

```

### VerticalGradient

Compares the current Y position, with a messy transition, just like the deepslate and bedrock transitions.

```
SurfaceCondition {
    VerticalGradient <String (seed)> <VerticalAnchor> <VerticalAnchor>
}

```


### YAbove

Checks if the current Y position is above a specified Y value.

```
SurfaceCondition {
    YAbove <VerticalAnchor> <Int> <"Add" | "Sub">
}
```

## Vertical Anchors

Vertical Anchors appear in a few places where a Y value is required. 
They can be provided in three forms:

```
<Int (Absolute Y Value)>
~<Int (Relative AboveBottom)>
~-<Int (Relative BelowTop)>
```


# Known Issues:

## Unqualified References are Inconsistent

Unqualified references currently have inconsistent behavior, and
sometimes reference the file's namespace, while at other times they reference the
minecraft namespace. It is recommended to always use qualified references.
