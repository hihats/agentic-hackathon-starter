<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let genres: string[] = [];
  export let selectedGenre: string = '';
  export let isLoading: boolean = false;

  const dispatch = createEventDispatcher<{
    change: { genre: string };
  }>();

  function handleGenreChange(event: Event) {
    const select = event.target as HTMLSelectElement;
    const genre = select.value;
    selectedGenre = genre;
    dispatch('change', { genre });
  }
</script>

<div class="genre-filter">
  <label for="genre-select" class="filter-label">ジャンル絞り込み</label>
  <select 
    id="genre-select" 
    bind:value={selectedGenre} 
    on:change={handleGenreChange}
    disabled={isLoading}
    class="genre-select"
  >
    <option value="">全ジャンル</option>
    {#each genres as genre}
      <option value={genre}>{genre}</option>
    {/each}
  </select>
</div>

<style>
  .genre-filter {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 16px;
  }

  .filter-label {
    font-weight: 500;
    color: #333;
    font-size: 14px;
    white-space: nowrap;
  }

  .genre-select {
    padding: 8px 12px;
    border: 2px solid #e1e5e9;
    border-radius: 8px;
    font-size: 14px;
    background-color: white;
    color: #333;
    cursor: pointer;
    transition: border-color 0.2s ease;
    min-width: 150px;
  }

  .genre-select:hover:not(:disabled) {
    border-color: #007bff;
  }

  .genre-select:focus {
    outline: none;
    border-color: #007bff;
    box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.1);
  }

  .genre-select:disabled {
    background-color: #f8f9fa;
    color: #6c757d;
    cursor: not-allowed;
  }
</style>