package repository

func Close() error {
	// Close all repositories
	if err := closeRepository(); err != nil {
		return err
	}

	return nil
}
