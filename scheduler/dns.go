package scheduler

import (
	"errors"

	"github.com/ansurfen/cushion/utils"
	"github.com/spf13/viper"
)

type asset struct {
	// url is concrete address of fetching assest
	URL string `json:"url"`
	// path points driver address in localhost, which be used to adapt local dns
	Path string `json:"path"`
}

type DNS struct {
	Driver  map[string]asset `json:"driver"`
	Plugin  map[string]asset `json:"plugin"`
	Version string           `json:"version"`
	file    *viper.Viper     `json:"-"`
}

const dnsBlank = `
{
	"driver": {},
	"plugin": {},
	"version": ""
}
`

func CreateDNS(path string) *DNS {
	utils.SafeWriteFile(path, []byte(dnsBlank))
	return OpenDNS(path)
}

func OpenDNS(path string) *DNS {
	v := viper.NewWithOptions(viper.KeyDelimiter("###"))
	v.SetConfigFile(path)
	dns := &DNS{
		file:   v,
		Plugin: make(map[string]asset),
		Driver: make(map[string]asset),
	}
	if err := dns.file.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := dns.file.Unmarshal(dns); err != nil {
		panic(err)
	}
	return dns
}

func (dns *DNS) GetPlugin(domain string) asset {
	return dns.Plugin[domain]
}

func (dns *DNS) PutPlugin(domain, url, path string) error {
	if _, ok := dns.Plugin[domain]; ok {
		return errors.New("plugin exist already")
	}
	dns.Plugin[domain] = asset{URL: url, Path: path}
	dns.file.Set("plugin", dns.Plugin)
	dns.file.WriteConfig()
	return nil
}

func (dns *DNS) UnputPlugin(domain string) {
	delete(dns.Plugin, domain)
}

func (dns *DNS) UpdatePlugin(domain, url, path string) {
	dns.Plugin[domain] = asset{URL: url, Path: path}
	dns.file.Set("plugin", dns.Plugin)
	dns.file.WriteConfig()
}

func (dns *DNS) GetDriver(domain string) asset {
	return dns.Driver[domain]
}

func (dns *DNS) PutDriver(domain, url, path string) error {
	if _, ok := dns.Driver[domain]; ok {
		return errors.New("domain exist already")
	}
	dns.Driver[domain] = asset{URL: url, Path: path}
	dns.file.Set("driver", dns.Driver)
	dns.file.WriteConfig()
	return nil
}

func (dns *DNS) UpdateDriver(domain, url, path string) {
	dns.Driver[domain] = asset{URL: url, Path: path}
	dns.file.Set("driver", dns.Driver)
	dns.file.WriteConfig()
}

func (dns *DNS) UnputDriver(domain string) {
	delete(dns.Driver, domain)
}

func (dns *DNS) AliasDriver(domain, alias string) error {
	if len(dns.GetDriver(alias).URL) > 0 {
		return errors.New("alias exist already")
	}
	if driver := dns.GetDriver(domain); len(driver.URL) > 0 {
		dns.PutDriver(alias, driver.URL, driver.Path)
	}
	return nil
}

func (dns *DNS) AliasPlugin(domain, alias string) error {
	if len(dns.GetPlugin(alias).URL) > 0 {
		return errors.New("alias exist already")
	}
	if plugin := dns.GetPlugin(domain); len(plugin.URL) > 0 {
		dns.PutPlugin(alias, plugin.URL, plugin.Path)
	}
	return nil
}
