package error

var flags map[string]string = map[string]string{
	"INVALID_LOGIN": "INVALID_LOGIN",
}

type FlagsType struct{}

var Flags = new(FlagsType)

func (FlagsType) Get(key string) string {
	if flags[key] == "" {
		panic("Flag not found: " + key)
	}
	return flags[key]
}
