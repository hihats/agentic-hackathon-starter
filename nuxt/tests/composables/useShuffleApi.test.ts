import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { useShuffleApi } from '../../composables/useShuffleApi'

describe('useShuffleApi', () => {
  const mockFetch = vi.fn()

  beforeEach(() => {
    vi.stubGlobal('fetch', mockFetch)
  })

  afterEach(() => {
    vi.unstubAllGlobals()
    mockFetch.mockReset()
  })

  it('should return shuffle function', () => {
    const { shuffle } = useShuffleApi()
    expect(typeof shuffle).toBe('function')
  })

  it('should call API with correct parameters', async () => {
    const mockResponse = {
      groups: [['Alice', 'Bob'], ['Charlie', 'Dave']],
    }
    mockFetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(mockResponse),
    })

    const { shuffle } = useShuffleApi()
    const result = await shuffle({
      participants: ['Alice', 'Bob', 'Charlie', 'Dave'],
      group_size: 2,
    })

    expect(mockFetch).toHaveBeenCalledWith('/api/shuffle', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        participants: ['Alice', 'Bob', 'Charlie', 'Dave'],
        group_size: 2,
      }),
    })

    expect(result).toEqual(mockResponse)
  })

  it('should call API with num_groups parameter', async () => {
    const mockResponse = {
      groups: [['Alice', 'Bob'], ['Charlie', 'Dave']],
    }
    mockFetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(mockResponse),
    })

    const { shuffle } = useShuffleApi()
    await shuffle({
      participants: ['Alice', 'Bob', 'Charlie', 'Dave'],
      num_groups: 2,
    })

    expect(mockFetch).toHaveBeenCalledWith('/api/shuffle', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        participants: ['Alice', 'Bob', 'Charlie', 'Dave'],
        num_groups: 2,
      }),
    })
  })

  it('should throw error on API failure', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: false,
      json: () => Promise.resolve({ message: '参加者は4名以上必要です' }),
    })

    const { shuffle } = useShuffleApi()

    await expect(
      shuffle({
        participants: ['Alice', 'Bob'],
        group_size: 2,
      })
    ).rejects.toThrow('参加者は4名以上必要です')
  })

  it('should throw default error message when API returns no message', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: false,
      json: () => Promise.resolve({}),
    })

    const { shuffle } = useShuffleApi()

    await expect(
      shuffle({
        participants: ['Alice', 'Bob'],
        group_size: 2,
      })
    ).rejects.toThrow('Shuffle failed')
  })
})
