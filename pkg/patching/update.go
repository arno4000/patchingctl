package patching

func UpdateOS() error {
	currentOS, err := DetectOS()
	if err != nil {
		return err
	}
	switch currentOS {
	case "debian":
		return updateDebian()
	case "rhel":
		return updateRHEL()
	}
	return nil
}

func updateDebian() error {
	err := ExecuteCommand("apt-get update")
	if err != nil {
		return err
	}
	err = ExecuteCommand("DEBIAN_FRONTEND=noninteractive apt-get upgrade -y")
	if err != nil {
		return err
	}
	return nil
}

func updateRHEL() error {
	err := ExecuteCommand("dnf upgrade -y")
	if err != nil {
		return err
	}
	return nil
}
