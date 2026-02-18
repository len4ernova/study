package main

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// bindAndSearchWithParams - связывание с сервером и получение данных.
func bindAndSearchWithParams(l *ldap.Conn, user, pass, basedn, filter string) (*ldap.SearchResult, error) {
	err := l.Bind("uid="+user+",cn=users,cn=accounts,dc=iifrf,dc=local", pass)
	if err != nil {
		//"Ошибка связывания"
		return nil, err
	}

	searchReq := ldap.NewSearchRequest(
		basedn,                 // Base DN
		ldap.ScopeWholeSubtree, // Scope
		ldap.NeverDerefAliases, // Aliases
		0,                      // SizeLimit
		0,                      // TimeLimit
		false,                  // TypesOnly
		filter,                 // Filter
		[]string{"cn"},         // Attributes to retrieve
		nil,                    // Controls
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	for _, entry := range result.Entries {
		fmt.Printf("Entry DN: %s\n", entry.DN)
		for _, attribute := range entry.Attributes {
			fmt.Printf("  Attribute: %s = %v\n", attribute.Name, attribute.Values)
		}
	}

	if len(result.Entries) > 0 {
		// result.Entries[0].Print()
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

// authLdap - аутентификация ldap, в случае ошибки вернётся err.
func authLdap(urlLDAP, user, pass, basedn, filter string) error {
	l, err := ldap.DialURL(urlLDAP)
	if err != nil {
		// "Ошибка подключения к ldap"
		return err
	}
	defer l.Close()

	_, err = bindAndSearchWithParams(l, user, pass, basedn, filter)
	if err != nil {
		return err
	}
	//result.Entries[0].Print()
	return nil
}

func main() {
	user := "user_name"
	pass := "pass.."
	basedn := "dc=local"
	filter := "(uid=user_name)"
	ldapURL := "ldap://xxx.xxx.xxx.xxx:389"

	err := authLdap(ldapURL, user, pass, basedn, filter)
	if err != nil {
		fmt.Println(":(")
	} else {
		fmt.Println(":)")
	}

}
