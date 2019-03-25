// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package main

import (
	"github.com/spf13/afero"

	"github.com/nettorta/pandora/cli"
	"github.com/nettorta/pandora/components/example/import"
	"github.com/nettorta/pandora/components/phttp/import"
	"github.com/nettorta/pandora/core/import"
)

func main() {
	// CLI don't know anything about components initially.
	// All extpoints constructors and default configurations should be registered, before CLI run.
	fs := afero.NewOsFs()
	coreimport.Import(fs)
	phttp.Import(fs)
	example.Import()

	cli.Run()
}
