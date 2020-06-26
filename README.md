<div style="width: 100% !important; align: right;">

![Go](https://github.com/hellgate75/go-plugins/workflows/Go/badge.svg?branch=master) &nbsp; &nbsp;

</div>

# go-plugins
Plugins management library, it allows to manage plugins and plugins loading.

As library it requires to be imported as following code sample:
```
import (
    plugins log "github.com/myaccount/myproject/myloggerpackage"
    plugins "github.com/hellgate75/go-plugins"
    "github.com/hellgate75/go-plugins/model"
    "fmt"
)

// Enables the plugins.
func init() {
    var myLogger = log.New()
    plugins.EnablePlugins(model.PluginsConfig{
        "my-root-folder",
        true
    }, myLogger)
}

// Collects and count plugins
func main() {
    var pluginsList = plgins.GetPlugins()
    fmt.Printf("Loaded %v plugins", len(pluginsList))
}

```

You need to implement in the library the method ```GetProxyStub()```  it must return an element of type 
``` model.ProxyStub```. Any plugin can be of different type and model information can be found here:

* [model/plugins.go](/model/plugins.go)

Main methods are present in this file:

* [go-plugins.go](/go-plugins.go)

Mandatory method is ```EnablePlugins```. 

The logic is easy and components into the plugin are generic. The ```model.PluginType``` defines the kind of plugin you are expecting from the input plugin libraries.



## Build the project

Build command sample :
```
go build -buildmode=shared github.com/hellgate75/go-plugins
```

## Test the project

```
go test -v github.com/hellgate75/go-plugins/...
```

Enjoy the experience.

## License

The library is licensed with [LGPL v. 3.0](/LICENSE) clauses, with prior authorization of author before any production or commercial use. Use of this library or any extension is prohibited due to high risk of damages due to improper use. No warranty is provided for improper or unauthorized use of this library or any implementation.

Any request can be prompted to the author [Fabrizio Torelli](https://www.linkedin.com/in/fabriziotorelli) at the following email address:

[hellgate75@gmail.com](mailto:hellgate75@gmail.com)
