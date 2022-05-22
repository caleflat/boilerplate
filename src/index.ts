/**
 * @returns the example promise
 */
function createExamplePromise(): Promise<boolean> {
  return new Promise((resolve, _reject) => {
    resolve(true);
  });
}
