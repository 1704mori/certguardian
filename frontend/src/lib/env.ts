import * as env from "$env/static/public";

type Config = {
  PUBLIC_PORT: number;
  PUBLIC_CRON_INTERVAL: string;
  PUBLIC_NEAR_EXPIRY_THRESHOLD: string;
};

export function _env() {
  const config: Config = {
    ...env,
    PUBLIC_PORT: parseInt(env.PUBLIC_PORT),
  };

  return Object.freeze(config);
}
