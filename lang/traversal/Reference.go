package traversal

import (
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammar"
)

func init() {
	Register(func(ref *grammar.ResourceReferenceContext, currentNamespace string, _ *Scope) Construct {
			parts := ref.AllIdentifier()
			switch len(parts) {
			default:
				return nil
			case 1:
				return Ref(currentNamespace, parts[0].GetText())
			case 2:
				return Ref(parts[0].GetText(), parts[1].GetText())
			}
		},
	)
}

type Reference struct {
	Construct
	name      string
	namespace string

	resolver func() error
}

func NewReference(name string, namespace string) *Reference {
	return &Reference{
		name:      name,
		namespace: namespace,
	}
}

// Ref Shorter alias for NewReference
func Ref(namespace string, name string) *Reference {
	return NewReference(name, namespace)
}

func (r Reference) GetName() string {
	return r.name
}

func (r Reference) GetNamespace() string {
	return r.namespace
}

func (r Reference) String() string {
	return fmt.Sprintf("%s:%s", r.namespace, r.name)
}

func (r Reference) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", r.String())), nil
}

func (r Reference) Equals(other Reference) bool {
	return r.String() == other.String()
}

func (r *Reference) SetResolver(replacer func() error) {
	r.resolver = replacer
}

func (r Reference) Resolve() error {
	if r.resolver != nil {
		return r.resolver()
	}
	return nil
}
