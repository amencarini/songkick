# Songkick

Wrapper for the Songkick API

## Usage

Get your API key at http://www.songkick.com/developer

```golang
client := songkick.NewClient(songkickApiKey)

// 24426 is the location id for London
// See http://www.songkick.com/developer/location-search
events, err := client.GetEvents("2015-07-25", "2015-07-26", 24426)
```

## TODO

- Complete API
- Add tests
- Add doc
