// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gateway

type gatewayOption func(gate *YockdGateWay) error

func OptionPolicy() gatewayOption {
	return func(gate *YockdGateWay) error {
		return nil
	}
}
