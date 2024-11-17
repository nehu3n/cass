package config

type ConfigStruct struct {
	DryRun bool `toml:"dry-run" comment:"View the commit before sending it."`
}
