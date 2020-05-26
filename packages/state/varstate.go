package state

import (
	"bytes"
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/database"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/variables"
	"github.com/mr-tron/base58"
	"io"
)

type variableState struct {
	stateIndex uint32
	empty      bool
	stateHash  hashing.HashValue
	vars       variables.Variables
}

// VariableState

func NewVariableState(varState VariableState) VariableState {
	if varState == nil {
		return &variableState{
			vars:  variables.New(nil),
			empty: true,
		}
	}
	return &variableState{
		stateIndex: varState.StateIndex(),
		stateHash:  varState.Hash(),
		vars:       variables.New(varState.Variables()),
	}
}

func (vs *variableState) StateIndex() uint32 {
	return vs.stateIndex
}

func (vs *variableState) IncStateIndex() {
	vs.stateIndex++
}

// applies batch of state updates. Increases state index
func (vs *variableState) ApplyBatch(batch Batch) error {
	if !vs.empty {
		if batch.StateIndex() != vs.stateIndex+1 {
			return fmt.Errorf("batch state index #%d can't be applied to the state #%d",
				batch.StateIndex(), vs.stateIndex)
		}
	} else {
		if batch.StateIndex() != 0 {
			return fmt.Errorf("batch state index #%d can't be applied to the empty state", batch.StateIndex())
		}
	}
	batch.ForEach(func(stateUpd StateUpdate) bool {
		vs.ApplyStateUpdate(stateUpd)
		return false
	})
	vs.IncStateIndex()
	return nil
}

// applies one state update. Doesn't change state index
func (vs *variableState) ApplyStateUpdate(stateUpd StateUpdate) {
	vs.Variables().Apply(stateUpd.Variables())

	vh := vs.Hash()
	sh := hashing.GetHashValue(stateUpd)
	vs.stateHash = *hashing.HashData(vh.Bytes(), sh.Bytes())
	vs.empty = false
}

func (vs *variableState) Hash() hashing.HashValue {
	return vs.stateHash
}

func (vs *variableState) Variables() variables.Variables {
	return vs.vars
}

func (vs *variableState) saveToDb(addr *address.Address) error {
	dbase, err := database.GetVariableStateDB()
	if err != nil {
		return err
	}

	key := database.DbKeyVariableState(addr)
	h := vs.Hash()
	log.Debugw("state saving to db",
		"addr", addr.String(),
		"state index", vs.StateIndex(),
		"dbkey", base58.Encode(key),
		"stateHash", h.String(),
	)

	return dbase.Set(database.Entry{
		Key:   key,
		Value: hashing.MustBytes(vs),
	})
}

func (vs *variableState) Write(w io.Writer) error {
	if _, err := w.Write(util.Uint32To4Bytes(vs.stateIndex)); err != nil {
		return err
	}
	if _, err := w.Write(vs.stateHash.Bytes()); err != nil {
		return err
	}
	if err := vs.vars.Write(w); err != nil {
		return err
	}
	return nil
}

func (vs *variableState) Read(r io.Reader) error {
	if err := util.ReadUint32(r, &vs.stateIndex); err != nil {
		return err
	}
	if _, err := r.Read(vs.stateHash[:]); err != nil {
		return err
	}
	if err := vs.vars.Read(r); err != nil {
		return err
	}
	return nil
}

// saves variable state to db together with the batch of state updates
func (vs *variableState) Commit(addr address.Address, b Batch) error {
	// TODO make it Badger-atomic transaction
	// TODO mark processed requests in db in separate index

	if err := b.(*batch).saveToDb(&addr); err != nil {
		return err
	}
	if err := vs.saveToDb(&addr); err != nil {
		return err
	}
	if err := MarkRequestsProcessed(b.RequestIds()); err != nil {
		return err
	}
	return nil
}

func StateExist(addr address.Address) (bool, error) {
	dbase, err := database.GetVariableStateDB()
	if err != nil {
		return false, err
	}
	key := database.DbKeyVariableState(&addr)
	log.Debugw("checking if state exist",
		"addr", addr.String(),
		"dbkey", base58.Encode(key))
	return dbase.Contains(key)
}

// loads variable state and corresponding batch
func LoadVariableState(addr address.Address) (VariableState, Batch, error) {
	dbase, err := database.GetVariableStateDB()
	if err != nil {
		return nil, nil, err
	}
	entry, err := dbase.Get(database.DbKeyVariableState(&addr))
	if err != nil {
		return nil, nil, err
	}
	varState := NewVariableState(nil).(*variableState)
	if err = varState.Read(bytes.NewReader(entry.Value)); err != nil {
		return nil, nil, err
	}
	batch, err := LoadBatch(addr, varState.StateIndex())
	if err != nil {
		return nil, nil, err
	}
	if varState.StateIndex() != batch.StateIndex() {
		return nil, nil, fmt.Errorf("inconsistent solid state: state indices must be equal")
	}
	return varState, batch, nil
}