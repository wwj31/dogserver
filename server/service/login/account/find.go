package account

func (s *Mgr) AccountByUId(uid uint64) *Account {
	return s.accountsByUId[uid]
}

func (s *Mgr) AccountByPlatform(platformId string) *Account {
	return s.accountsByPlatformId[platformId]
}

func (s *Mgr) AccountByName(name string) *Account {
	return s.accountsByName[name]
}
