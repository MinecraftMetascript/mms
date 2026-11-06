// Every form the overhaul adds, written untidily.
extern Template font_width(text: string):int
JigsawHole a,b ,  c
Type LangText=string|{string: string}
Type Mode = "rigid"|"terrain_matching"
/** A wave. */
Type Wave{mob: Id<entity_type>,count: int=1
  /** Spawn delay. */
  delay: int? }
const   woods=[oak,spruce]
const n: int=3
let shared: [string]=[]
for w in woods{Recipe ${w}_crate=Shaped(["PPP"],{P="${w}_planks"},barrel)
  ItemTag ${w}_things=[...[for x in woods if x!=w: "${x}_log"], #${w}_logs]}
for k,v in pairs { Lang ${k} { "a.${v}"=v } }
Template Pick(a: int=1,b: Model? =none): Model=a??b
Blueprint Thing{name: string
  as Id<item> = stick
  emit{primary CustomItem ${self.path}{base=stick,name=if(name!=none) name else "x"}
  OnTick{
  /say hi
  }}}
Test t{setup{game.wave=1}
run next_wave()
expect game.wave==2}
Function f{if(x)error(x,"y") else warn(x,"z")
  let r=run{
  /say a
}
  let s=(if (a) b else c)
}
