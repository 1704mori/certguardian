import { host } from "$lib/env";

type ApiResponse<T = any> = { message: T; error?: string; };
export type DomainList = { commonName: string; issuer: string; validFrom: string; validTo: string; isExpired: boolean };

async function addDomain(domain: string): Promise<ApiResponse<string>> {
  const response = await fetch(`${host()}/v1/domain`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ commonName: domain }),
  });

  const data: ApiResponse<string> = await response.json();

  if (response.status !== 200) {
    console.error(data);
    // return [null as any, new Error(data.error)];
    throw new Error(data.error);
  }

  return data;
}

async function listDomains(): Promise<ApiResponse<DomainList[]>> {
  const response = await fetch(`${host()}/v1/domain`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  const data: ApiResponse<DomainList[]> = await response.json();

  if (response.status !== 200) {
    console.error(data);
    // return [null as any, new Error(data.error)];
    throw new Error(data.error);
  }

  return data;
}

async function deleteDomain(domain: string): Promise<ApiResponse<string>> {
  const response = await fetch(`${host()}/v1/domain/${domain}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  const data: ApiResponse<string> = await response.json();

  if (response.status !== 200) {
    console.error(data);
    throw new Error(data.error);
  }

  return data
}

export {
  addDomain,
  listDomains,
  deleteDomain,
}