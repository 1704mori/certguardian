export async function addDomain(domain: string): Promise<void> {
  const response = await fetch(`http://localhost:7070/v1/`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ commonName: domain }),
  });
  if (!response.ok) {
    throw new Error(response.statusText);
  }
  
  const data = await response.json();
  return data;
}