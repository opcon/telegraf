// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types.go

package fieldsystem

const (
	FieldSystemVersion = "9.12.8"
)

type (
	Bank_set_mon struct {
		Active_bank struct {
			Active_bank [33]byte
			pad_cgo_0   [3]byte
			State       M5state
		}
		Active_vsn struct {
			Active_vsn [33]byte
			pad_cgo_0  [3]byte
			State      M5state
		}
		Inactive_bank struct {
			Inactive_bank [33]byte
			pad_cgo_0     [3]byte
			State         M5state
		}
		Inactive_vsn struct {
			Inactive_vsn [33]byte
			pad_cgo_0    [3]byte
			State        M5state
		}
	}
	Bbc_cmd struct {
		Freq   int64
		Source int32
		Bw     [2]int32
		Bwcomp [2]int32
		Gain   struct {
			Mode  int32
			Value [2]int32
			Old   int32
		}
		Avper int32
	}
	Bbc_mon struct {
		Lock   int32
		Pwr    [2]uint32
		Serno  int32
		Timing int32
	}
	Bs struct {
		Image_reject_filter uint32
		Level               Servo
		Offset              Servo
		Magn_stats          uint32
		Flip_64MHz_out      uint32
		Digital_format      uint32
		Flip_input          uint32
		P_hilbert_no        uint8
		N_hilbert_no        uint8
		pad_cgo_0           [2]byte
		Sub_band            uint32
		Q_fir_no            uint8
		I_fir_no            uint8
		Clock_decimation    int8
		pad_cgo_1           [1]byte
		Add_sub             Mux
		Usb_mux             Mux
		Lsb_mux             Mux
		Usb_threshold       uint8
		Lsb_threshold       uint8
		pad_cgo_2           [2]byte
		Usb_servo           Servo
		Lsb_servo           Servo
		Flip_usb            uint32
		Flip_lsb            uint32
		Monitor             Mux
		Digout              Digout
	}
	Calrx_cmd struct {
		File      [65]byte
		pad_cgo_0 [3]byte
		Type      int32
		Lo        [2]float64
	}
	Capture_mon struct {
		Qa struct {
			Drive int32
			Chan  int32
		}
		General struct {
			Word1 uint32
			Word2 uint32
		}
		Time struct {
			Word3 uint32
			Word4 uint32
		}
	}
	Clock_set_cmd struct {
		Freq struct {
			Freq  int32
			State M5state
		}
		Source struct {
			Source    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Clock_gen struct {
			Clock_gen float64
			State     M5state
		}
	}
	Cmd_ds struct {
		Name      *int8
		Equal     int8
		pad_cgo_0 [7]byte
		Argv      [100]*int8
	}
	Das struct {
		Ds_mnem          [3]byte
		pad_cgo_0        [5]byte
		Ifp              [2]Ifp
		Voltage_p5V_ifp1 float32
		Voltage_p5V_ifp2 float32
		Voltage_m5d2V    float32
		Voltage_p9V      float32
		Voltage_m9V      float32
		Voltage_p15V     float32
		Voltage_m15V     float32
		pad_cgo_1        [4]byte
	}
	Data_check_mon struct {
		Missing struct {
			Missing int64
			State   M5state
		}
		Mode struct {
			Mode      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0 [4]byte
		Submode   struct {
			Submode   [33]byte
			pad_cgo_0 [7]byte
			First     int64
			State     M5state
		}
		Time struct {
			Time  M5time
			Bad   int64
			State M5state
		}
		Offset struct {
			Offset int64
			Size   int64
			State  M5state
		}
		Period struct {
			Period M5time
			State  M5state
		}
		Bytes struct {
			Bytes int64
			State M5state
		}
		Source struct {
			Source    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_1 [4]byte
		Start     struct {
			Start M5time
			State M5state
		}
		Code struct {
			Code  int32
			State M5state
		}
		Frames struct {
			Frames int32
			State  M5state
		}
		Header struct {
			Header M5time
			State  M5state
		}
		Total struct {
			Total float32
			State M5state
		}
		pad_cgo_2 [4]byte
		Byte      struct {
			Byte  int64
			State M5state
		}
	}
	Data_valid_cmd struct {
		User_dv   int32
		Pb_enable int32
	}
	Dbbc_cont_cal_cmd struct {
		Mode    int32
		Samples int32
	}
	Dbbcform_cmd struct {
		Mode int32
		Test int32
	}
	Dbbcgain_cmd struct {
		Bbc    int32
		State  int32
		GainU  int32
		GainL  int32
		Target int32
	}
	Dbbcgain_mon struct {
		State  int32
		Target int32
	}
	Dbbcifx_cmd struct {
		Input       int32
		Att         int32
		Agc         int32
		Filter      int32
		Target_null int32
		Target      uint32
	}
	Dbbcifx_mon struct {
		Tp uint32
	}
	Dbbcnn_cmd struct {
		Freq      uint64
		Source    int32
		Bw        int32
		Avper     int32
		pad_cgo_0 [4]byte
	}
	Dbbcnn_mon struct {
		Agc   int32
		Gain  [2]int32
		Tpon  [2]uint32
		Tpoff [2]uint32
	}
	Digout struct {
		Setting  uint32
		Mode     uint32
		Tristate uint32
	}
	Disk2file_cmd struct {
		Scan_label struct {
			Scan_label [65]byte
			pad_cgo_0  [3]byte
			State      M5state
		}
		Destination struct {
			Destination [129]byte
			pad_cgo_0   [3]byte
			State       M5state
		}
		Start struct {
			Start     [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		End struct {
			End       [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Options struct {
			Options   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Disk2file_mon struct {
		Option struct {
			Option    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0  [4]byte
		Start_byte struct {
			Start_byte int64
			State      M5state
		}
		End_byte struct {
			End_byte int64
			State    M5state
		}
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_1 [4]byte
		Current   struct {
			Current int64
			State   M5state
		}
		Scan_number struct {
			Scan_number int32
			State       M5state
		}
		pad_cgo_2 [4]byte
	}
	Disk_pos_mon struct {
		Record struct {
			Record int64
			State  M5state
		}
		Play struct {
			Play  int64
			State M5state
		}
		Stop struct {
			Stop  int64
			State M5state
		}
	}
	Disk_record_cmd struct {
		Record struct {
			Record int32
			State  M5state
		}
		Label struct {
			Label     [65]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Disk_record_mon struct {
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0 [4]byte
		Scan      struct {
			Scan  int64
			State M5state
		}
	}
	Disk_serial_mon struct {
		Count  int32
		Serial [16]struct {
			Serial    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Dist_cmd struct {
		Atten [2]int32
		Input [2]int32
		Avper int32
		Old   [2]int32
	}
	Dist_mon struct {
		Serial int32
		Timing int32
		Totpwr [2]uint32
	}
	Dot_mon struct {
		Time struct {
			Time      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		FHG_status struct {
			FHG_status [33]byte
			pad_cgo_0  [3]byte
			State      M5state
		}
		OS_time struct {
			OS_time   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		DOT_OS_time_diff struct {
			DOT_OS_time_diff [33]byte
			pad_cgo_0        [3]byte
			State            M5state
		}
	}
	Dqa_chan struct {
		Bbc      int32
		Track    int32
		Amp      float32
		Phase    float32
		Parity   uint64
		Crcc_a   uint64
		Crcc_b   uint64
		Resync   uint64
		Nosync   uint64
		Num_bits uint64
	}
	Dqa_cmd struct {
		Dur int32
	}
	Dqa_mon struct {
		A Dqa_chan
		B Dqa_chan
	}
	Ds_cmd struct {
		Type      uint16
		Mnem      [3]byte
		pad_cgo_0 [1]byte
		Cmd       uint16
		Data      uint16
	}
	Ds_mon struct {
		Resp uint16
		Data [2]byte
	}
	Flux_ds struct {
		Name      [11]byte
		Type      int8
		Fmin      float32
		Fmax      float32
		Fcoeff    [3]float32
		Size      float32
		Model     int8
		pad_cgo_0 [3]byte
		Mcoeff    [6]float32
	}
	Form4_cmd struct {
		Mode      int32
		Rate      int32
		Enable    [2]uint64
		Aux       [2]uint32
		Codes     [64]int32
		Bits      int32
		Fan       int32
		Barrel    int32
		Modulate  int32
		Last      int32
		Synch     int32
		Roll      [16][64]int32
		Start_map int32
		End_map   int32
		A2d       [16][64]int32
	}
	Form4_mon struct {
		Status   int32
		Error    int32
		Rack_ids int32
		Version  int32
	}
	Fscom struct {
		Iclbox      int64
		Iclopr      int64
		Nums        [40]int64
		AZOFF       float32
		DECOFF      float32
		ELOFF       float32
		Ibmat       int32
		Ibmcb       int32
		ICAPTP      [2]int32
		IRDYTP      [2]int32
		IRENVC      int32
		ILOKVC      int32
		ITRAKA      [2]int32
		ITRAKB      [2]int32
		pad_cgo_0   [4]byte
		TPIVC       [15]uint64
		ISTPTP      [2]float32
		ITACTP      [2]float32
		KHALT       int32
		KECHO       int32
		KENASTK     [2][2]int32
		INEXT       [3]int32
		RAOFF       float32
		XOFF        float32
		YOFF        float32
		LLOG        [8]byte
		LNEWPR      [8]byte
		LNEWSK      [8]byte
		LPRC        [8]byte
		LSTP        [8]byte
		LSKD        [8]byte
		LEXPER      [8]byte
		LFEET_FS    [2][6]byte
		Lgen        [2][2]int16
		ICHK        [23]int32
		Tempwx      float32
		Humiwx      float32
		Preswx      float32
		Speedwx     float32
		Directionwx int32
		Ep1950      float32
		Epoch       float32
		Cablev      float32
		Height      float32
		pad_cgo_1   [4]byte
		Ra50        float64
		Dec50       float64
		Radat       float64
		Decdat      float64
		Alat        float64
		Wlong       float64
		Systmp      [36]float32
		Ldsign      int32
		Lfreqv      [90]byte
		Lnaant      [8]byte
		Lsorna      [10]byte
		Idevant     [64]byte
		Idevgpib    [64]byte
		Idevlog     [64][5]byte
		Ndevlog     int64
		Imodfm      int32
		Ipashd      [2][2]int32
		Iratfm      int32
		Ispeed      [2]int32
		Idirtp      [2]int32
		Cips        [2]int64
		Bit_density [2]int32
		Ienatp      [2]int32
		Inp1if      int32
		Inp2if      int32
		Ionsor      int32
		Imaxtpsd    [2]int32
		Iskdtpsd    [2]int32
		Motorv      [2]float32
		Inscint     [2]float32
		Inscsl      [2]float32
		Outscint    [2]float32
		Outscsl     [2]float32
		Itpthick    [2]int32
		Wrvolt      [2]float32
		Capstan     [2]int32
		Go          struct {
			Allocated int32
			Name      [32][5]byte
		}
		Sem struct {
			Allocated int32
			Name      [32][5]byte
		}
		pad_cgo_2 [4]byte
		Check     struct {
			Bbc       [16]int32
			Bbc_time  [16]int64
			Dist      [2]int32
			Vform     int32
			pad_cgo_0 [4]byte
			Fm_cn_tm  int64
			Rec       [2]int32
			Vkrepro   [2]int32
			Vkenable  [2]int32
			Vkmove    [2]int32
			Systracks [2]int32
			Rc_mv_tm  [2]int64
			Vklowtape [2]int32
			Vkload    [2]int32
			Rc_ld_tm  [2]int64
			S2rec     S2rec_check
			K4rec     K4rec_check
			Ifp       [4]int32
			pad_cgo_1 [4]byte
			Ifp_time  [4]int64
		}
		Stcnm   [4][2]byte
		Stchk   [4]int32
		Dist    [2]Dist_cmd
		Bbc     [16]Bbc_cmd
		Tpi     [36]int64
		Tpical  [36]int64
		Tpizero [36]int64
		Equip   struct {
			Rack         int32
			Drive        [2]int32
			Drive_type   [2]int32
			Rack_type    int32
			Wx_met       int32
			Wx_host      [65]byte
			pad_cgo_0    [3]byte
			Mk4sync_dflt int32
		}
		Klvdt_fs     [2]int32
		Vrepro       [2]Vrepro_cmd
		Vform        Vform_cmd
		Venable      [2]Venable_cmd
		Systracks    [2]Systracks_cmd
		Dqa          Dqa_cmd
		User_info    User_info_cmd
		S2st         S2st_cmd
		S2_rec_state int32
		Rec_mode     Rec_mode_cmd
		Data_valid   [2]Data_valid_cmd
		S2label      S2label_cmd
		pad_cgo_3    [7]byte
		Form4        Form4_cmd
		Diaman       float32
		Slew1        float32
		Slew2        float32
		Lolim1       float32
		Lolim2       float32
		Uplim1       float32
		Uplim2       float32
		Refreq       float32
		I70kch       int32
		I20kch       int32
		Time         struct {
			Rate       [2]float32
			Offset     [2]int64
			Epoch      [2]int64
			Span       [2]int64
			Secs_off   int64
			Index      int32
			Icomputer  [2]int32
			Model      int8
			pad_cgo_0  [3]byte
			Ticks_off  uint64
			Usecs_off  int64
			Init_error int32
			Init_errno int32
		}
		Posnhd       [2][2]float32
		Class_count  int32
		Horaz        [30]float32
		Horel        [30]float32
		Mcb_dev      [64]byte
		Hwid         uint8
		pad_cgo_4    [3]byte
		Iw_motion    int32
		Lowtp        [2]int32
		Form_version int32
		Sterp        int32
		Wrhd_fs      [2]int32
		Vfm_xpnt     int32
		Actual       struct {
			S2rec [2]struct {
				Rstate         int32
				Rstate_valid   int32
				Position       int64
				Posvar         int64
				Position_valid int32
				pad_cgo_0      [4]byte
			}
			S2rec_inuse int32
			pad_cgo_0   [4]byte
		}
		Freqvc                   [15]float32
		Ibwvc                    [15]int32
		Ifp2vc                   [16]int32
		Cwrap                    [8]byte
		Vacsw                    [2]int32
		Motorv2                  [2]float32
		Itpthick2                [2]int32
		Thin                     [2]int32
		Vac4                     [2]int32
		Wrvolt2                  [2]float32
		Wrvolt4                  [2]float32
		Wrvolt42                 [2]float32
		User_dev1_name           [2]byte
		User_dev2_name           [2]byte
		pad_cgo_5                [4]byte
		User_dev1_value          float64
		User_dev2_value          float64
		Rvac                     [2]Rvac_cmd
		Wvolt                    [2]Wvolt_cmd
		Lo                       Lo_cmd
		Pcalform                 Pcalform_cmd
		Pcald                    Pcald_cmd
		Extbwvc                  [15]float32
		pad_cgo_6                [4]byte
		Freqif3                  int64
		Imixif3                  int32
		Pcalports                Pcalports_cmd
		K4_rec_state             int32
		K4st                     K4st_cmd
		K4tape_sqn               [9]byte
		pad_cgo_7                [3]byte
		K4vclo                   K4vclo_cmd
		K4vc                     K4vc_cmd
		K4vcif                   K4vcif_cmd
		K4vcbw                   K4vcbw_cmd
		K3fm                     K3fm_cmd
		Reccpu                   [2]int32
		K4label                  K4label_cmd
		pad_cgo_8                [3]byte
		K4rec_mode               K4rec_mode_cmd
		K4recpatch               K4recpatch_cmd
		K4pcalports              K4pcalports_cmd
		Select                   int32
		Rdhd_fs                  [2]int32
		Knewtape                 [2]int32
		Ihdmndel                 [2]int32
		Scan_name                Scan_name_cmd
		Tacd                     Tacd_shm
		Iat1if                   int32
		Iat2if                   int32
		Iat3if                   int32
		Erchk                    int32
		Ifd_set                  int32
		If3_set                  int32
		Bbc_tpi                  [16][2]uint64
		Vifd_tpi                 [4]uint64
		Mifd_tpi                 [3]uint64
		Cablevl                  float32
		Cablediff                float32
		Imk4fmv                  int32
		Tpicd                    Tpicd_cmd
		ITPIVC                   [15]int32
		Tpigain                  [36]int64
		Iapdflg                  int32
		K4rec_mode_stat          int32
		Onoff                    Onoff_cmd
		Rxgain                   [20]Rxgain_ds
		Iswif3_fs                [4]int32
		Ipcalif3                 int32
		Flux                     [100]Flux_ds
		pad_cgo_9                [4]byte
		Tpidiff                  [36]int64
		Tpidiffgain              [36]int64
		Caltemps                 [36]float32
		Calrx                    Calrx_cmd
		Ibds                     int32
		Ds_dev                   [64]byte
		N_das                    uint8
		Lba_image_reject_filters uint8
		pad_cgo_10               [2]byte
		Lba_digital_input_format uint32
		pad_cgo_11               [4]byte
		Das                      [2]Das
		Ifp_tpi                  [4]uint64
		M_das                    uint8
		Mk5vsn                   [33]byte
		pad_cgo_12               [6]byte
		Mk5vsn_logchg            int64
		Logchg                   int64
		User_device              User_device_cmd
		Disk_record              Disk_record_cmd
		Monit5                   struct {
			Pong      int32
			pad_cgo_0 [4]byte
			Ping      [2]Monit5_ping
		}
		Disk2file Disk2file_cmd
		In2net    In2net_cmd
		Abend     struct {
			Normal_end  int32
			Other_error int32
		}
		pad_cgo_13        [4]byte
		S2bbc             [4]S2bbc_data
		S2das             S2das_check
		Ntp_synch_unknown int32
		Last_check        struct {
			String    [256]byte
			Ip2       int32
			Who       [3]byte
			pad_cgo_0 [1]byte
		}
		Mk5host         [129]byte
		pad_cgo_14      [3]byte
		Mk5b_mode       Mk5b_mode_cmd
		Vsi4            Vsi4_cmd
		Holog           Holog_cmd
		Satellite       Satellite_cmd
		pad_cgo_15      [4]byte
		Ephem           [14400]Satellite_ephem
		Satoff          Satoff_cmd
		Tle             Tle_cmd
		Dbbcnn          [16]Dbbcnn_cmd
		Dbbcifx         [4]Dbbcifx_cmd
		Dbbcform        Dbbcform_cmd
		Dbbcddcv        int32
		Dbbcpfbv        int32
		Dbbc_cond_mods  int32
		Dbbc_cont_cal   Dbbc_cont_cal_cmd
		Dbbc_if_factors [4]int32
		Dbbcgain        Dbbcgain_cmd
		M5b_crate       int32
		Dbbcddcvl       [1]byte
		Dbbcddcvs       [16]byte
		pad_cgo_16      [3]byte
		Dbbcddcvc       int32
		Mk6_units       [2]int32
		Mk6_active      [2]int32
		Mk6_record      [3]Mk6_record_cmd
		Mk6_last_check  [2]struct {
			String    [256]byte
			Ip2       int32
			Who       [3]byte
			What      [3]byte
			pad_cgo_0 [2]byte
		}
		Rdbe_units     [4]int32
		Rdbe_active    [4]int32
		Rdbe_tsys_data [4]Rdbe_tsys_data
		Rdbehost       [4][129]byte
		Rdbe_atten     [5]Rdbe_atten_cmd
		Rdtcn          [4]Rdtcn
		Fserr_cls      Fserr_cls
		Dbbc_defined   int32
		Dbbc2_defined  int32
		Rdbe_equip     struct {
			Rms_t     float32
			Rms_min   float32
			Rms_max   float32
			Pcal_amp  [1]byte
			pad_cgo_0 [3]byte
		}
		Monit6     Monit6
		pad_cgo_17 [4]byte
	}
	Fserr_cls struct {
		Buf       [125]byte
		pad_cgo_0 [3]byte
		Nchars    int32
	}
	Ft struct {
		Sync             uint32
		pad_cgo_0        [4]byte
		Nco_centre_value uint64
		Nco_offset_value uint64
		Nco_phase_value  uint64
		Nco_timer_value  uint64
		Nco_test         uint32
		Nco_use_offset   uint32
		Nco_sync_reset   uint32
		Nco_use_timer    uint32
		Q_fir_no         uint8
		I_fir_no         uint8
		Clock_decimation int8
		pad_cgo_1        [1]byte
		Add_sub          Mux
		Usb_mux          Mux
		Lsb_mux          Mux
		Usb_threshold    uint8
		Lsb_threshold    uint8
		pad_cgo_2        [2]byte
		Usb_servo        Servo
		Lsb_servo        Servo
		Flip_usb         uint32
		Flip_lsb         uint32
		Monitor          Mux
		Digout           Digout
		pad_cgo_3        [4]byte
	}
	Holog_cmd struct {
		Az           float32
		El           float32
		Azp          int32
		Elp          int32
		Ical         int32
		Proc         [33]byte
		pad_cgo_0    [3]byte
		Stop_request int32
		Setup        int32
		Wait         int32
	}
	Ifp struct {
		Frequency      float64
		Bandwidth      uint32
		Filter_mode    uint32
		Flip_usb       uint32
		Flip_lsb       uint32
		Format         uint32
		Magn_stats     uint32
		Corr_type      uint32
		Corr_source    [2]uint32
		At_clock_delay int8
		pad_cgo_0      [3]byte
		Ft_lo          float64
		Ft_filter_mode uint32
		pad_cgo_1      [4]byte
		Ft_offs        float64
		Ft_phase       float64
		Track          [2]byte
		Initialised    int8
		pad_cgo_2      [1]byte
		Source         int32
		Filter_output  uint32
		Bs             Bs
		pad_cgo_3      [4]byte
		Ft             Ft
		Out            Out
		Temp_analog    float32
		Pll_ld         float32
		Pll_vc         float32
		Ref_err        uint8
		Sync_err       uint8
		pad_cgo_4      [2]byte
		Temp_digital   float32
		Processing     uint8
		Clk_err        uint8
		Blank          uint8
		pad_cgo_5      [5]byte
	}
	In2net_cmd struct {
		Control struct {
			Control int32
			State   M5state
		}
		Destination struct {
			Destination [33]byte
			pad_cgo_0   [3]byte
			State       M5state
		}
		Options struct {
			Options   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Last_destination [33]byte
		pad_cgo_0        [3]byte
	}
	In2net_mon struct {
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0 [4]byte
		Received  struct {
			Received int64
			State    M5state
		}
		Buffered struct {
			Buffered int64
			State    M5state
		}
	}
	K3fm_cmd struct {
		Mode      int32
		Rate      int32
		Input     int32
		Aux       [12]byte
		Synch     int32
		Aux_start int32
		Output    int32
	}
	K3fm_mon struct {
		Daytime [15]byte
		Status  [3]uint8
	}
	K4label_cmd struct {
		Label [9]byte
	}
	K4pcalports_cmd struct {
		Ports [2]int32
	}
	K4pcalports_mon struct {
		Amp   [2]float32
		Phase [2]float32
	}
	K4rec_check struct {
		Check int32
		State int32
		Mode  int32
		Pca   int32
		Pcb   int32
		Drm   int32
		Synch int32
		Aux   int32
	}
	K4rec_mode_cmd struct {
		Bw int32
		Bt int32
		Ch int32
		Im int32
		Nm int32
	}
	K4rec_mode_mon struct {
		Ts int32
		Fm int32
		Ta int32
		Pb int32
	}
	K4recpatch_cmd struct {
		Ports [16]int32
	}
	K4st_cmd struct {
		Record int32
	}
	K4vcbw_cmd struct {
		Bw [2]int32
	}
	K4vc_cmd struct {
		Lohi [16]int32
		Att  [16]int32
		Loup [16]int32
	}
	K4vcif_cmd struct {
		Att [4]int32
	}
	K4vclo_cmd struct {
		Freq [16]int64
	}
	K4vclo_mon struct {
		Yes  [16]byte
		Lock [16]byte
	}
	K4vc_mon struct {
		Yes    [16]byte
		Usbpwr [16]int32
		Lsbpwr [16]int32
	}
	Lo_cmd struct {
		Lo       [8]float64
		Sideband [8]int32
		Pol      [8]int32
		Spacing  [8]float64
		Offset   [8]float64
		Pcal     [8]int32
	}
	M5state struct {
		Known int32
		Error int32
	}
	M5time struct {
		Year              int32
		Day               int32
		Hour              int32
		Minute            int32
		Seconds           float64
		Seconds_precision int32
		pad_cgo_0         [4]byte
	}
	Mcb_cmd struct {
		Device    [2]byte
		pad_cgo_0 [2]byte
		Addr      uint32
		Data      uint32
		Cmd       int32
	}
	Mcb_mon struct {
		Data uint32
	}
	Mk5b_mode_cmd struct {
		Source struct {
			Source int32
			State  M5state
		}
		pad_cgo_0 [4]byte
		Mask      struct {
			Mask  uint64
			State M5state
		}
		Decimate struct {
			Decimate int32
			State    M5state
		}
		Fpdp struct {
			Fpdp  int32
			State M5state
		}
		Disk struct {
			Disk  int32
			State M5state
		}
		pad_cgo_1 [4]byte
	}
	Mk6_disk_pos_mon struct {
		Record struct {
			Record int64
			State  M5state
		}
		Play struct {
			Play  int64
			State M5state
		}
		Stop struct {
			Stop  int64
			State M5state
		}
	}
	Mk6_record_cmd struct {
		Action struct {
			Action    [22]byte
			pad_cgo_0 [2]byte
			State     M5state
		}
		Duration struct {
			Duration int32
			State    M5state
		}
		Size struct {
			Size  int32
			State M5state
		}
		Scan struct {
			Scan      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Experiment struct {
			Experiment [9]byte
			pad_cgo_0  [3]byte
			State      M5state
		}
		Station struct {
			Station   [9]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Mk6_record_mon struct {
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Group struct {
			Group int32
			State M5state
		}
		Number struct {
			Number int32
			State  M5state
		}
		Name struct {
			Name      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Mk6_scan_check_mon struct {
		Scan struct {
			Scan  int64
			State M5state
		}
		Label struct {
			Label     [65]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Type struct {
			Type      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Code struct {
			Code  int32
			State M5state
		}
		pad_cgo_0 [4]byte
		Start     struct {
			Start M5time
			State M5state
		}
		Length struct {
			Length M5time
			State  M5state
		}
		Total struct {
			Total float32
			State M5state
		}
		pad_cgo_1 [4]byte
		Missing   struct {
			Missing int64
			State   M5state
		}
		Error struct {
			Error     [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_2 [4]byte
	}
	Monit5_ping struct {
		Active    int32
		pad_cgo_0 [4]byte
		Bank      [2]struct {
			Vsn       [33]byte
			pad_cgo_0 [7]byte
			Seconds   float64
			Gb        float64
			Percent   float64
			Itime     [6]int32
		}
	}
	Monit6 struct {
		Tsys [2][4]int32
		Pcal [2][4]int32
	}
	Mux struct {
		Setting   uint8
		pad_cgo_0 [3]byte
		Mode      uint32
	}
	Onoff_cmd struct {
		Rep          int32
		Intp         int32
		Cutoff       float32
		Step         float32
		Wait         int32
		Ssize        float32
		Proc         [33]byte
		pad_cgo_0    [7]byte
		Devices      [134]Onoff_devices
		Itpis        [134]int32
		Fwhm         float32
		Stop_request int32
		Setup        int32
		pad_cgo_1    [4]byte
	}
	Onoff_devices struct {
		Lwhat     [4]byte
		Pol       int8
		pad_cgo_0 [3]byte
		Ifchain   int32
		Flux      float32
		Corr      float32
		pad_cgo_1 [4]byte
		Center    float64
		Fwhm      float32
		Tcal      float32
		Dpfu      float32
		Gain      float32
	}
	Out struct {
		S2_lo            S2_out
		S2_hi            S2_out
		Atmb_corr_source uint32
		Mb_corr_2_source uint32
		At_clock_delay   uint8
		pad_cgo_0        [3]byte
	}
	Pcald_cmd struct {
		Continuous   int32
		Bits         int32
		Integration  int32
		Stop_request int32
		Count        [2][16]int32
		Freqs        [2][16][17]float64
	}
	Pcalform_cmd struct {
		Count  [2][16]int32
		Which  [2][16][17]int32
		Tones  [2][16][17]int32
		Strlen [2][16][17]int32
		Freqs  [2][16][17]float64
	}
	Pcalports_cmd struct {
		Bbc [2]int32
	}
	Pps_source_cmd struct {
		Source struct {
			Source    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Rclcn_req_buf struct {
		Count       int32
		pad_cgo_0   [4]byte
		Class_fs    int64
		Nchars      int32
		Prev_nchars int32
		Buf         [512]uint8
	}
	Rclcn_res_buf struct {
		Class_fs int32
		Count    int32
		Ifc      int32
		Nchars   int32
		Buf      [512]uint8
	}
	Rdbe_atten_cmd struct {
		Ifc struct {
			Ifc   int32
			State M5state
		}
		Atten struct {
			Atten int32
			State M5state
		}
		Target struct {
			Target float32
			State  M5state
		}
	}
	Rdbe_atten_mon struct {
		Ifc [2]struct {
			Ifc struct {
				Ifc   int32
				State M5state
			}
			Atten struct {
				Atten int32
				State M5state
			}
			RMS struct {
				RMS   float32
				State M5state
			}
		}
	}
	Rdbe_dot_mon struct {
		Time struct {
			Time      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Status struct {
			Status    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		OS_time struct {
			OS_time   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		DOT_OS_time_diff struct {
			DOT_OS_time_diff [33]byte
			pad_cgo_0        [3]byte
			State            M5state
		}
		Actual_DOT_time struct {
			Actual_DOT_time [33]byte
			pad_cgo_0       [3]byte
			State           M5state
		}
		Vdif_epoch struct {
			Vdif_epoch int32
			State      M5state
		}
	}
	Rdbe_tsys_cycle struct {
		Epoch      [14]byte
		pad_cgo_0  [2]byte
		Epoch_vdif int32
		Tsys       [18][2]float32
		Pcal_amp   [1024]float32
		Pcal_phase [1024]float32
		Pcal_ifx   int32
		Sigma      float32
		Raw_ifx    int32
		Dot2gps    float64
		Pcaloff    float64
	}
	Rdbe_tsys_data struct {
		Data      [2]Rdbe_tsys_cycle
		Iping     int32
		pad_cgo_0 [4]byte
	}
	Rdtcn struct {
		Control [2]Rdtcn_control
		Iping   int32
	}
	Rdtcn_control struct {
		Continuous   int32
		Cycle        int32
		Stop_request int32
		Data_valid   Data_valid_cmd
	}
	Rec_mode_cmd struct {
		Mode       [21]byte
		pad_cgo_0  [3]byte
		Group      int32
		Roll       int32
		Num_groups int32
	}
	Regs struct {
		Error   uint8
		Warning uint8
	}
	Req_buf struct {
		Count     int32
		pad_cgo_0 [4]byte
		Class_fs  int64
		Nchars    int32
		Buf       [512]uint8
		pad_cgo_1 [4]byte
	}
	Req_rec struct {
		Type      int32
		Device    [2]byte
		pad_cgo_0 [2]byte
		Addr      uint32
		Data      uint32
	}
	Res_buf struct {
		Class_fs int32
		Count    int32
		Ifc      int32
		Nchars   int32
		Buf      [512]uint8
	}
	Res_rec struct {
		State int32
		Code  int32
		Data  uint32
		Array [24]uint8
	}
	Rtime_mon struct {
		Seconds struct {
			Seconds float64
			State   M5state
		}
		Gb struct {
			Gb    float64
			State M5state
		}
		Percent struct {
			Percent float64
			State   M5state
		}
		Total_rate struct {
			Total_rate float64
			State      M5state
		}
		Mode struct {
			Mode      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Sub_mode struct {
			Sub_mode  [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Track_rate struct {
			Track_rate float64
			State      M5state
		}
		Source struct {
			Source    [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0 [4]byte
		Mask      struct {
			Mask  uint64
			State M5state
		}
		Decimate struct {
			Decimate int32
			State    M5state
		}
		pad_cgo_1 [4]byte
	}
	Rvac_cmd struct {
		Inches float32
		Set    int32
	}
	Rvac_mon struct {
		Volts float32
	}
	Rxgain_ds struct {
		Type      int8
		pad_cgo_0 [3]byte
		Lo        [2]float32
		Year      int32
		Month     int32
		Day       int32
		Fwhm      struct {
			Model     int8
			pad_cgo_0 [3]byte
			Coeff     float32
		}
		Pol       [2]byte
		pad_cgo_1 [2]byte
		Dpfu      [2]float32
		Gain      struct {
			Form      int8
			Type      int8
			pad_cgo_0 [2]byte
			Coeff     [10]float32
			Ncoeff    int32
			Opacity   int8
			pad_cgo_1 [3]byte
		}
		Tcal_ntable int32
		Tcal_npol   [2]int32
		Tcal        [600]struct {
			Pol       int8
			pad_cgo_0 [3]byte
			Freq      float32
			Tcal      float32
		}
		Trec         [2]float32
		Spill_ntable int32
		Spill        [20]struct {
			El float32
			Tk float32
		}
	}
	S2bbc_data struct {
		Freq      uint64
		Tpiavg    uint16
		Ifsrc     int8
		Bw        [2]byte
		Agcmode   int8
		Init      int8
		pad_cgo_0 [1]byte
	}
	S2das_check struct {
		Check     uint64
		Agc       int8
		Encode    int8
		Mode      [21]byte
		FSstatus  int8
		SeqName   [25]byte
		BW        int8
		pad_cgo_0 [6]byte
	}
	S2label_cmd struct {
		Tapeid   [21]byte
		Tapetype [7]byte
		Format   [33]byte
	}
	S2_out struct {
		Source uint32
		Format uint32
	}
	S2rec_check struct {
		Check     int32
		User_info struct {
			Label [4]int32
			Field [4]int32
		}
		Speed    int32
		State    int32
		Mode     int32
		Group    int32
		Roll     int32
		Dv       int32
		Tapeid   int32
		Tapetype int32
	}
	S2st_cmd struct {
		Dir    int32
		Speed  int32
		Record int32
	}
	Satellite_cmd struct {
		Name      [25]byte
		Tlefile   [65]byte
		pad_cgo_0 [2]byte
		Mode      int32
		Wrap      int32
		Satellite int32
		Tle0      [25]byte
		Tle1      [70]byte
		Tle2      [70]byte
		pad_cgo_1 [3]byte
	}
	Satellite_ephem struct {
		T  int64
		Az float64
		El float64
	}
	Satoff_cmd struct {
		Seconds   float64
		Cross     float64
		Hold      int32
		pad_cgo_0 [4]byte
	}
	Scan_check_mon struct {
		Scan struct {
			Scan  int64
			State M5state
		}
		Label struct {
			Label     [65]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_0 [4]byte
		Start     struct {
			Start M5time
			State M5state
		}
		Length struct {
			Length M5time
			State  M5state
		}
		Missing struct {
			Missing int64
			State   M5state
		}
		Mode struct {
			Mode      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Submode struct {
			Submode   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Rate struct {
			Rate  float32
			State M5state
		}
		Type struct {
			Type      [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Code struct {
			Code  int32
			State M5state
		}
		Total struct {
			Total float32
			State M5state
		}
		Error struct {
			Error     [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		pad_cgo_1 [4]byte
	}
	Scan_name_cmd struct {
		Name       [17]byte
		Session    [17]byte
		Station    [17]byte
		pad_cgo_0  [5]byte
		Duration   int64
		Continuous int64
	}
	Servo struct {
		Setting   uint16
		pad_cgo_0 [2]byte
		Mode      uint32
		Readout   int32
	}
	Systracks_cmd struct {
		Track [4]int32
	}
	Tacd_shm struct {
		Day                 int32
		Day_frac            int32
		Msec_counter        float32
		Usec_correction     float32
		Usec_bias           float32
		Cooked_correction   float32
		Pc_v_utc            float32
		Utc_correction_nsec float32
		Utc_correction_sec  int32
		Day_a               int32
		Day_frac_a          int32
		Rms                 float32
		Usec_average        float32
		Max                 float32
		Min                 float32
		Day_frac_old        int32
		Day_frac_old_a      int32
		Continuous          int32
		Nsec_accuracy       int32
		Sec_average         int32
		Stop_request        int32
		Port                int32
		Check               int32
		Display             int32
		Hostpc              [80]byte
		Oldnew              [8]byte
		Oldnew_a            [11]byte
		File                [40]byte
		Status              [8]byte
		Tac_ver             [20]byte
		pad_cgo_0           [1]byte
	}
	Tape_cmd struct {
		Set   int32
		Reset int32
	}
	Tape_mon struct {
		Foot    int32
		Sense   int32
		Vacuum  int32
		Chassis int32
		Stat    int32
		Error   int32
	}
	Tle_cmd struct {
		Tle0      [25]byte
		Tle1      [70]byte
		Tle2      [70]byte
		pad_cgo_0 [3]byte
		Catnum    [3]int64
	}
	Tpicd_cmd struct {
		Continuous   int32
		Cycle        int32
		Stop_request int32
		Itpis        [36]int32
		Ifc          [36]int32
		Lwhat        [36][2]byte
		Tsys_request int32
	}
	User_device_cmd struct {
		Lo       [6]float64
		Sideband [6]int32
		Pol      [6]int32
		Center   [6]float64
	}
	User_info_cmd struct {
		Labels [4][17]byte
		Field1 [17]byte
		Field2 [17]byte
		Field3 [33]byte
		Field4 [49]byte
	}
	User_info_parse struct {
		Field     int32
		Label     int32
		String    [49]byte
		pad_cgo_0 [3]byte
	}
	Venable_cmd struct {
		General int32
		Group   [8]int32
	}
	Vform_cmd struct {
		Mode   int32
		Rate   int32
		Format int32
		Enable struct {
			Low    uint32
			High   uint32
			System uint32
		}
		Aux        [28][4]uint32
		Codes      [32]int32
		Fan        int32
		Barrel     int32
		Tape_clock int32
		Qa         struct {
			Drive int32
			Chan  int32
		}
		Last int32
	}
	Vform_mon struct {
		Version int32
		Sys_st  int32
		Mcb_st  int32
		Hdw_st  int32
		Sfw_st  int32
		Int_st  int32
	}
	Vrepro_cmd struct {
		Mode      [2]int32
		Track     [2]int32
		Head      [2]int32
		Equalizer [2]int32
		Bitsynch  int32
	}
	Vsi4 struct {
		Value int32
		Set   int32
	}
	Vsi4_cmd struct {
		Config Vsi4
		Pcalx  Vsi4
		Pcaly  Vsi4
	}
	Vsi4_mon struct {
		Version int32
	}
	Vsn_mon struct {
		Vsn struct {
			Vsn       [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Check struct {
			Check     [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
		Disk struct {
			Disk  int32
			State M5state
		}
		Original_vsn struct {
			Original_vsn [33]byte
			pad_cgo_0    [3]byte
			State        M5state
		}
		New_vsn struct {
			New_vsn   [33]byte
			pad_cgo_0 [3]byte
			State     M5state
		}
	}
	Vst_cmd struct {
		Dir   int32
		Speed int32
		Cips  uint32
		Rec   int32
	}
	Wvolt_cmd struct {
		Volts [2]float32
		Set   [2]int32
	}
)
