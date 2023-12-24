import type { CertInfo } from "$lib/api";
import { addDomain, deleteDomain as apiDeleteDomain, listDomains } from "$lib/api/domain";
import { writable } from "svelte/store";

export const domains = writable<CertInfo[]>([]);

async function initializeDomains() {
  domains.set([]);
  const response = await listDomains();
  if (!response.error) {
    domains.set(response.message);
  }
}

async function addNewDomain(domain: string) {
  const response = await addDomain(domain);
  if (!response.error) {
    await initializeDomains();
  }

  return response
}

async function deleteDomain(domain: string) {
  const response = await apiDeleteDomain(domain);
  if (!response.error) {
    await initializeDomains();
  }

  return response
}

initializeDomains();

export default {
  subscribe: domains.subscribe,
  addNewDomain,
  deleteDomain,
};