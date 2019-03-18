package modbusAntenna

var antennas = map[string]map[byte][]register{
	"patriot12m": patriot12m,
	"mgo12m":     mgo12m,
}
