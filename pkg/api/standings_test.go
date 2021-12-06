package api

import (
	"context"
	"errors"
	"testing"

	"github.com/AidanFogarty/go-ergast/pkg/testutil"
)

const (
	successDrivingStandingsRespXML = `
	<?xml version="1.0" encoding="utf-8"?>
	<?xml-stylesheet type="text/xsl" href="/schemas/mrd-1.4.xsl"?>
	<MRData xmlns="http://ergast.com/mrd/1.4" series="f1" url="http://ergast.com/api/f1/2021/driverstandings" limit="30" offset="0" total="21">
		<StandingsTable season="2021">
			<StandingsList season="2021" round="17">
				<DriverStanding position="1" positionText="1" points="287.5" wins="8">
					<Driver driverId="max_verstappen" code="VER" url="http://en.wikipedia.org/wiki/Max_Verstappen">
						<PermanentNumber>33</PermanentNumber>
						<GivenName>Max</GivenName>
						<FamilyName>Verstappen</FamilyName>
						<DateOfBirth>1997-09-30</DateOfBirth>
						<Nationality>Dutch</Nationality>
					</Driver>
					<Constructor constructorId="red_bull" url="http://en.wikipedia.org/wiki/Red_Bull_Racing">
						<Name>Red Bull</Name>
						<Nationality>Austrian</Nationality>
					</Constructor>
				</DriverStanding>
				<DriverStanding position="2" positionText="2" points="275.5" wins="5">
					<Driver driverId="hamilton" code="HAM" url="http://en.wikipedia.org/wiki/Lewis_Hamilton">
						<PermanentNumber>44</PermanentNumber>
						<GivenName>Lewis</GivenName>
						<FamilyName>Hamilton</FamilyName>
						<DateOfBirth>1985-01-07</DateOfBirth>
						<Nationality>British</Nationality>
					</Driver>
					<Constructor constructorId="mercedes" url="http://en.wikipedia.org/wiki/Mercedes-Benz_in_Formula_One">
						<Name>Mercedes</Name>
						<Nationality>German</Nationality>
					</Constructor>
				</DriverStanding>
			</StandingsList>
		</StandingsTable>
	</MRData>
	`
	noDriverStandingsAvailableRespXML = `
	<?xml version="1.0" encoding="utf-8"?>
	<?xml-stylesheet type="text/xsl" href="/schemas/mrd-1.4.xsl"?>
	<MRData xmlns="http://ergast.com/mrd/1.4" series="f1" url="http://ergast.com/api/f1/1900/driverstandings" limit="30" offset="0" total="0">
		<StandingsTable season="1900">
		</StandingsTable>
	</MRData>
	`
)

func TestErgast_DriverStandings(t *testing.T) {
	type args struct {
		ctx  context.Context
		year int
	}
	tests := []struct {
		name           string
		args           args
		response       string
		wantErr        bool
		expectedErr    error
		expectedLength int
	}{
		{
			name:           "Successful Driver Standings Response",
			response:       successDrivingStandingsRespXML,
			args:           args{context.TODO(), 2021},
			wantErr:        false,
			expectedLength: 2,
		},
		{
			name:           "No Driver Standings Response returns Error",
			response:       noDriverStandingsAvailableRespXML,
			args:           args{context.TODO(), 1900},
			wantErr:        true,
			expectedErr:    errors.New("no races returned for season provided"),
			expectedLength: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := newTestServer(test.response)

			ergast := Ergast{
				BaseURL:    srv.URL,
				HTTPClient: srv.Client(),
			}

			standings, err := ergast.DriverStandings(test.args.ctx, test.args.year)

			testutil.Equals(t, test.expectedLength, len(standings))
			testutil.Equals(t, test.expectedErr, err)

			if (err != nil) != test.wantErr {
				t.Errorf("Ergast.DriverStandings() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
