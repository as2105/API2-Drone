package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/utils"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

const (
	versionDelimiter         = "-"
	initialResourceVersionID = "0" + versionDelimiter + "0"
)

func getResourceID(req *http.Request) (uuid.UUID, error) {
	uuidStr := mux.Vars(req)["resourceID"]
	if uuidStr == "" {
		return nil, errors.New("no resource ID was provided")
	}
	uuidObj := uuid.Parse(uuidStr)
	if uuidObj == nil {
		return nil, errors.New("unable to parse UUID from resource ID")
	}
	return uuidObj, nil
}

func getRequestParameters(req *http.Request, h resourceHandler) (*models.Parameters, error) {
	log := h.getLogger()
	params := &models.Parameters{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(params); err != nil {
		return nil, err
	}
	log.Debug(spew.Sdump(params))
	return params, nil
}

// TODO: support mode, profile
func getValidationParameters(req *http.Request, h resourceHandler) (interface{}, error) {
	params, err := getRequestParameters(req, h)
	if err != nil {
		return nil, errors.Wrap(err, "could not get validation parameters")
	}
	for _, p := range params.Parameter {
		if p.Name == "resource" {
			return p.Resource, nil
		}
	}
	return nil, errors.New("unable to find resource parameter")
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
		return errors.New("resource failed FHIR validation")
	}
	return json.Unmarshal(bArr, target)
}

func resourceCreated(
	rndr *render.Render,
	rw http.ResponseWriter,
	resourceID uuid.UUID,
	versionID string,
	lastModified time.Time,
	resource interface{},
) {
	location := fmt.Sprintf("/fhir/%s/%s", utils.GetBaseTypeName(resource), resourceID.String())
	if versionID != "" {
		location = fmt.Sprintf("%s/_history/%s", location, versionID)
		rw.Header().Set("Etag", generateETag(versionID))
	}
	rw.Header().Set("Last-Modified", lastModified.Format(http.TimeFormat))
	rw.Header().Set("Location", location)
	rndr.JSON(rw, http.StatusCreated, resource)
}

func generateETag(versionID string) string {
	return fmt.Sprintf(`W/"%s"`, versionID)
}

func resourceRead(
	rndr *render.Render,
	rw http.ResponseWriter,
	req *http.Request,
	status int,
	versionID string,
	lastModified interface{},
	resource interface{},
	conditional bool,
) error {
	var ts time.Time
	switch t := lastModified.(type) {
	case time.Time:
		ts = t
	case string:
		parsedTS, err := time.Parse(time.RFC3339, t)
		if err != nil {
			ts = time.Now()
		} else {
			ts = parsedTS
		}
	}
	skip := false
	if conditional {
		if since := req.Header.Get("If-Modified-Since"); since != "" {
			sinceTs, err := time.Parse(http.TimeFormat, since)
			if err != nil {
				return errors.Wrap(err, "unable to parse If-Modified-Since timestamp")
			}
			if ts.Equal(sinceTs) || ts.Before(sinceTs) {
				skip = true
			}
		}
		if etag := req.Header.Get("If-None-Match"); etag != "" && versionID != "" {
			if etag == generateETag(versionID) {
				skip = true
			}
		}
	}
	if skip {
		rw.WriteHeader(http.StatusNotModified)
		return nil
	}
	if versionID != "" {
		rw.Header().Set("Etag", generateETag(versionID))
	}
	rw.Header().Set("Last-Modified", ts.Format(http.TimeFormat))
	rndr.JSON(rw, status, resource)
	return nil
}

func validateJSONResource(h resourceHandler) http.Handler {
	log := h.getLogger()
	rndr := h.getRenderer()
	validator := h.getJSONValidator()

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resource, err := getValidationParameters(req, h)
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
		issues := []*models.OperationOutcomeIssue{}
		if valid {
			issues = append(issues, &models.OperationOutcomeIssue{
				Severity:    models.OperationOutcomeIssueSeverityInformation,
				Code:        models.OperationOutcomeIssueCodeInformational,
				Diagnostics: "no issues detected",
			})
		} else {
			for _, e := range vErrs {
				issues = append(issues, &models.OperationOutcomeIssue{
					Severity:    models.OperationOutcomeIssueSeverityError,
					Code:        models.OperationOutcomeIssueCodeInvalid,
					Diagnostics: e.String(),
				})
			}
		}
		outcome := &models.OperationOutcome{}
		outcome.Issue = issues
		rndr.JSON(rw, http.StatusOK, outcome)
	})
}

// generateResourceVersionID creates a version string for a resource based on an update nonce and the block number of the previous version
// must adhere to "id" regexp: https://www.hl7.org/fhir/datatypes.html#id
func generateResourceVersionID(updateCount uint, previousBlockNumber *big.Int) string {
	return fmt.Sprintf("%d%s%s", updateCount, versionDelimiter, previousBlockNumber.String())
}

func getUpdateCountFromVersionID(versionID string) (uint, error) {
	c, err := strconv.ParseUint(strings.Split(versionID, versionDelimiter)[0], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to parse update count from version string %q", versionID)
	}
	return uint(c), nil
}
