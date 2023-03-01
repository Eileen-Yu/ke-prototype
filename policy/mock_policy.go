package policy

/*
package policy

import (
	"encoding/json"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Suibian struct {
	TargetResources []string
	ForbidRules     []string
}

var SuibianPolicy Suibian = Suibian{
	TargetResources: []string{"pod"},
	ForbidRules:     []string{"file I/O"},
}

type Policy struct {
	Spec Spec `json:"spec" yaml:"spec"`
}

type Spec struct {
	Rules []Rule `json:"rules,omitempty" yaml:"rules,omitempty"`
}

type Rule struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	MatchResources MatchResources `json:"match,omitempty" yaml:"match,omitempty"`

	Restriction Restriction `json:"restrict,omitempty" yaml:"restrict,omitempty"`
}

type MatchResources struct {
	ResourceDescription `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type ResourceDescription struct {
	Kinds []string `json:"kinds,omitempty" yaml:"kinds,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type Restriction struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	RawPattern *apiextv1.JSON `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

var policyA = []byte(`
{
  "spec": {
    "rules": [
      {
        "name": "foo",
        "match": {
          "resources": {
          "kinds": [
            "Pod"
          ],
          }
        }
        "Restrict": {
          "message": "bar",
          "pattern": [
            {
            "spec": {
              "containers": [
                {
                  "name": "A"
                }
              ],
              "behavior": [
                "volume-mount"
              ],
            }
            }
          ]
        }
      }
    ]
}
`)

var mockPolicy *Policy
var err = json.Unmarshal(policyA, &mockPolicy)
*/
