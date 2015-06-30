package widther

import (
	"strings"
	"testing"
)

func TestWidthing(t *testing.T) {
	var texts = []string{
		"func main() {      func Widthen(text______string, delimiter rune, width uint)\t(splits []string) {\nlength\t = \tlen(text)\t\nif length <= width {splits = append\n(splits, text)\nreturn   }}   }",
		"International donations are gratefully accepted, but we cannot make\nany statements concerning tax treatment of donations received from\noutside the United States.  U.S. laws alone swamp our small staff.\n\nPlease check the Project Gutenberg Web pages for current donation\nmethods and addresses.  Donations are accepted in a number of other\nways including including checks, online payments and credit card\ndonations.  To donate, please visit: http://pglaf.org/donate\n\n\nSection 5.  General Information About Project Gutenberg-tm electronic\nworks.\n\nProfessor Michael S. Hart is the originator of the Project Gutenberg-tm\nconcept of a library of electronic works that could be freely shared\nwith anyone.  For thirty years, he produced and distributed Project\nGutenberg-tm eBooks with only a loose network of volunteer support.\n\nProject Gutenberg-tm eBooks are often created from several printed\neditions, all of which are confirmed as Public Domain in the U.S.\nunless a copyright notice is included.  Thus, we do not necessarily\nkeep eBooks in compliance with any particular paper edition.\n\nMost people start at our Web site which has the main PG search facility:\n\n     http://www.gutenberg.net\n\nThis Web site includes information about Project Gutenberg-tm,\nincluding how to make donations to the Project Gutenberg Literary\nArchive Foundation, how to help produce our new eBooks, and how to\nsubscribe to our email newsletter to hear about new eBooks.",
	}

	expectations := []struct {
		delimiters []uint8
		widths     []int
	}{
		{
			delimiters: []uint8{' ', '_'},
			widths:     []int{10, 20, 1, 0, -1, 90, 900},
		},
		{
			delimiters: []uint8{'X', '_'},
			widths:     []int{5, 80, -9, 29, -1, 90, -15},
		},
		{
			delimiters: []uint8{'_', 'p'},
			widths:     []int{5, 80, -9, 29, -1, 90, -15},
		},
	}

	for _, exp := range expectations {
		widths := exp.widths

		for _, width := range widths {
			for _, text := range texts {
				splits := Widthen(text, width, exp.delimiters...)
				reconstructed := strings.Join(splits, "")

				if reconstructed != text {
					t.Errorf("expected \"%s\" got \"%s\"", text, reconstructed)
				}
			}
		}
	}
}
