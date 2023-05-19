package cli

type mockCLI struct {
	actualArgs  []string
	actualError error
}

func (mcli mockCLI) Exec(args []string) error {
	mcli.actualArgs = args
	mcli.actualError = nil
	return nil
}
