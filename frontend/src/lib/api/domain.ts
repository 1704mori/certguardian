type Response<T> = [T, Error | null];
type ApiResponse = { message: string | any; error: string; };

export async function addDomain(domain: string): Promise<Response<ApiResponse>> {
  const response = await fetch(`http://localhost:7070/v1/`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ commonName: domain }),
  });

  const data: ApiResponse = await response.json();

  if (response.status !== 200) {
    console.error(data);
    return [null as any, new Error(data.error)];
  }

  return [data, null];
}