<template>
  <div class="participant-input">
    <label for="participants">参加者 (改行またはカンマ区切り)</label>
    <textarea
      id="participants"
      v-model="inputText"
      :placeholder="placeholder"
      rows="10"
      @input="handleInput"
    />
    <p class="count">{{ participantCount }} 名が入力されています</p>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  modelValue: string[];
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string[]): void;
}>();

const inputText = ref('');
const error = ref('');

const placeholder = `名前を入力してください:
山田太郎
佐藤花子
鈴木一郎
または: 山田太郎, 佐藤花子, 鈴木一郎`;

const participantCount = computed(() => props.modelValue.length);

const handleInput = () => {
  // Parse input: split by newlines or commas
  const parsed = inputText.value
    .split(/[\n,]/)
    .map(name => name.trim())
    .filter(name => name.length > 0);

  emit('update:modelValue', parsed);

  // Validation
  if (parsed.length > 0 && parsed.length < 4) {
    error.value = '4名以上の参加者が必要です';
  } else {
    error.value = '';
  }
};
</script>

<style scoped>
.participant-input {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: #333;
}

textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-family: inherit;
  font-size: 1rem;
  resize: vertical;
  box-sizing: border-box;
}

textarea:focus {
  outline: none;
  border-color: #00dc82;
  box-shadow: 0 0 0 2px rgba(0, 220, 130, 0.2);
}

.count {
  color: #666;
  font-size: 0.875rem;
  margin-top: 0.5rem;
  margin-bottom: 0;
}

.error {
  color: #dc3545;
  font-size: 0.875rem;
  margin-top: 0.25rem;
  margin-bottom: 0;
}
</style>
