package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SnowUser struct {
	Result []struct {
		CalendarIntegration    string `json:"calendar_integration"`
		Country                string `json:"country"`
		UserPassword           string `json:"user_password"`
		LastLoginTime          string `json:"last_login_time"`
		Source                 string `json:"source"`
		SysUpdatedOn           string `json:"sys_updated_on"`
		Building               string `json:"building"`
		WebServiceAccessOnly   string `json:"web_service_access_only"`
		Notification           string `json:"notification"`
		EnableMultifactorAuthn string `json:"enable_multifactor_authn"`
		SysUpdatedBy           string `json:"sys_updated_by"`
		SysCreatedOn           string `json:"sys_created_on"`
		SysDomain              struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		State              string `json:"state"`
		Vip                string `json:"vip"`
		SysCreatedBy       string `json:"sys_created_by"`
		Zip                string `json:"zip"`
		HomePhone          string `json:"home_phone"`
		TimeFormat         string `json:"time_format"`
		LastLogin          string `json:"last_login"`
		DefaultPerspective string `json:"default_perspective"`
		Active             string `json:"active"`
		SysDomainPath      string `json:"sys_domain_path"`
		CostCenter         struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"cost_center"`
		Phone                   string `json:"phone"`
		Name                    string `json:"name"`
		EmployeeNumber          string `json:"employee_number"`
		PasswordNeedsReset      string `json:"password_needs_reset"`
		Gender                  string `json:"gender"`
		City                    string `json:"city"`
		FailedAttempts          string `json:"failed_attempts"`
		UserName                string `json:"user_name"`
		Roles                   string `json:"roles"`
		Title                   string `json:"title"`
		SysClassName            string `json:"sys_class_name"`
		SysID                   string `json:"sys_id"`
		FederatedID             string `json:"federated_id"`
		InternalIntegrationUser string `json:"internal_integration_user"`
		LdapServer              string `json:"ldap_server"`
		MobilePhone             string `json:"mobile_phone"`
		Street                  string `json:"street"`
		Company                 struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"company"`
		Department struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"department"`
		FirstName         string `json:"first_name"`
		Email             string `json:"email"`
		Introduction      string `json:"introduction"`
		PreferredLanguage string `json:"preferred_language"`
		Manager           string `json:"manager"`
		LockedOut         string `json:"locked_out"`
		SysModCount       string `json:"sys_mod_count"`
		LastName          string `json:"last_name"`
		Photo             string `json:"photo"`
		Avatar            string `json:"avatar"`
		MiddleName        string `json:"middle_name"`
		SysTags           string `json:"sys_tags"`
		TimeZone          string `json:"time_zone"`
		Schedule          string `json:"schedule"`
		DateFormat        string `json:"date_format"`
		Location          struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"location"`
	} `json:"result"`
}

type AddSnowUser struct {
	UserName       string `json:"user_name"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	EmployeeNumber string `json:"employee_number"`
	Password       string `json:"password"`
}

type NewSnowUser struct {
	Result struct {
		CalendarIntegration    string `json:"calendar_integration"`
		Country                string `json:"country"`
		UserPassword           string `json:"user_password"`
		LastLoginTime          string `json:"last_login_time"`
		Source                 string `json:"source"`
		SysUpdatedOn           string `json:"sys_updated_on"`
		Building               string `json:"building"`
		WebServiceAccessOnly   string `json:"web_service_access_only"`
		Notification           string `json:"notification"`
		EnableMultifactorAuthn string `json:"enable_multifactor_authn"`
		SysUpdatedBy           string `json:"sys_updated_by"`
		SysCreatedOn           string `json:"sys_created_on"`
		SysDomain              struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		State                   string `json:"state"`
		Vip                     string `json:"vip"`
		SysCreatedBy            string `json:"sys_created_by"`
		Zip                     string `json:"zip"`
		HomePhone               string `json:"home_phone"`
		TimeFormat              string `json:"time_format"`
		LastLogin               string `json:"last_login"`
		DefaultPerspective      string `json:"default_perspective"`
		Active                  string `json:"active"`
		SysDomainPath           string `json:"sys_domain_path"`
		CostCenter              string `json:"cost_center"`
		Phone                   string `json:"phone"`
		Name                    string `json:"name"`
		EmployeeNumber          string `json:"employee_number"`
		PasswordNeedsReset      string `json:"password_needs_reset"`
		Gender                  string `json:"gender"`
		City                    string `json:"city"`
		FailedAttempts          string `json:"failed_attempts"`
		UserName                string `json:"user_name"`
		Roles                   string `json:"roles"`
		Title                   string `json:"title"`
		SysClassName            string `json:"sys_class_name"`
		SysID                   string `json:"sys_id"`
		FederatedID             string `json:"federated_id"`
		InternalIntegrationUser string `json:"internal_integration_user"`
		LdapServer              string `json:"ldap_server"`
		MobilePhone             string `json:"mobile_phone"`
		Street                  string `json:"street"`
		Company                 string `json:"company"`
		Department              string `json:"department"`
		FirstName               string `json:"first_name"`
		Email                   string `json:"email"`
		Introduction            string `json:"introduction"`
		PreferredLanguage       string `json:"preferred_language"`
		Manager                 string `json:"manager"`
		LockedOut               string `json:"locked_out"`
		SysModCount             string `json:"sys_mod_count"`
		LastName                string `json:"last_name"`
		Photo                   string `json:"photo"`
		Avatar                  string `json:"avatar"`
		MiddleName              string `json:"middle_name"`
		SysTags                 string `json:"sys_tags"`
		TimeZone                string `json:"time_zone"`
		Schedule                string `json:"schedule"`
		DateFormat              string `json:"date_format"`
		Location                string `json:"location"`
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

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
	fmt.Println(err)
	}

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

// CreateUser - Create a  New User
func (s *UserClient) CreateUser(n AddSnowUser) (NewSnowUser, error) {

	u := NewSnowUser{}

	userJSON, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return u, err
	}

	url := s.Instance + "/api/now/table/sys_user"

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

// UpdateUser - update users
func (s *UserClient) UpdateUser(id string, n AddSnowUser) (NewSnowUser, error) {

	u := NewSnowUser{}

	userJSON, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return u, err
	}

	url := s.Instance + "/api/now/table/sys_user/"+id

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(userJSON))
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