// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var LIF_Setting struct {
	LIFFID          string
	LIFFRedirectURL string
}

func main() {
	LIF_Setting.LIFFID = os.Getenv("YOUR_LIFF_ID")
	LIF_Setting.LIFFRedirectURL = os.Getenv("YOUR_REDIRECT_URL")

	//Web APIs
	http.HandleFunc("/", liff)

	//provide by Heroku
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func liff(w http.ResponseWriter, r *http.Request) {
	// The user accesses the linking URL.
	tmpl := template.Must(template.ParseFiles("login.tmpl"))
	// The web server displays the login screen.
	if err := tmpl.Execute(w, LIF_Setting); err != nil {
		log.Println("Template err:", err)
	}
}
