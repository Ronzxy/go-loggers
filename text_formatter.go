package logger

import (
	"fmt"
	"github.com/skygangsta/go-utils"
	"strings"
	"time"
)

type TextFormatter struct {
	Format string
}

var (
	DefaultFormat = "%{Name} - %{Time:Y-m-d H:M:S.ms} - %{Level:5} - %{File}:%{Line:3} - %{Message}"
)

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{
		Format: DefaultFormat,
	}
}

func (this *TextFormatter) SetFormat(format string) {
	this.Format = format
}

func (this *TextFormatter) Message(data map[string]interface{}, args ...interface{}) string {
	var (
		varPattern string
		varName    string
		vars       = make([]string, 2)
		str        = this.Format
	)
	for {
		varPattern, varName = Variable("%", "([a-zA-Z_][0-9a-zA-Z\\s\\._/:-]*)", str)
		if varName == "" {
			// no variable, exit loop
			break
		}

		if strings.Contains(varName, ":") {
			vars = strings.Split(varName, ":")
		} else {
			vars[0] = varName
			vars[1] = ""
		}

		switch strings.ToUpper(vars[0]) {
		case "NAME":
			{
				varName = fmt.Sprintf("%v", data["Name"])
			}
		case "TIME":
			{
				if vars[1] != "" {
					varName = util.NewDate().Format(time.Now(), strings.Join(vars[1:], ":"))
				} else {
					varName = time.Now().Format(DefaultLogTimeFormat)
				}
			}
		case "LEVEL":
			{
				format := "%s"
				if vars[1] != "" {
					format = fmt.Sprintf("%%-%ss", vars[1])
				}

				varName = fmt.Sprintf(format, data["Level"])
			}
		case "FILE":
			{
				format := "%s"
				if vars[1] != "" {
					format = fmt.Sprintf("%%-%ss", vars[1])
				}

				varName = fmt.Sprintf(format, data["File"])
			}
		case "LINE":
			{
				format := "%s"
				if vars[1] != "" {
					format = fmt.Sprintf("%%-%sd", vars[1])
				}

				varName = fmt.Sprintf(format, data["Line"])
			}
		case "MESSAGE":
			{
				//format := fmt.Sprintf("%s", data["Format"])
				varName = fmt.Sprint(args...)
			}
		default:
			{
				Errorf("unsupported log format %s", varName)
			}
		}

		str = strings.Replace(str, varPattern, varName, -1)
	}

	return str
}
