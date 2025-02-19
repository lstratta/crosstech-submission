package models

func SetupOneTrackWithFiveSignals() []Track {
	return []Track{
		{
			TrackId: 55,
			Source:  "Acton Central",
			Target:  "Willesden Junction",
			SignalIds: []Signal{
				{
					SignalId:   453,
					SignalName: "SIG:AW148(CO) ACTON WELLS JCN",
					ELR:        "LPC5",
					Mileage:    3.1745,
				},
				{
					SignalId:   2848,
					SignalName: "SIG:SN169(CO) IECC PDRF14 LOC R3/107",
					ELR:        "ONM1",
					Mileage:    4.2126,
				},
				{
					SignalId:   2849,
					SignalName: "SIG:SN173(CO) IECC PDMN02 LOC M3/144",
					ELR:        "MNO1",
					Mileage:    5.6889,
				},
				{
					SignalId:   13717,
					SignalName: "SIG:WS16(CO)WILLESDEN JCN",
					ELR:        "KXD",
					Mileage:    7.8087,
				},
				{
					SignalId:   13720,
					SignalName: "SIG:WS22(CO)WILLESDEN JCN",
					ELR:        "KDX",
					Mileage:    6.4412,
				},
			},
		},
	}
}

func SetupOneTrack() []Track {
	return []Track{
		{
			TrackId: 55,
			Source:  "Acton Central",
			Target:  "Willesden Junction",
		},
	}
}
