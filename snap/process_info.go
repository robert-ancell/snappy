// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2019 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package snap

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/snapcore/snapd/sandbox/cgroup"
)

var cgroupProcGroup = cgroup.ProcGroup

func NameFromPid(pid int) (string, error) {
	if cgroup.IsUnified() {
		// not supported
		return "", fmt.Errorf("not supported")
	}

	group, err := cgroupProcGroup(pid, cgroup.MatchV1Controller("freezer"))
	if err != nil {
		return "", fmt.Errorf("cannot determine cgroup path of pid %v: %v", pid, err)
	}

	if strings.HasPrefix(group, "/snap.") {
		snap := strings.SplitN(filepath.Base(group), ".", 2)[1]
		return snap, nil
	}

	return "", fmt.Errorf("cannot find a snap for pid %v", pid)
}
