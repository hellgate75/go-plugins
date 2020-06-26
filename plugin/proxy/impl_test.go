package proxy

import (
	"github.com/hellgate75/go-plugins/model"
	"testing"
)

type samplePlugin struct {
	content		map[string]interface{}
	ptype		model.PluginType
}

func (p *samplePlugin) PluginType() model.PluginType {
	return p.ptype
}

func (p *samplePlugin) Content() map[string]interface{} {
	return p.content
}

var pluginType1 = model.PluginType(1)
var pluginType2 = model.PluginType(2)

func getProxy() Proxy {
	return &proxy{
		plugins: []model.Plugin{
			&samplePlugin{
				ptype: pluginType1,
				content: make(map[string]interface{}),
			},
			&samplePlugin{
				ptype: pluginType2,
				content: make(map[string]interface{}),
			},
		},
	}
}

func TestProxyFunctionGetPlugins(t *testing.T) {
	var pr = getProxy()
	var plugins = pr.GetPlugins()
	assert(t,"Plugins must be 2", 2, len(plugins))
}

func TestProxyFunctionGetPluginsByType(t *testing.T) {
	var pr = getProxy()
	plugins, _ := pr.GetPluginsByType(pluginType1)
	assert(t,"Plugins of type 1 must be 1", 1, len(plugins))
}

func TestProxyFunctionGetFilteredPluginsByType(t *testing.T) {
	var pr = getProxy()
	plugins, _ := pr.GetFilteredPluginsByType(pluginType1, func(pl model.Plugin) bool {
		return byte(pl.PluginType()) == byte(pluginType1)
	})
	assert(t,"Plugins of type 1 and filtered by type 1 must be 1", 1, len(plugins))
}


func TestProxyFunctionGetFilteredPlugins(t *testing.T) {
	var pr = getProxy()
	plugins, _ := pr.GetFilteredPlugins(func(pl model.Plugin) bool {
		return byte(pl.PluginType()) == byte(pluginType1)
	})
	assert(t,"Plugins filtered by type 1 must be 1", 1, len(plugins))
}