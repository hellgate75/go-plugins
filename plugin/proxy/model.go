package proxy

import (
	"fmt"
	"github.com/hellgate75/go-plugins/log"
	"github.com/hellgate75/go-plugins/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"strings"
)

// Use custom plugins loading proxies
var UsePlugins = false

// Use custom plugins loading proxies
var ForcePluginsReload = false
// Use custom plugins folder to seek for libraries
var BaseAppFolder = ""
// Plugin Modules sub folder name
var PluginsModulesFolder = "modules"
// Assume this extension name for ;loading the libraries (we hope in future windows will allow plugins)
var PluginLibrariesExtension = "so"

// Message logger interface
var Logger log.Logger

// Define Behaviors of a Proxy Component
type Proxy interface {
	// Get all available plugins loaded from shared / static libraries
	GetPlugins() []*model.Plugin
	// Get available plugins of a specific plugin type
	GetPluginsByType(pt model.PluginType) ([]*model.Plugin, error)
	// Get available plugins of a specific plugin type, filtered by user defined function
	GetFilteredPluginsByType(pt model.PluginType, filter func(model.Plugin) bool) ([]*model.Plugin, error)
	// Get available plugins filtered by user defined function
	GetFilteredPlugins(filter func(model.Plugin) bool) ([]*model.Plugin, error)
}

// Collects plugins once from proxy and return the list
// It required in the libraries the presence of method 'GetPluginProxy'
// that return the model.PluginProxy object
func getPlugins() []model.Plugin {
	if len(pluginDataList) > 0 && ! ForcePluginsReload {
		return pluginDataList
	} else {
		pluginDataList = make([]model.Plugin, 0)
	}
	var outList = make([]model.Plugin, 0)
	if UsePlugins {
		if Logger != nil {
			Logger.Debug("modules.proxy.getPlugins() -> Loading library for map modules")
		}
		forEachPluginDo(func(pluginsList []model.Plugin) {
			if len(pluginsList) > 0 {
				outList = append(outList, pluginsList...)
			}
		})
	}
	return outList
}

// Filters file files vy plugins extension name
func filterByExtension(fileName string) bool {
	n := len(PluginLibrariesExtension)
	fileNameLen := len(fileName)
	posix := fileNameLen - n - 1
	return posix > 0 && strings.ToLower(fileName[posix:]) == strings.ToLower("." + PluginLibrariesExtension)
}

// Lists all library files in a given folder path
func listLibrariesInFolder(dirName string) []string {
	var out = make([]string, 0)
	_, err0 := os.Stat(dirName)
	if err0 == nil {
		lst, err1 := ioutil.ReadDir(dirName)
		if err1 == nil {
			for _,file := range lst {
				if file.IsDir() {
					fullDirPath := dirName + string(os.PathSeparator) + file.Name()
					newList := listLibrariesInFolder(fullDirPath)
					out = append(out, newList...)
				} else {
					if filterByExtension(file.Name()) {
						fullFilePath := dirName + string(os.PathSeparator) + file.Name()
						out = append(out, fullFilePath)
						
					}
				}
			}
		}
	}
	return out
}

// Recovers plugins and submits plugins to given callback function
func forEachPluginDo(callback func([]model.Plugin)())  {
	var pluginsList = make([]model.Plugin, 0)
	dirName := getDefaultPluginsFolder()
	_, err0 := os.Stat(dirName)
	if err0 == nil {
		libraries := listLibrariesInFolder(dirName)
		for _,libraryFullPath := range libraries {
			if Logger != nil {
				Logger.Debugf("modules.proxy.forEachSenderInPlugins() -> Loading help from library: %s", libraryFullPath)
			}
			pluginObj, err := plugin.Open(libraryFullPath)
			if err == nil {
				sym, err2 := pluginObj.Lookup("GetPluginProxy")
				if err2 != nil {
					module := sym.(func() model.PluginProxy)()
					plugins, err := module.GetPlugins()
					if err != nil {
						if Logger != nil {
							Logger.Errorf("Errors initializing plugin library: %s, details: %v", libraryFullPath, err)
						} else {
							fmt.Printf("Errors initializing plugin library: %s, details: %v\n", libraryFullPath, err)
						}
					} else {
						pluginsList = append(pluginsList, plugins...)
					}
				}
			}
		}
	}
	callback(pluginsList)
}

// Local Plugins cache variable
var pluginDataList = make([]model.Plugin, 0)

// Creates a New Proxy filled wit all available Built-In and Custom Modules, loaded just on first call
func NewProxy() Proxy {
	if len(pluginDataList) == 0 {
		pluginDataList = getPlugins()
	}
	return &proxy{
		plugins: pluginDataList,
	}
}

// Initializes the plugins proxy system
func InitPlugins(config model.PluginsConfig, logger log.Logger) {
	BaseAppFolder = config.BaseFolderPath
	ForcePluginsReload = config.ForcePluginsReload
	Logger = logger
	if runtime.GOOS == `windows` {
		PluginLibrariesExtension = "dll"
	} else {
		PluginLibrariesExtension = "so"
	}
	if len(pluginDataList) == 0 || ForcePluginsReload {
		pluginDataList = getPlugins()
	}
}

// Calculates the plugins module folder
func getDefaultPluginsFolder() string {
	if strings.TrimSpace(BaseAppFolder) != "" {
		return filepath.Dir(strings.TrimSpace(BaseAppFolder ))
	}
	execPath, err := os.Executable()
	if err != nil {
		pwd, errPwd := os.Getwd()
		if errPwd != nil {
			return filepath.Dir(".") + string(os.PathSeparator) + PluginsModulesFolder
		}
		return filepath.Dir(pwd) + string(os.PathSeparator) + PluginsModulesFolder
	}
	f, _ := filepath.Split(execPath)
	return filepath.Dir(f) + string(os.PathSeparator) + PluginsModulesFolder
}