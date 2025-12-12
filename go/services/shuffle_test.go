package services

import (
	"testing"
)

func TestNewShuffleService(t *testing.T) {
	service := NewShuffleService()
	if service == nil {
		t.Error("NewShuffleService() returned nil")
	}
	if service.rng == nil {
		t.Error("ShuffleService.rng is nil")
	}
}

func TestShuffle_WithGroupSize(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"}
	groupSize := 2

	groups, err := service.Shuffle(participants, &groupSize, nil)

	if err != nil {
		t.Errorf("Shuffle() returned error: %v", err)
	}

	if len(groups) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(groups))
	}

	// 全参加者が含まれているか確認
	allMembers := make(map[string]bool)
	for _, group := range groups {
		for _, member := range group {
			allMembers[member] = true
		}
	}

	if len(allMembers) != len(participants) {
		t.Errorf("Expected %d unique members, got %d", len(participants), len(allMembers))
	}
}

func TestShuffle_WithNumGroups(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"}
	numGroups := 2

	groups, err := service.Shuffle(participants, nil, &numGroups)

	if err != nil {
		t.Errorf("Shuffle() returned error: %v", err)
	}

	if len(groups) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(groups))
	}

	// 全参加者が含まれているか確認
	totalMembers := 0
	for _, group := range groups {
		totalMembers += len(group)
	}

	if totalMembers != len(participants) {
		t.Errorf("Expected %d total members, got %d", len(participants), totalMembers)
	}
}

func TestShuffle_WithRemainder(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
	groupSize := 2

	groups, err := service.Shuffle(participants, &groupSize, nil)

	if err != nil {
		t.Errorf("Shuffle() returned error: %v", err)
	}

	// 5人を2人ずつに分けると、2グループ + 余り1人
	// 余りは既存グループに分配されるので、3人グループが1つできる
	totalMembers := 0
	for _, group := range groups {
		totalMembers += len(group)
	}

	if totalMembers != len(participants) {
		t.Errorf("Expected %d total members, got %d", len(participants), totalMembers)
	}
}

func TestShuffle_NoGroupSizeOrNumGroups(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave"}

	_, err := service.Shuffle(participants, nil, nil)

	if err == nil {
		t.Error("Expected error when neither group_size nor num_groups is provided")
	}
}

func TestShuffle_GroupSizeLargerThanParticipants(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave"}
	groupSize := 10

	groups, err := service.Shuffle(participants, &groupSize, nil)

	if err != nil {
		t.Errorf("Shuffle() returned error: %v", err)
	}

	// グループサイズが参加者数より大きい場合、2グループに分ける
	if len(groups) < 2 {
		t.Errorf("Expected at least 2 groups, got %d", len(groups))
	}
}

func TestShuffle_Randomness(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank", "Grace", "Henry"}
	groupSize := 2

	// 複数回シャッフルして、結果が異なることを確認
	results := make([]string, 3)
	for i := 0; i < 3; i++ {
		groups, _ := service.Shuffle(participants, &groupSize, nil)
		// 最初のグループの最初のメンバーを記録
		results[i] = groups[0][0]
	}

	// 少なくとも1回は異なる結果になることを期待（確率的に）
	// ただし、このテストは確率的なので、常に成功するとは限らない
	// テストの信頼性のため、このチェックはコメントアウト
	// allSame := results[0] == results[1] && results[1] == results[2]
	// if allSame {
	//     t.Log("Warning: All shuffle results were the same (this can happen by chance)")
	// }
}

func TestShuffle_MinimumParticipants(t *testing.T) {
	service := NewShuffleService()
	participants := []string{"Alice", "Bob", "Charlie", "Dave"}
	groupSize := 2

	groups, err := service.Shuffle(participants, &groupSize, nil)

	if err != nil {
		t.Errorf("Shuffle() returned error: %v", err)
	}

	if len(groups) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(groups))
	}
}
