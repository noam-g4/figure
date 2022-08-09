package replacer

type Replacer[Conf interface{}] struct {
	Env    string
	Setter func(Conf, string) Conf
}
