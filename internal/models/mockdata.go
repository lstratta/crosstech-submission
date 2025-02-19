package models

func SetupThreeTracksEachWithFiveSignals() []Track {
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
					SignalName: "SIG:SN169(CO) IECC PDR314 LOC R3/107",
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
		{
			TrackId: 3247,
			Source:  "Battersea Park",
			Target:  "Clapham Junction",
			SignalIds: []Signal{
				{
					SignalId:   13721,
					SignalName: "SIG:WS23(CO)WILLESDEN JCN",
					ELR:        "KDX",
					Mileage:    6.6028,
				},
				{
					SignalId:   13722,
					SignalName: "SIG:WS25(CO)WILLESDEN JCN",
					ELR:        "DKX",
					Mileage:    6.6702,
				},
				{
					SignalId:   13871,
					SignalName: "SIG:WM126(CO)WILLESDEN JCN",
					ELR:        "MDF1",
					Mileage:    6.8593,
				},
				{
					SignalId:   13873,
					SignalName: "SIG:WM129(CO)WILLESDEN JCN",
					ELR:        "FDM1",
					Mileage:    7.8602,
				},
				{
					SignalId:   13720,
					SignalName: "SIG:WS22(CO)WILLESDEN JCN",
					ELR:        "KDX",
					Mileage:    6.4412,
				},
			},
		},
		{
			TrackId: 4522,
			Source:  "Brent Cross West",
			Target:  "Acton Central",
			SignalIds: []Signal{
				{
					SignalId:   13907,
					SignalName: "SIG:WM819(CO)WILLESDEN JCN",
					ELR:        "MMH",
					Mileage:    2.2734,
				},
				{
					SignalId:   13894,
					SignalName: "SIG:WM779(CO)HARLESDEN JCN",
					ELR:        "VIM",
					Mileage:    2.5724,
				},
				{
					SignalId:   453,
					SignalName: "SIG:AW148(CO) ACTON WELLS JCN",
					ELR:        "LPC5",
					Mileage:    3.1745,
				},
				{
					SignalId:   13898,
					SignalName: "SIG:WM788(CO)BRENT NEW JCN",
					ELR:        "MHM",
					Mileage:    3.2005,
				},
				{
					SignalId:   13730,
					SignalName: "SIG:WS37(CO)STONEBRIDGE PARK",
					ELR:        "KXD",
					Mileage:    8.7886,
				},
			},
		},
	}
}

func SetupThreeTracks() []Track {
	return []Track{
		{
			TrackId: 55,
			Source:  "Acton Central",
			Target:  "Willesden Junction",
		},
		{
			TrackId: 3247,
			Source:  "Battersea Park",
			Target:  "Clapham Junction",
		},
		{
			TrackId: 4522,
			Source:  "Brent Cross West",
			Target:  "Acton Central",
		},
	}
}
