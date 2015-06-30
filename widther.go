package widther

type delimiter func(uint8) bool

func spaceDelimiter(c uint8) bool {
	return c == ' '
}

func isWhiteSpace(c uint8) bool {
	return c == ' ' || c == '\t'
}

func Widthen(text string, width int, delimiters ...uint8) (splits []string) {
	ch := map[uint8]bool{}
	for _, delim := range delimiters {
		ch[delim] = true
	}

	fn := func(c uint8) bool {
		_, ok := ch[c]
		return ok
	}

	return widthen(text, fn, width)
}

func WidthenByLimit(text string, width int) (splits []string) {
	return widthen(text, spaceDelimiter, width)
}

func widthen(text string, fn delimiter, width int) (splits []string) {
	length := len(text)
	if length <= width || width < 1 {
		splits = append(splits, text)
		return
	}

	splitsChan := make(chan string)

	go func() {
		defer close(splitsChan)

		i := 0
		end := 0

		processing := true
		for processing {
			end = i + width

			if end >= length {
				end = length
				processing = false
			}

			splitsChan <- text[i:end]
			i = end
		}
	}()

	for split := range splitsChan {
		stop := len(split) - 1
		endIndex := stop
		sediments := []string{}

		for endIndex >= 0 {
			if fn(split[endIndex]) {
				last := split[endIndex:stop]
				stop = endIndex
				sediments = append(sediments, last)
			}

			if endIndex <= width {
				sediments = append(sediments, split[:endIndex+1])
				break
			}

			if !isWhiteSpace(split[endIndex]) {
				endIndex -= 1
			} else {
				endIndex -= 4
			}
		}

		reversed := []string{}
		for i := len(sediments) - 1; i >= 0; i-- {
			reversed = append(reversed, sediments[i])
		}
		splits = append(splits, reversed...)
	}

	return
}
