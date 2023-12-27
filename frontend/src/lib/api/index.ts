import { host } from "$lib/env";

export type CertInfo = {
  commonName: string;
  issuer: string;
  validFrom: string;
  validTo: string;
  isExpired: boolean;
  isNearExpiry: boolean;
};

export type ApiResponse<T = any> = { message: T; error?: string };

type Method = "GET" | "POST" | "PUT" | "DELETE";
type Endpoint = "domain" | "cert";

type Params = {
  method: Method;
  body?: Record<string, any>;
  params?: Record<string, any>;
  query?: Record<string, any>;
};

export async function base<T = any>(
  endpoint: Endpoint,
  { body, method, params, query }: Params,
): Promise<ApiResponse<T>> {
  let url = new URL(`${host()}/v1/${endpoint}`);

  for (const q in query) {
    url.searchParams.set(q, query[q]);
  }

  for (const param in params) {
    url.pathname += `/${params[param]}`;
  }

  const response = await fetch(url, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  }).catch(() => null);

  if (!response) {
    console.log("skipping error");
    return {
      message: [] as T,
    };
  }

  const data: ApiResponse<T> = await response.json();

  if (response.status !== 200) {
    console.error(data);
    // return [null as any, new Error(data.error)];
    throw new Error(data.error);
  }

  return data;
}
