// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package internal

var mapExcludeAddrs = map[string]bool{
	// TODO: Hardcode exclude address for disable receive funds
	//"anam19ad4tprcf9ew4577qph3jfzpf9slcrkpmxwvah": true,
}

func MergeExcludeAddrs(m map[string]bool) map[string]bool {
	for k, v := range mapExcludeAddrs {
		m[k] = v
	}
	return m
}
