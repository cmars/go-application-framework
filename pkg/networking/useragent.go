package networking

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/snyk/go-application-framework/pkg/configuration"
)

type UserAgentInfo struct {
	App                           string
	AppVersion                    string
	Integration                   string
	IntegrationVersion            string
	IntegrationEnvironment        string
	IntegrationEnvironmentVersion string
	OS                            string
	Arch                          string
	ProcessName                   string
}

func UserAgentFromConfig(config configuration.Configuration, app string, appVersion string) UserAgentInfo {
	processName, _ := os.Executable()
	integration := config.GetString(configuration.INTEGRATION_NAME)
	integrationVersion := config.GetString(configuration.INTEGRATION_VERSION)
	integrationEnvironment := config.GetString(configuration.INTEGRATION_ENVIRONMENT)
	integrationEnvironmentVersion := config.GetString(configuration.INTEGRATION_ENVIRONMENT_VERSION)
	return UserAgentInfo{
		App:                           app,
		AppVersion:                    appVersion,
		OS:                            runtime.GOOS,
		Arch:                          runtime.GOARCH,
		ProcessName:                   filepath.Base(processName),
		Integration:                   integration,
		IntegrationVersion:            integrationVersion,
		IntegrationEnvironment:        integrationEnvironment,
		IntegrationEnvironmentVersion: integrationEnvironmentVersion,
	}
}

// ToUserAgentHeader returns a string that can be used as a User-Agent header.
// The string is following this format:
// <app>/<appVer> (<os>;<arch>;<procName>) <integration>/<integrationVersion> (<integrationEnv>/<integrationEnvVersion>)
// Everything other than the app, app version and system information (os/arch/process name) is optional.
// The integration environment is only added if the integration is set.
func (s UserAgentInfo) String() string {
	str := fmt.Sprint(
		s.App, "/", s.AppVersion,
		" (", s.OS, ";", s.Arch, ";", s.ProcessName, ")",
	)
	if s.Integration != "" {
		str += fmt.Sprint(" ", s.Integration, "/", s.IntegrationVersion)
		if s.IntegrationEnvironment != "" {
			str += fmt.Sprint(" (", s.IntegrationEnvironment, "/", s.IntegrationEnvironmentVersion, ")")
		}
	}

	return str
}
