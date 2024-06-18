package runtime

import "os"

type RunEnv string

const (
	RunLocalEnv = "local"
	RunTestEnv  = "test"
	RunProdEnv  = "prod"
)

func (e RunEnv) FromStr(env string) RunEnv {
	if env == "prod" {
		return RunProdEnv
	} else if env == "test" {
		return RunTestEnv
	} else {
		return RunLocalEnv
	}
}

// Env 全局环境变量
type Env struct {
	AppName      string
	RunEnv       RunEnv
	WorkDir      string
	BuildVersion string
}

func (e *Env) LoadRunEnv() {
	var env = os.Getenv(e.AppName)
	if len(env) <= 0 {
		e.RunEnv = RunLocalEnv
	}
	e.RunEnv = RunEnv.FromStr(RunLocalEnv, env)
}
