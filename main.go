package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func GetValidValue(pl []string) []float64 {
	var flis = make([]float64, 0, 4)
	for i, v := range pl{
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil
		}
		if f < -180 || f > 180 {
			return nil
		}
		if f < -90 || f > 90 {
			if i == 1 || i == 3 {
				return nil
			}
		}
		flis = append(flis, f)
	}
	return flis
}

func main() {
	a := app.New()
	w := a.NewWindow("Calc Distance")
	res := widget.NewEntry()
	res.Disable()

	lngA := widget.NewEntry()
	latA := widget.NewEntry()
	lngB := widget.NewEntry()
	latB := widget.NewEntry()

	form := widget.NewForm(
		&widget.FormItem{Text: "Result:", Widget: res},
		&widget.FormItem{Text: "longitude A", Widget: lngA},
		&widget.FormItem{Text: "latitude A", Widget: latA},
		&widget.FormItem{Text: "longitude B", Widget: lngB},
		&widget.FormItem{Text: "latitude B", Widget: latB},
		)

	form.OnSubmit = func() {
		validValue := GetValidValue([]string{lngA.Text, latA.Text, lngB.Text, latB.Text})
		if validValue == nil {
			res.SetText("wrong value!")
			return
		}
		d := Distance(P{validValue[0], validValue[1]}, P{validValue[2], validValue[3]})
		res.SetText(strconv.FormatFloat(d, 'f', 6, 64))
	}
	w.SetContent(form)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
