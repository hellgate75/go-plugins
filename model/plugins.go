package model

// Plugin Type
type PluginType	byte

//  Plugin component behaviour interface
type Plugin interface {
	PluginType() PluginType
	Content() map[string]interface{}
}

// Interface used by plugin to enumerate multiple plugins
type PluginProxy interface{
	GetPlugins() ([]Plugin, error)
}

// Plugins Configuration Helper structure
type PluginsConfig struct {
	// Define base folder that contains 'module' sub-folder within all plugins
	BaseFolderPath		string		`yaml:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// Force reload despite the cache all plugins from the module folder at each request.
	ForcePluginsReload	bool		`yaml:"force,omitempty" json:"force,omitempty" xml:"force,omitempty"`
}
