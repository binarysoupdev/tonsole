/*
The rand package provides helpers for utilizing random generated values in automated test cases.

To ensure tests remain deterministic, a constant seed should be used within the test.

	const SEED = 42
	r := rand.New(SEED)

The motivation behind the rand package is to make tests easier to read by abstracting away arbitrary values.
In the following example, we clearly see password is an ASCII string of 30 characters.
Its exact value, however, is unimportant for the purpose of testing, so the value remains hidden.

	password := r.ASCII(30)
*/
package rand
