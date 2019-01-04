# goscrptr

**Example**

`./goscrptr -eval "Good morning, it's {{timeSpeech}}, the weather in {{city 90210}} is {{temperature 90210}} degrees and weather conditions are {{weatherCondition 90210}}"`

**Result**

`Good morning, it's 4 33 PM, the weather in Beverly Hills is 74.1 degrees and weather conditions are Clear`

**Keywords**

- `time` - Current time (ie. "6:47 PM")
- `timeSpeech` - Speech-friendly time format (ie. "6 47 PM")
- `city <ZIP>` - City name from ZIP code
- `temperature <ZIP>` - Temprature at ZIP code
- `weatherCondition <ZIP>` - Weather conditions at ZIP code
- `spotifyTrack` - Current Spotify track
- `spotifyArtist` - Current Spotify artist