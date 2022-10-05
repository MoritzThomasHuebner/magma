/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reindex_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"magma/orc8r/cloud/go/clock"
	configurator_test_init "magma/orc8r/cloud/go/services/configurator/test_init"
	configurator_test "magma/orc8r/cloud/go/services/configurator/test_utils"
	device_test_init "magma/orc8r/cloud/go/services/device/test_init"
	directoryd_types "magma/orc8r/cloud/go/services/directoryd/types"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/cloud/go/services/state/indexer"
	"magma/orc8r/cloud/go/services/state/indexer/reindex"
	state_test_init "magma/orc8r/cloud/go/services/state/test_init"
	state_test "magma/orc8r/cloud/go/services/state/test_utils"
)

func TestSingletonRunSuccess00(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail00(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess01(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail01(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess02(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail02(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess03(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail03(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess04(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail04(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess05(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail05(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess06(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail06(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess07(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail07(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess08(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail08(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess09(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail09(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess10(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail10(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess11(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail11(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess12(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail12(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess13(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail13(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess14(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail14(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess15(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail15(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess16(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail16(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess17(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail17(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess18(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail18(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess19(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail19(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess20(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail20(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess21(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail21(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess22(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail22(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess23(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail23(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess24(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail24(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess25(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail25(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess26(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail26(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess27(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail27(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess28(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail28(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess29(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail29(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess30(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail30(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess31(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail31(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess32(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail32(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess33(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail33(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess34(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail34(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess35(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail35(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess36(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail36(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess37(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail37(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess38(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail38(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess39(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail39(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess40(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail40(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess41(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail41(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess42(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail42(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess43(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail43(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess44(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail44(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess45(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail45(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess46(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail46(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess47(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail47(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess48(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail48(t *testing.T)    { TestSingletonRunFail(t) }
func TestSingletonRunSuccess49(t *testing.T) { TestSingletonRunSuccess(t) }
func TestSingletonRunFail49(t *testing.T)    { TestSingletonRunFail(t) }

func TestSingletonRunSuccess(t *testing.T) {
	// Make nullimpotent calls to handle code coverage indeterminacy
	reindex.TestHookReindexSuccess()
	reindex.TestHookReindexDone()

	// Writes to channel after completing a job
	reindexSuccessNum, reindexDoneNum := 0, 0
	ch := make(chan interface{})

	reindex.TestHookReindexSuccess = func() {
		reindexSuccessNum += 1
	}
	defer func() { reindex.TestHookReindexSuccess = func() {} }()

	reindex.TestHookReindexDone = func() {
		reindexDoneNum += 1
		ch <- nil
	}
	defer func() { reindex.TestHookReindexDone = func() {} }()

	clock.SkipSleeps(t)
	defer clock.ResumeSleeps(t)

	r := initSingletonReindexTest(t)

	ctx, cancel := context.WithCancel(context.Background())
	go r.Run(ctx)

	// Single indexer
	idx0 := getIndexer(id0, zero, version0, true)
	idx0.On("GetTypes").Return(allTypes).Once()
	// Register indexers
	register(t, idx0)

	// Check
	recvCh(t, ch)
	recvNoCh(t, ch)

	idx0.AssertExpectations(t)
	require.Equal(t, 1, reindexSuccessNum)
	require.Equal(t, 1, reindexDoneNum)

	// Bump existing indexer version
	idx0a := getIndexerNoIndex(id0, version0, version0a, false)
	idx0a.On("GetTypes").Return(gwStateType).Once()
	idx0a.On("Index", mock.Anything, mock.Anything).Return(nil, nil).Times(nNetworks)
	// Register indexers
	register(t, idx0a)

	// Check
	recvCh(t, ch)
	recvNoCh(t, ch)

	idx0a.AssertExpectations(t)
	require.Equal(t, 2, reindexSuccessNum)
	require.Equal(t, 2, reindexDoneNum)

	// Test that a network/hardware pair that has been added after Run
	// will have its states reindexed as well

	// reportAdditionalState reports enough directory records to cause 3 batches per network
	// (with the +1 gateway status per network). It adds an extra network from 3 -> 4,
	// so numBatches following this method will be 3 * 4 = 12
	reportAdditionalState(t, nid3, hwid3, 3)

	idx5 := getIndexerNoIndex(id5, zero, version5, true)
	idx5.On("GetTypes").Return(allTypes).Once()
	idx5.On("Index", mock.Anything, mock.Anything).Return(nil, nil).Times(newNBatches)
	// Register indexers
	register(t, idx5)

	// Check
	recvCh(t, ch)
	recvNoCh(t, ch)

	idx5.AssertExpectations(t)
	require.Equal(t, 3, reindexSuccessNum)
	require.Equal(t, 3, reindexDoneNum)
	cancel()
	select {
	case <-ctx.Done():
		indexer.DeregisterAllForTest(t)
	case <-time.After(defaultTestTimeout):
		indexer.DeregisterAllForTest(t)
		t.Fatal("Timed out waiting for context to cancel.")
	}
}

func TestSingletonRunFail(t *testing.T) {
	// Make nullimpotent calls to handle code coverage indeterminacy
	reindex.TestHookReindexSuccess()
	reindex.TestHookReindexDone()

	// Writes to channel after completing a job
	reindexSuccessNum, reindexDoneNum := 0, 0
	ch := make(chan interface{})

	reindex.TestHookReindexSuccess = func() {
		reindexSuccessNum += 1
	}
	defer func() { reindex.TestHookReindexSuccess = func() {} }()

	reindex.TestHookReindexDone = func() {
		ch <- nil
		reindexDoneNum += 1
	}
	defer func() { reindex.TestHookReindexDone = func() {} }()

	clock.SkipSleeps(t)
	defer clock.ResumeSleeps(t)

	r := initSingletonReindexTest(t)
	reportAdditionalState(t, nid3, hwid3, 3)

	// Indexer returns err => reindex jobs fail

	// Fail1 at PrepareReindex
	fail1 := getBasicIndexer(id1, version1)
	fail1.On("GetTypes").Return(allTypes).Once()
	fail1.On("PrepareReindex", zero, version1, true).Return(someErr1).Once()

	// Fail2 at first Reindex
	fail2 := getBasicIndexer(id2, version2)
	fail2.On("GetTypes").Return(allTypes).Once()
	fail2.On("PrepareReindex", zero, version2, true).Return(nil).Once()
	fail2.On("Index", mock.Anything, mock.Anything).Return(nil, someErr2).Once()

	// Fail3 at CompleteReindex
	fail3 := getBasicIndexer(id3, version3)
	fail3.On("GetTypes").Return(allTypes).Once()
	fail3.On("PrepareReindex", zero, version3, true).Return(nil).Once()
	fail3.On("Index", mock.Anything, mock.Anything).Return(nil, nil).Times(newNBatches)
	fail3.On("CompleteReindex", zero, version3).Return(someErr3).Once()

	// Register indexers
	register(t, fail1, fail2, fail3)

	// Only run the reindexer after all indexers have been registered.
	// This is to ensure that the reindexer will not start with only one
	// or two indexers registered, which would cause the reindexer to
	// run four or five times, respectively.
	ctx, cancel := context.WithCancel(context.Background())
	go r.Run(ctx)

	// Check
	recvCh(t, ch)
	recvCh(t, ch)
	recvCh(t, ch)
	cancel()
	select {
	case <-ctx.Done():
		indexer.DeregisterAllForTest(t)
		recvNoCh(t, ch)
		fail1.AssertExpectations(t)
		fail2.AssertExpectations(t)
		fail3.AssertExpectations(t)
		require.Equal(t, 0, reindexSuccessNum)
		require.Equal(t, 3, reindexDoneNum)
	case <-time.After(defaultTestTimeout):
		indexer.DeregisterAllForTest(t)
		t.Fatal("Timed out waiting for context to cancel.")
	}
}

// initSingletonReindexTest reports enough directory records to cause 3 batches per network
// (with the +1 gateway status per network). It creates 3 networks,
// so numBatches following this method will be 3 * 3 = 9
func initSingletonReindexTest(t *testing.T) reindex.Reindexer {
	indexer.DeregisterAllForTest(t)

	configurator_test_init.StartTestService(t)
	device_test_init.StartTestService(t)

	reindexer := state_test_init.StartTestSingletonServiceInternal(t)

	// Report enough directory records to cause 3 batches per network (with the +1 gateway status per network)
	reportAdditionalState(t, nid0, hwid0, 0)
	reportAdditionalState(t, nid1, hwid1, 1)
	reportAdditionalState(t, nid2, hwid2, 2)
	return reindexer
}

// reportMoreState reports enough directory records to cause 3 batches per network
// (with the +1 gateway status per network). It adds an extra network from 3 -> 4,
// so numBatches following this method will be 3 * 4 = 12
func reportAdditionalState(t *testing.T, nid string, hwid string, networkNumber int) {
	configurator_test.RegisterNetwork(t, nid, fmt.Sprintf("Network %v for reindex test", networkNumber))
	configurator_test.RegisterGateway(t, nid, hwid, &models.GatewayDevice{HardwareID: hwid})

	ctx := state_test.GetContextWithCertificate(t, hwid)

	var records []*directoryd_types.DirectoryRecord
	var deviceIDs []string
	for i := 0; i < directoryRecordsPerNetwork; i++ {
		hwidStr := fmt.Sprintf("hwid%d", i)
		imsiStr := fmt.Sprintf("imsiStr%d", i)
		records = append(records, &directoryd_types.DirectoryRecord{LocationHistory: []string{hwidStr}})
		deviceIDs = append(deviceIDs, imsiStr)
	}
	reportDirectoryRecord(t, ctx, deviceIDs, records)

	// Report one gateway status per network
	gwStatus := &models.GatewayStatus{Meta: map[string]string{"foo": "bar"}}
	reportGatewayStatus(t, ctx, gwStatus)
}
