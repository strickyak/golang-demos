package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func Markup(s []byte) []string {
	t1 := blackfriday.MarkdownCommon(s)
	t2 := bluemonday.UGCPolicy().SanitizeBytes(t1)
	t3 := bytes.Split(t2, []byte("\n"))

	var lines []string
	page := 1
	for _, t := range t3 {
		s := string(t)
		if strings.HasPrefix(s, "<h1>") {
			s = fmt.Sprintf("<h1><a name=0><small>0.</small></a> &nbsp; %s", s[4:])
		}
		if strings.HasPrefix(s, "<h2>") {
			s = fmt.Sprintf("<br><br><br><br><br><br><br><br><br><br><hr><br><h2><a name=%d><small>%d.</small></a> &nbsp; %s", page, page, s[4:])
			page++
		}
		lines = append(lines, s)
	}
	return lines
}

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(HEAD)
	for _, line := range Markup(raw) {
		fmt.Println(line)
	}
	fmt.Println(TAIL)
}

const HEAD = `

<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

    <style>
			body {
				background-color: #330033;
				/* color: #FFFFFF; */
				color: #4444FF;
				font-family: "Trebuchet Ms", Verdana, sans-serif;
				/* font-size: 500%; */
				/* font-family: courier, monospace; */
				/* font-family: monospace;*/
			}

			:visited {
				color: #9900ff;
				font-family: monospace;
				text-decoration: none;
			}
			:link {
				color: #9900ff;
				font-family: monospace;
				text-decoration: none;
			}
			:hover {
				color: #9900ff;
				font-family: monospace;
				text-decoration: none;
			}

		  pre {
			  width: 95%;
				padding: 0.6em;
				background-color: #000000;
			}
			code {
				background-color: #000000;
				color: #44cc44;
				font-family: monospace;
			}
    </style>

    <script src="./jquery-3.3.1.min.js"></script>

    <script>
      Page = -1;

      // Thanks https://stackoverflow.com/questions/13735912/anchor-jumping-by-using-javascript
      function jump(h) {
    location.href = "#" + h;
      }

      $(document).keypress(function(e){
    kc = e.keyCode & 31;
    if (kc == 14) {  // Next
      ++Page;
      jump(Page);
    } else if (kc == 16) { // Prev
      --Page;
      if (Page < 0) Page = 0;
      jump(Page);
    } else if (kc == 8) { // Home
      Page = 0;
      jump(Page);
    } else if (kc == 20) { // Ten
      Page += 10;
      jump(Page);
    }
      });

    </script>
  </head>
  <body>
	  <small>
	    Use "n" for next slide; "p" for previous slide; "h" to go home to first slide.
	  </small>
`

const TAIL = `
<br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
That's all, folks!
<br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
<br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
  </body>
</html>
`
