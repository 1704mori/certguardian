type Config = {
  port: number;
}

export function env() {
  // const text = document.querySelector("script#config__json")?.textContent || "{}";

  // const pageConfig = JSON.parse(text);

  // const config: Config = {
  //   ...pageConfig.env,
  // };

  // return Object.freeze(config);
  return Object.freeze({
    port: 7070
  });
}

/**
 * 
 * @returns http://localhost:[port]
 */
export function host() {
  return `http://localhost:${env().port}`
}