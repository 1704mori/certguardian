import { addCertDir, deleteDir as apiDeleteDir, listCerts, type CertList } from "$lib/api/certificates";
import { writable } from "svelte/store";

export const certificates = writable<CertList>();

async function initializeCerts() {
  certificates.set({} as CertList);
  const response = await listCerts();
  if (!response.error) {
    certificates.set(response.message);
  }
}

async function addNewDir(dirs: string[]) {
  const response = await addCertDir(dirs);
  if (!response.error) {
    await initializeCerts();
  }

  return response
}

async function deleteDir(dir: string) {
  const response = await apiDeleteDir(dir);
  if (!response.error) {
    await initializeCerts();
  }

  return response
}

initializeCerts();

export default {
  subscribe: certificates.subscribe,
  addNewDir,
  deleteDir,
};