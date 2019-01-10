package nborm

import "fmt"

type Validator func(Field) error

var validatorMap = make(map[string]Validator)

func RegisterValidator(name string, v Validator) {
	if _, ok := validatorMap[name]; ok {
		panic(fmt.Errorf("nborm.RegisterValidator() error: validator already exists (%s)", name))
	}
	validatorMap[name] = v
}

func nullValidator(f Field) error {
	if !f.isNullable() && f.getDefVal() == nil && f.IsNull() {
		return fmt.Errorf("nborm.nullValidator() error: %s.%s cannot be null", f.dbName(), f.tabName())
	}
	return nil
}

func init() {
	RegisterValidator("nullValidator", nullValidator)
}
