package modbusAntenna

var antennas = map[string]map[byte][]register{
	"patriot12m":      patriot12m,
	"intertronics12m": intertronic12m,
}
