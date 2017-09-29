// Copyright (c) 2017 Pantheon technologies s.r.o.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//Package sync_test contains tests for sync utilities
package sync_test

import (
	"github.com/stretchr/testify/assert"
	"pantheon.tech/ligato-bgp/agent/utils/sync"
	"testing"
	"time"
)

const expected = uint32(1)

func TestSynchronizer(t *testing.T) {
	assert.New(t).Nil(initializeCounterAndRunSync(500 * time.Millisecond))
}

func TestSynchronizerFail(t *testing.T) {
	assert.New(t).NotNil(initializeCounterAndRunSync(2 * time.Second))
}

func initializeCounterAndRunSync(d time.Duration) error {
	counter := uint32(0)

	go func(i *uint32) {
		time.Sleep(d)
		counter++
	}(&counter)
	return sync.WaitCounterMatch(1*time.Second, expected, &counter)
}
