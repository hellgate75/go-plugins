package proxy

import (
	"errors"
	"fmt"
	"github.com/hellgate75/go-plugins/model"
)

type proxy struct {
	plugins		[]model.Plugin
}

func (p *proxy) GetPlugins() []*model.Plugin {
	var out = make([]*model.Plugin, 0)
	for _, p := range p.plugins {
		out = append(out, &p)
	}
	return out
}

func (p *proxy) GetPluginsByType(pt model.PluginType) ([]*model.Plugin, error) {
	var err error
	defer func(){
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("PluginProxy.GetPluginsByType() - Error: %v", r))
		}
	}()
	var out = make([]*model.Plugin, 0)
	for idx, pl := range p.plugins {
		if Logger != nil {
			Logger.Tracef("PluginProxy.GetPluginsByType() - Evaluating plugin: %v", pl)
		}
		if byte(pl.PluginType()) == byte(pt) {
			fmt.Printf("Found with type: %v\n", pl.PluginType())
			out = append(out, &(p.plugins[idx]))
		}
	}
	return out, err
}

func (p *proxy) GetFilteredPluginsByType(pt model.PluginType, filter func(model.Plugin) bool) ([]*model.Plugin, error) {
	var err error
	defer func(){
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("PluginProxy.GetPluginsByType() - Error: %v", r))
		}
	}()
	var out = make([]*model.Plugin, 0)
	lst, errL := p.GetPluginsByType(pt)
	if errL != nil {
		return out, errL
	}
	for _, p := range lst {
		if filter(*p) {
			out = append(out, p)
		}
	}
	return out, err
}


func (p *proxy) GetFilteredPlugins(filter func(model.Plugin) bool) ([]*model.Plugin, error) {
	var err error
	defer func(){
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("PluginProxy.GetPluginsByType() - Error: %v", r))
		}
	}()
	var out = make([]*model.Plugin, 0)
	for _, p := range p.plugins {
		if filter(p) {
			out = append(out, &p)
		}
	}
	return out, err
}

