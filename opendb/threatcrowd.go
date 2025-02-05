/*

=======================
Scilla - Information Gathering Tool
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/scilla

	@Author:      edoardottt, https://www.edoardoottavianelli.it

*/

package opendb

import (
	"encoding/json"
	"net/http"
	"time"
)

//ThreatcrowdSubdomains retrieves from the below url some known subdomains.
func ThreatcrowdSubdomains(domain string) []string {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	result := make([]string, 0)
	url := "https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=" + domain
	wrapper := struct {
		Records []string `json:"subdomains"`
	}{}
	resp, err := client.Get(url)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	dec.Decode(&wrapper)
	if err != nil {
		return result
	}
	result = append(result, wrapper.Records...)
	return result
}
