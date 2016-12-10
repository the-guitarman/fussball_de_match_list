[![Build Status](https://travis-ci.org/the-guitarman/fussball_de_match_list.svg?branch=master)](https://travis-ci.org/the-guitarman/fussball_de_match_list)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# fussball.de Match List Parser
This is a microservice which parses a match list of a german soccer team from fussball.de

# Usage

Run the service:

````
go run *.go
````

Or compile it and run the executable:

````
go build *.go
./main
````

Now you can use it in your browser:

````
http://localhost:3333/match-list?url=http://www.fussball.de/mannschaft/spvgg-blau-weiss-chemnitz02-spvgg-blau-weiss-chemnitz-02-sachsen/-/saison/1516/team-id/011MIF6PMK000000VTVG0001VTR8C1K7#!/section/stage
````

Returns something like this:

````
{
  "team_name": "Spvgg. Blau-Weiß Chemnitz02",
  "matches": [
    {
      "start_at": "2017-02-26T16:00:00+01:00",
      "competition": "Kreisfreundschaftsspiele",
      "home": "FSV Zschopau/​Krumhermersdorf",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-03-05T14:00:00+01:00",
      "competition": "Kreisoberliga",
      "home": "Post SV Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-03-12T14:00:00+01:00",
      "competition": "Kreispokal",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "FSV Grün-Weiß Klaffenbach"
    },
    {
      "start_at": "2017-03-19T15:00:00+01:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Adorf"
    },
    {
      "start_at": "2017-03-26T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "FSV Grün-Weiß Klaffenbach",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-04-02T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Eiche Reichenbrand"
    },
    {
      "start_at": "2017-04-09T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "USG Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-04-23T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "VfL Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-04-30T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Viktoria 03 Einsiedel"
    },
    {
      "start_at": "2017-05-07T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "FSV Grüna-Mittelbach",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    }
  ]
}
````

"team_name" holds the name of the team to which this list belongs to. You can use it to determine wether a match is a home match or not.

# License
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
