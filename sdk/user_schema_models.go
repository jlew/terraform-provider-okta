package sdk

import "encoding/json"

type (
	UserSchema struct {
		Schema      string                 `json:"$schema,omitempty"`
		Created     string                 `json:"created,omitempty"`
		Definitions *UserSchemaDefinitions `json:"definitions,omitempty"`
		ID          string                 `json:"id,omitempty"`
		LastUpdated string                 `json:"lastUpdated,omitempty"`
		Name        string                 `json:"name,omitempty"`
		Properties  *UserSchemaProperties  `json:"properties,omitempty"`
		Title       string                 `json:"title,omitempty"`
		Type        string                 `json:"type,omitempty"`
	}

	UserSchemaPermission struct {
		Action    string `json:"action,omitempty"`
		Principal string `json:"principal,omitempty"`
	}

	UserSchemaPropertyProfile struct {
		AllOf []*UserSchemaRef `json:"allOf,omitempty"`
	}

	UserSchemaDefinitions struct {
		Base   *UserSubSchemaProperties `json:"base,omitempty"`
		Custom *UserSubSchemaProperties `json:"custom,omitempty"`
	}

	UserSchemaItem struct {
		Enum  []string          `json:"enum,omitempty"`
		OneOf []*UserSchemaEnum `json:"oneOf,omitempty"`
		Type  string            `json:"type,omitempty"`
	}

	UserSchemaMaster struct {
		Type string `json:"type,omitempty"`
	}

	UserSchemaEnum struct {
		Const string `json:"const,omitempty"`
		Title string `json:"title,omitempty"`
	}

	UserSubSchema struct {
		Description       string                  `json:"description,omitempty"`
		Enum              []string                `json:"enum,omitempty"`
		Format            string                  `json:"format,omitempty"`
		Items             *UserSchemaItem         `json:"items,omitempty"`
		Master            *UserSchemaMaster       `json:"master,omitempty"`
		MaxLength         *int                    `json:"maxLength,omitempty"`
		MinLength         *int                    `json:"minLength,omitempty"`
		Mutability        string                  `json:"mutability,omitempty"`
		OneOf             []*UserSchemaEnum       `json:"oneOf,omitempty"`
		Pattern           *string                 `json:"pattern,omitempty"`
		Permissions       []*UserSchemaPermission `json:"permissions,omitempty"`
		Required          *bool                   `json:"required,omitempty"`
		Scope             string                  `json:"scope,omitempty"`
		Title             string                  `json:"title,omitempty"`
		Type              string                  `json:"type,omitempty"`
		Union             string                  `json:"union,omitempty"`
		Unique            string                  `json:"unique,omitempty"`
		ExternalName      string                  `json:"externalName,omitempty"`
		ExternalNamespace string                  `json:"externalNamespace,omitempty"`
		IsLogin           bool                    `json:"-"`
	}

	UserSubSchemaProperties struct {
		ID         string                    `json:"id,omitempty"`
		Properties map[string]*UserSubSchema `json:"properties,omitempty"`
		Required   []interface{}             `json:"required,omitempty"`
		Type       string                    `json:"type,omitempty"`
	}

	UserSchemaProperties struct {
		Profile *UserSchemaPropertyProfile `json:"profile,omitempty"`
	}

	UserSchemaRef struct {
		Ref string `json:"$ref,omitempty"`
	}
)

// This is workaround for issue, when we should set 'pattern' to 'null' explicitly to use default Email Format for login
// but for other cases `null` causes 500 error
func (u *UserSubSchema) MarshalJSON() ([]byte, error) {
	type localIDX UserSubSchema
	m, err := json.Marshal((*localIDX)(u))
	if !u.IsLogin {
		return m, err
	}
	if err != nil {
		return nil, err
	}
	var a interface{}
	_ = json.Unmarshal(m, &a)
	b := a.(map[string]interface{})
	p := b["pattern"]
	if p == nil || p.(string) == "" {
		b["pattern"] = nil
	}
	ret, err := json.Marshal(b)
	return ret, err
}
