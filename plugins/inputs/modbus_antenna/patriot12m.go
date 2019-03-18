package modbusAntenna

import (
	"bytes"
	"encoding/binary"
)

var patriot12m = map[byte][]register{
	0: []register{
		{18215, "PowerContactorAux", integer},
		{18401, "ConnectionFilter", integer},
		{18402, "AzimuthStowEnable", integer},
		{18403, "UT1TimeOffset", integer},
		{18410, "AntennaLatitude", fixedPoint},
		{18411, "AntennaLongitude", fixedPoint},
		{18412, "AntennaSiteElevation", integer},
		{18413, "AzimuthStowAngle", fixedPoint},
		{18414, "ElevationStowAngle", fixedPoint},
		{18415, "PointCorr1", integer},
		{18416, "PointCorr2", integer},
		{18417, "PointCorr3", integer},
		{18418, "PointCorr4", integer},
		{18419, "PointCorr5", integer},
		{18420, "PointCorr6", integer},
		{18421, "PointCorr7", integer},
		{18422, "PointCorr8", integer},
		{18423, "PointCorr9", integer},
		{23383, "CentralStatus", p12Central},
		{23384, "PowerSwitch", integer},
		{23385, "RunControl", integer},
		{23386, "RunMode", integer},
		{23387, "OffsetMode", integer},
		{23388, "Reset", integer},
		{23389, "Reboot", integer},
		{23390, "Stow", integer},
		{23391, "CorrectionDisable", integer},
		{23394, "AzRaTrackPoint", fixedPoint},
		{23395, "ElDecTrackPoint", fixedPoint},
		{23397, "TimeTrackPoint", integer},
		{23398, "DayTrackPoint", integer},
		{23399, "DataMode", integer},
		{23400, "AzTrackStartTurn", integer},
		{23401, "TrackDataSource", integer},
		{23402, "RADecOffsetMode", integer},
		{23403, "NumberOfUnexpiredPoints", integer},
		{23404, "AzimuthPosition", fixedPoint},
		{23405, "ElevationPosition", fixedPoint},
		{23406, "AzRAPosition", fixedPoint},
		{23407, "ElDecPosition", fixedPoint},
		{23408, "RAOffset", fixedPoint},
		{23409, "DecOffset", fixedPoint},
		{23410, "RAVirtualAxis", fixedPoint},
		{23411, "DecVirtualAxis", fixedPoint},
		{23412, "FlushTrackArray", integer},
		{23418, "RampAzRA", fixedPoint},
		{23419, "RampElDec", fixedPoint},
		{23420, "RampAzRAVelocity", fixedPoint},
		{23421, "RampElDecVelocity", fixedPoint},
		{23422, "RampEpochTime", integer},
		{23423, "RampEpochDay", integer},
		{23424, "AzElOffsetMode", integer},
		{23425, "AzOffset", fixedPoint},
		{23426, "ElOffset", fixedPoint},
		{23428, "TimeoutReset", integer},
		{23429, "TimeoutDisable", integer},
		{23434, "SNTPTime", integer},
		{23435, "SystemClockms", integer},
		{23436, "SystemClockMJDay", integer},
		{23584, "AzimuthPositionCorrected", fixedPoint},
		{23585, "AzimuthVelocity", fixedPoint},
		{23586, "AzimuthPosOffset", fixedPoint},
		{23588, "AzimuthVirtualAxis", fixedPoint},
		{23604, "ElevationPositionCorrected", fixedPoint},
		{23605, "ElevationVelocity", fixedPoint},
		{23606, "ElevationPosOffset", fixedPoint},
		{23608, "ElevationVirtualAxis", fixedPoint},
		{23683, "AzimuthMasterStatus", p12Azmaster},
		{23684, "Azimuth1", fixedPoint},
		{23685, "AzimuthError1", fixedPoint},
		{23686, "AzimuthFeedbackVelocity", fixedPoint},
		{23687, "AzimuthMotorCurrent", integer},
		{23693, "AzimuthSlaveStatus", p12Azslave},
		{23694, "AzimuthSlaveMotorCurrent", integer},
		{23703, "ElevationStatus", p12Elmaster},
		{23704, "Elevation1", fixedPoint},
		{23705, "ElevationError1", fixedPoint},
		{23706, "ElevationFeedbackVelocity", fixedPoint},
		{23707, "ElevationMotorCurrent", integer},
		{25184, "RestartCentral", integer},
	},
}

// Register Filters specific to antenna

// Interpret registers as a real values stored as a fixed-point number,
// encoded in a big-endian 32bit with a scaling factor of 1/10000
func fixedPoint(name string, rawval []byte) map[string]interface{} {
	var val int32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}
	return map[string]interface{}{name: float64(val) / 10000.0}
}

// Azimuth Master Drive status bits
func p12Azmaster(name string, rawval []byte) map[string]interface{} {
	const (
		AzimuthDriveHealthy = 1 << iota
		AzimuthDriveEnergized
		AzimuthDriveControl
		AzimuthBrakeRelease
		AzimuthLowSoftLimit
		AzimuthHighSoftLimit
		AzimuthLowHardLimit
		AzimuthHighHardLimit
		AzimuthMainProfilerAtPosition
		AzimuthMainProfilerAtSpeed
		AzimuthOffsetProfilerAtPosition
		AzimuthOffsetProfilerAtSpeed
		AzimuthDigitalLockToVirtualAxis
		AzimuthPositionDemandLow
		AzimuthPositionDemandHigh
		AzimuthTurnsCountError
		AzimuthDriveEnable
		AzimuthRunPermit
		_ // byte 2, bit 2 is unused
		AzimuthVirtualSpeedLimit
		_ // byte 2, bit 4 is unused
		AzimithMotorBrakeAlarm
		AzimithMotorBrakeIndicator
		//Remainder of array unused
	)
	var val uint32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}

	return map[string]interface{}{
		"AzimuthMasterStatus":             val,
		"AzimuthDriveHealthy":             val&AzimuthDriveHealthy != 0,
		"AzimuthDriveEnergized":           val&AzimuthDriveEnergized != 0,
		"AzimuthDriveControl":             val&AzimuthDriveControl != 0,
		"AzimuthBrakeRelease":             val&AzimuthBrakeRelease != 0,
		"AzimuthLowSoftLimit":             val&AzimuthLowSoftLimit != 0,
		"AzimuthHighSoftLimit":            val&AzimuthHighSoftLimit != 0,
		"AzimuthLowHardLimit":             val&AzimuthLowHardLimit != 0,
		"AzimuthHighHardLimit":            val&AzimuthHighHardLimit != 0,
		"AzimuthMainProfilerAtPosition":   val&AzimuthMainProfilerAtPosition != 0,
		"AzimuthMainProfilerAtSpeed":      val&AzimuthMainProfilerAtSpeed != 0,
		"AzimuthOffsetProfilerAtPosition": val&AzimuthOffsetProfilerAtPosition != 0,
		"AzimuthOffsetProfilerAtSpeed":    val&AzimuthOffsetProfilerAtSpeed != 0,
		"AzimuthDigitalLockToVirtualAxis": val&AzimuthDigitalLockToVirtualAxis != 0,
		"AzimuthPositionDemandLow":        val&AzimuthPositionDemandLow != 0,
		"AzimuthPositionDemandHigh":       val&AzimuthPositionDemandHigh != 0,
		"AzimuthTurnsCountError":          val&AzimuthTurnsCountError != 0,
		"AzimuthDriveEnable":              val&AzimuthDriveEnable != 0,
		"AzimuthRunPermit":                val&AzimuthRunPermit != 0,
		"AzimuthVirtualSpeedLimit":        val&AzimuthVirtualSpeedLimit != 0,
		"AzimithMotorBrakeAlarm":          val&AzimithMotorBrakeAlarm != 0,
		"AzimithMotorBrakeIndicator":      val&AzimithMotorBrakeIndicator != 0,
	}
}

// Azimuth Slave Drive status bits
func p12Azslave(name string, rawval []byte) map[string]interface{} {
	const (
		AzimuthSlaveDriveHealthy = 1 << iota
		AzimuthSlaveDriveEnergized
		AzimuthSlaveDriveEnabled
		AzimuthSlaveMotorBrakeAlarm
		AzimuthSlaveMotorBrakeIndicator
	)
	var val uint32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}

	return map[string]interface{}{
		"AzimuthSlaveStatus":              val,
		"AzimuthSlaveDriveHealthy":        val&AzimuthSlaveDriveHealthy != 0,
		"AzimuthSlaveDriveEnergized":      val&AzimuthSlaveDriveEnergized != 0,
		"AzimuthSlaveDriveEnabled":        val&AzimuthSlaveDriveEnabled != 0,
		"AzimuthSlaveMotorBrakeAlarm":     val&AzimuthSlaveMotorBrakeAlarm != 0,
		"AzimuthSlaveMotorBrakeIndicator": val&AzimuthSlaveMotorBrakeIndicator != 0,
	}
}

// Central status bits
func p12Central(name string, rawval []byte) map[string]interface{} {
	const (
		//byte 0
		CentralPowerContactorAux = 1 << iota
		CentralControl
		CentralDrivesSummary
		CentralClockInitialised
		CentralSNTPResponse
		Central30sTimeout
		CentralTrackArray10
		CentralAutoStow
		//byte 1
		CentralStowInProgress
		CentralTrackModeCoordinateSystem
		CentralRADecTrackDataSource, CentralRADecTrackDataSourceShift   = 3 << iota, iota
		_, _                                                            //above consumes 2 bits
		Central30sDisable                                               = 1 << iota
		CentralAzimuthTrackStartTurn, CentralAzimuthTrackStartTurnShift = 3 << iota, iota
		_, _                                                            //above consumes 2 bits
		CentralAzimuthTrackStartTurnOutOfRange                          = 1 << iota
		//byte 2
		CentralTrackArrayReinitialisation = 1 << iota
		_                                 //byte 2 bit 1 unused
		CentralElevationOnline
		_ //byte 2 bit 3 unused
		CentralAzimuthMasterOnline
		CentralAzimuthSlaveOnline
		CentralRunMode, CentralRunModeShift = 3 << iota, iota
		_, _                                //above consumes 2 bits
		//byte 3
		CentralOffsetMode = 1 << iota
		_                 //byte 3 bit 1 unused
		CentralRADecOffsetMode
		CentralAzElOffsetMode
		CentralCorrectionDisable, CentralCorrectionDisableShift = 3 << iota, iota
		_, _                                                    //above consumes 2 bits
		_, _                                                    //byte 3 bit 6 unused
		CentralConnectionFilter                                 = 1 << iota
	)

	var val uint32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}

	return map[string]interface{}{
		"CentralStatus":                          val,
		"CentralPowerContactorAux":               val&CentralPowerContactorAux != 0,
		"CentralControl":                         val&CentralControl != 0,
		"CentralDrivesSummary":                   val&CentralDrivesSummary != 0,
		"CentralClockInitialised":                val&CentralClockInitialised != 0,
		"CentralSNTPResponse":                    val&CentralSNTPResponse != 0,
		"Central30sTimeout":                      val&Central30sTimeout != 0,
		"CentralTrackArray10":                    val&CentralTrackArray10 != 0,
		"CentralAutoStow":                        val&CentralAutoStow != 0,
		"CentralStowInProgress":                  val&CentralStowInProgress != 0,
		"CentralTrackModeCoordinateSystem":       val&CentralTrackModeCoordinateSystem != 0,
		"CentralRADecTrackDataSource":            val & CentralRADecTrackDataSource >> CentralRADecTrackDataSourceShift,
		"Central30sDisable":                      val&Central30sDisable != 0,
		"CentralAzimuthTrackStartTurn":           val & CentralAzimuthTrackStartTurn >> CentralAzimuthTrackStartTurnShift,
		"CentralAzimuthTrackStartTurnOutOfRange": val&CentralAzimuthTrackStartTurnOutOfRange != 0,
		"CentralTrackArrayReinitialisation":      val&CentralTrackArrayReinitialisation != 0,
		"CentralElevationOnline":                 val&CentralElevationOnline != 0,
		"CentralAzimuthMasterOnline":             val&CentralAzimuthMasterOnline != 0,
		"CentralAzimuthSlaveOnline":              val&CentralAzimuthSlaveOnline != 0,
		"CentralRunMode":                         val & CentralRunMode >> CentralRunModeShift,
		"CentralOffsetMode":                      val&CentralOffsetMode != 0,
		"CentralRADecOffsetMode":                 val&CentralRADecOffsetMode != 0,
		"CentralAzElOffsetMode":                  val&CentralAzElOffsetMode != 0,
		"CentralCorrectionDisable":               val & CentralCorrectionDisable >> CentralCorrectionDisableShift,
		"CentralConnectionFilter":                val&CentralConnectionFilter != 0,
	}
}

// Elevation Drive status bits
func p12Elmaster(name string, rawval []byte) map[string]interface{} {
	const (
		// byte 0
		ElevationDriveHealthy = 1 << iota
		ElevationDriveEnergized
		ElevationDriveControl
		ElevationBrakeRelease
		ElevationLowSoftLimit
		ElevationHighSoftLimit
		ElevationLowHardLimit
		ElevationHighHardLimit
		//byte 1
		ElevationMainProfilerAtPosition
		ElevationMainProfilerAtSpeed
		ElevationOffsetProfilerAtPosition
		ElevationOffsetProfilerAtSpeed
		ElevationDigitalLockToVirtualAxis
		ElevationPositionDemandLow
		ElevationPositionDemandHigh
		_ // byte 1 bit 7 unused
		// byte 2
		ElevationDriveEnable
		ElevationRunPermit
		ElevationMainBrakeAuxContact
		ElevationVirtualSpeedLimit
		ElevationMainBrakeAlarm
		ElevationMotorBrakeAlarm
		ElevationMotorBrakeIndicator
		_ // byte 2 bite 7 unused
		// byte 3 unused
	)

	var val uint32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}

	return map[string]interface{}{
		"ElevationStatus":                   val,
		"ElevationDriveHealthy":             val&ElevationDriveHealthy != 0,
		"ElevationDriveEnergized":           val&ElevationDriveEnergized != 0,
		"ElevationDriveControl":             val&ElevationDriveControl != 0,
		"ElevationBrakeRelease":             val&ElevationBrakeRelease != 0,
		"ElevationLowSoftLimit":             val&ElevationLowSoftLimit != 0,
		"ElevationHighSoftLimit":            val&ElevationHighSoftLimit != 0,
		"ElevationLowHardLimit":             val&ElevationLowHardLimit != 0,
		"ElevationHighHardLimit":            val&ElevationHighHardLimit != 0,
		"ElevationMainProfilerAtPosition":   val&ElevationMainProfilerAtPosition != 0,
		"ElevationMainProfilerAtSpeed":      val&ElevationMainProfilerAtSpeed != 0,
		"ElevationOffsetProfilerAtPosition": val&ElevationOffsetProfilerAtPosition != 0,
		"ElevationOffsetProfilerAtSpeed":    val&ElevationOffsetProfilerAtSpeed != 0,
		"ElevationDigitalLockToVirtualAxis": val&ElevationDigitalLockToVirtualAxis != 0,
		"ElevationPositionDemandLow":        val&ElevationPositionDemandLow != 0,
		"ElevationPositionDemandHigh":       val&ElevationPositionDemandHigh != 0,
		"ElevationDriveEnable":              val&ElevationDriveEnable != 0,
		"ElevationRunPermit":                val&ElevationRunPermit != 0,
		"ElevationMainBrakeAuxContact":      val&ElevationMainBrakeAuxContact != 0,
		"ElevationVirtualSpeedLimit":        val&ElevationVirtualSpeedLimit != 0,
		"ElevationMainBrakeAlarm":           val&ElevationMainBrakeAlarm != 0,
		"ElevationMotorBrakeAlarm":          val&ElevationMotorBrakeAlarm != 0,
		"ElevationMotorBrakeIndicator":      val&ElevationMotorBrakeIndicator != 0,
	}
}
