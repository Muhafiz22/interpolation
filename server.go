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
	MatrixTable [][]float32 `json:"matrixTable"`
	ApproxValue float32     `json:"approxValue"`
	Err         error       `json:"err"`
	Terms       []int       `json:"terms"`
}

var data MyData = MyData{X: []float32{}, Y: []float32{}}

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", data.index)
	mux.HandleFunc("/validate-x", data.getXValues)
	mux.HandleFunc("/validate-y", data.getYValues)
	mux.HandleFunc("/get-approx", data.getAproxValue)
	mux.HandleFunc("/get-table", data.getTable)
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
		nStr := strings.TrimSpace(str)
		val, err := strconv.ParseFloat(nStr, 32)
		if err != nil {
			tmpl.Execute(w, nil)
			data.Err = errors.New("cannot process invalid format/input")
		} else {
			data.Err = nil
			dataX = append(dataX, float32(val))
			data.X = dataX
			data.Terms = pushIter(len(data.X))
		}
	}
}

func (data *MyData) getYValues(w http.ResponseWriter, r *http.Request) {
	var dataY []float32
	tmpl := template.Must(template.ParseFiles("./static/error.html"))
	a := r.PostFormValue("data-y")
	strs := strings.Split(a, ",")
	for _, str := range strs {
		nStr := strings.TrimSpace(str)
		val, err := strconv.ParseFloat(nStr, 32)
		if err != nil {
			tmpl.Execute(w, nil)
			data.Err = errors.New("cannot process invalid format/input")
		} else {

			data.Err = nil
			dataY = append(dataY, float32(val))
			data.Y = dataY
		}
	}
}

func (data *MyData) getAproxValue(w http.ResponseWriter, r *http.Request) {
	x := r.PostFormValue("data-approx")
	xStr := strings.TrimSpace(x)
	dataApprox, err := strconv.ParseFloat(xStr, 32)
	if err != nil {
		tmpl := template.Must(template.ParseFiles("./static/error.html"))
		tmpl.Execute(w, nil)
		data.Err = errors.New("cannot process invalid aprox")
	} else {
		data.ApproxValue = float32(dataApprox)
		data.Err = nil
	}
}

func (data *MyData) getTable(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/table.html"))

	tmpl.Execute(w, data)
}

// TODO: add validation if len(x) == len(y)
// add validation for data
func (data *MyData) processData(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("data").Parse("{{ . }}"))
	yx, table := calculateBackwardDiffInterpolation(data.X, data.Y, data.ApproxValue)
	if data.Err != nil {
		w.Write([]byte("invalid data format"))
	} else {
		data.MatrixTable = table
		data.Terms = pushIter(len(data.X))
		tmpl.Execute(w, fmt.Sprintf("f(x) = %f", yx))
	}
}

func pushIter(len int) []int {
	output := make([]int, len)
	for idx := range output {
		output[idx] = idx + 1
	}
	return output
}
