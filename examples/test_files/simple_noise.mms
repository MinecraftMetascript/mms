Namespace test_files {
    Noise {
        NoAmp = Noise(-1)
        SingleArg = Noise(-1).Amplitudes(1)
        ManyArgs = Noise(-1).Amplitudes(1,2,5,-7,0.5)
    }
}