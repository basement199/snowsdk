package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UHrBusinessPartner struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `u_hr_business_partner` is not an object
}

func (b *UHrBusinessPartner) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type PartnerObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj PartnerObj
	if err := json.Unmarshal(data, &obj); err == nil {
		b.Link = obj.Link
		b.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	b.Raw = raw
	return nil
}

type Building struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `building` is not an object
}

func (b *Building) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type BuildingObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj BuildingObj
	if err := json.Unmarshal(data, &obj); err == nil {
		b.Link = obj.Link
		b.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	b.Raw = raw
	return nil
}

type Location struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `location` is not an object
}

func (l *Location) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type LocationObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj LocationObj
	if err := json.Unmarshal(data, &obj); err == nil {
		l.Link = obj.Link
		l.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	l.Raw = raw
	return nil
}

type Department struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `department` is not an object
}

func (d *Department) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type DepartmentObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj DepartmentObj
	if err := json.Unmarshal(data, &obj); err == nil {
		d.Link = obj.Link
		d.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	d.Raw = raw
	return nil
}

type Company struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `company` is not an object
}

func (c *Company) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type CompanyObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj CompanyObj
	if err := json.Unmarshal(data, &obj); err == nil {
		c.Link = obj.Link
		c.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.Raw = raw
	return nil
}

type Manager struct {
	Link  string
	Value string
	Raw   string // Stores the raw string if `manager` is not an object
}

func (m *Manager) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an object
	type ManagerObj struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	}
	var obj ManagerObj
	if err := json.Unmarshal(data, &obj); err == nil {
		m.Link = obj.Link
		m.Value = obj.Value
		return nil
	}

	// Otherwise, assume it's a string
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	m.Raw = raw
	return nil
}

type SnowUser struct {
	Result []struct {
		CalendarIntegration           string             `json:"calendar_integration"`
		UResourceEnded                string             `json:"u_resource_ended"`
		UUserprincipalname            string             `json:"u_userprincipalname"`
		SysUpdatedOn                  string             `json:"sys_updated_on"`
		Building                      Building           `json:"building"`
		UMailCounter                  string             `json:"u_mail_counter"`
		SsoSource                     string             `json:"sso_source"`
		UFullDomain                   string             `json:"u_full_domain"`
		UObjectsid                    string             `json:"u_objectsid"`
		State                         string             `json:"state"`
		Vip                           string             `json:"vip"`
		SysCreatedBy                  string             `json:"sys_created_by"`
		Zip                           string             `json:"zip"`
		UJobFunction                  string             `json:"u_job_function"`
		TimeFormat                    string             `json:"time_format"`
		UStageDn                      string             `json:"u_stage_dn"`
		LastLogin                     string             `json:"last_login"`
		USubdivision                  string             `json:"u_subdivision"`
		Active                        string             `json:"active"`
		UCcbCostCenter                string             `json:"u_ccb_cost_center"`
		UCubicleNo                    string             `json:"u_cubicle_no"`
		UNetworkLoginEnabled          string             `json:"u_network_login_enabled"`
		UItResource                   string             `json:"u_it_resource"`
		ULastLogonTimeException       string             `json:"u_last_logon_time_exception"`
		SysDomainPath                 string             `json:"sys_domain_path"`
		TransactionLog                string             `json:"transaction_log"`
		UOrgMgrs                      string             `json:"u_org_mgrs"`
		Costcenter                    string             `json:"costcenter"`
		Phone                         string             `json:"phone"`
		EmployeeNumber                string             `json:"employee_number"`
		UWorkforceCategory            string             `json:"u_workforce_category"`
		XMobiPWelcomeNotificationSent string             `json:"x_mobi_p_welcome_notification_sent"`
		Gender                        string             `json:"gender"`
		UIgnoreByMigration            string             `json:"u_ignore_by_migration"`
		City                          string             `json:"city"`
		ULocationType                 string             `json:"u_location_type"`
		UserName                      string             `json:"user_name"`
		ULdapLastLoginStamp           string             `json:"u_ldap_last_login_stamp"`
		UDsktpAdminAssigned           string             `json:"u_dsktp_admin_assigned"`
		UResourceStarted              string             `json:"u_resource_started"`
		SysClassName                  string             `json:"sys_class_name"`
		UHrBusinessPartner            UHrBusinessPartner `json:"u_hr_business_partner"`
		UWorkforceType                string             `json:"u_workforce_type"`
		UOtherPhone                   string             `json:"u_other_phone"`
		UContTeEligFlag               string             `json:"u_cont_te_elig_flag"`
		UOffice                       string             `json:"u_office"`
		UFgSecid                      string             `json:"u_fg_secid"`
		UActivesync                   string             `json:"u_activesync"`
		Email                         string             `json:"email"`
		URegion                       string             `json:"u_region"`
		UHrForHr                      string             `json:"u_hr_for_hr"`
		Manager                       Manager            `json:"manager"`
		LockedOut                     string             `json:"locked_out"`
		UAuthorityLevel               string             `json:"u_authority_level"`
		LastName                      string             `json:"last_name"`
		Photo                         string             `json:"photo"`
		UUpdateSources                string             `json:"u_update_sources"`
		Avatar                        string             `json:"avatar"`
		UBadgeExpiryDate              string             `json:"u_badge_expiry_date"`
		URoom                         string             `json:"u_room"`
		UAccNvrExpires                string             `json:"u_acc_nvr_expires"`
		CorrelationID                 string             `json:"correlation_id"`
		DateFormat                    string             `json:"date_format"`
		UBusinessUnit                 string             `json:"u_business_unit"`
		Country                       string             `json:"country"`
		UEmployeeClass                string             `json:"u_employee_class"`
		LastLoginTime                 string             `json:"last_login_time"`
		UDivision                     string             `json:"u_division"`
		XPdIntegrationPagerdutyID     string             `json:"x_pd_integration_pagerduty_id"`
		ULastDayWorked                string             `json:"u_last_day_worked"`
		Source                        string             `json:"source"`
		UPreferredPronoun             string             `json:"u_preferred_pronoun"`
		UServiceAcct                  string             `json:"u_service_acct"`
		UPersonalMobilePhone          string             `json:"u_personal_mobile_phone"`
		XMobiCExternalUserLastName    string             `json:"x_mobi_c_external_user_last_name"`
		WebServiceAccessOnly          string             `json:"web_service_access_only"`
		UExtEmplID                    string             `json:"u_ext_empl_id"`
		UNotifyVip                    string             `json:"u_notify_vip"`
		SysCreatedOn                  string             `json:"sys_created_on"`
		UItilUser                     string             `json:"u_itil_user"`
		UWfqResourceID                string             `json:"u_wfq_resource_id"`
		UBudgetCodes                  string             `json:"u_budget_codes"`
		UPhotoUploaded                string             `json:"u_photo_uploaded"`
		HomePhone                     string             `json:"home_phone"`
		UContProjectedEndDate         string             `json:"u_cont_projected_end_date"`
		UEmployeeResourceGroups       string             `json:"u_employee_resource_groups"`
		USeniorityDate                string             `json:"u_seniority_date"`
		AverageDailyFte               string             `json:"average_daily_fte"`
		UHrPersonnel                  string             `json:"u_hr_personnel"`
		UPimNumber                    string             `json:"u_pim_number"`
		UUserid                       string             `json:"u_userid"`
		UHrManagerPath                string             `json:"u_hr_manager_path"`
		Name                          string             `json:"name"`
		XViinBfCustAllCatalog         string             `json:"x_viin_bf_cust_all_catalog"`
		FailedAttempts                string             `json:"failed_attempts"`
		UCompanyDivision              string             `json:"u_company_division"`
		Title                         string             `json:"title"`
		SysID                         string             `json:"sys_id"`
		UPwdNvrExprs                  string             `json:"u_pwd_nvr_exprs"`
		FederatedID                   string             `json:"federated_id"`
		UDirExclude                   string             `json:"u_dir_exclude"`
		InternalIntegrationUser       string             `json:"internal_integration_user"`
		MobilePhone                   string             `json:"mobile_phone"`
		Street                        string             `json:"street"`
		UAltEmail                     string             `json:"u_alt_email"`
		Company                       Company            `json:"company"`
		Department                    Department         `json:"department"`
		UType                         string             `json:"u_type"`
		FirstName                     string             `json:"first_name"`
		UPreferredName                string             `json:"u_preferred_name"`
		XMobiCExternalPhoneNumber     string             `json:"x_mobi_c_external_phone_number"`
		Introduction                  string             `json:"introduction"`
		PreferredLanguage             string             `json:"preferred_language"`
		UNotes                        string             `json:"u_notes"`
		XMobiCExternalUserFirstName   string             `json:"x_mobi_c_external_user_first_name"`
		UCluster                      string             `json:"u_cluster"`
		UNumber                       string             `json:"u_number"`
		UGalName                      string             `json:"u_gal_name"`
		UFloor                        string             `json:"u_floor"`
		SysModCount                   string             `json:"sys_mod_count"`
		UHumanMailbox                 string             `json:"u_human_mailbox"`
		ULockOutDate                  string             `json:"u_lock_out_date"`
		MiddleName                    string             `json:"middle_name"`
		SysTags                       string             `json:"sys_tags"`
		TimeZone                      string             `json:"time_zone"`
		ULocalUser                    string             `json:"u_local_user"`
		UWorkAtHome                   string             `json:"u_work_at_home"`
		XMobiCMdmUserName             string             `json:"x_mobi_c_mdm_user_name"`
		Location                      Location           `json:"location"`
		UNameID                       string             `json:"u_name_id"`
		UObjectguid                   string             `json:"u_objectguid"`
		UAssignmentCategory           string             `json:"u_assignment_category"`
	} `json:"result"`
}

// UserClient defines the Okta Client
type UserClient struct {
	Admin         string
	AdminPassword string
	Instance      string
}

// NewUserClient - defines a new Okta Api client
func NewUserClient(admin string, password string, instance string) *UserClient {
	return &UserClient{
		Admin:         admin,
		AdminPassword: password,
		Instance:      instance,
	}
}

// GetUserId search for user by SNOW  object ID
func (s *UserClient) GetUserId(userId string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=sys_id=" + userId

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetUserName search for user by user_name
func (s *UserClient) GetUserName(userName string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=user_name=" + userName

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetUPN search for user by user_
func (s *UserClient) GetUPN(upn string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=u_userprincipalname=" + upn

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

//GetObjectGUID - search for user by objectGUID

func (s *UserClient) GetObjectGUID(objectguid string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=u_objectguid=" + objectguid

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetEmail search for user by email
func (s *UserClient) GetEmail(email string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=email=" + email

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetEmployeeNumber search for user by employee_number
func (s *UserClient) GetEmployeeNumber(employeeNumber string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user?sysparm_query=employee_number=" + employeeNumber

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetAll - get all users
func (s *UserClient) GetAll() (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance + "/api/now/table/sys_user"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}

	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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

// GetUpdatedUsers - get all users update since a specified date
func (s *UserClient) GetUpdatedUsers(sinceDate string) (SnowUser, error) {

	u := SnowUser{}

	url := s.Instance+"/api/now/table/sys_user?"+"sysparm_query=sys_updated_on>=javascript:gs.dateGenerate("+"'"+sinceDate+"'"+",'00:00:00')"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return u, err
	}
	
	// Set the basic authentication header
	req.SetBasicAuth(s.Admin, s.AdminPassword)

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
