package ohbemgo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOhbem_QueryPvPRank(t *testing.T) {
	o := Ohbem{Leagues: leagues, LevelCaps: levelCaps}
	_ = o.FetchPokemonData()
	p, _ := o.QueryPvPRank(1, 0, 0, 0, 10, 5, 0, 22.5)

	j, _ := json.Marshal(p)
	fmt.Printf("%s", j)

	for i := 0; i < 100000; i++ {
		_, _ = o.QueryPvPRank(1, 0, 0, 0, 10, 5, 0, 22.5)
	}

}
