package main

import (
	"github.com/marcusolsson/tui-go"
	"github.com/robbmue/GoLoytec/client"
)

var logo = `
           _                 _            
          | |               | |           
__ _  ___ | |     ___  _   _| |_ ___  ___ 
/ _  |/ _ \| |    / _ \| | | | __/ _ \/ __|
| (_| | (_) | |___| (_) | |_| | ||  __/ (__ 
\__, |\___/|______\___/ \__, |\__\___|\___|
__/ |                   __/ |             
|___/                   |___/
`

func textUserInterface(clientInstance *client.Client) {
	namebox := tui.NewLabel("loytec controller")
	namebox.SetSizePolicy(tui.Expanding, tui.Maximum)

	logoBox := tui.NewLabel(logo)

	logoBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	header := tui.NewVBox(
		tui.NewHBox(
			namebox,
			tui.NewLabel("Press ESC to exit"),
		),
		logoBox,
	)

	discoButton := tui.NewButton("[toggle]")
	discoButton.OnActivated(
		func(btn *tui.Button) {
			go discoMode(clientInstance)
		},
	)

	lightOnButton := tui.NewButton("[On]")
	lightOnButton.OnActivated(
		func(btn *tui.Button) {
			clientInstance.Light(1, 100)
		},
	)

	lightOffButton := tui.NewButton("[Off]")
	lightOffButton.OnActivated(
		func(btn *tui.Button) {
			clientInstance.Light(0, 0)
		},
	)

	blendTopButton := tui.NewButton("[Top]")
	blendTopButton.OnActivated(
		func(btn *tui.Button) {
			clientInstance.Sunblind(client.Top)
		},
	)

	blendBottomButton := tui.NewButton("[Bottom]")
	blendBottomButton.OnActivated(
		func(btn *tui.Button) {
			clientInstance.Sunblind(client.Bottom)
		},
	)

	body := tui.NewVBox(
		tui.NewHBox(
			tui.NewVBox(
				tui.NewLabel("Disco"),
				tui.NewHBox(
					discoButton,
				),
			),
			tui.NewVBox(
				tui.NewLabel("Light"),
				tui.NewHBox(
					lightOnButton,
					lightOffButton,
				),
			),
			tui.NewVBox(
				tui.NewLabel("Blends"),
				tui.NewHBox(
					blendTopButton,
					blendBottomButton,
				),
			),
		),
	)
	body.SetBorder(true)
	body.SetSizePolicy(tui.Expanding, tui.Maximum)

	box := tui.NewVBox(
		header,
		body,
	)
	box.SetBorder(true)
	box.SetSizePolicy(tui.Expanding, tui.Maximum)

	tui.DefaultFocusChain.Set(discoButton, lightOnButton, lightOffButton, blendTopButton, blendBottomButton)

	ui, err := tui.New(box)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })
	ui.SetKeybinding("c", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
