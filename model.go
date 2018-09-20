package gogremlin

type Vertex struct {
	Type  string      `json:"@type"`
	Value VertexValue `json:"@value"`
}

type VertexValue struct {
	ID         ID             `json:"id"`
	Label      string         `json:"label"`
	Properties PropertyDetail `json:"properties,omitempty"`
}
type ID struct {
	Type  string `json:"@type"`
	Value int64  `json:"@value"`
}

type Data struct {
	Label      string
	Name       string
	ID         string
	Properties map[string]string
}

type APIData struct {
	Label      string            `json:"dbName"`
	Vertex     string            `json:"tableName"`
	Properties map[string]string `json:"cols"`
}

type Property struct {
	Type  string `json:@type`
	Value string `json:"@value"`
}

type PropertyDetail map[string][]PropertyInfo

type PropertyInfo struct {
	Type  string        `json:"@type"`
	Value PropertyValue `json:"@value"`
}
type PropertyValue struct {
	ID    PropertyID `json:"id"`
	Value string     `json:"value"`
	Label string     `json:"label"`
}

type PropertyID struct {
	Type  string          `json::"@type"`
	Value PropertyIDValue `json:"@value"`
}

type PropertyIDValue struct {
	RelationID string `json:"relationId"`
}

type Edge struct {
}
