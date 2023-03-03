package excav

import "testing"

func TestRepository_GetName(t *testing.T) {
	repo := &Repository{
		Name: "/hello/world/",
	}

	name := repo.GetName()
	if name != "/hello/world" {
		t.FailNow()
	}

	repo2 := &Repository{
		Name: "/hello/world",
	}

	name = repo2.GetName()
	if name != "/hello/world" {
		t.FailNow()
	}
}

func TestRepository_HasTag(t *testing.T) {
	// given repo with tags
	repo := Repository{
		Tags: []string{"tag1", "tag2"},
	}

	// when we check if repo has 'tag1'
	res := repo.HasTag("tag1")

	// then result is true
	if !res {
		t.FailNow()
	}
}

func TestRepository_HasTagNegative(t *testing.T) {
	// given repo with tags
	repo := Repository{
		Tags: []string{"tag1", "tag2"},
	}

	// when we try to check non-existing tag
	res := repo.HasTag("unknown")

	// then the result is false
	if res {
		t.FailNow()
	}
}

func TestRepository_HasTags(t *testing.T) {
	// given repo with tags
	repo := Repository{
		Tags: []string{"tag1", "tag2", "tag3"},
	}

	// when we try to check if repo has tag2 and tag3
	res := repo.HasTags("tag2", "tag3")

	// then the result is true
	if !res {
		t.FailNow()
	}
}

func TestRepository_HasTagsNegative(t *testing.T) {
	// given repo with tags
	repo := Repository{
		Tags: []string{"tag1", "tag2", "tag3"},
	}

	// when we try to check if repo has tag3 and unknown
	res := repo.HasTags("unknown", "tag3")

	// then the result is false
	if res {
		t.FailNow()
	}
}
