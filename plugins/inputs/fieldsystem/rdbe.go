package fieldsystem

func watchrdbe(rdbeindex int, fscom *fs.Fscom, pointch chan *client.Point) {
	rdbeid := string('a' + rdbeindex)

	pcalfields := map[string]interface{}{}
	tsysfields := map[string]interface{}{}

	tags := map[string]string{
		"antenna": ANTENNA,
		"rdbe":    rdbeid,
		"if":      "0",
	}

	for {
		iping := fscom.Rdbe_tsys_data[rdbeindex].Iping
		rdbedata := fscom.Rdbe_tsys_data[rdbeindex].Data[iping]

		// Tsys data. Both IFs are updated every seconds
		for i := 0; i < len(rdbedata.Tsys[0]); i++ {
			tags["if"] = fmt.Sprintf("%d", i)
			for j := 0; j < len(rdbedata.Tsys); j++ {
				if rdbedata.Tsys[j][i] < -8e6 ||
					IsNan32(rdbedata.Tsys[j][i]) ||
					IsInf32(rdbedata.Tsys[j][i]) {
					delete(tsysfields, fmt.Sprintf("chan%d", j))
					// tsysfields[fmt.Sprintf("chan%d", j)] = -1
					continue
				}
				tsysfields[fmt.Sprintf("chan%d", j)] = rdbedata.Tsys[j][i]
			}
			pt, err := client.NewPoint(TSYS_MEASURMENT, tags, tsysfields, time.Now())
			if err != nil {
				log.Println(err)
				// probably all tsys fields overflowed, so ignore
				continue
			}
			pointch <- pt
		}

		// Pcal data alternates between IF every second
		tags["if"] = fmt.Sprintf("%d", rdbedata.Pcal_ifx)

		for i, amp := range rdbedata.Pcal_amp {
			if amp < -8e6 || IsNan32(amp) || IsInf32(amp) {
				delete(pcalfields, fmt.Sprintf("%d", i))
				continue
			}
			pcalfields[fmt.Sprintf("%d", i)] = amp
		}
		pt, err := client.NewPoint(PCALAMP_MEASURMENT, tags, pcalfields, time.Now())
		if err != nil {
			panic(err)
		}
		pointch <- pt

		for i, phase := range rdbedata.Pcal_phase {
			if phase < -8e6 ||
				IsNan32(phase) ||
				IsInf32(phase) {
				delete(pcalfields, fmt.Sprintf("chan%d", i))
				continue
			}
			pcalfields[fmt.Sprintf("%d", i)] = phase
		}
		pt, err = client.NewPoint(PCALPHASE_MEASURMENT, tags, pcalfields, time.Now())
		if err != nil {
			panic(err)
		}
		pointch <- pt

		for iping == fscom.Rdbe_tsys_data[rdbeindex].Iping {
			time.Sleep(POLL_TIME)
		}
	}

}
