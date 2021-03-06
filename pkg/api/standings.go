package api

import (
	"context"
	"errors"
	"fmt"
)

// DriverStandings retrieves the driver standings for a given season `year`.
func (ergast *Ergast) DriverStandings(ctx context.Context, year int) ([]DriverStanding, error) {

	url := fmt.Sprintf("%s/%d/driverStandings", ergast.BaseURL, year)
	driverStandings, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}

	if len(driverStandings.StandingsTable.DriverStandings) == 0 {
		return nil, errors.New("no races returned for season provided")
	}

	return driverStandings.StandingsTable.DriverStandings, nil
}

// ConstructorStandings retrieves the constructors standings for a given season `year`.
func (ergast *Ergast) ConstructorStandings(ctx context.Context, year int) ([]ConstructorStanding, error) {

	url := fmt.Sprintf("%s/%d/constructorStandings", ergast.BaseURL, year)
	constructorStandings, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}

	if len(constructorStandings.StandingsTable.ConstructorStandings) == 0 {
		return nil, errors.New("no races returned for season provided")
	}

	return constructorStandings.StandingsTable.ConstructorStandings, nil
}
