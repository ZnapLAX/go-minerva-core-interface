package options

const (
	RSAKey     = "rsa"
	Ed25519Key = "ed25519"

	DefaultRSALen = 4096
)

type KeyGenerateSettings struct {
	Algorithm string
	Size      int
}

type KeyRenameSettings struct {
	Force bool
}

type KeyGenerateOption func(*KeyGenerateSettings) error
type KeyRenameOption func(*KeyRenameSettings) error

func KeyGenerateOptions(opts ...KeyGenerateOption) (*KeyGenerateSettings, error) {
	options := &KeyGenerateSettings{
		Algorithm: RSAKey,
		Size:      -1,
	}
	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func KeyRenameOptions(opts ...KeyRenameOption) (*KeyRenameSettings, error) {
	options := &KeyRenameSettings{
		Force: false,
	}
	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type keyOpts struct{}

var Key keyOpts

func (keyOpts) Type(algorithm string) KeyGenerateOption {
	return func(settings *KeyGenerateSettings) error {
		settings.Algorithm = algorithm
		return nil
	}
}

func (keyOpts) Size(size int) KeyGenerateOption {
	return func(settings *KeyGenerateSettings) error {
		settings.Size = size
		return nil
	}
}

func (keyOpts) Force(force bool) KeyRenameOption {
	return func(settings *KeyRenameSettings) error {
		settings.Force = force
		return nil
	}
}
