/**
 * @returns a promise that resolves to true 
 */
export function createExamplePromise(): Promise<boolean> {
  return new Promise((resolve, _reject) => {
    resolve(true);
  });
}