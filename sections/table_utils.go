package sections

import (
	"errors"
	"github.com/rpaloschi/dxf-go/core"
)

func TableEntryTags(tags core.TagSlice) ([]core.TagSlice, error) {
	groups := core.TagGroups(tags, 0)
	lastIndex := len(tags) - 1
	first := groups[0][0].Value.ToString()
	last := groups[lastIndex][0].Value.ToString()

	if first != "TABLE" || last != "ENDTAB" {
		return nil, errors.New("Invalid table. Missing TABLE AND/OR ENDTAB tags.")
	}

	return groups[1:lastIndex], nil
}

func SplitTagChunks(tags core.TagSlice, stopTag core.Tag, chunkDelimiter core.Tag) []core.TagSlice {
	chunks := make([]core.TagSlice, 0)

	tagIndex := 0
	for {
		if *tags[tagIndex] == stopTag {
			break
		}

		chunk := make([]*core.Tag, 1)
		chunk[0] = tags[tagIndex]

		for ; *tags[tagIndex] != chunkDelimiter; tagIndex++ {
			chunk = append(chunk, tags[tagIndex])
		}

		chunks = append(chunks, chunk)
	}

	return chunks
}
