// Sound("sounds/embers.ogg") copies the file to assets/embers/sounds/ and
// adds an embers:embers entry to the generated sounds.json. The id it
// stands for is what sound_event needs.

/** The song the Embers disc plays in a jukebox. */
JukeboxSong embers = {
  sound_event = Sound("sounds/embers.ogg")
  description = { translate = "jukebox_song.embers.embers" }
  length_in_seconds = 8
  comparator_output = 7
}
