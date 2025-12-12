<template>
  <div class="result-display" v-if="groups.length > 0">
    <h2>シャッフル結果</h2>
    <div class="groups">
      <div
        v-for="(group, index) in groups"
        :key="index"
        class="group"
      >
        <h3>グループ {{ getGroupLabel(index) }}</h3>
        <ul>
          <li v-for="member in group" :key="member">{{ member }}</li>
        </ul>
        <p class="member-count">{{ group.length }}名</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  groups: string[][];
}>();

const getGroupLabel = (index: number): string => {
  return String.fromCharCode(65 + index); // A, B, C, ...
};
</script>

<style scoped>
.result-display {
  margin-top: 2rem;
}

h2 {
  color: #333;
  margin-bottom: 1rem;
  font-size: 1.5rem;
}

.groups {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.group {
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  border: 1px solid #dee2e6;
  border-radius: 12px;
  padding: 1.25rem;
  transition: transform 0.2s, box-shadow 0.2s;
}

.group:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.group h3 {
  color: #00dc82;
  margin: 0 0 0.75rem 0;
  font-size: 1.1rem;
  font-weight: bold;
}

.group ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.group li {
  padding: 0.4rem 0;
  border-bottom: 1px solid #e9ecef;
  color: #333;
}

.group li:last-child {
  border-bottom: none;
}

.member-count {
  margin: 0.75rem 0 0 0;
  font-size: 0.8rem;
  color: #888;
  text-align: right;
}
</style>
