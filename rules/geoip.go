package rules

import (
	C "github.com/Dreamacro/clash/constant"

	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
)

var mmdb *geoip2.Reader

func init() {
	var err error
	mmdb, err = geoip2.Open(C.MMDBPath)
	if err != nil {
		log.Fatalf("Can't load mmdb: %s", err.Error())
	}
}

type GEOIP struct {
	country string
	adapter string
}

func (g *GEOIP) RuleType() C.RuleType {
	return C.GEOIP
}

func (g *GEOIP) IsMatch(addr *C.Addr) bool {
	if addr.IP == nil {
		return false
	}
	record, _ := mmdb.Country(*addr.IP)
	return record.Country.IsoCode == g.country
}

func (g *GEOIP) Adapter() string {
	return g.adapter
}

func (g *GEOIP) Payload() string {
	return g.country
}

func NewGEOIP(country string, adapter string) *GEOIP {
	return &GEOIP{
		country: country,
		adapter: adapter,
	}
}
