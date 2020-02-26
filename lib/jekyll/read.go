package jekyll

import (
	"encoding/json"
	"io/ioutil"
)

const pageJsonPath = "_site/api/v1/pages.json"

func ReadPageJson() (out *Pages) {
	file, err := ioutil.ReadFile(pageJsonPath)
	if err != nil {
		panic(err)
	}

	out = new(Pages)

	err = json.Unmarshal(file, out)
	if err != nil {
		panic(err)
	}

	return
}
