package cmd

func ExecuteBf(contents []byte) (string, error) {
	program, err := CompileBf(string(contents))
	if err != nil {
		return "", err
	}
	return program.Execute(), nil
}
