package excav

import "testing"

func TestLoad(t *testing.T) {
	// When we load inventory from file
	inventory, err := OpenInventory("testdata/inventory.yaml")

	// Then the load must be without errors
	if err != nil {
		t.FailNow()
	}

	// ... and all repositories are in inventory
	if len(inventory.XRepositories) != 4 {
		t.FailNow()
	}

	// .. and /org/team2/repo4 have parameters
	repo := inventory.Get("/org/team2/repo4")
	param1Value := repo.RepoParams["param1"]
	if param1Value != "val1" {
		t.FailNow()
	}
}

func TestInventory_GetByTags(t *testing.T) {
	// Given inventory with several repositories
	inventory, err := OpenInventory("testdata/inventory.yaml")
	if err != nil {
		t.FailNow()
	}

	// When we get 'prod' repositories
	repos := inventory.GetByTags("prod")

	// Then we should have all 2 prod repos
	if len(repos) != 2 {
		t.FailNow()
	}

	// ... and any repo must be 'prod'
	if !repos[0].HasTag("prod") {
		t.FailNow()
	}
}

func TestInventory_GetByTagsMultiple(t *testing.T) {
	// Given inventory with several repositories
	inventory, err := OpenInventory("testdata/inventory.yaml")
	if err != nil {
		t.FailNow()
	}

	// When we get 'sandbox' of 'team2' repositories
	repo := inventory.GetByTags("sandbox", "team2")

	// Then we should have only one repo
	if len(repo) != 1 {
		t.FailNow()
	}

	// ... and is '' which matching  both 'sandbox' and 'team2' tags
	if repo[0].Name != "/org/team2/repo3" {
		t.FailNow()
	}

	if !repo[0].HasTag("sandbox") || !repo[0].HasTags("team2") {
		t.FailNow()
	}
}
