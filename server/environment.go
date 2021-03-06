package server

import (
	"os"
)

// Defines the basic interface that all environments need to implement so that
// a server can be properly controlled.
type Environment interface {
	// Returns the name of the environment.
	Type() string

	// Determines if the environment is currently active and running a server process
	// for this specific server instance.
	IsRunning() (bool, error)

	// Performs an update of server resource limits without actually stopping the server
	// process. This only executes if the environment supports it, otherwise it is
	// a no-op.
	InSituUpdate() error

	// Runs before the environment is started. If an error is returned starting will
	// not occur, otherwise proceeds as normal.
	OnBeforeStart() error

	// Starts a server instance. If the server instance is not in a state where it
	// can be started an error should be returned.
	Start() error

	// Stops a server instance. If the server is already stopped an error should
	// not be returned.
	Stop() error

	// Waits for a server instance to stop gracefully. If the server is still detected
	// as running after seconds, an error will be returned, or the server will be terminated
	// depending on the value of the second argument.
	WaitForStop(seconds int, terminate bool) error

	// Determines if the server instance exists. For example, in a docker environment
	// this should confirm that the container is created and in a bootable state. In
	// a basic CLI environment this can probably just return true right away.
	Exists() (bool, error)

	// Terminates a running server instance using the provided signal. If the server
	// is not running no error should be returned.
	Terminate(signal os.Signal) error

	// Destroys the environment removing any containers that were created (in Docker
	// environments at least).
	Destroy() error

	// Returns the exit state of the process. The first result is the exit code, the second
	// determines if the process was killed by the system OOM killer.
	ExitState() (uint32, bool, error)

	// Creates the necessary environment for running the server process. For example,
	// in the Docker environment create will create a new container instance for the
	// server.
	Create() error

	// Attaches to the server console environment and allows piping the output to a
	// websocket or other internal tool to monitor output. Also allows you to later
	// send data into the environment's stdin.
	Attach() error

	// Follows the output from the server console and will begin piping the output to
	// the server's emitter.
	FollowConsoleOutput() error

	// Sends the provided command to the running server instance.
	SendCommand(string) error

	// Reads the log file for the process from the end backwards until the provided
	// number of bytes is met.
	Readlog(int64) ([]string, error)

	// Polls the given environment for resource usage of the server when the process
	// is running.
	EnableResourcePolling() error

	// Disables the polling operation for resource usage and sets the required values
	// to 0 in the server resource usage struct.
	DisableResourcePolling() error
}
