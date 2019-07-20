package modbusAntenna

import (
	"bytes"
	"encoding/binary"
)

var isiDegrees = fixedpoint(0.0001)
var isiDegreesPerSec = fixedpoint(0.00000001)

var intertronic12m = map[byte][]register{
	0: []register{
		{18403, "20.20", "ut1_offset", integer},
	},
	// Registers less than 23383
	3: []register{
		{23383, "70.00", "central_status", isiCentral},
		{23384, "70.01", "cmd_power", boolean},
		{23385, "70.02", "cmd_operate", boolean},
		{23386, "70.03", "cmd_run_mode", enum("stop", "position", "velocity", "stow/service", "track")},
		{23387, "70.04", "cmd_offset_mode", enum("stop", "position")},
		{23388, "70.05", "cmd_reset", boolean},
		{23389, "70.06", "cmd_reboot", integer},
		{23390, "70.07", "cmd_stow_service", enum("none", "stow", "service")},
		{23391, "70.08", "cmd_correction", enum("refraction+model", "model", "refraction", "none")},

		{23399, "70.16", "data_mode", enum("horizon", "equatorial")},

		{23401, "70.18", "track_data_source", enum("INVALID", "array", "position+rate")},
		{23402, "70.19", "ra_dec_offset_enabled", boolean},
		{23403, "70.20", "unexpired_points", integer},
		{23404, "70.21", "position_target_az", isiDegrees},
		{23405, "70.22", "position_target_el", isiDegrees},
		{23406, "70.23", "track_target_azra", isiDegrees},
		{23407, "70.24", "track_target_eldec", isiDegrees},
		{23408, "70.25", "ra_offset", isiDegrees},
		{23409, "70.26", "dec_offset", isiDegrees},

		{23410, "70.27", "ra_virtual_axis", isiDegrees},
		{23411, "70.28", "dec_virtual_axis", isiDegrees},

		{23413, "70.30", "loaded_track_points", integer},

		{23424, "70.41", "track_offsets_enabled", boolean},
		{23425, "70.42", "track_offsets_az", isiDegrees},
		{23426, "70.43", "track_offsets_el", isiDegrees},

		{23434, "70.51", "unix_time", integer},

		{23484, "71.01", "part_ramp_azra", isiDegrees},
		{23485, "71.02", "part_ramp_eldec", isiDegrees},
		{23486, "71.03", "part_ramp_azra_vel", isiDegreesPerSec},
		{23487, "71.04", "part_ramp_eldec_vel", isiDegreesPerSec},
		{23488, "71.05", "part_ramp_day", integer},
		{23489, "71.06", "part_ramp_ms", integer},

		{23884, "75.01", "azimuth_error", isiDegrees},
		{23885, "75.02", "azimuth", isiDegrees},
		{23886, "75.03", "azimuth_master_current", integer},
		{23887, "75.04", "mjd_day", integer},
		{23888, "75.05", "mjd_ms", integer},
		{23889, "75.06", "elevation", isiDegrees},
		{23890, "75.07", "elevation_error", isiDegrees},
		{23891, "75.08", "elevation_master_current", integer},
	},
}

// Central status bits
func isiCentral(name string, rawval []byte) (map[string]interface{}, error) {
	const (
		//byte 0
		CentralPowerContactorAux = 1 << iota
		CentralControl
		CentralDrivesSummary
		CentralClockInitialised
		CentralSNTPResponse
		Central30sTimeout
		_ //unused
		CentralAutoStow

		//byte 1
		CentralStowInProgress
		CentralTrackModeAZEL
		CentralTrackModePART
		_ // unsused
		Central30sDisable
		CentralTrackStatus, CentralTrackStatusShift = 3 << iota, iota
		_, _                                        //above consumes 2 bits
		CentralParameterSaveAllowed                 = 1 << iota

		//byte 2
		CentralTrackTableFlushing = 1 << iota
		_                         //byte 2 bit 1 unused
		CentralElevationMasterOnline
		CentralElevationSlaveOnline
		CentralAzimuthMasterOnline
		CentralAzimuthSlaveOnline
		CentralRunMode, CentralRunModeShift = 3 << iota, iota
		_, _                                //above consumes 2 bits

		//byte 3
		CentralHorizonPositionOffset = 1 << iota
		CentralMaintenanceMode
		CentralEquatorialPositionOffset
		CentarlTrackingOffset
		CentralRefractionCorrectionsEnabled
		CentralModelCorrectionsEnabled
		CentralServiceInProgress
		_ //byte 3 bit 7 unused
	)

	var val uint32
	reader := bytes.NewReader(rawval)
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		return nil, err
	}

	return map[string]interface{}{
		name: val,
		//byte 0
		"status_power":             val&CentralPowerContactorAux == 0,
		"status_remote":            val&CentralControl == 0,
		"status_operate":           val&CentralDrivesSummary == 0,
		"status_clock_initialized": val&CentralClockInitialised == 0,
		"status_sntp_active":       val&CentralSNTPResponse == 0,
		"status_timeout_triggered": val&Central30sTimeout == 0,
		//unused
		"status_auto_stow": val&CentralAutoStow != 0,

		// byte 1
		"status_stow_in_progress": val&CentralStowInProgress != 0,
		"status_coord_azel":       val&CentralTrackModeAZEL == 0,
		"status_track_part":       val&CentralTrackModePART != 0,
		//unused
		"status_timeout_enabled": val&Central30sDisable == 0,
		"status_track_status":    val & CentralTrackStatus >> CentralTrackStatusShift,
		// two bits
		"status_parameter_save_allowed": val&CentralParameterSaveAllowed != 0,

		// byte 2
		"status_track_table_flushing": val&CentralTrackTableFlushing != 0,
		// unused
		"status_elevation_master": val&CentralElevationMasterOnline != 0,
		"status_elevation_slave":  val&CentralElevationSlaveOnline != 0,
		"status_azimuth_master":   val&CentralAzimuthMasterOnline != 0,
		"status_azimuth_slave":    val&CentralAzimuthSlaveOnline != 0,
		"status_run_mode":         val & CentralRunMode >> CentralRunModeShift,
		// two bits

		// byte 3
		"status_azel_offset":            val&CentralHorizonPositionOffset != 0,
		"status_maintenance":            val&CentralMaintenanceMode == 0,
		"status_radec_offset":           val&CentralEquatorialPositionOffset != 0,
		"status_refraction_corrections": val&CentralRefractionCorrectionsEnabled == 0,
		"status_model_corrections":      val&CentralModelCorrectionsEnabled == 0,
		"status_serivce_in_progress":    val&CentralServiceInProgress != 0,
		// unused
	}, nil
}
