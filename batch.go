package batch

type abort string

func (e abort) Error() string {
	return string(e)
}

// ErrAbort is a sentinel error which indicates a batch
// operation should abort early.
const ErrAbort = abort("done")

// Func is called for each batch.
// Any error will cancel the batching operation but returning Abort
// indicates it was deliberate, and not an error case.
type Func func(start, end int) error

// Batch calls fn for all items up to count.
// Returns any error from fn but for ErrAbort returns nil.
func Batch(count, batchSize int, fn Func) error {
	for i := 0; i < count; i += batchSize {
		end := i + batchSize - 1
		if end > count {
			end = count - 1
		}

		err := fn(i, end)
		if err != nil {
			if err == ErrAbort {
				return nil
			}
			return err
		}
	}
	return nil
}
