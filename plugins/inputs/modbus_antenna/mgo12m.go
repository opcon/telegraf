package modbusAntenna

var mgo12m = map[byte][]register{
	0: []register{
		{18402, "AzimuthStowEnable", integer},
		{18403, "UT1TimeOffset", integer},
	},
	3: []register{
		{23885, "Azimuth1", fixedPoint},
		{23887, "SystemClockMJDay", integer},
		{23888, "SystemClockms", integer},
		{23889, "Elevation1", fixedPoint},
	},
}

/*
ReadAddresses[0]  = 23436; // mjd
ReadAddresses[1]  = 23435; // milliseconds
ReadAddresses[2]  = 23704; // Elevation1
ReadAddresses[3]  = 23684; // Azimith 1
ReadAddresses[4]  = 23383; // Central Status
ReadAddresses[5]  = 23683; // Az Master Status
ReadAddresses[6]  = 23693; // Az Slave Status
ReadAddresses[7]  = 23703; // El Status
ReadAddresses[8]  = 23687; // Az Motor Current
ReadAddresses[9]  = 23694; // Az Slave Motor Current
ReadAddresses[10] = 23707; // El Motor Current
ReadAddresses[11] = 18403; // UT1 Time Offset
ReadAddresses[12] = 18402; // AzimuthStowEnable
*/
