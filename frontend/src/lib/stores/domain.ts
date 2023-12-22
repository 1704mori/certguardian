import { addDomain as _addDomain } from "$lib/api/domain";
import { writable } from "svelte/store";

export function domain() {
  const { subscribe, update, set } = writable([]);

  async function addDomain(domain: string) {
    
    // listDomains();
    return _addDomain(domain);
  }

  return {
    subscribe,
    addDomain,
  };
}