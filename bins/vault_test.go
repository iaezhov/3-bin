// bins/vault_test.go
package bins_test

import (
	"encoding/json"
	"testing"

	"hw/3/bins"
)

type mockDb struct {
	data []byte
}

func (m *mockDb) Read() ([]byte, error) {
	if m.data == nil {
		return nil, nil
	}
	return m.data, nil
}

func (m *mockDb) Write(content []byte) error {
	m.data = content
	return nil
}

func TestVaultAddBin(t *testing.T) {
	db := &mockDb{}
	vault := bins.NewVault(db)

	bin := bins.NewBin("123", "my bin", false)
	vault.AddBin(*bin)

	if len(vault.Bins) != 1 {
		t.Fatalf("Expected 1 bin, got %d", len(vault.Bins))
	}
	if vault.Bins[0].ID != "123" {
		t.Errorf("Expected ID '123', got '%s'", vault.Bins[0].ID)
	}

	var stored bins.Vault
	if err := json.Unmarshal(db.data, &stored); err != nil {
		t.Fatalf("Failed to unmarshal stored data: %v", err)
	}
	if len(stored.Bins) != 1 {
		t.Errorf("Stored vault should have 1 bin, got %d", len(stored.Bins))
	}
}

func TestVaultDeleteBin(t *testing.T) {
	db := &mockDb{}
	vault := bins.NewVault(db)

	vault.AddBin(*bins.NewBin("id-1", "first", false))
	vault.AddBin(*bins.NewBin("id-2", "second", false))

	success := vault.DeleteBin("id-1")
	if !success {
		t.Fatal("DeleteBin should return true")
	}
	if len(vault.Bins) != 1 {
		t.Fatalf("Expected 1 bin, got %d", len(vault.Bins))
	}
	if vault.Bins[0].ID != "id-2" {
		t.Errorf("Expected remaining bin ID 'id-2', got '%s'", vault.Bins[0].ID)
	}

	success = vault.DeleteBin("id-3")
	if success {
		t.Fatal("DeleteBin should return false for non-existent ID")
	}
}

func TestVaultLoadFromDB(t *testing.T) {
	db := &mockDb{
		data: []byte(`{"bins":[{"id":"existing","name":"test","private":false}]}`),
	}
	vault := bins.NewVault(db)

	if len(vault.Bins) != 1 {
		t.Fatalf("Expected 1 bin loaded from DB, got %d", len(vault.Bins))
	}
	if vault.Bins[0].ID != "existing" {
		t.Errorf("Expected ID 'existing', got '%s'", vault.Bins[0].ID)
	}
}

func TestVaultSave(t *testing.T) {
	db := &mockDb{}
	vault := bins.NewVault(db)

	vault.AddBin(*bins.NewBin("save-test", "test", false))

	var loaded bins.Vault
	if err := json.Unmarshal(db.data, &loaded); err != nil {
		t.Fatalf("Failed to unmarshal saved data: %v", err)
	}

	if len(loaded.Bins) != 1 {
		t.Fatalf("Expected 1 bin in saved data, got %d", len(loaded.Bins))
	}
	if loaded.Bins[0].ID != "save-test" {
		t.Errorf("Expected saved ID 'save-test', got '%s'", loaded.Bins[0].ID)
	}
}
