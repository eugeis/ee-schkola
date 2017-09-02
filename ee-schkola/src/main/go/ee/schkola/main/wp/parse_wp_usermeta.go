package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"encoding/json"
	"regexp"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"sort"
	"strings"
)

var work = "/Users/ee/Documents/BSS-Verwaltung/BSS-2017-2018"

var de = map[string]string{
	"scool_year":         "Klasse",
	"biblikum_1":         "Biblikum 1 geschrieben?",
	"biblikum_2":         "Biblikum 2 geschrieben?",
	"paided_last_years":  "Gebühr der Vorjahren bezahlt?",
	"first_name":         "Vorname",
	"last_name":          "Nachname",
	"birth_date":         "Geburtsdatum",
	"gender":             "Geschlecht",
	"phone_number":       "Telefon",
	"church":             "Gemeinde",
	"church_commitment":  "Gemeinde einverstanden?",
	"church_member":      "Mitglied welcher Gemeinde?",
	"church_services":    "Gemeindedienste",
	"church_responsible": "Pastor / Leitungskreis",
	"church_contact":     "Telefon von Pastor / Leitungskreis",
	"user_email":         "E-Mail",
	"address":            "Adresse",
	"plz":                "PLZ",
	"city":               "Ort",
	"job":                "Beruf",
	"education":          "Bildung",
	"marital_state":      "Familienstand",
	"spirit":             "Geistlicher_Werdegang______________________________________________________Warum_Bibelschule?",
}

func main() {
	parseJson()
}

func parseJson() {
	file, e := ioutil.ReadFile(fmt.Sprintf("%v/wp_usermeta.json", work))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var data []map[string]interface{}
	e = json.Unmarshal(file, &data)
	if e != nil {
		fmt.Printf("JSON unmarshal error: %v\n", e)
		os.Exit(1)
	}

	keys := make(map[string]string)

	u := make(map[string]map[string]interface{})

	toDec, _ := regexp.Compile("^[a-z]\\:.*")

	for _, d := range data {
		user := d["user_id"].(string)
		key := d["meta_key"].(string)
		str := fmt.Sprintf("%v", d["meta_value"])

		userMap := u[user]
		if u[user] == nil {
			userMap = make(map[string]interface{})
			u[user] = userMap
		}

		if toDec.MatchString(str) {
			if dec, err := php_serialize.NewUnSerializer(str).Decode(); err == nil && dec != nil {
				fill(key, dec, userMap, keys)
			} else {
				println(dec, err)
			}
		} else {
			keys[key] = key
			userMap[key] = prepareString(key, str)
		}
	}

	users := make(map[string]map[string]interface{})
	users1 := make([]string, 0)
	users2 := make([]string, 0)
	users3 := make([]string, 0)
	users4 := make([]string, 0)
	users5 := make([]string, 0)

	for _, v := range u {
		key := fmt.Sprintf("%v.%v.%v.%v", v["scool_year"], v["last_name"], v["first_name"], v["birth_date"])
		users[key] = v
		switch v["scool_year"] {
		case "1. Klasse":
			users1 = append(users1, key)
		case "2. Klasse (B1)":
			users2 = append(users2, key)
		case "3. Klasse (B1)":
			users3 = append(users3, key)
		case "4. Klasse (B1,B2)":
			users4 = append(users4, key)
		case "Zusatzjahr (B1,B2)":
			users5 = append(users5, key)
		}
	}

	writeCsvEmailGroup(users1, users2, users3, users4, users5, users)
	writeInsertInto(users1, users2, users3, users4, users5, users)
	writeCsvEmailContacts(users1, users2, users3, users4, users5, users)
	writeHtmlReport(users1, users2, users3, users4, users5, users)
}

func writeInsertInto(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/users.sql", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	writeInsertIntoForGroup(users1, users, "2017 Klasse 1", w)
	writeInsertIntoForGroup(users2, users, "2016 1. Klasse", w)
	writeInsertIntoForGroup(users3, users, "2016 2. Klasse", w)
	writeInsertIntoForGroup(users4, users, "2016 3. Klasse", w)
	writeInsertIntoForGroup(users5, users, "2016 4. Klasse", w)
	w.Flush()
}

func writeInsertIntoForGroup(userKeys []string, users map[string]map[string]interface{}, group string, w *bufio.Writer) {
	//write data
	w.WriteString("insert into wp_users_bs (user_email, gr) values ")
	for _, userKey := range userKeys {
		user := users[userKey]
		//w.WriteString(fmt.Sprintf(`<%v %v>%v,`, user["first_name"], user["last_name"], user["user_email"]))
		w.WriteString(fmt.Sprintf("('%v', '%v'),", user["user_email"], group))
	}
	w.WriteString(";\n")
}

func writeCsvEmailGroup(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/emails.txt", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	writeCsvEmailsForGroup(users1, users, "2017 Klasse 1", w)
	writeCsvEmailsForGroup(users2, users, "2017 Klasse 2", w)
	writeCsvEmailsForGroup(users3, users, "2017 Klasse 3", w)
	writeCsvEmailsForGroup(users4, users, "2017 Klasse 4", w)
	writeCsvEmailsForGroup(users5, users, "2017 Zusatzklasse", w)
	w.Flush()
}

func writeCsvEmailsForGroup(userKeys []string, users map[string]map[string]interface{}, group string, w *bufio.Writer) {
	//write data
	w.WriteString(group)
	w.WriteString(": ")
	for _, userKey := range userKeys {
		user := users[userKey]
		//w.WriteString(fmt.Sprintf(`<%v %v>%v,`, user["first_name"], user["last_name"], user["user_email"]))
		w.WriteString(fmt.Sprintf(`%v,`, user["user_email"]))
	}
	w.WriteString("\n\n\n\n")
}

func writeCsvEmailContacts(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/googleContacts.csv", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("Last Name;First Name;Birthday;Gender;Mobile Phone:Company;E-mail Address;Categories\n")
	writeCsvEmailContactsForGroup(users1, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 1"}, w)
	writeCsvEmailContactsForGroup(users2, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 2"}, w)
	writeCsvEmailContactsForGroup(users3, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 3"}, w)
	writeCsvEmailContactsForGroup(users4, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 4"}, w)
	writeCsvEmailContactsForGroup(users5, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Zusatzklasse"}, w)
	w.Flush()
}

func writeCsvEmailContactsForGroup(userKeys []string, users map[string]map[string]interface{}, columns []string, w *bufio.Writer) {
	//write data
	for _, userKey := range userKeys {
		user := users[userKey]
		for i, k := range columns {
			cell := user[k]
			if cell == nil {
				cell = k
			}
			if i > 0 {
				w.WriteString(";")
			}
			w.WriteString(fmt.Sprintf(`%v`, cell))
		}
		w.WriteString("\n")
	}
}

func writeHtmlReport(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/bewerbungen.html", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(`<html>
		<head>
			<style>
				table {
					font-size:16px;
					font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
					border-collapse: collapse;
					border-spacing: 0;
					width: 100%;
				}

				td, th {
					border: 1px solid #ddd;
					text-align: left;
					padding: 8px;
				}

				tr:nth-child(even){background-color: #f2f2f2}

				th {
					padding-top: 11px;
					padding-bottom: 11px;
					background-color: #4CAF50;
					color: white;
				}

				.spirit {
					width: 100px;
				}

				.photo {
					height:	200px;
				}
			</style>
		</head>`)
	w.WriteString("<body>")
	w.WriteString("<h2>1. Klasse</h2>")
	writeHtmlReportForGroup(users1, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "photo", "spirit"}, w)
	w.WriteString("<h2>2. Klasse</h2>")
	writeHtmlReportForGroup(users2, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "paided_last_years", "photo", "spirit"}, w)
	w.WriteString("<h2>3. Klasse</h2>")
	writeHtmlReportForGroup(users3, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "paided_last_years", "photo", "spirit"}, w)
	w.WriteString("<h2>4. Klasse</h2>")
	writeHtmlReportForGroup(users4, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "biblikum_2", "paided_last_years", "photo", "spirit"}, w)
	w.WriteString("<h2>Zusatzjahr</h2>")
	writeHtmlReportForGroup(users5, users, []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "biblikum_2", "paided_last_years", "photo", "spirit"}, w)
	w.WriteString("</body>")
	w.WriteString("</html>")
	w.Flush()
}

func writeHtmlReportForGroup(userKeys []string, users map[string]map[string]interface{}, columns []string, w *bufio.Writer) {

	sort.Strings(userKeys)

	w.WriteString("<table>")
	w.WriteString("<thead>")
	//write header
	w.WriteString("<th>Nr.</th>")
	for _, column := range columns {
		w.WriteString(fmt.Sprintf("<th class=\"%v\">%v</th>", column, de[column]))
	}
	w.WriteString("</thead>")

	w.WriteString("\n")

	//write data
	for i, userKey := range userKeys {
		user := users[userKey]
		w.WriteString("<tr>")
		w.WriteString(fmt.Sprintf("<td>%v</td>", i+1))
		for _, k := range columns {
			cell := user[k]
			if cell == nil {
				w.WriteString("<td></td>")
			} else if k == "photo" {
				w.WriteString(fmt.Sprintf("<td><img src='img/%v' class='photo'></td>", cell))
			} else {
				w.WriteString(fmt.Sprintf("<td>%v</td>", cell))
			}

		}
		w.WriteString("</tr>")
		w.WriteString("\n")
	}
	w.WriteString("</table>")
}

func toString(value interface{}) string {
	return fmt.Sprint(value)
}

func prepareString(key, str string) (ret string) {
	ret = strings.Replace(strings.Replace(str, "\r\n", "<br>", -1), "\n", "<br>", -1)
	return
}

func fill(mainKey string, m php_serialize.PhpValue, user map[string]interface{}, keys map[string]string) {
	if a, ok := m.(php_serialize.PhpArray); ok {
		for kk, vv := range a {
			kv := fmt.Sprintf("%v", kk)
			if subMap, ok := vv.(php_serialize.PhpArray); ok {
				fill(kv, subMap, user, keys)
			} else {
				if kv == "0" {
					kv = mainKey
				}
				user[kv] = prepareString(kv, toString(vv))
			}
		}
	} else {
		println(a)
	}
}
