package subprocess

type watchLog struct {
	cpu    bool
	memory bool
}

// Details is simply the struct method-wrapper for the "exec" package
type Details struct {
	env      []string
	command  []string
	watchLog watchLog
}

// New creates a new exec.Details struct
func New() *Details {
	new := &Details{}
	new.watchLog = watchLog{cpu: false, memory: false}

	return new
}

// WithEnviron creates a new exec.Details struct with the os.Environ object copied-in
func (s *Details) WithEnviron(env []string) *Details {
	clone := *s // This does a shallow clone

	clone.env = env

	return &clone
}

// WithCommand creates a new exec.Details struct with the command []string copied-in
func (s *Details) WithCommand(command []string) *Details {
	clone := *s // This does a shallow clone

	clone.command = command

	return &clone
}

// WithWatchLog creates a new exec.Details struct with CPU or Memory watch loggers enabled
func (s *Details) WithWatchLog(cpu, memory bool) *Details {
	clone := *s // This does a shallow clone

	clone.watchLog.cpu = cpu
	clone.watchLog.memory = memory

	return &clone
}
