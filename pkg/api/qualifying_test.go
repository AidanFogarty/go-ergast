package api

import (
	"context"
	"testing"

	"github.com/AidanFogarty/go-ergast/pkg/testutil"
)

const (
	successQualifyingRespXML = `
	<?xml version="1.0" encoding="utf-8"?>
	<?xml-stylesheet type="text/xsl" href="/schemas/mrd-1.4.xsl"?>
	<MRData xmlns="http://ergast.com/mrd/1.4" series="f1" url="http://ergast.com/api/f1/2021/last/qualifying" limit="30" offset="0" total="10">
		<RaceTable season="2021" round="21">
			<Race season="2021" round="21" url="http://en.wikipedia.org/wiki/2021_Saudi_Arabian_Grand_Prix">
				<RaceName>Saudi Arabian Grand Prix</RaceName>
				<Circuit circuitId="jeddah" url="http://en.wikipedia.org/wiki/Jeddah_Street_Circuit">
					<CircuitName>Jeddah Street Circuit</CircuitName>
					<Location lat="21.5433" long="39.1728">
						<Locality>Jeddah</Locality>
						<Country>Saudi Arabia</Country>
					</Location>
				</Circuit>
				<Date>2021-12-05</Date>
				<Time>17:30:00Z</Time>
				<QualifyingList>
					<QualifyingResult number="44" position="1">
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
						<Q1>1:28.466</Q1>
						<Q2>1:27.712</Q2>
						<Q3>1:27.511</Q3>
					</QualifyingResult>
					<QualifyingResult number="77" position="2">
						<Driver driverId="bottas" code="BOT" url="http://en.wikipedia.org/wiki/Valtteri_Bottas">
							<PermanentNumber>77</PermanentNumber>
							<GivenName>Valtteri</GivenName>
							<FamilyName>Bottas</FamilyName>
							<DateOfBirth>1989-08-28</DateOfBirth>
							<Nationality>Finnish</Nationality>
						</Driver>
						<Constructor constructorId="mercedes" url="http://en.wikipedia.org/wiki/Mercedes-Benz_in_Formula_One">
							<Name>Mercedes</Name>
							<Nationality>German</Nationality>
						</Constructor>
						<Q1>1:28.057</Q1>
						<Q2>1:28.054</Q2>
						<Q3>1:27.622</Q3>
					</QualifyingResult>
					<QualifyingResult number="33" position="3">
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
						<Q1>1:28.285</Q1>
						<Q2>1:27.953</Q2>
						<Q3>1:27.653</Q3>
					</QualifyingResult>
					<QualifyingResult number="16" position="4">
						<Driver driverId="leclerc" code="LEC" url="http://en.wikipedia.org/wiki/Charles_Leclerc">
							<PermanentNumber>16</PermanentNumber>
							<GivenName>Charles</GivenName>
							<FamilyName>Leclerc</FamilyName>
							<DateOfBirth>1997-10-16</DateOfBirth>
							<Nationality>Monegasque</Nationality>
						</Driver>
						<Constructor constructorId="ferrari" url="http://en.wikipedia.org/wiki/Scuderia_Ferrari">
							<Name>Ferrari</Name>
							<Nationality>Italian</Nationality>
						</Constructor>
						<Q1>1:28.310</Q1>
						<Q2>1:28.459</Q2>
						<Q3>1:28.054</Q3>
					</QualifyingResult>
					<QualifyingResult number="11" position="5">
						<Driver driverId="perez" code="PER" url="http://en.wikipedia.org/wiki/Sergio_P%C3%A9rez">
							<PermanentNumber>11</PermanentNumber>
							<GivenName>Sergio</GivenName>
							<FamilyName>Pérez</FamilyName>
							<DateOfBirth>1990-01-26</DateOfBirth>
							<Nationality>Mexican</Nationality>
						</Driver>
						<Constructor constructorId="red_bull" url="http://en.wikipedia.org/wiki/Red_Bull_Racing">
							<Name>Red Bull</Name>
							<Nationality>Austrian</Nationality>
						</Constructor>
						<Q1>1:28.021</Q1>
						<Q2>1:27.946</Q2>
						<Q3>1:28.123</Q3>
					</QualifyingResult>
					<QualifyingResult number="10" position="6">
						<Driver driverId="gasly" code="GAS" url="http://en.wikipedia.org/wiki/Pierre_Gasly">
							<PermanentNumber>10</PermanentNumber>
							<GivenName>Pierre</GivenName>
							<FamilyName>Gasly</FamilyName>
							<DateOfBirth>1996-02-07</DateOfBirth>
							<Nationality>French</Nationality>
						</Driver>
						<Constructor constructorId="alphatauri" url="http://en.wikipedia.org/wiki/Scuderia_AlphaTauri">
							<Name>AlphaTauri</Name>
							<Nationality>Italian</Nationality>
						</Constructor>
						<Q1>1:28.401</Q1>
						<Q2>1:28.314</Q2>
						<Q3>1:28.125</Q3>
					</QualifyingResult>
					<QualifyingResult number="4" position="7">
						<Driver driverId="norris" code="NOR" url="http://en.wikipedia.org/wiki/Lando_Norris">
							<PermanentNumber>4</PermanentNumber>
							<GivenName>Lando</GivenName>
							<FamilyName>Norris</FamilyName>
							<DateOfBirth>1999-11-13</DateOfBirth>
							<Nationality>British</Nationality>
						</Driver>
						<Constructor constructorId="mclaren" url="http://en.wikipedia.org/wiki/McLaren">
							<Name>McLaren</Name>
							<Nationality>British</Nationality>
						</Constructor>
						<Q1>1:28.338</Q1>
						<Q2>1:28.344</Q2>
						<Q3>1:28.180</Q3>
					</QualifyingResult>
					<QualifyingResult number="22" position="8">
						<Driver driverId="tsunoda" code="TSU" url="http://en.wikipedia.org/wiki/Yuki_Tsunoda">
							<PermanentNumber>22</PermanentNumber>
							<GivenName>Yuki</GivenName>
							<FamilyName>Tsunoda</FamilyName>
							<DateOfBirth>2000-05-11</DateOfBirth>
							<Nationality>Japanese</Nationality>
						</Driver>
						<Constructor constructorId="alphatauri" url="http://en.wikipedia.org/wiki/Scuderia_AlphaTauri">
							<Name>AlphaTauri</Name>
							<Nationality>Italian</Nationality>
						</Constructor>
						<Q1>1:28.503</Q1>
						<Q2>1:28.222</Q2>
						<Q3>1:28.442</Q3>
					</QualifyingResult>
					<QualifyingResult number="31" position="9">
						<Driver driverId="ocon" code="OCO" url="http://en.wikipedia.org/wiki/Esteban_Ocon">
							<PermanentNumber>31</PermanentNumber>
							<GivenName>Esteban</GivenName>
							<FamilyName>Ocon</FamilyName>
							<DateOfBirth>1996-09-17</DateOfBirth>
							<Nationality>French</Nationality>
						</Driver>
						<Constructor constructorId="alpine" url="http://en.wikipedia.org/wiki/Alpine_F1_Team">
							<Name>Alpine F1 Team</Name>
							<Nationality>French</Nationality>
						</Constructor>
						<Q1>1:28.752</Q1>
						<Q2>1:28.574</Q2>
						<Q3>1:28.647</Q3>
					</QualifyingResult>
					<QualifyingResult number="99" position="10">
						<Driver driverId="giovinazzi" code="GIO" url="http://en.wikipedia.org/wiki/Antonio_Giovinazzi">
							<PermanentNumber>99</PermanentNumber>
							<GivenName>Antonio</GivenName>
							<FamilyName>Giovinazzi</FamilyName>
							<DateOfBirth>1993-12-14</DateOfBirth>
							<Nationality>Italian</Nationality>
						</Driver>
						<Constructor constructorId="alfa" url="http://en.wikipedia.org/wiki/Alfa_Romeo_in_Formula_One">
							<Name>Alfa Romeo</Name>
							<Nationality>Swiss</Nationality>
						</Constructor>
						<Q1>1:28.899</Q1>
						<Q2>1:28.616</Q2>
						<Q3>1:28.754</Q3>
					</QualifyingResult>
				</QualifyingList>
			</Race>
		</RaceTable>
	</MRData>
	`
)

func TestErgast_QualifyingResults(t *testing.T) {
	type args struct {
		ctx   context.Context
		year  int
		round string
	}
	tests := []struct {
		name        string
		args        args
		response    string
		want        []QualifyingResult
		expectedLen int
		wantErr     bool
	}{
		{
			name:        "Successful Qualifying Response",
			args:        args{context.TODO(), 2021, "last"},
			response:    successQualifyingRespXML,
			wantErr:     false,
			expectedLen: 10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := newTestServer(test.response)

			ergast := &Ergast{
				BaseURL:    srv.URL,
				HTTPClient: srv.Client(),
			}
			got, err := ergast.QualifyingResults(test.args.ctx, test.args.year, test.args.round)

			testutil.Equals(t, test.expectedLen, len(got))

			if (err != nil) != test.wantErr {
				t.Errorf("Ergast.QualifyingResults() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
