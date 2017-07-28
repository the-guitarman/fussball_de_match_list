[![Build Status](https://travis-ci.org/the-guitarman/fussball_de_match_list.svg?branch=master)](https://travis-ci.org/the-guitarman/fussball_de_match_list)
[![Code Climate](https://codeclimate.com/github/the-guitarman/fussball_de_match_list/badges/gpa.svg)](https://codeclimate.com/github/the-guitarman/fussball_de_match_list)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# fussball.de Match List Parser

This is a microservice which parses a match list of a german soccer team from fussball.de

# Usage

Compile it and run the executable:

````
go build *.go
````

Run to get some help:

````
./main --help
````

Run with url option:

````
./main --url http://www.fussball.de/mannschaft/spvgg-blau-weiss-chemnitz02-spvgg-blau-weiss-chemnitz-02-sachsen/-/saison/1718/team-id/011MIF6PMK000000VTVG0001VTR8C1K7
````

Or run it in server mode:

````
./main --serve
````

Now you can use it in your browser:

````
http://localhost:3333/match-list?url=http://www.fussball.de/mannschaft/spvgg-blau-weiss-chemnitz02-spvgg-blau-weiss-chemnitz-02-sachsen/-/saison/1718/team-id/011MIF6PMK000000VTVG0001VTR8C1K7
````

You may use another port:

````
./main --serve --port 4444
````

Returns something like this:

````
{
  "team_name": "Spvgg. Blau-Weiß Chemnitz02",
  "matches": [
    {
      "start_at": "2017-07-30T15:00:00+02:00",
      "competition": "Kreisfreundschaftsspiele",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "FSV Limbach-Oberfrohna"
    },
    {
      "start_at": "2017-08-06T15:00:00+02:00",
      "competition": "Sachsen-Pokal",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Liebertwolkwitz"
    },
    {
      "start_at": "2017-08-13T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "USG Chemnitz"
    },
    {
      "start_at": "2017-08-20T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "SG Neukirchen/​E.",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-08-27T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "BSC Rapid Chemnitz 2"
    },
    {
      "start_at": "2017-09-10T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "VTB Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-09-17T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Eiche Reichenbrand"
    },
    {
      "start_at": "2017-09-24T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "SV Adorf",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "2017-10-01T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "FSV Grün-Weiß Klaffenbach"
    },
    {
      "start_at": "2017-10-15T15:00:00+02:00",
      "competition": "Kreisoberliga",
      "home": "VfL Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    }
  ]
}
````

"team_name" holds the name of the team to which this list belongs to. You can use it to determine wether a match is a home match or not.

# License
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
