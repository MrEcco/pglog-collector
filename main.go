package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/hpcloud/tail"
)

func main() {

	// Flags
	filename := flag.String("f", "postgresql-example.csv", "name of CSV log file")
	flag.Parse()

	entryParts := make([]string, 0)

	var timestampRegexp *regexp.Regexp
	timestampRegexp, _ = regexp.Compile("^\\d{4}-\\d{2}-\\d{2}\\ \\d{2}:\\d{2}:\\d{2}\\.\\d{3}")

	csvlog, err := tail.TailFile(*filename, tail.Config{Follow: true})
	if err != nil {
		panic(err.Error())
	}
	for tailLine := range csvlog.Lines {
		line := tailLine.Text
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err.Error())
			}
		}
		if err != nil {
			panic(err.Error())
		}
		if i := timestampRegexp.FindIndex([]byte(line)); i != nil {
			// Start of log entry
			values := splitLogEntry(entryParts)
			l := len(values)
			if l == 0 {
				// Skip first or empty message
			} else if l == 22 {
				fmt.Println(jsonizeFromMap(convertToMap(values)))
			} else {
				// Brocken entry
				fmt.Println(jsonizeFromMap(convertBrockenEntryToMap(values)))
			}
			entryParts = make([]string, 0)
			entryParts = append(entryParts, line)
		} else {
			// Part of previous log entry
			entryParts = append(entryParts, line)
		}
	}

}

func splitLogEntry(parts []string) []string {

	values := make([]string, 0)

	// Parse parts
	remainPart := ""
	for _, p := range parts {

		isOpenedDoubleQuotes := false
		indexPointer := 0

		p = remainPart + p

		for j, c := range p {
			if c == '"' {
				isOpenedDoubleQuotes = !isOpenedDoubleQuotes
			} else {
				if !isOpenedDoubleQuotes && c == ',' && j != 0 {
					if p[indexPointer] == ',' {
						values = append(values, repairDoubleQuotes(p[indexPointer+1:j]))
					} else {
						values = append(values, repairDoubleQuotes(p[indexPointer:j]))
					}
					indexPointer = j
				}
			}
		}

		remainPart = p[indexPointer:]
	}

	return values
}

func repairDoubleQuotes(str string) string {
	l := len(str)

	// Remove start/end
	if l < 2 {
		return str
	} else if str == "\"\"" {
		return ""
	} else {
		if str[0] == '"' && str[l-1] == '"' {
			return strings.Replace(str[1:l-1], "\"\"", "\"", -1) // doublequotes deduplication
		}
		return strings.Replace(str, "\"\"", "\"", -1) // doublequotes deduplication
	}
}

func convertToMap(values []string) map[string]string {
	retmap := make(map[string]string)
	// This field just crush log message
	// retmap["time"] = values[0]                   // log_time timestamp(3) with time zone,
	retmap["user_name"] = values[1]              // user_name text,
	retmap["database_name"] = values[2]          // database_name text,
	retmap["process_id"] = values[3]             // process_id integer,
	retmap["connection_from"] = values[4]        // connection_from text,
	retmap["session_id"] = values[5]             // session_id text,
	retmap["session_line_num"] = values[6]       // session_line_num bigint,
	retmap["command_tag"] = values[7]            // command_tag text,
	retmap["session_start_time"] = values[8]     // session_start_time timestamp with time zone,
	retmap["virtual_transaction_id"] = values[9] // virtual_transaction_id text,
	retmap["transaction_id"] = values[10]        // transaction_id bigint,
	retmap["error_severity"] = values[11]        // error_severity text,
	retmap["sql_state_code"] = values[12]        // sql_state_code text,
	retmap["log_message"] = values[13]           // detail text,
	retmap["hint"] = values[14]                  // hint text,
	retmap["internal_query"] = values[15]        // internal_query text,
	retmap["internal_query_pos"] = values[16]    // internal_query_pos integer,
	retmap["context"] = values[17]               // context text,
	retmap["query"] = values[18]                 // query text,
	retmap["query_pos"] = values[19]             // query_pos integer,
	retmap["location"] = values[20]              // location text,
	retmap["application_name"] = values[21]      // application_name text,
	return retmap
}

func convertBrockenEntryToMap(parts []string) map[string]string {
	retmap := make(map[string]string)
	retmap["error"] = "broken log entry"
	retmap["log_entry"] = strings.Join(parts, "")
	return retmap
}

func jsonizeFromMap(inputmap map[string]string) string {
	if bytebuf, err := json.Marshal(inputmap); err != nil {
		panic(err.Error())
	} else {
		return string(bytebuf)
	}
}
