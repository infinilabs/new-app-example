package main

import (
        "infini.sh/framework"
        "infini.sh/framework/core/module"
        "infini.sh/framework/core/util"
        "infini.sh/framework/modules/api"
        "infini.sh/new_app/config"
)

func main() {

        terminalHeader := ("     __               _      ___  ___ \n")
        terminalHeader += ("  /\\ \\ \\_____      __/_\\    / _ \\/ _ \\\n")
        terminalHeader += (" /  \\/ / _ \\ \\ /\\ / //_\\\\  / /_)/ /_)/\n")
        terminalHeader += ("/ /\\  /  __/\\ V  V /  _  \\/ ___/ ___/ \n")
        terminalHeader += ("\\_\\ \\/ \\___| \\_/\\_/\\_/ \\_/\\/   \\/     \n\n")

        terminalFooter := ("Goodbye~")
        app := framework.NewApp("new_app", "Make a golang application is such easy~.",
                util.TrimSpaces(config.Version), util.TrimSpaces(config.BuildNumber), util.TrimSpaces(config.LastCommitLog), util.TrimSpaces(config.BuildDate), util.TrimSpaces(config.EOLDate), terminalHeader, terminalFooter)
        app.IgnoreMainConfigMissing()
        app.Init(nil)
        defer app.Shutdown()
        if app.Setup(func() {
                module.RegisterSystemModule(&api.APIModule{})
                module.Start()
        }, func() {
        }, nil) {
                app.Run()
        }
}
