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
				/* color: #FFFFFF; */
				color: #4444FF;
				font-family: "Trebuchet Ms", Verdana, sans-serif;
				/* font-size: 500%; */
				/* font-family: courier, monospace; */
				/* font-family: monospace;*/
			}

			:visited {
				XXcolor: #FFFF22;
				color: #8800ee;
				font-family: monospace;
				text-decoration: none;
			}
			:link {
				XXcolor: #FFFF22;
				color: #8800ee;
				font-family: monospace;
				text-decoration: none;
			}
			:hover {
				XXcolor: #FFFF22;
				color: #8800ee;
				font-family: monospace;
				text-decoration: none;
			}

		  pre {
			  width: 95%;
				padding: 0.6em;
				/*  background-color: #111133; */
				background-color: #000022;
			}
			code {
				background-color: #000022;
				/* background-color: #111111; */
				color: #44cc44;
				/* margin: 0.2em; */
				/* font-size: 500%; */
				/* font-family: courier, monospace; */
				font-family: monospace;
			}



			.XXXgreen {
				color: #00FF00;
				font-size: 500%;
				/* font-family: courier, monospace; */
				font-family: monospace;
			}
			.XXXyellow {
				color: #FFFF00;
				font-size: 500%;
				font-family: monospace;
			}
			.XXXbig-green {
				color: #00FF00;
				font-size: 750%;
				font-family: monospace;
			}
			.XXXbig-yellow {
				color: #FFFF00;
				font-size: 650%;
				font-family: monospace;
			}
			.XXXtitle {
				color: #aaaaaa;
				font-size: 250%;
				font-family: Arial, sans-serif;
				text-align: center;
			}
			.XXXprose {
				color: #FFFF00;
				font-size: 375%;
				font-family: Arial, sans-serif;
			}
			.XXXgloss {
				color: #bbbbbb;
				font-size: 150%;
				font-family: Arial, sans-serif;
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
      Page = 10;
      jump(Page);
    }
      });

    </script>
  </head>
  <body bgcolor="#220022"  was="#000055">
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
