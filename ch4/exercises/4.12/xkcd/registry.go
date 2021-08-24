package xkcd

import "strconv"

type Registry map[int64]*Comic

func (registry *Registry) add(comic *Comic) {
	(*registry)[comic.Number] = comic
}

func (registry *Registry) get(comicId string) (*Comic, bool) {
	id, _ := strconv.ParseInt(comicId, 10, 64)
	comic, ok := (*registry)[id]
	return comic, ok
}

func newRegistry() Registry {
	return make(Registry)
}
