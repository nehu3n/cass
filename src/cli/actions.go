package cli

import "cass/src/config"

func initAction() error {
	if err := config.WriteConfigFile(); err != nil {
		return err
	}

	return nil
}
