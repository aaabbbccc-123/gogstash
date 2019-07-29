// Code generated by Protoconf (https://version.uuzu.com/Merlion/protoconf). DO NOT EDIT.
// Code generated at 26 Jul 19 19:05 +08, with protoconf version 0.2.1
// This code is designed to work with protoconf-go (https://github.com/yoozoo/protoconf-go) v1.0

package prometheus

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"

	protoconfGo "github.com/yoozoo/protoconf_go"
)

type watchListener func(string)

type _watchKeyInterface interface {
	watchKey(key string, listener watchListener, isLeaf bool)
}

type _watchNodeStruct struct {
	listener watchListener
	key      string
}

type Config struct {
	address     string
	app_configs map[string]*AppConfig

	_keys          []string
	_watchLeafList map[string]watchListener
	_watchNodeList []*_watchNodeStruct
	_prefix        string
	_locker        sync.Mutex
}

var _ protoconfGo.Configuration = (*Config)(nil)

// make sure strconv is used
var _, _ = strconv.Atoi("-42")

func newConfig() *Config {
	p := new(Config)
	p.address = "8080"
	p.app_configs = make(map[string]*AppConfig)
	p._keys = []string{
		"address",
		"app_configs/MAP_ENTRY/app_name",
		"app_configs/MAP_ENTRY/metrics/MAP_ENTRY/filters/MAP_ENTRY/field",
		"app_configs/MAP_ENTRY/metrics/MAP_ENTRY/filters/MAP_ENTRY/regex",
		"app_configs/MAP_ENTRY/metrics/MAP_ENTRY/metric_name",
		"app_configs/MAP_ENTRY/metrics/MAP_ENTRY/metric_type",
	}
	p._watchLeafList = make(map[string]watchListener)
	return p
}

var instance *Config = newConfig()

func GetInstance() *Config {
	return instance
}

func (p *Config) ApplicationName() string {
	return "北美平台-gogstash"
}

func (p *Config) ValidKeys() []string {
	return p._keys
}

func (p *Config) watchKey(key string, listener watchListener, isLeaf bool) {
	p._locker.Lock()
	defer p._locker.Unlock()

	if isLeaf {
		p._watchLeafList[key] = listener
	} else {
		p._watchNodeList = append(p._watchNodeList, nil)
		i := sort.Search(len(p._watchNodeList), func(i int) bool { return p._watchNodeList[i] == nil || len(p._watchNodeList[i].key) < len(key) })
		copy(p._watchNodeList[i+1:], p._watchNodeList[i:])
		p._watchNodeList[i] = &_watchNodeStruct{listener: listener, key: key}
	}
}

func (p *Config) NotifyValueChange(key string, value string) {
	p._locker.Lock()
	defer p._locker.Unlock()

	listener, ok := p._watchLeafList[key]
	if !ok {
		for _, n := range p._watchNodeList {
			if strings.HasPrefix(key, n.key) {
				listener = n.listener
				ok = true
				break
			}
		}
	}
	if ok {
		p.SetValue(key, value)
		listener(value)
	}
}

func (p *Config) DeleteKey(key string) {
	p._locker.Lock()
	defer p._locker.Unlock()

	for _, n := range p._watchNodeList {
		if strings.HasPrefix(key, n.key) {
			if p.DeleteMapObject(key) {
				n.listener("")
			}
			return
		}
	}
}

func (p *Config) Lock() {
	p._locker.Lock()
}
func (p *Config) Unlock() {
	p._locker.Unlock()
}

func (p *Config) GetAddress() string {
	return p.address
}
func (p *Config) GetApp_configs() map[string]*AppConfig {
	return p.app_configs
}
func (p *Config) WatchApp_configs(l watchListener) {
	p.watchKey("app_configs", l, false)
}

func (p *Config) DeleteMapObject(key string) bool {
	keys := strings.SplitN(key, "/", 3)
	if len(keys) < 2 {
		return false
	}

	switch keys[0] {
	case "app_configs":
		if len(keys) < 3 {
			return false
		}
		if obj, ok := p.app_configs[keys[1]]; ok {
			if !obj.DeleteMapObject(keys[2]) {
				delete(p.app_configs, keys[1])
			}
		}
		return true
	}
	return false
}

func (p *Config) DefaultValue(key string) *string {
	switch key {
	case "address":
		tmp := "8080"
		return &tmp
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "app_configs":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					if obj, ok := p.app_configs[keys[0]]; ok {
						return obj.DefaultValue(keys[1])
					}
				}
			}
		}
	}
	return nil
}

func (p *Config) SetValue(key string, value string) error {
	switch key {
	case "address":
		p.address = value
		return nil
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "app_configs":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					obj, ok := p.app_configs[keys[0]]
					if !ok {
						obj = NewAppConfig(p, p._prefix+"app_configs/"+keys[0])
						p.app_configs[keys[0]] = obj
					}
					return obj.SetValue(keys[1], value)
				}
			}
		}
	}
	return errors.New("Unknown key:" + key)
}

type Filter struct {
	field string
	regex string

	_parent _watchKeyInterface
	_prefix string
}

func NewFilter(parent _watchKeyInterface, prefix string) *Filter {
	p := new(Filter)

	p._parent = parent
	p._prefix = prefix + "/"
	p.field = "field name"
	p.regex = ""

	return p
}

func (p *Filter) watchKey(key string, l watchListener, isLeaf bool) {
	p._parent.watchKey(p._prefix+key, l, isLeaf)
}

func (p *Filter) GetField() string {
	return p.field
}
func (p *Filter) GetRegex() string {
	return p.regex
}

func (p *Filter) DeleteMapObject(key string) bool {
	keys := strings.SplitN(key, "/", 3)
	if len(keys) < 2 {
		return false
	}

	switch keys[0] {
	}
	return false
}

func (p *Filter) DefaultValue(key string) *string {
	switch key {
	case "field":
		tmp := "field name"
		return &tmp
	case "regex":
		tmp := ""
		return &tmp
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			}
		}
	}
	return nil
}

func (p *Filter) SetValue(key string, value string) error {
	switch key {
	case "field":
		p.field = value
		return nil
	case "regex":
		p.regex = value
		return nil
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			}
		}
	}
	return errors.New("Unknown key:" + key)
}

type Metric struct {
	metric_name string
	metric_type int64
	filters     map[string]*Filter

	_parent _watchKeyInterface
	_prefix string
}

func NewMetric(parent _watchKeyInterface, prefix string) *Metric {
	p := new(Metric)

	p._parent = parent
	p._prefix = prefix + "/"
	p.metric_name = "metric name"
	p.metric_type, _ = strconv.ParseInt("0", 10, 64)
	p.filters = make(map[string]*Filter)

	return p
}

func (p *Metric) watchKey(key string, l watchListener, isLeaf bool) {
	p._parent.watchKey(p._prefix+key, l, isLeaf)
}

func (p *Metric) GetMetric_name() string {
	return p.metric_name
}
func (p *Metric) GetMetric_type() int64 {
	return p.metric_type
}
func (p *Metric) GetFilters() map[string]*Filter {
	return p.filters
}

func (p *Metric) DeleteMapObject(key string) bool {
	keys := strings.SplitN(key, "/", 3)
	if len(keys) < 2 {
		return false
	}

	switch keys[0] {
	case "filters":
		if len(keys) < 3 {
			return false
		}
		if obj, ok := p.filters[keys[1]]; ok {
			if !obj.DeleteMapObject(keys[2]) {
				delete(p.filters, keys[1])
			}
		}
		return true
	}
	return false
}

func (p *Metric) DefaultValue(key string) *string {
	switch key {
	case "metric_name":
		tmp := "metric name"
		return &tmp
	case "metric_type":
		tmp := "0"
		return &tmp
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "filters":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					if obj, ok := p.filters[keys[0]]; ok {
						return obj.DefaultValue(keys[1])
					}
				}
			}
		}
	}
	return nil
}

func (p *Metric) SetValue(key string, value string) error {
	switch key {
	case "metric_name":
		p.metric_name = value
		return nil
	case "metric_type":
		tmp, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		p.metric_type = tmp
		return nil
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "filters":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					obj, ok := p.filters[keys[0]]
					if !ok {
						obj = NewFilter(p, p._prefix+"filters/"+keys[0])
						p.filters[keys[0]] = obj
					}
					return obj.SetValue(keys[1], value)
				}
			}
		}
	}
	return errors.New("Unknown key:" + key)
}

type AppConfig struct {
	app_name string
	metrics  map[string]*Metric

	_parent _watchKeyInterface
	_prefix string
}

func NewAppConfig(parent _watchKeyInterface, prefix string) *AppConfig {
	p := new(AppConfig)

	p._parent = parent
	p._prefix = prefix + "/"
	p.app_name = "app name"
	p.metrics = make(map[string]*Metric)

	return p
}

func (p *AppConfig) watchKey(key string, l watchListener, isLeaf bool) {
	p._parent.watchKey(p._prefix+key, l, isLeaf)
}

func (p *AppConfig) GetApp_name() string {
	return p.app_name
}
func (p *AppConfig) GetMetrics() map[string]*Metric {
	return p.metrics
}

func (p *AppConfig) DeleteMapObject(key string) bool {
	keys := strings.SplitN(key, "/", 3)
	if len(keys) < 2 {
		return false
	}

	switch keys[0] {
	case "metrics":
		if len(keys) < 3 {
			return false
		}
		if obj, ok := p.metrics[keys[1]]; ok {
			if !obj.DeleteMapObject(keys[2]) {
				delete(p.metrics, keys[1])
			}
		}
		return true
	}
	return false
}

func (p *AppConfig) DefaultValue(key string) *string {
	switch key {
	case "app_name":
		tmp := "app name"
		return &tmp
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "metrics":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					if obj, ok := p.metrics[keys[0]]; ok {
						return obj.DefaultValue(keys[1])
					}
				}
			}
		}
	}
	return nil
}

func (p *AppConfig) SetValue(key string, value string) error {
	switch key {
	case "app_name":
		p.app_name = value
		return nil
	default:
		keys := strings.SplitN(key, "/", 2)
		if len(keys) >= 2 {
			newKey := keys[0]
			switch newKey {
			case "metrics":
				keys = strings.SplitN(keys[1], "/", 2)
				if len(keys) >= 2 {
					obj, ok := p.metrics[keys[0]]
					if !ok {
						obj = NewMetric(p, p._prefix+"metrics/"+keys[0])
						p.metrics[keys[0]] = obj
					}
					return obj.SetValue(keys[1], value)
				}
			}
		}
	}
	return errors.New("Unknown key:" + key)
}
