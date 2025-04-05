package ecs

import "time"

type Thread struct {
	// ID of the Thread.
	ID int64 `json:"id,omitempty"`

	// Name of the Thread.
	Name string `json:"name,omitempty"`
}

// Process defines the properties for capturing information about a process and help correlate metrics information with
// a process id/name from a log message.
//
// The Process.PID often stays in the metric itself and is copied to the global field for correlation.
type Process struct {
	// Args represents the arguments used when the process was started, beginning with the absolute path to the
	// executable.
	//
	// May be filtered to protect sensitive information.
	Args []string `json:"args,omitempty"`

	// ArgsCount represents the length of the list Process.Args.
	//
	// This field can be useful for querying or performing bucket analysis on how many arguments were provided to start
	// a process. More arguments may be an indication of suspicious activity.
	ArgsCount int64 `json:"args_count,omitempty"`

	// CommandLine that started the Process, including the absolute path to the executable, and all arguments.
	//
	// Some arguments may be filtered to protect sensitive information.
	CommandLine string `json:"command_line"`

	// End the time the Process ended.
	End time.Time `json:"end,omitempty"`

	// EntityID is a unique identifier for the Process.
	//
	// The implementation of this is specified by the data source, but some examples of what could be used here are a
	// process-generated UUID, Sysmon Process GUIDs, or a hash of some uniquely identifying components of a process.
	//
	// Constructing a globally unique identifier is a common practice to mitigate PID reuse as well as to identify a
	// specific process over time, across multiple monitored hosts.
	EntityID string `json:"entity_id,omitempty"`

	// Executable is the absolute path to the Process executable.
	Executable string `json:"executable,omitempty"`

	// ExitCode is the exit code of the Process for termination events.
	//
	// The field should be absent if there is no exit code for the event (e.g. Process.Start).
	ExitCode int64 `json:"exit_code,omitempty"`

	// Parent process.
	Parent *Process `json:"parent,omitempty"`

	// PGID is the identifier of the group of processes the process belongs to.
	PGID int64 `json:"pgid,omitempty"`

	// Process id.
	PID int64 `json:"pid,omitempty"`

	// Process name.
	// Sometimes called program name or similar.
	Name string `json:"name,omitempty"`

	// Thread represents the thread running the Process.
	Thread *Thread `json:"thread,omitempty"`

	// Title is the title of the Process.
	Title string `json:"title,omitempty"`

	// Start is the time the Process started.
	Start time.Time `json:"start,omitempty"`

	// Uptime the time in seconds the Process has been up.
	Uptime int64 `json:"uptime,omitempty"`

	// WorkingDirectory is the working directory of the Process.
	WorkingDirectory string `json:"working_directory,omitempty"`
}
