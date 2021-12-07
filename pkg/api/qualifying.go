package api

import (
	"context"
	"fmt"
)

func (ergast *Ergast) QualifyingResults(ctx context.Context, year int, round string) ([]QualifyingResult, error) {

	url := fmt.Sprintf("%s/%d", ergast.BaseURL, year)
	qualifyingResults, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}
	return qualifyingResults.RaceTable.Races[0].QualifyingResult, nil
}
