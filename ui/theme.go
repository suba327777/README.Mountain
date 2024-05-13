package ui

type Theme struct {
	Name                  string
	TitleColor            string
	IconColor             string
	MountainIconColor     string
	TriangleMountainColor string
	TextColor             string
	BgColor               string
	BorderColor           string
}

var Themes = []Theme{
	{
		Name:                  "default",
		TitleColor:            "#434d58",
		IconColor:             "#2f80ed",
		MountainIconColor:     "#ADFF2F",
		TriangleMountainColor: "#ADFF2F",
		TextColor:             "#434d58",
		BgColor:               "#FFFFeF",
		BorderColor:           "#e4e2e2",
	},
	{
		Name:                  "dark",
		TitleColor:            "#FFFFFF",
		IconColor:             "#FFFFFF",
		MountainIconColor:     "#ADFF2F",
		TriangleMountainColor: "#ADFF2F",
		TextColor:             "#FFFFFF",
		BgColor:               "#141321",
		BorderColor:           "#FFFFFF",
	},
	{
		Name:                  "sakura",
		TitleColor:            "#FFFFFF",
		IconColor:             "#FFB6A1",
		MountainIconColor:     "#AAD378",
		TriangleMountainColor: "#AAD378",
		TextColor:             "#FEFFFF",
		BgColor:               "#FDD9D9",
		BorderColor:           "#FEFFFF",
	},
	{
		Name:                  "maple",
		TitleColor:            "#D88352",
		IconColor:             "#ACBC65",
		MountainIconColor:     "#E58C43",
		TriangleMountainColor: "#FBAC33",
		TextColor:             "#D88352",
		BgColor:               "#FCC796",
		BorderColor:           "#D88352",
	},
}

func getTheme(name string) Theme {
	for _, theme := range Themes {
		if theme.Name == name {
			return theme
		}
	}
	return Themes[0]
}
