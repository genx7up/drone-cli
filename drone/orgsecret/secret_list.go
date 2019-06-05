package orgsecret

import (
	"html/template"
	"os"

	"github.com/urfave/cli"

	"github.com/drone/drone-cli/drone/internal"
)

var secretListCmd = cli.Command{
	Name:      "ls",
	Usage:     "list secrets",
	ArgsUsage: "",
	Action:    secretList,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "format",
			Usage:  "format output",
			Value:  tmplSecretList,
			Hidden: true,
		},
	},
}

func secretList(c *cli.Context) error {
	format := c.String("format") + "\n"
	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}
	list, err := client.OrgSecretListAll()
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Parse(format)
	if err != nil {
		return err
	}
	for _, secret := range list {
		tmpl.Execute(os.Stdout, secret)
	}
	return nil
}

// template for secret list items
var tmplSecretList = "\x1b[33m{{ .Name }} \x1b[0m" + `
Organization:       {{ .Namespace }}
Pull Request Read:  {{ .PullRequest }}
Pull Request Write: {{ .PullRequestPush }}
`
