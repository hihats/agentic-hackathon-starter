import type { ShuffleRequest, ShuffleResponse, ErrorResponse } from '~/types/shuffle';

export const useShuffleApi = () => {
  const shuffle = async (request: ShuffleRequest): Promise<ShuffleResponse> => {
    // Nuxtのサーバーサイドプロキシを使用
    const response = await fetch('/api/shuffle', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(request),
    });

    if (!response.ok) {
      const error: ErrorResponse = await response.json();
      throw new Error(error.message || 'Shuffle failed');
    }

    return response.json();
  };

  return { shuffle };
};
