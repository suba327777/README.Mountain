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
		Name:                  "onedark",
		TitleColor:            "#E4BF7A",
		IconColor:             "#8EB573",
		MountainIconColor:     "#ADFF2F",
		TriangleMountainColor: "#ADFF2F",
		TextColor:             "#DF6D74",
		BgColor:               "#282C34",
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
	{
		Name:                  "solarized",
		TitleColor:            "#268BD2",
		IconColor:             "#B58900",
		MountainIconColor:     "#859900",
		TriangleMountainColor: "#859900",
		TextColor:             "#586E75",
		BgColor:               "#FDF6E3",
		BorderColor:           "#586E75",
	},
	{
		Name:                  "solarized_dark",
		TitleColor:            "#268BD2",
		IconColor:             "#B58900",
		MountainIconColor:     "#859900",
		TriangleMountainColor: "#859900",
		TextColor:             "#839496",
		BgColor:               "#073642",
		BorderColor:           "#839496",
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
