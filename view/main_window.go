package view

import (
	"fatture75/controller"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetMainWindowView(w fyne.Window) *fyne.Container {

	content := container.New(layout.NewVBoxLayout())

	antenoreFilePathTextBox := widget.NewEntry()
	fillCalcSheetTextBox := widget.NewEntry()
	generateQuoteTextBox := widget.NewEntry()

	checkApply75Discount := widget.NewCheck("", func(value bool) {})
	checkApply75Discount.Checked = true

	fillCalcSheetForm := container.New(layout.NewFormLayout(),
		widget.NewLabel("Percorso del file Antenore: "), antenoreFilePathTextBox,

		widget.NewLabel("Nome del file generato: "), fillCalcSheetTextBox,

		layout.NewSpacer(), widget.NewButton("Esegui",
			func() {
				if validateEmptyTextBox(antenoreFilePathTextBox.Text, w) &&
					validateEmptyTextBox(fillCalcSheetTextBox.Text, w) {

					err := controller.FillCalcSheet(
						antenoreFilePathTextBox.Text, fillCalcSheetTextBox.Text,
					)

					if err != nil {
						dialog.ShowError(err, w)
					} else {
						dialog.ShowInformation("Notifica", "Il foglio di calcolo è stato compilato con successo!", w)
					}

				}
			},
		),
	)

	openFileDialogButton := widget.NewButton("Scegli file",
		func() {

			dialog.ShowFileOpen(
				func(uc fyne.URIReadCloser, err error) {
					if uc != nil {
						generateQuoteTextBox.Text = uc.URI().Path()
						generateQuoteTextBox.Refresh()
					}
				}, w)
		},
	)

	generateQuoteForm := container.New(layout.NewFormLayout(),
		widget.NewLabel("Nome del file di calcolo: "), generateQuoteTextBox,
		widget.NewLabel("Oppure seleziona il file: "), openFileDialogButton,
		widget.NewLabel("Applica sconto 75: "), checkApply75Discount,
		layout.NewSpacer(), widget.NewButton("Esegui",
			func() {
				if validateEmptyTextBox(generateQuoteTextBox.Text, w) {
					err := controller.GenerateNewQuote(generateQuoteTextBox.Text, checkApply75Discount.Checked)

					if err != nil {
						dialog.ShowError(err, w)
					} else {
						dialog.ShowInformation("Notifica", "Il preventivo è stato generato con successo sulla piattaforma Fatture in Cloud", w)
					}
				}
			},
		),
	)

	fillCalcSheetForm.Hide()
	generateQuoteForm.Hide()

	fillCalcSheetButton := widget.NewButton("Riempi foglio di calcolo",
		func() {
			fillCalcSheetForm.Show()
			generateQuoteForm.Hide()
		},
	)
	generateQuoteButton := widget.NewButton("Genera preventivo",
		func() {
			fillCalcSheetForm.Hide()
			generateQuoteForm.Show()
		},
	)

	// add buttons
	content.Add(container.New(layout.NewHBoxLayout(), fillCalcSheetButton, generateQuoteButton))

	content.Add(fillCalcSheetForm)

	content.Add(generateQuoteForm)

	return content

}

func validateEmptyTextBox(text string, parent fyne.Window) bool {
	cond := text == ""
	if cond {
		dialog.ShowError(
			fmt.Errorf("Tutte le caselle di testo vanno compilate!"), parent,
		)
	}
	return !cond
}
