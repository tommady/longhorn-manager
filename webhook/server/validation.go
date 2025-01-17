package server

import (
	"net/http"

	"github.com/rancher/wrangler/pkg/webhook"

	"github.com/longhorn/longhorn-manager/webhook/admission"
	"github.com/longhorn/longhorn-manager/webhook/client"
	"github.com/longhorn/longhorn-manager/webhook/resources/setting"
)

func Validation(client *client.Client) (http.Handler, []admission.Resource, error) {
	resources := []admission.Resource{}
	validators := []admission.Validator{
		setting.NewValidator(client.Datastore),
	}

	router := webhook.NewRouter()
	for _, v := range validators {
		addHandler(router, admission.AdmissionTypeValidation, admission.NewValidatorAdapter(v))
		resources = append(resources, v.Resource())
	}

	return router, resources, nil
}
