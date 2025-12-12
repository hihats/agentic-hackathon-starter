package services

import (
	"errors"
	"math/rand"
	"time"
)

// ShuffleService handles the shuffle logic
type ShuffleService struct {
	rng *rand.Rand
}

// NewShuffleService creates a new shuffle service
func NewShuffleService() *ShuffleService {
	return &ShuffleService{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Shuffle takes participants and group_size/num_groups, returns grouped participants
func (s *ShuffleService) Shuffle(participants []string, groupSize *int, numGroups *int) ([][]string, error) {
	n := len(participants)

	// Calculate actual group size
	var actualGroupSize int
	if groupSize != nil && *groupSize > 0 {
		actualGroupSize = *groupSize
	} else if numGroups != nil && *numGroups > 0 {
		actualGroupSize = n / *numGroups
		if actualGroupSize < 1 {
			actualGroupSize = 1
		}
	} else {
		return nil, errors.New("グループサイズまたはグループ数を指定してください")
	}

	// Ensure at least 2 groups can be formed
	if actualGroupSize >= n {
		actualGroupSize = n / 2
	}

	// Shuffle participants randomly
	shuffled := make([]string, len(participants))
	copy(shuffled, participants)
	s.rng.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// Create groups
	var groups [][]string
	for i := 0; i < len(shuffled); i += actualGroupSize {
		end := i + actualGroupSize
		if end > len(shuffled) {
			// Distribute remaining members to existing groups
			remainder := shuffled[i:]
			for j, member := range remainder {
				groups[j%len(groups)] = append(groups[j%len(groups)], member)
			}
			break
		}
		group := make([]string, actualGroupSize)
		copy(group, shuffled[i:end])
		groups = append(groups, group)
	}

	return groups, nil
}
