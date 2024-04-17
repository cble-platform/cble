package graph

import "entgo.io/ent/dialect/sql"

func IDFuzzySearch(search string) func(*sql.Selector) {
	return func(s *sql.Selector) {
		s.Where(sql.P(func(b *sql.Builder) {
			b.WriteString("id::text").WriteOp(sql.OpLike).Arg("%" + search + "%")
		}))
	}
}
