package exchange

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCOSS_SetPairs(t *testing.T) {
	coss := COSS{}
	_ = coss.RefreshPairs()
	pairs := coss.GetPairs()

	assert.Contains(t, pairs, &Pair{"ETH", "USD"})
	assert.Contains(t, pairs, &Pair{"ETH", "BTC"})
}

func TestCOSS_GetResponse(t *testing.T) {
	coss := COSS{}
	price, err := coss.GetResponse("ETH", "BTC")
	if err != nil {
		log.Fatal(err)
	}
	assert.True(t, price.Price > 0, "price from COSS isn't greater than 0")
	assert.True(t, price.Volume > 0, "volume from COSS isn't greater than 0")
}
