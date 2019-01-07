package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type JSONSchemaArray []*JSONSchema

type JSONDiscriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping"`
}

type JSONSchema struct {
	Schema string `json:"$schema"`
	ID04   string `json:"id,omitempty"`
	ID06   string `json:"$id,omitempty"`
	Ref    string `json:"$ref,omitempty"`

	Description   string                 `json:"description"`
	Discriminator *JSONDiscriminator     `json:"discriminator,omitempty"`
	Pattern       string                 `json:"pattern"`
	Definitions   map[string]*JSONSchema `json:"definitions"`
	AllOf         JSONSchemaArray        `json:"allOf,omitempty"`
	OneOf         JSONSchemaArray        `json:"oneOf,omitempty"`

	TypeValue interface{} `json:"type,omitempty"`
	Enum      []string    `json:"enum,omitempty"`
	Const     string      `json:"const,omitempty"`

	// Properties, Required and AdditionalProperties describe an object's child instances.
	// http://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5
	Properties           map[string]*JSONSchema
	Required             []string
	AdditionalProperties JSONSchemaArray

	// Default can be used to supply a default JSON value associated with a particular schema.
	// http://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.2
	Default interface{}

	// Items represents the types that are permitted in the array.
	// http://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4
	Items *JSONSchema
}

// ID returns the schema URI id.
func (s *JSONSchema) ID() string {
	if s.ID06 == "" && s.ID04 != "" {
		return s.ID04
	}
	return s.ID06
}

const (
	packagename = "models"
	fname       = "../../api/fhir.schema.r4.json"
)

var (
	imports = []string{"regexp", "encoding/json"}
	outfile = os.Stdout
)

func FHIRType(property *JSONSchema) string {

	if typ, ok := property.TypeValue.(string); ok {
		switch typ {
		case "array":
			return fmt.Sprintf("[]%s", FHIRType(property.Items))
		case "boolean":
			return "bool"
		case "number":
			switch property.Pattern {
			case `^[0]|([1-9][0-9]*)$`:
				return "uint64"
			case `^[1-9][0-9]*$`:
				return "uint64"
			case `^-?([0]|([1-9][0-9]*))$`:
				return "int64"
			case `^-?([0]|([1-9][0-9]*))(\.[0-9]+)?$`:
				return "float64"
			case `^-?(0|[1-9][0-9]*)(\.[0-9]+)?([eE][+-]?[0-9]+)?$`:
				return "float64"
			default:
				return "UNKNUM"
			}
		case "string":
			switch property.Pattern {
			// Date
			case "-?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?":
				return "time.Time"
			// DateTime
			case "-?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\\.[0-9]+)?(Z|(\\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?":
				return "time.Time"
			// Time
			case "([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\\.[0-9]+)?":
				return "time.Time"
			default:
				return "string"
			}
		default:
			return fieldName(typ)
		}
	} else if len(property.Ref) > 0 {
		switch property.Ref {
		case "#/definitions/string":
			return "string"
		case "#/definitions/boolean":
			return "bool"
		case "#/definitions/xhtml":
			return "string"
		case "#/definitions/integer":
			return "int64"
		case "#/definitions/decimal":
			return "float64"
		case "#/definitions/unsignedInt":
			return "uint64"
		case "#/definitions/positiveInt":
			return "uint64"
		case "#/definitions/date":
			return "string"
		case "#/definitions/time":
			return "string"
		case "#/definitions/dateTime":
			return "string"
		case "#/definitions/instant":
			return "string"
		case "#/definitions/canonical":
			return "string"
		case "#/definitions/uri":
			return "string"
		case "#/definitions/code":
			return "string"
		case "#/definitions/uuid":
			return "string"
		case "#/definitions/url":
			return "string"
		case "#/definitions/markdown":
			return "string"
		case "#/definitions/id":
			return "string"
		case "#/definitions/oid":
			return "string"
		default:
			t := property.Ref[strings.LastIndex(property.Ref, "/")+1:]
			r, _ := utf8.DecodeRuneInString(t)
			if unicode.IsUpper(r) {
				return fmt.Sprintf("*%s", fieldName(t))
			}
			return fmt.Sprintf("%s", fieldName(t))
		}

	} else if len(property.Const) > 0 {
		return "const"
	} else if len(property.Enum) > 0 {
		return "enum"
	}

	return "UNKTYPE"
}

func commentify(pre, s string) string {
	max := 85
	comments := []string{}

	split := strings.Split(strings.Replace(s, "\r", "\n", -1), "\n")
	for i := range split {
		str := split[i]
		for {
			if len(str) <= max {
				comments = append(comments, fmt.Sprintf("%s// %s", pre, str))
				break
			}
			clip := str[:max]
			i := strings.LastIndexFunc(clip, unicode.IsSpace)
			if i == -1 {
				comments = append(comments, fmt.Sprintf("%s// %s", pre, str))
				break
			}
			comments = append(comments, fmt.Sprintf("%s// %s", pre, str[:i]))
			str = str[i+1:]
		}
	}

	return strings.Join(comments, "\n")
}

func BuildType(typeName string, definition *JSONSchema) {
	typeName = fieldName(typeName)
	if len(definition.Pattern) > 0 && definition.TypeValue == "string" {
		fmt.Fprintf(outfile, "// %s is %s\n", typeName, commentify("", definition.Description)[3:])
		fmt.Fprintf(outfile, "type %s %s\n", typeName, definition.TypeValue)
		fmt.Fprintf(outfile, "// %[1]sPattern ...\n\tvar %[1]sPattern = regexp.MustCompile(`%[2]s`)\n", typeName, definition.Pattern)
		fmt.Fprintf(outfile, "// Validate ...\n\tfunc (t *%s) Validate() bool {\n\treturn %sPattern.MatchString(string(*t))\n}\n", typeName, typeName)
	} else {
		constants := make(map[string]string)
		enums := make(map[string][]string)

		fmt.Fprintf(outfile, "// %s is %s\n", typeName, commentify("", definition.Description)[3:])
		fmt.Fprintf(outfile, "type %s struct {\n", typeName)
		for prop := range definition.Properties {
			property := definition.Properties[prop]

			var propertyString string
			if prop[0] == '_' {
				propertyString = fmt.Sprintf("\t%sExt ", fieldName(prop[1:]))
			} else {
				propertyString = fmt.Sprintf("\t%s ", fieldName(prop))
			}

			omit := ""
			for _, req := range definition.Required {
				if prop == req {
					omit = ",omitempty"
					break
				}
			}
			ftype := FHIRType(property)
			if ftype == "const" {
				constants[prop] = property.Const
				continue
			} else if ftype == "enum" {
				enums[fieldName(prop)] = property.Enum
				ftype = fmt.Sprintf("%s%s", typeName, fieldName(prop))
			} else if ftype == "[]enum" {
				enums[fieldName(prop)] = property.Enum
				ftype = fmt.Sprintf("[]%s%s", typeName, fieldName(prop))
			}

			fmt.Fprintln(outfile, commentify("\t", property.Description))
			if len(property.Pattern) > 0 {
				fmt.Fprintf(outfile, "\t// pattern %s\n", property.Pattern)
			}
			fmt.Fprintf(outfile, "%s %s `json:\"%s%s\"`\n", propertyString, ftype, prop, omit)
		}

		fmt.Fprint(outfile, "}\n\n")
		for k, _ := range constants {
			fmt.Fprintf(outfile, "// %[2]s ...\n\tfunc (t *%[1]s) %[2]s() string {\n\treturn \"%[3]s\"\n}\n", typeName, fieldName(k), constants[k])
		}

		// Generate MarshalJSON code
		if len(constants) > 0 {
			fmt.Fprintf(outfile, "// MarshalJSON ...\n\tfunc (t *%s) MarshalJSON() ([]byte, error) {\n\treturn json.Marshal(struct {\n\t\t%s\n", typeName, typeName)
			for k, _ := range constants {
				fmt.Fprintf(outfile, "\t\t%s string `json:\"%s\"`\n", fieldName(k), k)
			}
			fmt.Fprintf(outfile, "\t}{\n\t\t%s: *t,\n", typeName)
			for k, _ := range constants {
				fmt.Fprintf(outfile, "\t\t%s: t.%s(),\n", fieldName(k), fieldName(k))
			}
			fmt.Fprint(outfile, "\t})\n}\n")
		}

		if len(enums) > 0 {
			enumarr := []string{}
			for k, _ := range enums {
				vals := enums[k]

				enumName := fmt.Sprintf("%s%s", typeName, k)
				fmt.Fprintf(outfile, "// %[1]s ...\n\ttype %[1]s string\n", enumName)

				for i := range vals {
					val := vals[i]
					titleVal := fieldName(val)
					enumarr = append(enumarr, fmt.Sprintf("// %[1]s%[2]s is a %[1]s value of \"%[3]s\"\n\t%[1]s%[2]s %[1]s = \"%[3]s\"", enumName, titleVal, val))
				}
			}
			fmt.Fprintf(outfile, "const (\n %s\n)\n\n", strings.Join(enumarr, "\n"))
		}
	}
}

var (
	replacer = strings.NewReplacer("Uuid", "UUID", "Uid", "UID", "Url", "URL", "Uri", "URI", "Xml", "XML",
		"Json", "JSON", "_", "", "Html", "HTML", "-", "", "/", "", ".", "", "<", "Lt", ">", "Gt", "=", "Eq", "!", "Not")
	idRegexp = regexp.MustCompile(`Id(\z|[A-Z])`)
)

func fieldName(in string) string {
	if in == "base64Binary" {
		return in
	}
	str := strings.Title(in)
	str = replacer.Replace(str)
	str = idRegexp.ReplaceAllString(str, "ID$2")
	return str
}

func main() {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	// var j map[string]interface{}
	var j JSONSchema
	json.Unmarshal([]byte(byteValue), &j)

	fmt.Fprintf(outfile, "package %s\n", packagename)
	fmt.Fprintln(outfile, commentify("", j.Description))
	fmt.Fprintf(outfile, "// Schema: %s\n", j.Schema)
	fmt.Fprint(outfile, "import (\n")
	for _, imp := range imports {
		fmt.Fprintf(outfile, "\t\"%s\"\n", imp)
	}
	fmt.Fprint(outfile, ")\n")

	fmt.Fprintf(outfile, "// FHIRVersion is the current FHIR version\n\tconst (\n\tFHIRVersion = \"%s\"\n)\n", j.ID()[strings.LastIndex(j.ID(), "/")+1:])

	fmt.Fprint(outfile, "// Resource ...\n\ttype Resource interface {\n\tResourceType() string\n}\n")
	fmt.Fprint(outfile, "// Validator ...\n\ttype Validator interface {\n\tValidate() bool\n}\n")

	if j.Discriminator != nil {
		fmt.Fprintf(os.Stderr, "has Discriminator: %s (%d)\n", j.Discriminator.PropertyName, len(j.Discriminator.Mapping))
	}

	if len(j.OneOf) > 0 {
		fmt.Fprintf(os.Stderr, "has OneOf: %d\n", len(j.OneOf))
	}

	for typeName, _ := range j.Discriminator.Mapping {
		definition := j.Definitions[typeName]
		BuildType(typeName, definition)
	}

	stringTypes := map[string]bool{
		"url":       true,
		"code":      true,
		"canonical": true,
		"uri":       true,
		"uuid":      true,
		"oid":       true,
		"markdown":  true,
		"id":        true,
	}
	_ = stringTypes
	omittedTypes := map[string]bool{
		"boolean":     true,
		"string":      true,
		"xhtml":       true,
		"decimal":     true,
		"integer":     true,
		"unsignedInt": true,
		"positiveInt": true,
		//======
		"date":     true,
		"time":     true,
		"dateTime": true,
		"instant":  true,
		//======
		"base64Binary": false,
	}

	for p := range j.Definitions {
		found := false
		for m := range j.Discriminator.Mapping {
			if m == p {
				found = true
			}
		}

		if !found && !omittedTypes[p] {
			definition := j.Definitions[p]
			BuildType(p, definition)
		}
	}

	fmt.Fprintf(os.Stderr, "omited %d of % d types\n", len(omittedTypes), len(j.Definitions))
}

// ============================
// type datePrecision int

// const (
// 	unknown datePrecision = iota
// 	year
// 	yearMonth
// 	yearMonthDay
// )

// type Date struct {
// 	time.Time
// 	precision datePrecision
// }

// // MarshalJSON implements the json.Marshaler interface.
// // The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
// func (t *Date) MarshalJSON() ([]byte, error) {
// 	format := time.RFC3339Nano
// 	switch t.precision {
// 	case year:
// 		format = "2006"
// 	case yearMonth:
// 		format = "2006-01"
// 	case yearMonthDay:
// 		format = "2006-01-02"
// 	default:
// 		return nil, errors.New("unknown precision")
// 	}

// 	if y := t.Year(); y < 0 || y >= 10000 {
// 		// RFC 3339 is clear that years are 4 digits exactly.
// 		// See golang.org/issue/4556#c15 for more discussion.
// 		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
// 	}

// 	b := make([]byte, 0, len(format)+2)
// 	b = append(b, '"')
// 	b = t.AppendFormat(b, format)
// 	b = append(b, '"')
// 	return b, nil
// }

// // UnmarshalJSON implements the json.Unmarshaler interface.
// // The time is expected to be a quoted string in RFC 3339 format.
// func (t *Date) UnmarshalJSON(data []byte) error {
// 	// "-?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?"
// 	// Ignore null, like in the main JSON package.
// 	if string(data) == "null" {
// 		return nil
// 	}
// 	// Fractional seconds are handled implicitly by Parse.
// 	var err error
// 	*t, err = Parse(`"`+time.RFC3339+`"`, string(data))
// 	return err
// }

// type DateTime struct {
// 	time.Time
// }

// type Time struct {
// 	time.Time
// }
