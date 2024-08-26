package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type MyData struct {
	X           []float32   `json:"x"`
	Y           []float32   `json:"y"`
	DeltaMatrix [][]float32 `json:"deltaMatrix"`
	ApproxValue float32     `json:"approxValue"`
	Err         error       `json:"err"`
}

var data MyData = MyData{X: []float32{}, Y: []float32{}}

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", data.index)
	mux.HandleFunc("/validate-x", data.getXValues)
	mux.HandleFunc("/validate-y", data.getYValues)
	mux.HandleFunc("/get-approx", data.getAproxValue)
	mux.HandleFunc("/process", data.processData)
	http.ListenAndServe(":8000", mux)
}

func (data *MyData) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, nil)
}

func (data *MyData) getXValues(w http.ResponseWriter, r *http.Request) {
	var dataX []float32
	tmpl := template.Must(template.ParseFiles("./static/error.html"))
	a := r.PostFormValue("data-x")
	strs := strings.Split(a, ",")
	for _, str := range strs {
		val, err := strconv.ParseFloat(str, 32)
		if err != nil {
			tmpl.Execute(w, nil)
			data.Err = errors.New("cannot process invalid format/input")
		} else {
			dataX = append(dataX, float32(val))
			data.X = dataX
		}
	}
}

func (data *MyData) getYValues(w http.ResponseWriter, r *http.Request) {
	var dataY []float32
	tmpl := template.Must(template.ParseFiles("./static/error.html"))
	a := r.PostFormValue("data-y")
	strs := strings.Split(a, ",")
	for _, str := range strs {
		val, err := strconv.ParseFloat(str, 32)
		if err != nil {
			tmpl.Execute(w, nil)
			data.Err = errors.New("cannot process invalid format/input")
		} else {
			dataY = append(dataY, float32(val))
			data.Y = dataY
		}
	}
}

func (data *MyData) getAproxValue(w http.ResponseWriter, r *http.Request) {
	x := r.PostFormValue("data-approx")
	dataApprox, err := strconv.ParseFloat(x, 32)
	if err != nil {
		tmpl := template.Must(template.ParseFiles("./static/error.html"))
		tmpl.Execute(w, nil)
		data.Err = errors.New("cannot process invalid aprox")
	} else {
		data.ApproxValue = float32(dataApprox)
	}
}

// TODO: add validation if len(x) == len(y)
func (data *MyData) processData(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("data").Parse("{{ . }}"))

	if len(data.X) != len(data.Y) {
		data.Err = errors.New("Length of X and Y does not match")
		w.Write([]byte(fmt.Sprintf("error, length of X and Y does not match")))

	}

	tmpl.Execute(w, data)
}
