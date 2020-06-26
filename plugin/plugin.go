package plugin

import (
	"errors"
	"fmt"
	"github.com/hellgate75/go-plugins/log"
	"github.com/hellgate75/go-plugins/model"
	"github.com/hellgate75/go-plugins/plugin/proxy"
)

// Message logger interface
var Logger log.Logger

// Interface that describe a Request in the Deploy Manager Plugin Explorer
type SeekRequest struct {
	Module string
	Symbol string
}

var proxyVia proxy.Proxy = nil

// Load plugins from the modules folder
func GetPlugins() ([]model.Plugin, error) {
	var errGlobal error = nil
	defer func() {
		if r := recover(); r != nil {
			if Logger != nil {
				Logger.Errorf("plugin.GetPlugins -> Error: %v", r)
			} else {
				fmt.Printf("plugin.GetPlugins -> Error: %v\n", r)
			}
			errGlobal = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	if proxyVia == nil {
		proxyVia = proxy.NewProxy()
	}
	return proxyVia.GetPlugins(), errGlobal
}

func SeekForPlugins(pluginType model.PluginType) ([]model.Plugin, error) {
	var errGlobal error = nil
	defer func() {
		if r := recover(); r != nil {
			if Logger != nil {
				Logger.Errorf("plugin.SeekForPlugins -> Error: %v", r)
			} else {
				fmt.Printf("plugin.SeekForPlugins -> Error: %v\n", r)
			}
			errGlobal = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	if proxyVia == nil {
		proxyVia = proxy.NewProxy()
	}
	var out = make([]model.Plugin, 0)
	plugins, err := proxyVia.GetPluginsByType(pluginType)
	if err == nil {
		out = append(out, plugins...)
	}
	if Logger != nil {
		Logger.Debugf("Plugins by type %v present: %v", pluginType, len(out))
	}
	if err != nil {
		return nil, err
	}
	return out, errGlobal
}

func FilteredPluginsByType(pluginType model.PluginType, filter func(model.Plugin) bool) ([]model.Plugin, error) {
	var errGlobal error = nil
	defer func() {
		if r := recover(); r != nil {
			if Logger != nil {
				Logger.Errorf("plugin.FilteredPluginsByType -> Error: %v", r)
			} else {
				fmt.Printf("plugin.FilteredPluginsByType -> Error: %v\n", r)
			}
			errGlobal = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var out = make([]model.Plugin, 0)
	plugins, err := SeekForPlugins(pluginType)
	if err != nil {
		return out, err
	}
	for _, p := range plugins {
		if filter(p) {
			out = append(out, p)
		}
	}
	if Logger != nil {
		Logger.Debugf("Plugins by type %v  filtered: %v", pluginType, len(out))
	}
	return out, errGlobal
}

func FilteredPlugins(filter func(model.Plugin) bool) ([]model.Plugin, error) {
	var errGlobal error = nil
	defer func() {
		if r := recover(); r != nil {
			if Logger != nil {
				Logger.Errorf("plugin.FilteredPlugins -> Error: %v", r)
			} else {
				fmt.Printf("plugin.FilteredPlugins -> Error: %v\n", r)
			}
			errGlobal = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var out = make([]model.Plugin, 0)
	if proxyVia == nil {
		proxyVia = proxy.NewProxy()
	}
	for _, p := range proxyVia.GetPlugins() {
		if filter(p) {
			out = append(out, p)
		}
	}
	if Logger != nil {
		Logger.Debugf("Plugins filtered: %v", len(out))
	}
	return out, errGlobal
}
