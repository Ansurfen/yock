// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "github.com/spf13/viper"

// OpenConfFromPath unmarshal file which located in disk to memory according to path
func OpenConf(path string, opts ...viper.Option) (*viper.Viper, error) {
	conf := viper.NewWithOptions(opts...)
	conf.SetConfigFile(path)
	return conf, conf.ReadInConfig()
}
