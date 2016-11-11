/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import "github.com/hyperledger/fabric/gossip/api"

// PKIidType defines the type that holds the PKI-id
// which is the security identifier of a peer
type PKIidType []byte

// MessageAcceptor is a predicate that is used to
// determine in which messages the subscriber that created the
// instance of the MessageAcceptor is interested in.
type MessageAcceptor func(interface{}) bool

// Payload defines an object that contains a ledger block
type Payload struct {
	ChainID api.ChainID // The channel's ID of the block
	Data    []byte      // The content of the message, possibly encrypted or signed
	Hash    string      // The message hash
	SeqNum  uint64      // The message sequence number
}
