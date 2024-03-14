package replace

import (
	"time"

	"github.com/HurdyGutty/go_office_automation/pkg/readExcel"
	"github.com/lukasjarosch/go-docx"
)

func Replace(person readExcel.Person, temp_dir, output_dir string) {
	replaceMap := docx.PlaceholderMap{}

	for k, v := range person {
		len := len(k)
		if t, ok := v.(time.Time); ok {
			v = t.Format("02/01/2006")
		}
		replaceMap[k[1:len-1]] = v
	}

	doc, err := docx.Open(temp_dir)
	if err != nil {
		panic(err)
	}

	// replace the keys with values from replaceMap
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}

	// write out a new file
	err = doc.WriteToFile(output_dir)
	if err != nil {
		panic(err)
	}
}
