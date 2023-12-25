import { base, type ApiResponse, type CertInfo } from ".";

export type CertList = {
  [dir: string]: {
    [certPath: string]: CertInfo
  }
};

async function addCertDir(directories: string[]): Promise<ApiResponse<string>> {
  return base<string>("cert", {
    method: "POST",
    body: {
      directories,
    }
  });
}

async function listCerts(): Promise<ApiResponse<CertList>> {
  return base<CertList>("cert", {
    method: "GET",
  })
}

async function deleteDir(dir: string): Promise<ApiResponse<string>> {
  return base<string>("cert", {
    method: "DELETE",
    params: {
      dir,
    }
  })
}

export {
  addCertDir,
  listCerts,
  deleteDir,
}