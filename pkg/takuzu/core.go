package takuzu

import "errors"

// Creating a new instance of Takuzu Core with given size
func NewCore(size int) (*Core, error) {

	if (size < 4 || size > 12) {
		return &Core{}, errors.New(incorrectSize)
	}

	return &Core{
		Size: size,
	}, nil
}

// Generating a new Takuzu field
func (tc *Core) Generate() {
	field, err := generateField(tc.Size)
	for err != nil {
		field, err = generateField(tc.Size)
	}

	_, errVerify := verify(field)
	for errVerify != nil || err != nil {
		field, err = generateField(tc.Size)
		_, errVerify = verify(field)
	}

	tc.Field = field
}

// Preparing a new Takuzu task based on already generated field with given fill factor
func (tc *Core) Prepare(fillFactor int) error {
	task, err := prepareField(tc.Field, fillFactor)
	if err != nil {
		return err
	}

	tc.Task = task

	return nil
}

// Verifying current Takuzu task
func (tc *Core) Verify() (*VerificationResult, error) {
	result, err := verify(tc.Task)
	return result, err
}