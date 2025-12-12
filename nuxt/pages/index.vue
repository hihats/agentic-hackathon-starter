<template>
  <div class="container">
    <h1>Shuffle Lunch</h1>
    <p class="subtitle">参加者をランダムにグループ分けします</p>

    <form @submit.prevent="handleShuffle">
      <ParticipantInput v-model="participants" />

      <GroupSettings
        @update:groupSize="groupSize = $event"
        @update:numGroups="numGroups = $event"
      />

      <button
        type="submit"
        class="shuffle-button"
        :disabled="!isValid || isLoading"
      >
        <span v-if="isLoading">シャッフル中...</span>
        <span v-else>シャッフル!</span>
      </button>

      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </form>

    <ResultDisplay :groups="groups" />
  </div>
</template>

<script setup lang="ts">
import type { ShuffleRequest } from '~/types/shuffle';

const { shuffle } = useShuffleApi();

const participants = ref<string[]>([]);
const groupSize = ref<number | undefined>(3);
const numGroups = ref<number | undefined>(undefined);
const groups = ref<string[][]>([]);
const isLoading = ref(false);
const errorMessage = ref('');

const isValid = computed(() => {
  return participants.value.length >= 4 && (groupSize.value || numGroups.value);
});

const handleShuffle = async () => {
  if (!isValid.value) return;

  isLoading.value = true;
  errorMessage.value = '';
  groups.value = [];

  try {
    const request: ShuffleRequest = {
      participants: participants.value,
    };

    if (groupSize.value) {
      request.group_size = groupSize.value;
    } else if (numGroups.value) {
      request.num_groups = numGroups.value;
    }

    const response = await shuffle(request);
    groups.value = response.groups;
  } catch (err) {
    errorMessage.value = err instanceof Error ? err.message : 'エラーが発生しました';
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

h1 {
  color: #00dc82;
  margin-bottom: 0.5rem;
  font-size: 2.5rem;
}

.subtitle {
  color: #666;
  margin-bottom: 2rem;
  font-size: 1.1rem;
}

.shuffle-button {
  width: 100%;
  padding: 1rem;
  font-size: 1.2rem;
  font-weight: bold;
  color: white;
  background: linear-gradient(135deg, #00dc82 0%, #00b368 100%);
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.shuffle-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 220, 130, 0.4);
}

.shuffle-button:active:not(:disabled) {
  transform: translateY(0);
}

.shuffle-button:disabled {
  background: #ccc;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.error-message {
  color: #dc3545;
  margin-top: 1rem;
  text-align: center;
  font-weight: 500;
}
</style>
