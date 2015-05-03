package params

import (
	"testing"
	"github.com/stretchr/testify/require"
//	"log"
)

type Params struct {
	local string
	BeginTime string						`param:"begin_time"`
	EndTime string							`param:"end_time"`
	Ninja string								`param:"_"`
	Number int									`param:"number"`
	Money float64								`param:"money"`	
}

type BadParams struct {
	M map[string]string
}

func TestParams(t *testing.T) {

	good := Params{
		local: "things",
		BeginTime: "2013-01-15T00:00:00Z",
		EndTime: "2013-01-31T00:00:00Z",
		Ninja: "things",
		Number: 999,
		Money: 999.999,
	}
	
	bad := BadParams{
		M: make(map[string]string),
	}
	
	s, err := Encode(&good)
	require.Equal(t, err, nil, "Should be able to encode good structs")	
	require.Equal(t, s, "begin_time=2013-01-15T00%3A00%3A00Z&end_time=2013-01-31T00%3A00%3A00Z&money=999.999&number=999", "Should generate correct value")
	
	_, err = Encode(&bad)
	require.NotEqual(t, err, nil, "Should return an error from structs with unsupported types")	
}