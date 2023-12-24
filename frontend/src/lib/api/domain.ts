import { base, type ApiResponse, type CertInfo } from ".";

async function addDomain(domain: string): Promise<ApiResponse<string>> {
  return base<string>("domain", {
    method: "GET",
    body: {
      commonName: domain
    }
  });
}

async function listDomains(): Promise<ApiResponse<CertInfo[]>> {
  return base<CertInfo[]>("domain", {
    method: "GET",
  })
}

async function deleteDomain(domain: string): Promise<ApiResponse<string>> {
  return base<string>("domain", {
    method: "DELETE",
    query: {
      domain,
    }
  })
}

export {
  addDomain,
  listDomains,
  deleteDomain,
}