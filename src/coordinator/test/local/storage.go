// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package local

import (
	"testing"
	"time"

	"github.com/m3db/m3db/src/coordinator/storage"
	"github.com/m3db/m3db/src/coordinator/storage/local"
	"github.com/m3db/m3db/src/dbnode/client"
	"github.com/m3db/m3x/ident"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	// TestNamespaceID is the namespace of the test unaggregated namespace
	// used by local storage.
	TestNamespaceID = "metrics"
	// TestRetention is the retention of the test unaggregated namespace
	// used by local storage.
	TestRetention = 30 * 24 * time.Hour
)

// NewStorageAndSession generates a new local storage and mock session
func NewStorageAndSession(
	t *testing.T,
	ctrl *gomock.Controller,
) (storage.Storage, *client.MockSession) {
	session := client.NewMockSession(ctrl)
	clusters, err := local.NewClusters(local.UnaggregatedClusterNamespaceDefinition{
		NamespaceID: ident.StringID(TestNamespaceID),
		Session:     session,
		Retention:   TestRetention,
	})
	require.NoError(t, err)
	storage := local.NewStorage(clusters, nil)
	return storage, session
}
