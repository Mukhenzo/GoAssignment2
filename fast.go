package main

import (
	"bufio"
	"bytes"
	json "encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

//easyjson:json
type userJson struct {
	Browsers []string `json:"browsers"`
	Email string `json:"email"`
	Name string `json:"name"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	// SlowSearch(out)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)

	user := &userJson{}
	var seenBrowsers []string
	uniqueBrowsers := 0

	fmt.Fprintln(out, "found users:")
	counter := 0
	for sc.Scan()  {
		counter++
		line := sc.Bytes()

		if bytes.Contains(line, []byte("Android")) == false && bytes.Contains(line, []byte("MSIE")) == false {
			continue
		}

		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			notSeenBefore := true
			toWrite := false

			if strings.Contains(browser, "Android") {
				isAndroid = true
				toWrite = true
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				toWrite = true
			}

			if toWrite {
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		fmt.Fprintf(out, "[%d] %s <%s>\n", counter-1, user.Name, email)
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBe44ce7eDecodeMukhenzoAssignment2Jsongg(in *jlexer.Lexer, out *userJson) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBe44ce7eEncodeMukhenzoAssignment2Jsongg(out *jwriter.Writer, in userJson) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v userJson) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBe44ce7eEncodeMukhenzoAssignment2Jsongg(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v userJson) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBe44ce7eEncodeMukhenzoAssignment2Jsongg(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *userJson) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBe44ce7eDecodeMukhenzoAssignment2Jsongg(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *userJson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBe44ce7eDecodeMukhenzoAssignment2Jsongg(l, v)
}
