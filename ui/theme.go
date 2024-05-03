package ui

type Theme struct {
	Name          string
	TitleColor    string
	IconColor     string
	MountainColor string
	TextColor     string
	BgColor       string
	BorderColor   string
}

var Themes = []Theme{
	{
		Name:          "default",
		TitleColor:    "#FFFFFF",
		IconColor:     "#FFFFFF",
		MountainColor: "#ADFF2F",
		TextColor:     "#FFFFFF",
		BgColor:       "#141321",
		BorderColor:   "#FFFFFF",
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
