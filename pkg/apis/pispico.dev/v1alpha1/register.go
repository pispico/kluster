package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	pispicodev "github.com/pispico/kluster/pkg/apis/pispico.dev"
)

var SchemeGroupVersion = schema.GroupVersion{Group: pispicodev.GroupName, Version: "v1alpha1"}

var (
	SchemeBuilder runtime.SchemeBuilder
	AddToScheme   = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &Kluster{}, &KlusterList{})
	return nil
}
