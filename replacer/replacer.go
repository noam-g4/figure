package replacer

func ReplaceConfigWithEnv[C interface{}](
	conf C,
	rs []Replacer[C],
) C {

	if len(rs) == 0 {
		return conf
	}

	ok, val := GetEnv(rs[0].Env)
	if !ok {
		return ReplaceConfigWithEnv(conf, rs[1:])
	}

	newConf := rs[0].Setter(conf, val)
	return ReplaceConfigWithEnv(newConf, rs[1:])

}
