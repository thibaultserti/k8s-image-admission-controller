package k8simageadmissioncontroller

import (
	"net/http"
)

func ValidatePod(w http.ResponseWriter, r *http.Request) {

	// ... Common code with the mutating webhook removed for
	// ... brevity. Refer to the past blog post for full code.

	// Create a response that either allows or rejects the pod creation
	// based off of the value of the hello label. Also, check to see if
	// we should supply a warning message even it is allowed.
	// admissionResponse := &admissionv1.AdmissionResponse{}
	// admissionResponse.Allowed = true

	// if value, ok := pod.Labels["hello"]; !ok {
	// 	admissionResponse.Allowed = false
	// 	admissionResponse.Result = &metav1.Status{
	// 		Message: "missing required hello label",
	// 	}
	// } else if value == "world" {
	// 	admissionResponse.Warnings = []string{"world will be deprecated for hello in the future"}
	// }

	// ... Common code with the mutating webhook removed for
	// ... brevity. Refer to the past blog post for full code.
}
