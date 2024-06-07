package customstorage

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

// Implement rest.Creater
func (s *Standalone) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	return obj, nil
}

// Implement rest.Getter
func (s *Standalone) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return nil, errors.NewNotFound(s.GroupResource(), name)
}

// Implement rest.Lister
func (s *Standalone) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	return nil, nil // Nothing is ready here xD
}

// Implement rest.Updater
func (s *Standalone) Update(ctx context.Context, name string, objif rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	obj, err := s.Get(ctx, name, nil)
	if err != nil {
		return nil, false, err
	}

	updated, err := objif.UpdatedObject(ctx, obj)
	if err != nil {
		return nil, false, err
	}

	return updated, false, nil
}

// Implement rest.Deleter
func (s *Standalone) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	return nil, false, errors.NewNotFound(s.GroupResource(), name)
}
