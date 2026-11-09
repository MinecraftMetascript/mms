package project

// WorldInstall is what building a pack into a world writes and what it
// asks of the player, worked out from the world's files before anything
// is written. `mms build --world` and the web editor (through the
// language server's mms/worldInstall) both install from it, so they write
// the same files and say the same next step.
type WorldInstall struct {
	// Data is the data pack, for <world>/datapacks/<namespace>. Its data/,
	// assets/ and pack.png are replaced; see WriteDir.
	Data []OutFile
	// Resources is how many resource pack files the build has.
	Resources int
	// Pack is the merged <world>/resourcepacks/resources.zip, or nil when
	// the build leaves that file alone: no resources, and no zip there.
	Pack *WorldPackResult
	// ResourcesChanged says the zip's bytes differ from the old ones.
	ResourcesChanged bool
	// Reset are the dimensions whose saved chunks to take out.
	Reset []string
	// Plan is what the player does next.
	Plan ApplyPlan
}

// PlanWorldInstall works out a world install of files (Emit's, with the
// assets read) for the project namespace ns. before is the data pack
// already in the world (ReadPackDir), and oldZip the world's
// resources.zip, nil when there is none. With reset, Reset holds
// resetIDs, or when that is nil the project's dimensions whose worldgen
// changed since before.
func PlanWorldInstall(files []OutFile, ns string, before map[string][]byte, oldZip []byte, reset bool, resetIDs []string) (*WorldInstall, error) {
	data, res := SplitPacks(files)
	in := &WorldInstall{Data: data, Resources: len(res)}
	if reset {
		if in.Reset = resetIDs; in.Reset == nil {
			in.Reset = ChangedDimensions(before, data)
		}
		for _, id := range in.Reset {
			if err := CheckDimensionID(id); err != nil {
				return nil, err
			}
		}
	}
	// With no resources, a pack already there still loses the files
	// this project put in it last time.
	if len(res) > 0 || oldZip != nil {
		m, err := MergeWorldPack(oldZip, ns, res)
		if err != nil {
			return nil, err
		}
		in.Pack = m
		in.ResourcesChanged = string(oldZip) != string(m.Zip)
	}
	in.Plan = PlanApply(before, data, in.ResourcesChanged)
	return in, nil
}
