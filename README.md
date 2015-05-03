# params
Create encoded url parameters from structs

## Install
`go get github.com/calce/params`

## Example
```go
type Params struct {
	local string		// unexported fields are ignored
	BeginTime string	`param:"begin_time"`
	EndTime string		`param:"end_time"`
	Ninja string		`param:"_"`	// fields with param tag "_" are ignored
	Number int			`param:"number"`
	Money float64		`param:"money"`	
}

p := Params{
	local: "things",
	BeginTime: "2013-01-15T00:00:00Z",
	EndTime: "2013-01-31T00:00:00Z",
	Ninja: "things",
	Number: 999,
	Money: 999.999,
}
s, err := Encode(&p)
// "begin_time=2013-01-15T00%3A00%3A00Z&end_time=2013-01-31T00%3A00%3A00Z&money=999.999&number=999"
```