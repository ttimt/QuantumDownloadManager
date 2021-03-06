package setting

// ConfigOption is the signature of functional option for Setting.
type ConfigOption func(u *Setting) error

// NewSetting returns a new instance of Setting.
func NewSetting(configurations ...ConfigOption) (*Setting, error) {
	setting := &Setting{}

	for _, configuration := range configurations {
		if err := configuration(setting); err != nil {
			return nil, err
		}
	}

	return setting, nil
}

// Functional options functions:

// NrOfConcurrentConnection allows setting the value of number of concurrent connection.
func NrOfConcurrentConnection(nrOfConcurrentConnection int) ConfigOption {
	return func(s *Setting) error {
		return s.SetNrOfConcurrentConnection(nrOfConcurrentConnection)
	}
}
