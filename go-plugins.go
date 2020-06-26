package go_plugins

/*
 It required in the plugin libraries the presence of method 'GetPluginProxy'
 that return the github.com/hellgate75/go-plugins/model/PluginProxy object.
 This feature allows developer to return more Plugins of different blends and types,
 stored within a single plugin library.
 Don't forget to initialize the Plugin library with package function 'EnablePlugins',
 that configures and pre-load plugins from the configured folder and start buffering plugins
 in the memory.
*/

import (
	"github.com/hellgate75/go-plugins/log"
	"github.com/hellgate75/go-plugins/model"
	"github.com/hellgate75/go-plugins/plugin"
	"github.com/hellgate75/go-plugins/plugin/proxy"
)

// Enable plugins package and start plugins scan features
func EnablePlugins(config model.PluginsConfig, logger log.Logger) {
	plugin.Logger = logger
	proxy.UsePlugins = true
	proxy.InitPlugins(config, logger)
}

// Collects plugins proxy and eventually load plugins if not already done
func GetPluginProxy() proxy.Proxy {
	return proxy.NewProxy()
}

// Collects all available plugins from required folder
func GetAllPlugins() ([]model.Plugin, error) {
	return plugin.GetPlugins()
}

// Collects all plugins matching with given plugin type
func GetPluginsByType(pluginType model.PluginType) ([]model.Plugin, error) {
	return plugin.SeekForPlugins(pluginType)
}

// Collects and filter (using given filter function) all plugins matching with given plugin type
func FilterPluginsByType(pluginType model.PluginType, filter func(model.Plugin) bool) ([]model.Plugin, error) {
	return plugin.FilteredPluginsByType(pluginType, filter)
}

// Collects and filter (using given filter function) all plugins
func FilterPlugins(filter func(model.Plugin) bool) ([]model.Plugin, error) {
	return plugin.FilteredPlugins(filter)
}
