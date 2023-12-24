/**
 * Utility function to resolve a promise and return an array [data, error].
 * If the promise resolves successfully, the error will be null.
 * If the promise is rejected, the data will be null.
 * 
 * @param promise The promise to be resolved.
 * @returns A tuple [data, error].
 */
export async function resolvePromise<T>(promise: Promise<T>): Promise<[T, Error | null]> {
  try {
    const data = await promise;
    return [data, null];
  } catch (error) {
    if (error instanceof Error) {
      return [null as any, error];
    } else {
      // If the caught error isn't an instance of Error, create a new Error with the given information
      return [null as any, new Error(String(error))];
    }
  }
}