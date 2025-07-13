package snowincidents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Incident struct {
	Result struct {
		Parent              string `json:"parent"`
		MadeSLA             string `json:"made_sla"`
		CausedBy            string `json:"caused_by"`
		WatchList           string `json:"watch_list"`
		UponReject          string `json:"upon_reject"`
		SysUpdatedOn        string `json:"sys_updated_on"`
		ChildIncidents      string `json:"child_incidents"`
		HoldReason          string `json:"hold_reason"`
		OriginTable         string `json:"origin_table"`
		TaskEffectiveNumber string `json:"task_effective_number"`
		ApprovalHistory     string `json:"approval_history"`
		Number              string `json:"number"`
		ResolvedBy          string `json:"resolved_by"`
		SysUpdatedBy        string `json:"sys_updated_by"`
		OpenedBy            struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"opened_by"`
		UserInput    string `json:"user_input"`
		SysCreatedOn string `json:"sys_created_on"`
		SysDomain    struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		State            string `json:"state"`
		RouteReason      string `json:"route_reason"`
		SysCreatedBy     string `json:"sys_created_by"`
		Knowledge        string `json:"knowledge"`
		Order            string `json:"order"`
		CalendarStc      string `json:"calendar_stc"`
		ClosedAt         string `json:"closed_at"`
		CmdbCi           string `json:"cmdb_ci"`
		DeliveryPlan     string `json:"delivery_plan"`
		Contract         string `json:"contract"`
		Impact           string `json:"impact"`
		Active           string `json:"active"`
		WorkNotesList    string `json:"work_notes_list"`
		BusinessService  string `json:"business_service"`
		BusinessImpact   string `json:"business_impact"`
		Priority         string `json:"priority"`
		SysDomainPath    string `json:"sys_domain_path"`
		Rfc              string `json:"rfc"`
		TimeWorked       string `json:"time_worked"`
		ExpectedStart    string `json:"expected_start"`
		OpenedAt         string `json:"opened_at"`
		BusinessDuration string `json:"business_duration"`
		GroupList        string `json:"group_list"`
		WorkEnd          string `json:"work_end"`
		CallerID         struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"caller_id"`
		ReopenedTime       string `json:"reopened_time"`
		ResolvedAt         string `json:"resolved_at"`
		ApprovalSet        string `json:"approval_set"`
		Subcategory        string `json:"subcategory"`
		WorkNotes          string `json:"work_notes"`
		UniversalRequest   string `json:"universal_request"`
		ShortDescription   string `json:"short_description"`
		CloseCode          string `json:"close_code"`
		CorrelationDisplay string `json:"correlation_display"`
		DeliveryTask       string `json:"delivery_task"`
		WorkStart          string `json:"work_start"`
		AssignmentGroup    struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"assignment_group"`
		AdditionalAssigneeList string `json:"additional_assignee_list"`
		BusinessStc            string `json:"business_stc"`
		Cause                  string `json:"cause"`
		Description            string `json:"description"`
		OriginID               string `json:"origin_id"`
		CalendarDuration       string `json:"calendar_duration"`
		CloseNotes             string `json:"close_notes"`
		Notify                 string `json:"notify"`
		ServiceOffering        string `json:"service_offering"`
		SysClassName           string `json:"sys_class_name"`
		ClosedBy               string `json:"closed_by"`
		FollowUp               string `json:"follow_up"`
		ParentIncident         string `json:"parent_incident"`
		SysID                  string `json:"sys_id"`
		ContactType            string `json:"contact_type"`
		ReopenedBy             string `json:"reopened_by"`
		IncidentState          string `json:"incident_state"`
		Urgency                string `json:"urgency"`
		ProblemID              string `json:"problem_id"`
		Company                string `json:"company"`
		ReassignmentCount      string `json:"reassignment_count"`
		ActivityDue            string `json:"activity_due"`
		AssignedTo             struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"assigned_to"`
		Severity             string `json:"severity"`
		Comments             string `json:"comments"`
		Approval             string `json:"approval"`
		SLADue               string `json:"sla_due"`
		CommentsAndWorkNotes string `json:"comments_and_work_notes"`
		DueDate              string `json:"due_date"`
		SysModCount          string `json:"sys_mod_count"`
		ReopenCount          string `json:"reopen_count"`
		SysTags              string `json:"sys_tags"`
		Escalation           string `json:"escalation"`
		UponApproval         string `json:"upon_approval"`
		CorrelationID        string `json:"correlation_id"`
		Location             string `json:"location"`
		Category             string `json:"category"`
	} `json:"result"`
}

type NewIncident struct {
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	CallerID         string `json:"caller_id"`
	Category         string `json:"category"`
	Subcategory      string `json:"subcategory"`
	Impact           string `json:"impact"`
	Urgency          string `json:"urgency"`
	AssignmentGroup  string `json:"assignment_group"`
	AssignedTo       string `json:"assigned_to"`
}

// IncidentClient defines the Okta Client
type IncidentClient struct {
	Admin         string
	AdminPassword string
	Instance      string
}

// NewIncidentClient - defines a new Okta Api client
func NewIncidentClient(admin string, password string, instance string) *IncidentClient {
	return &IncidentClient{
		Admin:         admin,
		AdminPassword: password,
		Instance:      instance,
	}
}

// CreateIncident - Create a new Incident
func (s *IncidentClient) CreateTicket(n NewIncident) (Incident, error) {

	u := Incident{}

	userJSON, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return u, err
	}

	url := s.Instance + "/api/now/table/incident"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	req.SetBasicAuth(s.Admin, s.AdminPassword)
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return u, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&u)

	if err != nil {
		return u, err
	}

	return u, err

}
