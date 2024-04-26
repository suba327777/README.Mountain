package ui

func GenerateCard() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100">
            <rect x="10" y="10" width="180" height="80" fill="lightblue" stroke="black" stroke-width="2"/>
            <text x="100" y="50" text-anchor="middle" alignment-baseline="middle" fill="black">Hello, SVG!</text>
        </svg>`
}
