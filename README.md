A Go implementation of the Rapid Automatic Keyword Extraction (RAKE) algorithm as described in: Rose, S., Engel, D., Cramer, N., & Cowley, W. (2010). Automatic Keyword Extraction from Individual Documents. In M. W. Berry & J. Kogan (Eds.), Text Mining: Theory and Applications: John Wiley & Sons.

Original Python implementation available at: https://github.com/aneesha/RAKE

The source code is released under the MIT License.

## Docs and Report Card
- godoc.org: https://godoc.org/github.com/afjoseph/RAKE.Go
- goreportcard.com: https://goreportcard.com/report/github.com/afjoseph/RAKE.Go

## Example Usage

```go
package main

import (
	"github.com/afjoseph/RAKE.go"
	"fmt"
)

func main() {
	text := `The growing doubt of human autonomy and reason has created a state of moral confusion where man is left without the guidance of either revelation or reason. The result is the acceptance of a relativistic position which proposes that value judgements and ethical norms are exclusively matters of arbitrary preference and that no objectively valid statement can be made in this realm... But since man cannot live without values and norms, this relativism makes him an easy prey for irrational value systems.`

	candidates := rake.RunRake(text)

	for _, candidate := range candidates {
		fmt.Printf("%s --> %f\n", candidate.Key, candidate.Value)
	}

	fmt.Printf("\nsize: %d\n", len(candidates))
}

<!---------------------------------------------------------->
<!--output-->
<!---------------------------------------------------------->
<!--objectively valid statement --> 9.000000-->
<!--exclusively matters --> 4.000000-->
<!--arbitrary preference --> 4.000000-->
<!--easy prey --> 4.000000-->
<!--relativistic position --> 4.000000-->
<!--human autonomy --> 4.000000-->
<!--relativism makes --> 4.000000-->
<!--growing doubt --> 4.000000-->
<!--moral confusion --> 4.000000-->
<!--ethical norms --> 3.500000-->
<!--norms --> 1.500000-->
<!--made --> 1.000000-->
<!--guidance --> 1.000000-->
<!--man --> 1.000000-->
<!--result --> 1.000000-->
<!--systems --> 1.000000-->
<!--values --> 1.000000-->
<!--realm --> 1.000000-->
<!--live --> 1.000000-->
<!--judgements --> 1.000000-->
<!--reason --> 1.000000-->
<!--left --> 1.000000-->
<!--proposes --> 1.000000-->
<!--irrational --> 1.000000-->
<!--created --> 1.000000-->
<!--acceptance --> 1.000000-->
<!--revelation --> 1.000000-->
<!--state --> 1.000000-->

<!--size: 28-->
```
