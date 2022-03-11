package users

import (
	"github.com/iotaledger/hive.go/configuration"
	"github.com/iotaledger/hive.go/node"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/users"
)

// PluginName is the name of the user plugin.
const PluginName = "Users"

var userMap map[string]*users.UserData

var config *configuration.Configuration

func Init(_config *configuration.Configuration) *node.Plugin {
	config = _config
	return node.NewPlugin(PluginName, node.Enabled, configure, run)
}

func run(_ *node.Plugin) {
}

func configure(plugin *node.Plugin) {
	err := loadUsersFromConfiguration()
	if err != nil {
		plugin.LogErrorf("Failed to pull users: {#err}")
	}

	users.InitUsers(userMap)
}

func loadUsersFromConfiguration() error {
	err := config.Unmarshal(parameters.UserList, &userMap)
	if err != nil {
		userMap = make(map[string]*users.UserData)
	}

	for username, userData := range userMap {
		userData.Username = username
	}

	return err
}