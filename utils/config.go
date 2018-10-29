package utils

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// SysConfig defines system configurations
type SysConfig struct {
	AuthMode                   string `yaml:"auth_mode" json:"auth_mode"`
	EmailFrom                  string `yaml:"email_from" json:"email_from"`
	EmailHost                  string `yaml:"email_host" json:"email_host"`
	EmailPort                  int    `yaml:"email_port" json:"email_port"`
	EmailIdentity              string `yaml:"email_identity" json:"email_identity"`
	EmailUsername              string `yaml:"email_username" json:"email_username"`
	EmailSsl                   bool   `yaml:"email_ssl" json:"email_ssl"`
	EmailInsecure              bool   `yaml:"email_insecure" json:"email_insecure"`
	LdapURL                    string `yaml:"ldap_url" json:"ldap_url"`
	LdapBaseDN                 string `yaml:"ldap_base_dn" json:"ldap_base_dn"`
	LdapFilter                 string `yaml:"ldap_filter" json:"ldap_filter"`
	LdapScope                  int    `yaml:"ldap_scope" json:"ldap_scope"`
	LdapUID                    string `yaml:"ldap_uid" jsonb:"ldap_uid"`
	LdapSearchDN               string `yaml:"ldap_search_dn" json:"ldap_search_dn"`
	LdapTimeout                int    `yaml:"ldap_timeout" json:"ldap_timeout"`
	ProjectCreationRestriction string `yaml:"project_creation_restriction" json:"project_creation_restriction"`
	SelfRegistration           bool   `yaml:"self_registration" json:"self_registration"`
	TokenExpiration            int    `yaml:"token_expiration" json:"token_expiration"`
	VerifyRemoteCert           bool   `yaml:"verify_remote_cert" json:"verify_remote_cert"`
	ScanAllPolicy              struct {
		Type      string `yaml:"type" json:"type"`
		Parameter struct {
			DailyTime int `yaml:"daily_time" json:"daily_time"`
		} `yaml:"parameter" json:"parameter"`
	} `yaml:"scan_all_policy" json:"scan_all_policy"`
}

// SysConfigLoad loads system configuration from conf/config.yaml.
func SysConfigLoad() (*SysConfig, error) {
	var config SysConfig

	dataBytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(dataBytes), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
