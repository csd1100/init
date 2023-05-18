package cli

type CLI interface {
	Exec(command string, args []string)
	IsInPath(command string)
	GetVersion(command string)
}
