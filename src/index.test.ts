import { createExamplePromise } from './index'

test('createExamplePromise() resolves to true', () => {
    expect(createExamplePromise()).resolves.toBe(true);
})