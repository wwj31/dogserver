package account

func (s *AccountMgr) AccountByUId(uid uint64) *Account {
	return s.accountsByUId[uid]
}

func (s *AccountMgr) AccountByPlatform(platformId string) *Account {
	return s.accountsByPlatformId[platformId]
}

func (s *AccountMgr) AccountByName(name string) *Account {
	return s.accountsByName[name]
}
