package api

import (
	"context"
	"fmt"
)

// FastestLap returns driver who finished the nth fastest lap.
// To get the fastest lap, fastestRank = 1.
func (ergast *Ergast) Fastest(ctx context.Context, year int, round int, fastestRank int) (interface{}, error) {

	url := fmt.Sprintf("%s/%d/%d/fastest/%d/results", ergast.BaseURL, year, round, fastestRank)
	results, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}

	fmt.Println(results)
	return nil, nil
}
