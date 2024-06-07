package customstorage

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Standalone struct {
	obj          runtime.Object
	objl         runtime.Object
	gvk          schema.GroupVersionKind
	namespaced   bool
	name         string
	singularName string
	shortNames   []string
}

func NewStandaloneInterface(obj runtime.Object, objl runtime.Object, gvk schema.GroupVersionKind, namespaced bool, name string, singular string, shortNames []string) *Standalone {
	return &Standalone{
		obj:          obj,
		objl:         objl,
		gvk:          gvk,
		namespaced:   namespaced,
		name:         name,
		singularName: singular,
		shortNames:   shortNames,
	}
}

func (s *Standalone) New() runtime.Object {
	obj := s.obj.DeepCopyObject()
	obj.GetObjectKind().SetGroupVersionKind(s.gvk)
	return obj
}

func (s *Standalone) NewList() runtime.Object {
	objl := s.objl.DeepCopyObject()
	objl.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   s.gvk.Group,
		Version: s.gvk.Version,
		Kind:    "List",
	})
	return objl
}

func (s *Standalone) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return nil, nil
}

func (s *Standalone) GroupResource() schema.GroupResource {
	return schema.GroupResource{
		Group:    s.gvk.Group,
		Resource: s.name,
	}
}

func (s *Standalone) NamespaceScoped() bool {
	return s.namespaced
}

func (s *Standalone) GetSingularName() string {
	return s.singularName
}

func (s *Standalone) ShortNames() []string {
	return s.shortNames
}

func (s *Standalone) Destroy() {
	// Nothing to do (right now!)
}
