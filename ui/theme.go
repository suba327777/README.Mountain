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
}

func getTheme(name string) Theme {
	for _, theme := range Themes {
		if theme.Name == name {
			return theme
		}
	}
	return Themes[0]
}
