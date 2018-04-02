// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package burrowmint

import (
	"fmt"

	version "github.com/hyperledger/burrow/version"
)

const (
	// Client identifier to advertise over the network
	burrowMintClientIdentifier = "burrowmint"
	// Major version component of the current release
	burrowMintVersionMajor = 0
	// Minor version component of the current release
	burrowMintVersionMinor = 17
	// Patch version component of the current release
	burrowMintVersionPatch = 0
)

// Define the compatible consensus engines this application manager
// is compatible and has been tested with.
var compatibleConsensus = [...]string{
	"tendermint-0.9",
}

func GetBurrowMintVersion() *version.VersionIdentifier {
	return version.New(burrowMintClientIdentifier, burrowMintVersionMajor,
		burrowMintVersionMinor, burrowMintVersionPatch)
}

func AssertCompatibleConsensus(consensusMinorVersion string) error {
	for _, supported := range compatibleConsensus {
		if consensusMinorVersion == supported {
			return nil
		}
	}
	return fmt.Errorf("BurrowMint (%s) is not compatible with consensus engine %s",
		GetBurrowMintVersion(), consensusMinorVersion)
}
