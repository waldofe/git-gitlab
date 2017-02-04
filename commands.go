package main

import (
	"log"
	"os"
	"os/exec"

	"fmt"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	command_clone,
	command_merge_request,
	command_merge,
	command_show,
}

var command_clone = cli.Command{
	Name:  "clone",
	Usage: "git lab clone <namescape/repository> [dir]",
	Description: `
`,
	Action: do_clone,
}

var command_merge_request = cli.Command{
	Name:  "merge-request",
	Usage: "create merge request.",
	Description: `
`,
	Action: do_merge_request,
}

var command_merge = cli.Command{
	Name:  "merge",
	Usage: "merge specified merge request.",
	Description: `
`,
	Action: do_merge,
}

var command_show = cli.Command{
	Name:  "show",
	Usage: "show issue, merge request, wiki and repository",
	Description: `
`,
	Action: do_show,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func do_clone(c *cli.Context) {
	remote := c.Args().Get(0)
	local := c.Args().Get(1)
	config := NewGlobalGitConfig()

	client, e := NewGitLabClient(config)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	ret, e := client.clone(remote, local)
	if e != nil {
		fmt.Println(e.Error())
	}
	if client != nil {
		fmt.Println(ret)
	}
}

func do_merge_request(c *cli.Context) {
}

func do_merge(c *cli.Context) {
}

func do_show(c *cli.Context) {
	issuablePath := c.Args().Get(0)

	config, e := NewLocalGitConfig()
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	projectPath, e := config.Project()
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	hostPath, e := config.Host()
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	completePath := (hostPath + "/" + projectPath + "/" + issuablePath)
	fmt.Println("Opening " + completePath)
	exec.Command("open", completePath).Output()
}
