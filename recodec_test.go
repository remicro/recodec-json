package recodecJson

import (
	"encoding/json"
	"github.com/remicro/trifle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type objectForTest struct {
	Label string
}

func TestCoder_Encode(t *testing.T) {
	exp := objectForTest{
		Label: trifle.String(),
	}
	data, err := Encoder().Encode(&exp)
	require.NoError(t, err)
	res := objectForTest{}
	require.NoError(t, json.Unmarshal(data, &res))
	assert.Equal(t, exp, res)
}

func TestCoder_Decode(t *testing.T) {
	exp := objectForTest{
		Label: trifle.String(),
	}
	data, err := Encoder().Encode(&exp)
	require.NoError(t, err)
	res := objectForTest{}
	require.NoError(t, Decoder().Decode(&res, data))
	assert.Equal(t, exp, res)
}
