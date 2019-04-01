package modbusAntenna

import (
	"bytes"
	"encoding/binary"
)

var p12Degrees = fixedpoint(0.0001)

var patriot12m = map[byte][]register{
	0: []register{
		{18215, "", "PowerContactorAux", integer},
		{18401, "", "ConnectionFilter", integer},
		{18402, "", "AzimuthStowEnable", integer},
		{18403, "", "UT1TimeOffset", integer},
		{18410, "", "AntennaLatitude", p12Degrees},
		{18411, "", "AntennaLongitude", p12Degrees},
		{18412, "", "AntennaSiteElevation", integer},
		{18413, "", "AzimuthStowAngle", p12Degrees},
		{18414, "", "ElevationStowAngle", p12Degrees},
		{18415, "", "PointCorr1", integer},
		{18416, "", "PointCorr2", integer},
		{18417, "", "PointCorr3", integer},
		{18418, "", "PointCorr4", integer},
		{18419, "", "PointCorr5", integer},
		{18420, "", "PointCorr6", integer},
		{18421, "", "PointCorr7", integer},
		{18422, "", "PointCorr8", integer},
		{18423, "", "PointCorr9", integer},
		{23383, "", "CentralStatus", p12Central},
		{23384, "", "PowerSwitch", integer},
		{23385, "", "RunControl", integer},
		{23386, "", "RunMode", integer},
		{23387, "", "OffsetMode", integer},
		{23388, "", "Reset", integer},
		{23389, "", "Reboot", integer},
		{23390, "", "Stow", integer},
		{23391, "", "CorrectionDisable", integer},
		{23394, "", "AzRaTrackPoint", p12Degrees},
		{23395, "", "ElDecTrackPoint", p12Degrees},
		{23397, "", "TimeTrackPoint", integer},
		{23398, "", "DayTrackPoint", integer},
		{23399, "", "DataMode", integer},
		{23400, "", "AzTrackStartTurn", integer},
		{23401, "", "TrackDataSource", integer},
		{23402, "", "RADecOffsetMode", integer},
		{23403, "", "NumberOfUnexpiredPoints", integer},
		{23404, "", "AzimuthPosition", p12Degrees},
		{23405, "", "ElevationPosition", p12Degrees},
		{23406, "", "AzRAPosition", p12Degrees},
		{23407, "", "ElDecPosition", p12Degrees},
		{23408, "", "RAOffset", p12Degrees},
		{23409, "", "DecOffset", p12Degrees},
		{23410, "", "RAVirtualAxis", p12Degrees},
		{23411, "", "DecVirtualAxis", p12Degrees},
		{23412, "", "FlushTrackArray", integer},
		{23418, "", "RampAzRA", p12Degrees},
		{23419, "", "RampElDec", p12Degrees},
		{23420, "", "RampAzRAVelocity", p12Degrees},
		{23421, "", "RampElDecVelocity", p12Degrees},
		{23422, "", "RampEpochTime", integer},
		{23423, "", "RampEpochDay", integer},
		{23424, "", "AzElOffsetMode", integer},
		{23425, "", "AzOffset", p12Degrees},
		{23426, "", "ElOffset", p12Degrees},
		{23428, "", "TimeoutReset", integer},
		{23429, "", "TimeoutDisable", integer},
		{23434, "", "SNTPTime", integer},
		{23435, "", "SystemClockms", integer},
		{23436, "", "SystemClockMJDay", integer},
		{23584, "", "AzimuthPositionCorrected", p12Degrees},
		{23585, "", "AzimuthVelocity", p12Degrees},
		{23586, "", "AzimuthPosOffset", p12Degrees},
		{23588, "", "AzimuthVirtualAxis", p12Degrees},
		{23604, "", "ElevationPositionCorrected", p12Degrees},
		{23605, "", "ElevationVelocity", p12Degrees},
		{23606, "", "ElevationPosOffset", p12Degrees},
		{23608, "", "ElevationVirtualAxis", p12Degrees},
		{23683, "", "AzimuthMasterStatus", p12Azmaster},
		{23684, "", "Azimuth1", p12Degrees},
		{23685, "", "AzimuthError1", p12Degrees},
		{23686, "", "AzimuthFeedbackVelocity", p12Degrees},
		{23687, "", "AzimuthMotorCurrent", integer},
		{23693, "", "AzimuthSlaveStatus", p12Azslave},
		{23694, "", "AzimuthSlaveMotorCurrent", integer},
		{23703, "", "ElevationStatus", p12Elmaster},
		{23704, "", "Elevation1", p12Degrees},
		{23705, "", "ElevationError1", p12Degrees},
		{23706, "", "ElevationFeedbackVelocity", p12Degrees},
		{23707, "", "ElevationMotorCurrent", integer},
		{25184, "", "RestartCentral", integer},
	},
}

// Register Filters specific to antenna

// Azimuth Master Drive status bits
func p12Azmaster(name string, rawval []byte) (map[string]interface{}, error) {
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
		return nil, err
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
	}, nil
}

// Azimuth Slave Drive status bits
func p12Azslave(name string, rawval []byte) (map[string]interface{}, error) {
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
		return nil, err
	}

	return map[string]interface{}{
		"AzimuthSlaveStatus":              val,
		"AzimuthSlaveDriveHealthy":        val&AzimuthSlaveDriveHealthy != 0,
		"AzimuthSlaveDriveEnergized":      val&AzimuthSlaveDriveEnergized != 0,
		"AzimuthSlaveDriveEnabled":        val&AzimuthSlaveDriveEnabled != 0,
		"AzimuthSlaveMotorBrakeAlarm":     val&AzimuthSlaveMotorBrakeAlarm != 0,
		"AzimuthSlaveMotorBrakeIndicator": val&AzimuthSlaveMotorBrakeIndicator != 0,
	}, nil
}

// Central status bits
func p12Central(name string, rawval []byte) (map[string]interface{}, error) {
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
		return nil, err
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
	}, nil
}

// Elevation Drive status bits
func p12Elmaster(name string, rawval []byte) (map[string]interface{}, error) {
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
		return nil, err
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
	}, nil
}
