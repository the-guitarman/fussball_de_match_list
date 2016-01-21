[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# fussball_de_match_list
Microservice that parses a match list of a german soccer team from fussball.de

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
http://localhost:3333/match-list?url=http://www.fussball.de/mannschaft/spvgg-blau-weiss-chemnitz02-spvgg-blau-weiss-chemnitz-02-sachsen/-/saison/1516/team-id/011MIF6PMK000000VTVG0001VTR8C1K7#!/section/stage`
````

Returns something like this:

````
{
  "team_name": "Spvgg. Blau-Weiß Chemnitz02",
  "matches": [
    {
      "start_at": "Sonntag, 24.01.2016 - 11:00 Uhr",
      "home": "SV Eiche Reichenbrand 2",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "Samstag, 13.02.2016 - 15:00 Uhr",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV 1990 Witzschdorf"
    },
    {
      "start_at": "Samstag, 27.02.2016 - 15:00 Uhr",
      "home": "TV Oberfrohna",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "Sonntag, 06.03.2016 - 14:00 Uhr",
      "home": "VfL Chemnitz",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "Sonntag, 13.03.2016 - 14:00 Uhr",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "CSV Siegmar"
    },
    {
      "start_at": "Sonntag, 20.03.2016 - 10:30 Uhr",
      "home": "BSC Rapid Chemnitz 2",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "Sonntag, 03.04.2016 - 15:00 Uhr",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "TSV IFA Chemnitz"
    },
    {
      "start_at": "Samstag, 09.04.2016 - 15:00 Uhr",
      "home": "SG Handwerk Rabenstein 2",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    },
    {
      "start_at": "Sonntag, 24.04.2016 - 15:00 Uhr",
      "home": "Spvgg. Blau-Weiß Chemnitz02",
      "guest": "SV Adorf"
    },
    {
      "start_at": "Sonntag, 01.05.2016 - 15:00 Uhr",
      "home": "FSV Grün-Weiß Klaffenbach",
      "guest": "Spvgg. Blau-Weiß Chemnitz02"
    }
  ]
}
````

# License
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
