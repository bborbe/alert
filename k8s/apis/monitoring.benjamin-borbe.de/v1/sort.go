// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import "strings"

type AlertsSorted []Alert

func (b AlertsSorted) Len() int {
	return len(b)
}

func (b AlertsSorted) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b AlertsSorted) Less(i, j int) bool {
	return CompareAlert(b[i], b[j]) < 0
}

func CompareAlert(a, b Alert) int {
	if result := strings.Compare(a.Namespace, b.Namespace); result != 0 {
		return result
	}
	if result := strings.Compare(a.Name, b.Name); result != 0 {
		return result
	}
	if result := strings.Compare(a.Spec.Name, b.Spec.Name); result != 0 {
		return result
	}
	if result := strings.Compare(a.Spec.For, b.Spec.For); result != 0 {
		return result
	}
	if result := strings.Compare(a.Spec.Expression, b.Spec.Expression); result != 0 {
		return result
	}
	return 0
}
