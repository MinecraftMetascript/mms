/**
 * Data packs cannot add items to the creative menu, so this gives one of
 * everything: /function embers:give
 */
Function give {
  /give @s amethyst_shard[item_model="embers:ember_shard",item_name={translate:"item.embers.ember_shard"}] 16
  /give @s music_disc_13[item_model="embers:ember_disc",item_name={translate:"item.embers.ember_disc"},jukebox_playable="embers:embers"]
  /give @s painting[painting/variant="embers:hearth"]
}
