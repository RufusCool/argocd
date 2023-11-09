package main

import (
	"net/http"
	"os"
	"strings"
	"math/rand"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func generateRandomColor() string {
	rand.Seed(time.Now().UnixNano())
	color := "#"
	for i := 0; i < 6; i++ {
		color += string("0123456789ABCDEF"[rand.Intn(16)])
	}
	return color
}

func main() {
	colors := []string{"#00FF00", "#FFA500", "#6495ED", "#FF1493"} // Adicione as cores que você deseja usar

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		color := generateRandomColor()
		index := rand.Intn(len(colors))
		if index == len(colors)-1 {
			index = 0
		} else {
			index++
		}
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
				<head>
					<style>
						body {
							display: flex;
							flex-direction: column;
							justify-content: center;
							align-items: center;
							height: 100vh;
							background-color: ` + colors[index] + `;
						}
						.circle {
							height: 100px;
							width: 100px;
							background-color: ` + color + `;
							border-radius: 50%;
							display: inline-block;
							margin: 20px;
						}
						h1 {
							text-align: center;
							color: #333;
						}
						p {
							text-align: center;
							color: #333;
						}
					</style>
				</head>
				<body>
					<h1>Workshop DevOps !!!!</h1>
					<p style="margin-top: 10px;">Nome do Pod: ` + hostname + `</p>
					<div class="circle"></div>
					<div class="circle"></div>
					<div class="circle"></div>
				</body>
			</html>
		`))
	})
	// add promhttp.Handler()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

