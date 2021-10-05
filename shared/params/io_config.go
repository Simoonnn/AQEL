package params

import (
	"os"
	"runtime"
	"time"
)

type IoConfig struct {
	ReadWritePermissions        os.FileMode
	ReadWriteExecutePermissions os.FileMode
	BoltTimeout                 time.Duration
}

var defaultIoConfig = &IoConfig{
	ReadWritePermissions:        0600,            //-rw------- Read and Write permissions for user
	ReadWriteExecutePermissions: 0700,            //-rwx------ Read Write and Execute (traverse) permissions for user
	BoltTimeout:                 1 * time.Second, // 1 second for the bolt DB to timeout.
}

var defaultWindowsIoConfig = &IoConfig{
	ReadWritePermissions:        0666,
	ReadWriteExecutePermissions: 0777,
	BoltTimeout:                 1 * time.Second,
}

// RaidoIoConfig returns the current io config for
// the raido blockchain.
func RaidoIoConfig() *IoConfig {
	if runtime.GOOS == "windows" {
		return defaultWindowsIoConfig
	}
	return defaultIoConfig
}
