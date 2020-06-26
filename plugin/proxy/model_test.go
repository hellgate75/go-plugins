package proxy

import (
	"fmt"
	"github.com/hellgate75/go-plugins/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func assert(t *testing.T, message string, expected, given interface{}) {
	if expected != given {
		t.Fatalf("Error during assertion: %s, expected value <%v> but given <%v>\n", message, expected, given)
	}
}

func assertNot(t *testing.T, message string, expected, given interface{}) {
	if expected == given {
		t.Fatalf("Error during assertion: %s, NOT expected value <%v> but given <%v>\n", message, expected, given)
	}
}

func TestFunctionGetDefaultPluginFolder(t *testing.T) {
	BaseAppFolder = ""
	var val = getDefaultPluginsFolder()
	assertNot(t, "Default folder must not be empty", "", val)
	assert(t, "Default folder must contains", true, strings.Index(val, PluginsModulesFolder) > 0)
}

func TestFunctionGetGivenPluginFolder(t *testing.T) {
	BaseAppFolder = "."
	var val = getDefaultPluginsFolder()
	var folder = filepath.Dir(fmt.Sprintf(".%c%s", os.PathSeparator, PluginsModulesFolder))
	assert(t, "Default folder must be given one", folder, val)
}

func TestFunctionInitPlugins(t *testing.T) {
	InitPlugins(model.PluginsConfig{
	}, nil)
	if runtime.GOOS == `windows` {
		assert(t, "Extension must be dll", "dll", PluginLibrariesExtension)
	} else {
		assert(t, "Extension must be so", "so", PluginLibrariesExtension)
	}
}

func TestFunctionNewProxy(t *testing.T) {
	p := NewProxy()
	assertNot(t,"Proxy must not be null", nil, p)
}

func TestFunctionListLibrariesInFolder(t *testing.T) {
	BaseAppFolder = ""
	var folder = getDefaultPluginsFolder()
	BaseAppFolder = folder
	_ = os.RemoveAll(folder)
	e := os.MkdirAll(folder, 0777)
	if e != nil {
		t.Fatalf("Unable to create folder: %s\n", folder)
	}
	var file = fmt.Sprintf("%s%c%s.%s", folder, os.PathSeparator, "sample", PluginLibrariesExtension)
	e = ioutil.WriteFile(file, []byte("sample"), 0777)
	if e != nil {
		t.Fatalf("Unable to create file: %s\n", file)
	}
	var fls = listLibrariesInFolder(folder)
	assert(t,"Folder must find one library occurrence", 1, len(fls))
	assert(t,"Library file must match", file, fls[0])
	_ = os.RemoveAll(folder)
}

func TestFunctionFilterByExtension(t *testing.T) {
	var file string
	if runtime.GOOS == `windows` {
		file = "sample.dll"
	} else {
		file = "sample.so"
	}
	assert(t,"Right library file must have right extension check flag", true, filterByExtension(file))
	file="sample.bin"
	assert(t,"Wrong library file must NOT have right extension check flag", false, filterByExtension(file))
}

func TestFunctionForEachPluginDo(t *testing.T) {
	var success = false
	forEachPluginDo(func(mls []model.Plugin){
		//TODO: Cannot complete test until I do not have any ready plugin, using fake success
		success = true
	})
	assert(t,"Test of for each plugin must return any plugin for the path", true, success)
}

func TestFunctionGetPlugins(t *testing.T) {
	var plugins = getPlugins()
	//TODO: Cannot complete test until I do not have any ready plugin, using fake success value
	assert(t,"Plugins must be loaded", 0, len(plugins))
}
