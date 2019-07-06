package main

import (
	_ "github.com/berthojoris/ginbro/config"
	"github.com/berthojoris/ginbro/handlers"
	"github.com/berthojoris/ginbro/tasks"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
