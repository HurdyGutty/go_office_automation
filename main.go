package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/HurdyGutty/go_office_automation/pkg/worker"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Office automation")
	myWindow.Resize(fyne.NewSize(500, 500))

	input_entry := widget.NewEntry()
	open1 := widget.NewButton("Excel file", func() {
		input_dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				input_entry.Refresh()
				return
			}
			input_entry.SetText(strings.TrimPrefix(reader.URI().String(), "file://"))
			input_entry.Refresh()
		}, myWindow)
		input_dialog.Resize(fyne.NewSize(500, 500))
		input_dialog.Show()
	})

	URI_list := []string{}
	template_entry := widget.NewList(func() int {
		return len(URI_list)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("templates")
	},
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(URI_list[i])
		})
	open2 := widget.NewButton("Template folder", func() {
		ouput_dialog := dialog.NewFolderOpen(func(reader fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				template_entry.Refresh()
				return
			}
			list, err := reader.List()
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			URI_list = []string{}
			for _, uri := range list {
				URI_list = append(URI_list, strings.TrimPrefix(uri.String(), "file://"))
			}
			template_entry.Refresh()
		}, myWindow)
		ouput_dialog.Resize(fyne.NewSize(500, 500))
		ouput_dialog.Show()

	})

	output_entry := widget.NewEntry()

	open3 := widget.NewButton("Output folder", func() {
		ouput_dialog := dialog.NewFolderOpen(func(reader fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				output_entry.Refresh()
				return
			}
			output_entry.SetText(reader.Path())
			output_entry.Refresh()
		}, myWindow)
		ouput_dialog.Resize(fyne.NewSize(500, 500))
		ouput_dialog.Show()
	})

	status_label := widget.NewLabel("")

	run_button := widget.NewButton("Run", func() {
		if input_entry.Text == "" {
			status_label.SetText("Please select an excel file")
			status_label.Refresh()
			return
		}
		if len(URI_list) == 0 {
			status_label.SetText("Please select a template folder")
			status_label.Refresh()
			return
		}
		if output_entry.Text == "" {
			status_label.SetText("Please select an output folder")
			status_label.Refresh()
			return
		}
		worker.CreateWorker(input_entry.Text, output_entry.Text, URI_list)
		status_label.SetText("Done")
		status_label.Refresh()
	})
	grid := container.New(layout.NewGridLayout(2), input_entry, open1, template_entry, open2, output_entry, open3, run_button, status_label)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
