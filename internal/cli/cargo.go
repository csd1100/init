package cli

type cargoExe struct {
	Command string
}

func (cargo cargoExe) GetCommand() string {
	return cargo.Command
}

func (cargo cargoExe) Exec(subcommand string, args []string) error {
	return execute(cargo, subcommand, args)
}

type cargoCLI struct {
	exe Executable
}

var Cargo = cargoCLI{
	exe: cargoExe{Command: "cargo"},
}

func (cargo cargoCLI) Init() error {
	return cargo.exe.Exec("init", []string{})
}

func (cargo cargoCLI) Check() error {
	return cargo.exe.Exec("check", []string{})
}

func (cargo cargoCLI) Fetch() error {
	return cargo.exe.Exec("fetch", []string{})
}

func (cargo cargoCLI) Clean() error {
	return cargo.exe.Exec("clean", []string{})
}

func (cargo cargoCLI) Sync(args map[string]string) error {
	err := cargo.Fetch()
	if err != nil {
		return err
	}

	err = cargo.Check()
	if err != nil {
		return err
	}

	return cargo.Clean()
}
