package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

func getResourceID(req *http.Request) string {
	return mux.Vars(req)["resourceID"]
}

func getRequestParameters(req *http.Request) (*models.Parameters, error) {
	params := &models.Parameters{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(params); err != nil {
		return nil, err
	}
	return params, nil
}

// TODO: support mode, profile
func getValidationParameters(req *http.Request) (interface{}, error) {
	params, err := getRequestParameters(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not get validation parameters")
	}
	for _, p := range params.Parameter {
		if p.Name == "resource" {
			return p.Resource, nil
		}
	}
	return nil, fmt.Errorf("unable to find resource parameter")
}

func loadResourceFromBody(target interface{}, req *http.Request, validator *models.JSONValidator) error {
	bArr, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	valid, _, err := validator.Validate(bArr)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("resource failed FHIR validation")
	}
	return json.Unmarshal(bArr, target)
}

func resourceCreated(rndr *render.Render, rw http.ResponseWriter, resourceID string, versionID string, lastModified time.Time, resource interface{}) {
	location := fmt.Sprintf("/fhir/%s/%s", utils.GetBaseTypeName(resource), resourceID)
	if versionID != "" {
		location = fmt.Sprintf("%s/_history/%s", location, versionID)
		rw.Header().Add("Etag", fmt.Sprintf(`W/"%s"`, versionID))
	}
	rw.Header().Add("Last-Modified", lastModified.Format(http.TimeFormat))
	rw.Header().Add("Location", location)
	rndr.JSON(rw, http.StatusCreated, resource)
}

func resourceRead(rndr *render.Render, rw http.ResponseWriter, status int, versionID string, lastModified time.Time, resource interface{}) {
	if versionID != "" {
		rw.Header().Add("Etag", fmt.Sprintf(`W/"%s"`, versionID))
	}
	rw.Header().Add("Last-Modified", lastModified.Format(http.TimeFormat))
	rndr.JSON(rw, status, resource)
}

func validateJSONResource(log *logging.Logger, rndr *render.Render, validator *models.JSONValidator) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resource, err := getValidationParameters(req)
		if err != nil {
			log.WithError(err).Panic("could not get validation parameters")
		}
		bytes, err := json.Marshal(resource)
		if err != nil {
			log.WithError(err).Panic("could not marshal resource")
		}
		valid, vErrs, err := validator.Validate(bytes)
		if err != nil {
			log.WithError(err).Panic("could not validate resource")
		}
		issues := []*models.OperationOutcome_Issue{}
		if valid {
			issues = append(issues, &models.OperationOutcome_Issue{
				Severity:    models.IssueSeverityInformation,
				Code:        models.IssueTypeInformational,
				Diagnostics: "no issues detected",
			})
		} else {
			for _, e := range vErrs {
				issues = append(issues, &models.OperationOutcome_Issue{
					Severity:    models.IssueSeverityError,
					Code:        models.IssueTypeInvalid,
					Diagnostics: e.String(),
				})
			}
		}
		outcome := models.NewOperationOutcome()
		outcome.Issue = issues
		rndr.JSON(rw, http.StatusOK, outcome)
	})
}
