export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  // Docker内部ではサービス名 "go" を使用
  const API_BASE = process.env.GO_API_URL || 'http://go:8080';

  try {
    const response = await fetch(`${API_BASE}/api/shuffle`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    if (!response.ok) {
      const error = await response.json();
      throw createError({
        statusCode: response.status,
        statusMessage: error.message || 'Shuffle failed',
      });
    }

    return await response.json();
  } catch (error: any) {
    if (error.statusCode) {
      throw error;
    }
    throw createError({
      statusCode: 500,
      statusMessage: error.message || 'Failed to connect to API',
    });
  }
});
