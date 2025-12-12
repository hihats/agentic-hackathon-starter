<template>
  <div class="group-settings">
    <label class="settings-label">グループ設定</label>
    <div class="mode-selector">
      <label class="radio-label">
        <input
          type="radio"
          value="group_size"
          v-model="mode"
          @change="updateValue"
        />
        <span>1グループあたりの人数を指定</span>
      </label>
      <label class="radio-label">
        <input
          type="radio"
          value="num_groups"
          v-model="mode"
          @change="updateValue"
        />
        <span>グループ数を指定</span>
      </label>
    </div>

    <div class="value-input">
      <label v-if="mode === 'group_size'" class="number-label">
        1グループあたりの人数:
        <input
          type="number"
          v-model.number="groupSize"
          min="2"
          @input="updateValue"
        />
      </label>
      <label v-else class="number-label">
        グループ数:
        <input
          type="number"
          v-model.number="numGroups"
          min="2"
          @input="updateValue"
        />
      </label>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { GroupingMode } from '~/types/shuffle';

const emit = defineEmits<{
  (e: 'update:groupSize', value: number | undefined): void;
  (e: 'update:numGroups', value: number | undefined): void;
}>();

const mode = ref<GroupingMode>('group_size');
const groupSize = ref(3);
const numGroups = ref(2);

const updateValue = () => {
  if (mode.value === 'group_size') {
    emit('update:groupSize', groupSize.value);
    emit('update:numGroups', undefined);
  } else {
    emit('update:groupSize', undefined);
    emit('update:numGroups', numGroups.value);
  }
};

// Initialize with default values
onMounted(() => {
  updateValue();
});
</script>

<style scoped>
.group-settings {
  margin-bottom: 1.5rem;
}

.settings-label {
  display: block;
  margin-bottom: 0.75rem;
  font-weight: bold;
  color: #333;
}

.mode-selector {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  color: #444;
}

.radio-label input[type="radio"] {
  accent-color: #00dc82;
}

.number-label {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: #444;
}

.number-label input[type="number"] {
  width: 80px;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 6px;
  font-size: 1rem;
}

.number-label input[type="number"]:focus {
  outline: none;
  border-color: #00dc82;
  box-shadow: 0 0 0 2px rgba(0, 220, 130, 0.2);
}
</style>
