package webTypeAdmin

import (
	"strings"
)

//a path to a type in whole type system
type Path []string

func (p Path) String() string {
	return strings.Join(p, ",")
}
func parsePath(ps string) Path {
	psa := strings.Split(ps, ",")
	pso := []string{}
	for _, v := range psa {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		pso = append(pso, v)
	}
	return Path(pso)
}
