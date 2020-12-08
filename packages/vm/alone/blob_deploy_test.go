// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0
package alone

import (
	"github.com/iotaledger/wasp/packages/vm/builtinvm/blob"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlobRepeatInit(t *testing.T) {
	e := New(t, false, false)
	req := NewCall(blob.Interface.Name, "init")
	_, err := e.PostRequest(req, nil)
	require.Error(t, err)
}

func TestBlobUpload(t *testing.T) {
	al := New(t, false, true)
	binary := []byte("supposed to be wasm")
	hwasm, err := al.UploadWasm(nil, binary)
	require.NoError(t, err)

	binBack, err := al.GetWasmBinary(hwasm)
	require.NoError(t, err)

	require.EqualValues(t, binary, binBack)
}

func TestBlobUploadTwice(t *testing.T) {
	al := New(t, false, false)
	binary := []byte("supposed to be wasm")
	hwasm1, err := al.UploadWasm(nil, binary)
	require.NoError(t, err)

	hwasm2, err := al.UploadWasm(nil, binary)
	require.NoError(t, err)

	require.EqualValues(t, hwasm1, hwasm2)

	binBack, err := al.GetWasmBinary(hwasm1)
	require.NoError(t, err)

	require.EqualValues(t, binary, binBack)
}

const wasmFile = "../../../tools/cluster/tests/wasptest_new/wasm/inccounter_bg.wasm"

func TestDeploy(t *testing.T) {
	al := New(t, false, true)
	hwasm, err := al.UploadWasmFromFile(nil, wasmFile)
	require.NoError(t, err)

	err = al.DeployContract(nil, "testInccounter", hwasm)
	require.NoError(t, err)
}

func TestDeployWasm(t *testing.T) {
	al := New(t, false, true)
	err := al.DeployWasmContract(nil, "testInccounter", wasmFile)
	require.NoError(t, err)
}

func TestDeployRubbish(t *testing.T) {
	al := New(t, false, false)
	name := "testInccounter"
	err := al.DeployWasmContract(nil, name, "blob_deploy_test.go")
	require.Error(t, err)

	_, err = al.FindContract(name)
	require.Error(t, err)
}